package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"anz.com/wankhadv/resubmitutil/cmd/types"
	"anz.com/wankhadv/resubmitutil/cmd/utils"
	"github.com/gocarina/gocsv"
)

var recordState []types.Record
var submissionRecords []types.Record
var threadSafeStates struct {
	lock         sync.Mutex
	states       []types.SubmissionState
	updatedCount int
}

// func (app *application) updateStates(records []types.SubmissionState, worker int, doneWorkerChan chan int) {
// 	// time.Sleep(time.Duration(record.RecordId*2) * time.Second)
// 	app.infoLogger.Printf("Worker %d started for %v\n", worker, records)
// 	for i := range records {
// 		_, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", records[i].RecordId))
// 		threadSafeStates.lock.Lock()
// 		if err != nil {
// 			records[i].State = "FAILED"
// 		} else {
// 			records[i].State = "SUCCESS"
// 		}
// 		threadSafeStates.updatedCount += 1
// 		threadSafeStates.lock.Unlock()
// 	}
// 	doneWorkerChan <- worker
// }

func (app *application) updateStates(records []types.SubmissionState, worker int) {
	// time.Sleep(time.Duration(record.RecordId*2) * time.Second)
	app.infoLogger.Printf("Worker %d started for %v\n", worker, records)
	for i := range records {
		_, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", records[i].RecordId))
		threadSafeStates.lock.Lock()
		if err != nil {
			records[i].State = "FAILED"
		} else {
			records[i].State = "SUCCESS"
		}
		threadSafeStates.updatedCount += 1
		threadSafeStates.lock.Unlock()
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	app.infoLogger.Println(r.Method, "at /")

	app.templates.home.Execute(w, nil)
}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.infoLogger.Println(r.Method, "at /submit")
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			app.errLogger.Fatal("Error parsing form values: ", err)
		}
		// fmt.Println(recordState)
		// app.infoLogger.Println("Submitting :", utils.GetSubmissionRecords(recordState, r.PostForm))
		submissionRecords = utils.GetSubmissionRecords(recordState, r.PostForm)
		threadSafeStates.states = utils.GetStates(submissionRecords)
		// doneWorkerChan := make(chan int)
		// numGroups := len(threadSafeStates.states) / (*app.workers)
		// if numGroups == 0 {
		// 	numGroups = 1
		// }
		// app.infoLogger.Println("Number of thread groups", numGroups)
		// extraThreadNeeded := false
		// if len(threadSafeStates.states)%(*app.workers) != 0 {
		// 	extraThreadNeeded = true
		// }
		// app.infoLogger.Println("Extra thread needed", extraThreadNeeded)
		// for i := 0; i < numGroups; i++ {
		// 	app.infoLogger.Println("loop")
		// 	lo := i * (*app.workers)
		// 	hi := lo + (*app.workers)
		// 	if hi > len(threadSafeStates.states) {
		// 		hi = len(threadSafeStates.states)
		// 	}
		// 	app.updateStates(threadSafeStates.states[lo:hi], i, doneWorkerChan)
		// }
		// if extraThreadNeeded {
		// 	app.infoLogger.Println("Extra thread")
		// 	extraWorker := <-doneWorkerChan
		// 	app.updateStates(threadSafeStates.states[numGroups*(*app.workers):], extraWorker, doneWorkerChan)
		// }
		for i, j := 0, 0; i < len(threadSafeStates.states); i, j = i+*app.workers, j+1 {
			hi := i + *app.workers
			if hi > len(threadSafeStates.states) {
				hi = len(threadSafeStates.states)
			}
			go app.updateStates(threadSafeStates.states[i:hi], j)
		}
		// app.infoLogger.Println(threadSafeStates.states)
		http.Redirect(w, r, "/submissionState", http.StatusSeeOther)
	}

	if r.Method == http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) update(w http.ResponseWriter, r *http.Request) {
	if threadSafeStates.updatedCount == len(threadSafeStates.states) {
		app.infoLogger.Println("Stop polling")
		w.WriteHeader(286)
	}
	app.infoLogger.Println("GET at /update")
	app.infoLogger.Println(threadSafeStates.states)
	templ := `
	{{range .}}
	<li>Submitting {{.RecordId}}... {{.State}} {{.NewEnquiryId}}</li>
	{{end}}
	`
	listTempl, err := template.New("update").Parse(templ)
	if err != nil {
		app.errLogger.Fatal("Error parsing update template : ", err)
	}
	err = listTempl.Execute(w, threadSafeStates.states)
	if err != nil {
		app.errLogger.Fatal("Error rendering submissions states : ", err)
	}
}

func (app *application) submissionState(w http.ResponseWriter, r *http.Request) {
	app.infoLogger.Println("GET at /submissionState")
	err := app.templates.submissions.Execute(w, threadSafeStates.states)
	if err != nil {
		app.errLogger.Fatal("Error rendering submissions states : ", err)
	}
}

func (app *application) list(w http.ResponseWriter, r *http.Request) {
	app.infoLogger.Println(r.Method, "at /list")
	if r.Method == http.MethodPost {
		csvFile, _, err := r.FormFile("csvfile")
		var records []types.Record
		if err != nil {
			app.errLogger.Fatal("Error receeiving multipart csv file", err)
		}
		defer csvFile.Close()
		err = gocsv.UnmarshalMultipartFile(&csvFile, &records)
		if err != nil {
			app.errLogger.Fatal("Error parsing csv", err)
		}
		// app.infoLogger.Println("Received :", records)
		recordState = records
		err = app.templates.list.Execute(w, records)
		if err != nil {
			app.errLogger.Fatal("Error rendering html : ", err)
		}
	}

	if r.Method == http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
