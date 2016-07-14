package main

import (
        "github.com/noamelya/mobilize/rabbit"
        "github.com/noamelya/mobilize/mailer"
)

func main() {
        rabbit.Listen("email_tasks", mailer.SendMail)
}

