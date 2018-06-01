package models

import (
	"fmt"
	"database/sql"
	"log"
	"github.com/metalscreame/GoToBoox/src/config"
)

type Category struct{
	Id int64
	Title string
}
func GetCategories() []Category{
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	Cats := []Category{}
	for rows.Next() {
		cat := new(Category)
		if err := rows.Scan(&cat.Id, &cat.Title); err != nil {
			log.Fatal(err)
		}
		Cats = append(Cats, *cat)
		fmt.Printf("id %d title %s\n", Cats[0].Id, Cats[0].Title)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return Cats
}

