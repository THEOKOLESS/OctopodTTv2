package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := NewClient()
	CountriesData, err := GetCountries(c)
	if err != nil {
		fmt.Println("Error while retreiving the country data:", err)
		os.Exit(0)
	}
	db, err := connectToDB()
	if err != nil {
		fmt.Println("Error while connecting to the db:", err)
		os.Exit(0)
	}
	insertOrUpdateCountries(db, CountriesData)

	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Updating the countries data...")
				CountriesData, err := GetCountries(c)
				if err != nil {
					fmt.Println("Error while retreiving the country data:", err)
					os.Exit(0)
				}
				insertOrUpdateCountries(db, CountriesData)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	<-done // wait for SIGINT or SIGTERM
	fmt.Println("Exiting...")
	close(quit) // shutdown goroutine
}
