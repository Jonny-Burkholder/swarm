package models

import "time"

type Result struct {
	Request
	StatusCode int
	Body       []byte
	Headers    map[string][]string
	Duration   time.Duration
	Assertions []Assertion
	Error      error
}
