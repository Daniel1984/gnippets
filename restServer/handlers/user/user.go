package handlers

import (
	"encoding/json"
	"github.com/go-pg/pg"
	"net/http"
	"strconv"
)

type User struct {
	Id   int
	Role string
	Name string
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	w.Header().Set("Content-Type", "application/json")

	var users []User
	if err := db.Model(&users).Select(); err != nil {
		http.Error(w, http.StatusText(403), http.StatusForbidden)
	} else {
		json, err := json.Marshal(users)
		if err != nil {
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}

		w.Write(json)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request, userId string, db *pg.DB) {
	w.Header().Set("Content-Type", "application/json")

	if intUid, err := strconv.Atoi(userId); err != nil {
		http.Error(w, http.StatusText(403), http.StatusForbidden)
	} else {
		user := &User{Id: intUid}
		if err := db.Select(user); err != nil {
			http.Error(w, http.StatusText(403), http.StatusForbidden)
		} else {
			json, err := json.Marshal(user)
			if err != nil {
				http.Error(w, http.StatusText(403), http.StatusForbidden)
				return
			}

			w.Write(json)
		}
	}
}
