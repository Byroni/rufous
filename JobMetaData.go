package main

type JobMetaData struct {
	jobID   string
	jobName string `dynamo:"jobName"`
}
