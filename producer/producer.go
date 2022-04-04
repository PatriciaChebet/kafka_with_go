package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	//Contacting the leader
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic-test", 0)
	//set a write deadline
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	//write something
	conn.WriteMessages(kafka.Message{Value: []byte("Message from me")})

}
