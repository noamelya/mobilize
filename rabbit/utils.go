package rabbit

import (
	"log"
	_ "github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}
