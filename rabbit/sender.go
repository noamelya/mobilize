package rabbit

import (
    "log"

    "github.com/streadway/amqp"
)

func Publish(queue string, body []string) {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    FailOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    FailOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        queue, // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    FailOnError(err, "Failed to declare a queue")


    for _, email := range body {
        err = ch.Publish(
            "",     // exchange
            q.Name, // routing key
            false,  // mandatory
            false,  // immediate
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(email),
            })
        log.Printf(" [x] Sent %s", email)
        FailOnError(err, "Failed to publish a message")
    }
}
