package handlers

import (
	"encoding/json"
	"github.com/go-pg/pg"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	Name string `json:"name"`
}

func forbidden(w http.ResponseWriter) {
	http.Error(w, http.StatusText(403), http.StatusForbidden)
}

func asJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	asJson(w)
	var users []User
	if err := db.Model(&users).Select(); err != nil {
		forbidden(w)
	} else {
		json, err := json.Marshal(users)
		if err != nil {
			forbidden(w)
			return
		}

		w.Write(json)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request, userId string, db *pg.DB) {
	asJson(w)
	if intUid, err := strconv.Atoi(userId); err != nil {
		forbidden(w)
	} else {
		user := &User{Id: intUid}
		if err := db.Select(user); err != nil {
			forbidden(w)
		} else {
			if json, err := json.Marshal(user); err != nil {
				forbidden(w)
			} else {
				w.Write(json)
			}
		}
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	if resBody, err := ioutil.ReadAll(r.Body); err != nil {
		forbidden(w)
	} else {
		user := &User{}
		json.Unmarshal(resBody, user)

		if err := db.Insert(user); err != nil {
			forbidden(w)
		} else {
			w.Write([]byte("OK!"))
		}
	}
}

func PatchUser(w http.ResponseWriter, r *http.Request, userId string, db *pg.DB) {
	asJson(w)
	if resBody, err := ioutil.ReadAll(r.Body); err != nil {
		forbidden(w)
	} else {
		if intUid, err := strconv.Atoi(userId); err != nil {
			forbidden(w)
		} else {
			user := &User{Id: intUid}
			if err := db.Select(user); err != nil {
				forbidden(w)
			} else {
				json.Unmarshal(resBody, user)
				if err := db.Update(user); err != nil {
					forbidden(w)
				}
			}
		}
	}
}
