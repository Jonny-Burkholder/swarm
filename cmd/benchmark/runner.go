package benchmark

import (
	"context"

	"github.com/jonny-burkholder/swarm/internal/models"
)

type Runner interface {
	Run(ctx *context.Context, collections []*models.Collection) error
}
