package cluster

import (
	"context"

	"github.com/determined-ai/determined/master/internal/authz"
	"github.com/determined-ai/determined/master/pkg/model"
)

// MiscAuthZ describes authz methods for misc actions.
type MiscAuthZ interface {
	/*
		- get master logs
		- get historical usage
		- get container associations. could be part of usage
		- view job queue
		- manipulate the job queue
	*/

	// CanUpdateAgents returns an error if the user is not authorized to manipulate agents.
	CanUpdateAgents(
		ctx context.Context, curUser *model.User,
	) (permErr error, err error)
}

// AuthZProvider is the authz registry for Notebooks, Shells, and Commands.
var AuthZProvider authz.AuthZProviderType[MiscAuthZ]
