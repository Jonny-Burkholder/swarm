package models

type Request struct {
	Method      string
	Path        string
	Auth        Auth
	Headers     map[string]string
	QueryParams map[string][]string // to be parsed if collection is http
	Body        []byte
	Assert      []Assertion
}
