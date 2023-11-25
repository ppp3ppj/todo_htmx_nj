package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/mattn/go-sqlite3"
    "net/http"
    "github.com/labstack/echo/v4"
)

func main() {
    fmt.Println("Welcome to Golang HTMX Todo app.")   
	db := initDB("./database.sqlite3")
	migrate(db)
	// close database connection before exiting program.
    defer db.Close()
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}

func initDB(filepath string) *sql.DB {
    if _, err := os.Stat(filepath); err != nil {
        file, err := os.Create(filepath)
        checkError(err)
        file.Close()
    }
	db, err := sql.Open("sqlite3", filepath)
    checkError(err)

	if db == nil { // if db is nil
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);`

	_, err := db.Exec(sql)
    checkError(err)

}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
