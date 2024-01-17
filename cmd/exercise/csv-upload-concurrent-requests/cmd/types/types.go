// Package types exports some types
package types

type Record struct {
	RecordId  int    `csv:"record_id"`
	FormName  string `csv:"form_name"`
	EnquiryId int    `csv:"enquiry_id"`
	SubmitUrl string `csv:"submit_url"`
	FirstName string `csv:"first_name"`
	LastName  string `csv:"last_name"`
	Email     string `csv:"email"`
}

type SubmissionState struct {
	RecordId     int
	State        string
	NewEnquiryId int
	Error        string
}
