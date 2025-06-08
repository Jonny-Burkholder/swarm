package models

type Run struct {
	ID      int
	Results []Result
	Error   error
}
