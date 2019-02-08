package models

type Config struct {

	Ratefrombet float64 `mapstructure:"RateFromBet"`
	Levelbonusrate float64	`mapstructure:"LevelBonusRate"`
	Totalquestpoints float64 `mapstructure:"TotalQuestPoints"`
	GameId int `mapstructure:"GameId"`
	Noofmilestones int `mapstructure:"NoOfMilestones"`
	Milestones []Milestone `mapstructure:"Milestones"`
	Db Database	`mapstructure:"database"`
}