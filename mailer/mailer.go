package mailer

import (
    "log"
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
    "github.com/noamelya/mobilize/config"
    "os"
)

func SendMail(to_email string) {
    log.Print("Trying to mail ", to_email)

    conf := config.ReadConfig()

    from := mail.NewEmail(conf.From_name, conf.From_email)
    subject := conf.Subject
    to := mail.NewEmail(conf.To_name, to_email)
    content := mail.NewContent("text/plain", conf.Content)
    m := mail.NewV3MailInit(from, subject, to, content)
    apiKey := os.Getenv("SENDGRID_API_KEY")
    request := sendgrid.GetRequest(apiKey, conf.Sg_route, conf.Sg_url)
    request.Method = "POST"
    request.Body = mail.GetRequestBody(m)
    response , err := sendgrid.API(request)
    if err != nil {
        log.Fatal(err)
    } else {
	log.Print(response.StatusCode)
	log.Print(response.Headers)
    }

}
