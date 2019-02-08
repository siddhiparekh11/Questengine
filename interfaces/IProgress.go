package interfaces

import (

	"github.com/siddhiparekh11/GoChallenge/models"
	"context"

)

type IProgress interface {

	UpdateProgress(ctx context.Context, progress models.Progress) (*models.Progress,*models.QError)
	GetPlayerState(ctx context.Context, playerId int) (*models.Progress,*models.QError)
	ConstructPlayerStateStruct(progress *models.Progress) (*models.PlayerState)
	GetAllQuestProgress(ctx context.Context, questId int, gameId int) ([]*models.Progress,*models.QError)

}