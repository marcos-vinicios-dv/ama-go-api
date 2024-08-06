package main

import (
	// "database/sql"
	// "fmt"
	// "log"

	"os/exec"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// db, err := sql.Open("postgres", "user=postgres dbname=ws sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// rows, err := db.Query("SELECT * from")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(rows)

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
