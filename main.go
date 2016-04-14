package main

import (
	"os"
	"fmt"
	"runtime"
	"time"
	"encoding/json"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sqs"
	"github.com/Alexander-Attar/golang-aws-sqs-example/worker"
)

	var accessKey string = os.Getenv("AWS_ACCESS_KEY_ID")
	var secretKey string = os.Getenv("AWS_SECRET_ACCESS_KEY")
	var queueName string = "https://sqs.us-east-1.amazonaws.com/168528444054/queue"

func Print(msg *sqs.Message) error {

	// Custom logic
	var data map[string]interface{}
    if err := json.Unmarshal([]byte(msg.Body), &data); err != nil {
        panic(err)
    }

	fmt.Println(fmt.Sprintf("[%s] Message ID : %v - %s", time.Now().Local(), msg.MessageId, msg.Body ))

	return nil
}

func main() {
	sleepTime := time.Millisecond * 1000 // 0.2 second
	receiveMessageNum := 10
	fmt.Println("===========================================")
	fmt.Println(fmt.Sprintf(" Use CPU(s) num      : %d", runtime.NumCPU()))
	fmt.Println(fmt.Sprintf(" Receive message num : %d", receiveMessageNum))
	fmt.Println(fmt.Sprintf(" Sleep time          : 1 second(s)"))
	fmt.Println("===========================================")

	runtime.GOMAXPROCS(runtime.NumCPU())
	auth := aws.Auth{AccessKey: accessKey, SecretKey: secretKey}
	mySqs := sqs.New(auth, aws.USEast)
	queue := &sqs.Queue{mySqs, queueName}
	worker.Start(queue, worker.HandlerFunc(Print), sleepTime, receiveMessageNum)
}
