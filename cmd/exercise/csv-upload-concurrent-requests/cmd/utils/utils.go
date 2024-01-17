// Package utils provides some utility functions
package utils

import (
	"net/url"
	"strconv"

	"anz.com/wankhadv/resubmitutil/cmd/types"
)

func GetSubmissionRecords(records []types.Record, values url.Values) []types.Record {
	submissionRecords := []types.Record{}
	for _, record := range records {
		if values.Has(strconv.FormatInt(int64(record.RecordId), 10)) {
			submissionRecords = append(submissionRecords, record)
		}
	}
	return submissionRecords
}

func GetStates(records []types.Record) []types.SubmissionState {
	states := []types.SubmissionState{}
	for _, record := range records {
		states = append(states, types.SubmissionState{RecordId: record.RecordId, State: "IN PROGRESS", NewEnquiryId: 0})
	}
	return states
}
