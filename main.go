package main


import (
	
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	 "github.com/siddhiparekh11/GoChallenge/repository"
	"github.com/siddhiparekh11/GoChallenge/controllers"
	"github.com/siddhiparekh11/GoChallenge/delivery"
	"github.com/siddhiparekh11/GoChallenge/models"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"net/http"
	"log"
	

)

type App struct {

	Router *mux.Router
	Conn *sql.DB

}

var config models.Config

func init() {

	viper.SetConfigFile(`Config.Quest.json`)
	err := viper.ReadInConfig()
	err = viper.Unmarshal(&config)
	log.Println(config)

	if err!=nil {

	}
}




func main() {

	app := App {
		Router : mux.NewRouter(),
	}

	conn , err := dbConnect()

	if err!=nil {
			fmt.Println(err)
	}

	app.Conn = conn

	gRepo := repository.NewGameRepository(app.Conn)
	gContr := controller.NewGameController(gRepo,config)
	delivery.NewGameHandler(app.Router,app.Conn,gContr)

	qRepo := repository.NewQuestRepository(app.Conn)
	qContr := controller.NewQuestController(qRepo)
	delivery.NewQuestHandler(app.Router,app.Conn,qContr)

	pRepo := repository.NewPlayerRepository(app.Conn)
	pContr := controller.NewPlayerController(pRepo)
	delivery.NewPlayerHandler(app.Router,app.Conn,pContr)

	prRepo := repository.NewProgressRepository(app.Conn)
	prContr := controller.NewProgressController(prRepo,qRepo,pRepo,gRepo,config)
	delivery.NewProgressHandler(app.Router,app.Conn,prContr)

	log.Fatal(http.ListenAndServe(":8000",app.Router))

}


func dbConnect() (*sql.DB,error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", viper.GetString(`database.user`), viper.GetString(`database.password`),viper.GetString(`database.host`),viper.GetString(`database.port`),viper.GetString(`database.name`))
	log.Println(connectionString)
	db,err := sql.Open("mysql",connectionString)
	if err!=nil {
		return nil,err
	}

	return db,nil

}




