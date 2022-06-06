// Copyright 2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package das

import (
	"context"
	"errors"
	"fmt"

	"github.com/offchainlabs/nitro/arbstate"
)

var ErrNotFound = errors.New("Not found")

type StorageService interface {
	arbstate.DataAvailabilityReader
	Put(ctx context.Context, data []byte, expirationTime uint64) error
	Sync(ctx context.Context) error
	Closer
	fmt.Stringer
	HealthCheck(ctx context.Context) error
}
