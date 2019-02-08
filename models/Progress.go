package models


import "time"

type Progress struct {
	GameId int  
	QuestId int 
	PlayerId int 
	QuestPointsEarned int `json:QuestPointsEarned"`
	TotQuestCompPer int `json:"TotalQuestPercentCompleted"`
	LastMilestoneIndex int `json:"LastMilestoneIndexCompleted"`
	LastMilestone CusMilestone
	Milestones []CusMilestone
	CreatedTimestamp time.Time 
	ChipAmountBet int
}


type PlayerState struct {

	PlayerLevel int `json:PlayerLevel`
	TotalQuestPercentCompleted int `json:"TotalQuestPercentCompleted"`
	LastMilestoneIndex int `json:"LastMilestoneIndex"`

}

type ProgressView struct {

	PlayerLevel int `json:PlayerLevel`
	QuestPointsEarned int `json:QuestPointsEarned"`
	TotQuestCompPer int `json:"TotalQuestPercentCompleted"`
	Milestones []CusMilestone `json:"MilestonesCompleted"`

}

type Prog struct {
	PlayerId int `json:"PlayerId"`
	PlayerLevel int `json:"PlayerLevel"`
	ChipAmountBet int `json:"ChipAmountBet"`
}