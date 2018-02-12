package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/joho/godotenv"
	"ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type PostBack struct {
	Method   string `json:"method"`
	Url      string `json:"url"`
	Mascot   string `json:"mascot"`
	Location string `json:"location"`
}

type Datum struct {
	Mascot   string `json:"mascot"`
	Location string `json:"location"`
}

type LogData struct {
	StartTime   time.Time `json:"starttime"`
	StatusCode  int       `json:"statuscode"`
	EndTime     time.Time `json:"endtime"`
	TimeElapsed int       `json:"timeelapsed"`
	Body        string    `json:"body"`
}

func logger(l LogData) {
	d := []byte{l.StartTime, l.StatusCode, l.EndTime, l.Body}
	err := ioutil.WriteFile("./deliveryAgent.go", d, 0644)
	if err != nil {
		panic(e)
	}
}

func performPB(pb PostBack) {

	client := &http.Client{}

	var l LogData

	method := ToUpper(pb.Method)
	str := pb.Url
	replacer := strings.NewReplacer("{mascot}", pb.Mascot, "{location}", pb.Location)
	url := replacer.Replace(str)

	l.startTime = time.Now()

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	l.EndTimendTime = time.Now()
	l.TimeElapsed = endTime.Sub(startTime)

	defer resp.Body.Close()

	l.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	l.statusCode, err = ioutil.ReadAll(resp.StatusCode)
	if err != nil {
		panic(err)
	}

	logger(l)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Connect
	conn, err := redis.Dial("tcp", "159.89.155.145:"+os.Getenv("REDISPORT"))
	if err != nil {
		panic(err)
	}

	response, err := conn.Do("AUTH", os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	_, err = conn.Do("PING")
	if err != nil {
		log.Fatal("Can't connect to the Redis database")
	}

	pb, err = redis.String(conn.Do("LPOP", "postback"))
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(2000 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			performPB(pb)
		}
	}()

}
