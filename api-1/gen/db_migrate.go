//go:generate go run db_migrate.go
package main

import "go-starter-api/pkg/db"

func main() {
	// fmt.Println("gen")
	// fmt.Printf("db server: '%v' \n", env.Env().DBServer)
	db.Migrate()
}
