package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

var thiscounter int

//var allcounter string

var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	thiscounter++
	client := redis.NewClient(&redis.Options{
		Addr:     "db:6379",
		Password: "",
		DB:       0,
	})
	err := client.Incr("allcount").Err()
	if err != nil {
		fmt.Println(err)
	}
	allcounter, err := client.Get("allcount").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "This is request "+strconv.Itoa(thiscounter)+" to this instance and request "+allcounter+" to the cluster")
	mutex.Unlock()
}
func main() {

	http.HandleFunc("/", incrementCounter)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
