package models

import (
	"fmt"
	"net/http"
)

type kind int

const (
	BasicAuth kind = iota
	BearerToken
	NoAuth
)

type Auth interface {
	Authenticate(req *http.Request) error
}

type DefaultAuth struct {
	Kind     kind
	Userame  string
	Password string
	Token    string
}

func (a *DefaultAuth) kind() string {
	switch a.Kind {
	case BasicAuth:
		return "BasicAuth"
	case BearerToken:
		return "BearerToken"
	case NoAuth:
		return "NoAuth"
	default:
		return "Unknown"
	}
}

func (a *DefaultAuth) Authenticate(req *http.Request) error {
	switch a.Kind {
	case BasicAuth:
		req.SetBasicAuth(a.Userame, a.Password)
	case BearerToken:
		req.Header.Set("Authorization", "Bearer "+a.Token)
	case NoAuth:
		// No authentication needed
	default:
		return fmt.Errorf("unknown authentication kind: %s", a.kind())
	}
	return nil
}
