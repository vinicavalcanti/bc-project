package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		ssmsvc := ssm.New(session.Must(session.NewSession(&aws.Config{})))
		keyname := os.Getenv("AWS_PARAMETER_STORE_KEY")

		param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
			Name:           &keyname,
			WithDecryption: aws.Bool(true),
		})

		if err != nil {
			log.Fatal(err)
		}

		value := aws.String(string(*param.Parameter.Value))

		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, string(*value))

	})
	http.ListenAndServe(":8080", nil)
}
