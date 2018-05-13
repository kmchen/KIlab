package main

//curl -d '{"Id":0, "Name":"kmchen"}' -H "Content-Type: application/json" -X POST http://localhost:9090/interviewer
import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"KILab/backend/model"
	"KILab/backend/operation"
	"KILab/backend/store"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var amaqAddr = flag.String("amaqAddr", "amqp://guest:guest@localhost:5672/", "amaq address")
var wsPort = flag.String("wsPort", ":3000", "wesocket port")
var redisPort = flag.String("redisPort", ":6379", "redis port")
var riderChanSize = flag.Int("riderChanSize", 10, "rider channel size")
var serverAddr = flag.String("serverAddr", "localhost:9090", "server address")

func getInterviewer(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys) < 1 {
		log.Println("Url Param id is missing")
		return
	}
	id := keys[0]
	log.Println("Url Param 'key' is: " + string(id))

	w.Write([]byte("pong"))
}

func createInterviewer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var interviewer model.Interviewer
	var id string
	err := decoder.Decode(&interviewer)
	id, err = ops.CreateInterviewer(interviewer)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(id))
}

var redisStore, _ = store.NewRedisStore(func() (redis.Conn, error) { return redis.Dial("tcp", *redisPort) })
var ops = operation.NewOperation(redisStore)

func main() {

	flag.Parse()

	//var err error

	router := mux.NewRouter()

	router.HandleFunc("/interviewer/{id}", getInterviewer).Methods("GET")
	router.HandleFunc("/interviewer", createInterviewer).Methods("POST")

	// Enable metrics monitoring
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(*serverAddr, router))
}
