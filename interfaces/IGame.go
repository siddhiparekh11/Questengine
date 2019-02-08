package interfaces

import (

	"github.com/siddhiparekh11/GoChallenge/models"
	"context"

)

type IGame interface {

	CreateGame(ctx context.Context, game models.Game) (bool,*models.QError)
	GetGames(ctx context.Context) ([] *models.Game, *models.QError)
	SetGameId(ctx context.Context, gameId int) (*models.Game,*models.QError)

}