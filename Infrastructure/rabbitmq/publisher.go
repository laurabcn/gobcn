package rabbitmq

import (
"fmt"
"github.com/streadway/amqp"
"log"
"os"
)

type Publisher struct {
}

func (p Publisher) Publish(message string) {

	queueName := os.Getenv("RABBIT_QUEUE_NAME")
	conn, errConnection := amqp.Dial(createRabbitUrl())
	failOnError(errConnection, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.
		Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	log.Printf(" [x] Sent %s", message)
	failOnError(err, "Failed to publish a message")
}

func createRabbitUrl() string {

	user := os.Getenv("RABBIT_USER")
	password := os.Getenv("RABBIT_PWD")
	host := os.Getenv("RABBIT_HOST")
	port := os.Getenv("RABBIT_PORT")
	vhost := os.Getenv("RABBIT_VHOST")

	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		vhost,
	)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

