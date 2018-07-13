package main

import (
	"log"
	"net/http"
	"os"

	rcapi "github.com/clio-paul-cisek/hackaton/lawpay/recurring/api"
	lredis "github.com/clio-paul-cisek/hackaton/repository/redis"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
	repository := lredis.New()

	err = repository.Connect()
	if err != nil {
		log.Fatal("Could not connect to redis.")
	}
	defer repository.Close()
	router := httprouter.New()
	lawPayCtrl := rcapi.New(repository)
	router.GET("/lawpay/:uuid", lawPayCtrl.Get)
	router.POST("/lawpay", lawPayCtrl.Create)
	router.POST("/record/create", lawPayCtrl.RecordCreate)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}
