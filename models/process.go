package models

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rapando/hash-engine/utils"
	"github.com/streadway/amqp"
)

var (
	q     *amqp.Connection
	qChan *amqp.Channel
)

func Process() {
	start := time.Now()
	defer func() {
		fmt.Printf("\nDONE : %s\n", time.Since(start))
	}()
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Printf("unable to load dotenv because %v", err)
		return
	}

	q, err = utils.QConnect(os.Getenv("Q_URI"))
	if err != nil {
		log.Printf("unable to connect to rabbitmq. exit")
		os.Exit(3)
	}
	qChan, err = q.Channel()
	if err != nil {
		log.Printf("unable to create a rabbitmq channel because %v", err)
		os.Exit(3)
	}

	var chars = "`1234567890-=\\][poiuytrewqasdfghjkl;'/.," +
		"mnbvcxz~!@#$%^&*()_+|}{POIUYTREWQASDFGHJKL:\"?><MNBVCXZ "
	chars = "randomcharacters"
	var n = int64(len(chars))

	combinations := GetNoOfCombinations(n)
	log.Printf("a string with %d characters has %d combinations", n, combinations.Int64())

	log.Println("getting combinations and pusblishing them to redis")

	GetCombinations(chars, combinations, qChan)
}
