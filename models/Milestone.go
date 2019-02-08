package models

type Milestone struct{
	Ind int `mapstructure:"Index"`
	PointsNeededToWin int `mapstructure:"PointsNeededToWin"`
	ChipsAwarded int `mapstructure:"ChipsAwarded"`
}


type CusMilestone struct {
	Ind int `json:"MilestoneIndex"`
	ChipsAwarded int `json:"ChipsAwarded"`
}