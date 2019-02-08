package interfaces

import (

	"github.com/siddhiparekh11/GoChallenge/models"
	"context"

)

type IPlayer interface {

	CreatePlayer(ctx context.Context, player models.Player) (bool,*models.QError)
	GetPlayers(ctx context.Context) ([] *models.Player, *models.QError)
	GetPlayer(ctx context.Context, playerId int) (*models.Player,*models.QError)

}