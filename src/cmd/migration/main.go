package main

import "pismo-transactions/src/database"

func main() {
	err := database.ExecuteMigration()
	if err != nil {
		panic(err)
	}
}
