package main

import "html/template"

func (app *application) templs() struct{ home, list, submissions *template.Template } {
	const home = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Form Resubmit</title>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	</head>
	<body>
		<main>
		<form action="/list" method="post" enctype="multipart/form-data">
			<input type="file" name="csvfile" accept=".csv" required>
			<input type="submit" value="uploadCSV">
		</form>
		</main>
	</body>
	</html>
	`

	const list string = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Form Resubmit</title>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	</head>
	<body>
		<main>
			<p>List of records</p>
			<form action="/submit" method="post">
			<table class="table-auto">
				<thead>
				<tr>
					<th>Select</th>
					<th>Record Id</th>
					<th>Form Name</th>
					<th>Enquiry Id</th>
					<th>First Name</th>
					<th>Last Name</th>
					<th>Email</th>
				</tr>
				</thead>
				<tbody>
				{{range .}}
					<tr>
						<td><input type="checkbox" id="{{.RecordId}}" name="{{.RecordId}}" value="{{.RecordId}}" /></td>
						<td>{{.RecordId}}</td>
						<td>{{.FormName}}</td>
						<td>{{.EnquiryId}}</td>
						<td>{{.FirstName}}</td>
						<td>{{.LastName}}</td>
						<td>{{.Email}}</td>
					</tr>
				{{end}}
				</tbody>
			</table>
			<input type="submit" value="submitForms">
			</form>
		</main>
	</body>
	</html>
	`
	const submissionsInProgress = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Form Resubmit</title>
		<meta charset="UTF-8">
		<script src="https://unpkg.com/htmx.org@1.9.10"></script>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	</head>
	<body>
		<main>
		<ul id="update" hx-trigger="every 1s" hx-get="/update">
		{{range .}}
		<li>Submitting {{.RecordId}}... {{.State}} {{.NewEnquiryId}}</li>
		{{end}}
		</ul>
		<button class="btn" hx-get="/update"
                        hx-target="#update">Update
		 </button>
		</main>
		<!--<script>
    setInterval(function() {
        htmx.trigger("#update", "load");
    }, 1000);
</script>-->
	</body>
	</html>`

	homeTemplate, err := template.New("home").Parse(home)
	if err != nil {
		app.errLogger.Fatal("Error parsing home template", err)
	}
	listTemplate, err := template.New("list").Parse(list)
	if err != nil {
		app.errLogger.Fatal("Error parsing list template", err)
	}
	submissionsTemplate, err := template.New("submissions").Parse(submissionsInProgress)
	if err != nil {
		app.errLogger.Fatal("Error parsing list template", err)
	}
	return struct{ home, list, submissions *template.Template }{homeTemplate, listTemplate, submissionsTemplate}
}
