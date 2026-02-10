package defaulthttp

import (
	"errors"
	"fmt"
)

var ErrCollection = errors.New("not all collections completed successfully")

type CollectionError string

func (e CollectionError) Error() string {
	msg := e
	return fmt.Sprintf("SWARM encountered a problem with collection #%s", msg)
}

type RunError int

func (e RunError) Error() string {
	id := e
	return fmt.Sprintf("SWARM encountered a problem with run #%d", id)
}
