package main

import (
	"flag"
	userHandler "github.com/gnippets/restServer/handlers/user"
	"github.com/gnippets/restServer/utils"
	"github.com/go-pg/pg"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	confPath := flag.String("conf", `./config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	conf, _ := utils.GetConfiguration(*confPath)

	db := pg.Connect(&pg.Options{
		User:     conf.User,
		Database: conf.Db,
	})

	defer db.Close()

	if err := utils.DbConnectionEstablished(db); err != nil {
		log.Fatal("Cant establish db connection", err)
	}

	router := httprouter.New()

	router.GET("/users", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		userHandler.GetAllUsers(w, r, db)
	})

	router.GET("/users/:userId", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userHandler.GetUserById(w, r, ps.ByName("userId"), db)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
