package main

import (
	"alisisService/internal/database"
	"context"
	"log"
	"time"
)

func main() {

	ctxBackgrond := context.Background()

	ctxToPing, cancel := context.WithTimeout(ctxBackgrond, 5*time.Second)
	defer cancel()

	conn, err := database.Connecton(ctxBackgrond)

	if err != nil {
		log.Fatalf("Failed to connected to the DataBase")
	} else {
		log.Println("Connection to DataBase succesfull ")

	}

	if err := conn.Ping(ctxToPing); err != nil {
		cancel()
		conn.Close(ctxBackgrond)
		log.Fatalf("Failed to ping to the DataBase for 5 second")
	} else {
		log.Println("ping to DataBase succesfull ")
	}

}
