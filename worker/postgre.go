package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectToDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=octopoddb password=wanadev sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	return db, err
}

func insertOrUpdateCountry(db *sql.DB, country CountryData) error {
	sqlStatement := `
    INSERT INTO countries (name_official_fr, flag_url, population)
    VALUES ($1, $2, $3)
	ON CONFLICT (name_official_fr)
	DO UPDATE SET flag_url = EXCLUDED.flag_url, population = EXCLUDED.population
	WHERE countries.flag_url != EXCLUDED.flag_url OR countries.population != EXCLUDED.population
    `

	_, err := db.Exec(sqlStatement, country.Name.Fra.Official, country.Flag, country.Population)
	return err
}

func insertOrUpdateCountries(db *sql.DB, countries []CountryData) {
	for _, country := range countries {
		err := insertOrUpdateCountry(db, country)
		if err != nil {
			fmt.Println("Error while insering to the db:", err)
		}
	}
}
