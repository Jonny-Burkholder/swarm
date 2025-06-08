package models

import "sync"

type Collection struct {
	Name     string
	Requests []Request
	Mu       *sync.Mutex
	Runs     []Run
}
