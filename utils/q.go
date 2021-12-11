package utils

import (
	"github.com/streadway/amqp"
)

func QConnect(qURI string) (conn *amqp.Connection, err error) {

	conn, err = amqp.Dial(qURI)
	if err != nil {
		return
	}
	return
}

func QPublish(channel *amqp.Channel, exchange, data string) (err error) {
	
	if err = channel.Publish(exchange, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(data),
	}); err != nil {
		return
	}
	return nil
}
