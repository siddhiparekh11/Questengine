package interfaces

import (

	"github.com/siddhiparekh11/GoChallenge/models"
	"context"

)

type IQuest interface {

	CreateQuest(ctx context.Context, quest models.Quest) (bool,*models.QError)
	GetQuests(ctx context.Context) ([] *models.Quest, *models.QError)
	GetQuest(ctx context.Context, questId int)(*models.Quest, *models.QError)

}