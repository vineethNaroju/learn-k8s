package demos

import (
	"context"
	"fmt"
	"learnk8s/redis"
	"net/http"

	"github.com/gorilla/mux"
)

var redisAddr = "localhost:6379"

var ctx = context.Background()

func DemoRedis() {
	redisClient := redis.StaticInstance(redisAddr).Client()

	fmt.Println(redisClient)

	router := mux.NewRouter()

	router.HandleFunc("/set/{key}/{val}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		val := vars["val"]

		if err := redisClient.Set(ctx, key, val, 0); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.String()))
		} else {
			rw.WriteHeader(http.StatusAccepted)
		}
	})

	router.HandleFunc("/get/{key}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		if val, err := redisClient.Get(ctx, key).Result(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		} else {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(val))
		}
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
