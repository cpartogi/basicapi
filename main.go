package main

import (
	"basicapi/db"
	"basicapi/handler"
	"basicapi/models"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("The is Server Running on localhost port 3000")

	db.Namedb["1"] = models.NameList{ID: 1, Name: "A"}
	db.Namedb["2"] = models.NameList{ID: 2, Name: "B"}
	db.Namedb["3"] = models.NameList{ID: 3, Name: "C"}

	http.HandleFunc("/", handler.GetData)

	err := http.ListenAndServe(":3000", nil)

	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
