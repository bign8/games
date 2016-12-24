package impl

import (
	"github.com/bign8/games"
	"github.com/bign8/games/impl/checkers"
	"github.com/bign8/games/impl/chess"
	"github.com/bign8/games/impl/connect4"
	gos "github.com/bign8/games/impl/go"
	"github.com/bign8/games/impl/mancala"
	"github.com/bign8/games/impl/ttt"
)

// TODO: figure out how to make this immutible
var Registry = map[string]games.Game{
	ttt.Game.Slug:      ttt.Game,
	chess.Game.Slug:    chess.Game,
	checkers.Game.Slug: checkers.Game,
	connect4.Game.Slug: connect4.Game,
	gos.Game.Slug:      gos.Game,
	mancala.Game.Slug:  mancala.Game,
}
