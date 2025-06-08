package defaulthttp

import (
	"errors"
	"fmt"
)

var ErrCollection = errors.New("Not all collections completed successfully")

type collectionError string

func (e collectionError) Error() string {
	return fmt.Sprintf("SWARM encountered a problem with collection #%s", e)
}

type runError int

func (e runError) Error() string {
	return fmt.Sprintf("SWARM encountered a problem with run #%d", e)
}
