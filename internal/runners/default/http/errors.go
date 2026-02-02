package defaulthttp

import (
	"errors"
	"fmt"
)

var ErrCollection = errors.New("not all collections completed successfully")

type collectionError string

func (e collectionError) Error() string {
	msg := e
	return fmt.Sprintf("SWARM encountered a problem with collection #%s", msg)
}

type runError int

func (e runError) Error() string {
	id := e
	return fmt.Sprintf("SWARM encountered a problem with run #%d", id)
}
