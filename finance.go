package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB
var configuration Configuration

type Card struct {
	Card_ID int64
	Card_Name string
	Description string
}

func GetCards() []Card {
	cards := []Card{}

	err := db.Select(&cards, "SELECT card_id, card_name, description FROM cards")
	if err != nil {
		log.Fatalln(err)
	}
	return cards
}

func InsertCard(card *Card) {
	_, err :=db.NamedExec(`INSERT INTO cards(card_name, description) VALUES (:card_name, :description)`, *card)
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateCardByName(card *Card) {
	_, err :=db.NamedExec(`UPDATE cards SET description = :description WHERE card_name = :card_name`, *card)
	if err != nil {
		log.Fatalln(err)
	}
}



func getHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	db, err = sqlx.Connect(configuration.DBType, configuration.ConnectionString)
	if err != nil {
		log.Fatalln(err)
	}

	//cardIns := Card {
	//	Card_Name: "dumm6666",
	//	Description: "Dummy card",
	//}
	//
	//db.NamedExec(`INSERT INTO cards(card_name, description) VALUES (:card_name, :description)`, cardIns)

	//card := Card{}
	//
	//rows, err := db.Queryx("SELECT card_id, card_name, description FROM cards")
	//
	//for rows.Next() {
	//	err := rows.StructScan(&card)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	fmt.Fprintf(w, "%v - %v - %v\n", card.Card_ID, card.Card_Name, card.Description)
	//}

	cardIns := Card {
		Card_Name: "**9753",
		Description: "Alfa-Bank",
	}

	InsertCard(&cardIns)

	cards := GetCards()

	for i, crd := range cards {
		fmt.Fprintf(w, "%v: %v - %v - %v\n", i + 1, crd.Card_ID, crd.Card_Name, crd.Description)
	}

	fmt.Fprintf(w, "Hello people")
}

func main() {

	configuration = LoadConfig()


	router := mux.NewRouter()
	router.HandleFunc("/", getHandler).Methods("GET")

	http.Handle("/", router)
	err := http.ListenAndServe(configuration.ListenTo, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
