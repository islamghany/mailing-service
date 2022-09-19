package main

import (
	"fmt"
	"log"
	"mailer-service/api"
	"mailer-service/mailer"
	"os"
	"strconv"
)

const webPort = 8080

func main() {

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	m := mailer.Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("ENCRYPTION"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
		FromName:    os.Getenv("FORM_NAME"),
	}

	//---	DEVELOPMENT
	// m := mailer.Mail{
	// 	Domain:      "localhost",
	// 	Host:        "localhost",
	// 	Port:        1025,
	// 	Username:    os.Getenv("MAIL_USERNAME"),
	// 	Password:    os.Getenv("MAIL_PASSWORD"),
	// 	Encryption:  "none",
	// 	FromAddress: "john.smith@example.com",
	// 	FromName:    "john Smith",
	// }
	server, err := api.NewServer(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Running the mail service")
	server.Start(webPort)
	if err != nil {
		log.Fatal(err)
	}
}
