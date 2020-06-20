package repository

import (
	"os"
	"strings"
	"time"
)

type Job struct {
	Code string
	Time time.Time

	State string `dynamo:"State"`
}

// TableName returns the table name
func (j *Job) TableName() string {
	isLambda := strings.ToUpper(os.Getenv("LAMBDA"))

	if isLambda == "TRUE" {
		return os.Getenv("JOB_TABLE")
	} else {
		return "jobs"
	}
}
