package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	_ "user-center/docs"

	serviceCenter "github.com/cyjme/service-center/service"
	"github.com/cyjme/service-center/util"
	"github.com/gin-gonic/gin"
	"github.com/cyjme/service-center/client"
	"github.com/cyjme/service-center/cmd/httpServer/router"
)

func main() {
	client.Init("localhost:2379")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.RegisterV1Router(r)

	bindPortAndRun(r)
}

func bindPortAndRun(r *gin.Engine) {
	port := getRandomPort()

	intranetIp, ipErr := util.GetIntranetIp()
	if ipErr != nil {
		log.Panic(ipErr)
	}

	//check server is already up
	go func() {
		for {
			time.Sleep(time.Second)
			log.Println("Checking server started...")
			resp, err := http.Get("http://localhost:" + port + "/ping")
			if err != nil {
				log.Println("Checking failed:", err)
				continue
			}
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				log.Println("Check server not ok", resp.StatusCode)
				continue
			}
			break
		}
		log.Println("SERVER UP AND RUNNING!")
		go serviceCenter.Register("service-center", intranetIp+":"+port)
	}()

	err := r.Run("0.0.0.0" + ":" + port) // listen and serve on 0.0.0.0:8080
	if strings.Contains(err.Error(), "address already in use") {
		bindPortAndRun(r)
		return
	}
}

func getRandomPort() string {
	min := 10000
	max := 11000

	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Intn(max-min) + min)
}
