package main

import (
	"fmt"
	"net/http"
	"github.com/Jeffail/gabs"
	"io/ioutil"
	"strings"
	"errors"
	"github.com/noamelya/mobilize/rabbit"
	"reflect"
)

func parseGhPost(w http.ResponseWriter, request *http.Request) {
	jsonBytes, _ := ioutil.ReadAll(request.Body)
	jsonParsed, _ := gabs.ParseJSON(jsonBytes)

	email := jsonParsed.Path("email").Data()

	var mailList []string
	switch reflect.TypeOf(email).Kind() {
	case reflect.String:
		mailList = strings.Fields(email.(string))
	case reflect.Slice:
		panic(errors.New("Multi mails are not supported"))
	}

	rabbit.Publish("email_tasks", mailList)
	fmt.Fprintf(w, "Email/s had been received and will be published")
}

func main() {
	http.HandleFunc("/", parseGhPost)
	http.ListenAndServe(":8080", nil)
}
