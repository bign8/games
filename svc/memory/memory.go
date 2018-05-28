package memory

import (
	"context"
	"errors"

	"github.com/bign8/games/svc"
)

type memory struct {
	players map[svc.PlayerID]interface{}
}

// New constructs a new memory persisted game service.
func New() svc.GameService {
	return &memory{
		players: nil,
	}
}

func (mem *memory) NewPlayer(context.Context, svc.GameSlug) (svc.PlayerID, error) {
	return "", errors.New("TODO")
}
