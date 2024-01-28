package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"github.com/ppp3ppj/todo_htmx_nj/services"
	"github.com/ppp3ppj/todo_htmx_nj/templates"
	"github.com/ppp3ppj/todo_htmx_nj/templates/components"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    fmt.Println("Welcome to Golang HTMX Todo app.")   
	db := initDB("./database.sqlite3")
	migrate(db)
	// close database connection before exiting program.
    defer db.Close()
    e := echo.New()

    memberService := &services.MemberService {
        DB: db,
    }


    e.Use(middleware.CORS())
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
      Format: "method=${method}, uri=${uri}, status=${status}\n",
    }))
    e.Use(middleware.Recover())
    e.GET("/", func(c echo.Context) error {
        members := memberService.GetAll()
        component := templates.Index(members)
        return component.Render(context.Background(), c.Response().Writer)
    })
    
    e.GET("/add_member_modal", func(c echo.Context) error {
        component := components.AddMemberModal(0, "", false)
        return component.Render(context.Background(), c.Response().Writer)
    }) 

    e.GET("/members/:id", func(c echo.Context) error {
        id := c.Param("id")
        if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
        }
        n, err := strconv.Atoi(id)
        if err != nil {
            fmt.Println(err)
        }
        info_member := memberService.Get(n)

        component := components.AddMemberModal(info_member.Id, info_member.Name, true)
        return component.Render(context.Background(), c.Response().Writer)
    }) 

    e.POST("/members", func(c echo.Context) error {
        data := c.FormValue("add-name-nj")
        if data == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Data")
        }
        memberService.Create(data)
        members := memberService.GetAll()
        componet := components.MemberCards(members)
        return componet.Render(context.Background(), c.Response().Writer) 
    })

    e.PUT("/members/:id", func(c echo.Context) error {
        id := c.Param("id")
        if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
        }
        n, err := strconv.Atoi(id)
        if err != nil {
            fmt.Println(err)
        }

        info_member := memberService.Get(n)
        
        text := c.FormValue("edit-name-nj")
        if text == "" {
            text = info_member.Name
        }

        memberService.Update(n, text)
        members := memberService.GetAll()
        component := components.MemberCards(members)
        return component.Render(context.Background(), c.Response().Writer)
    })

    e.DELETE("/members/:id", func(c echo.Context) error {
        id := c.Param("id")
        if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
        }
        n, err := strconv.Atoi(id)
        if err != nil {
            fmt.Println(err)
        }

        memberService.Delete(n)
        members := memberService.GetAll()
        component := components.MemberCards(members)
        return component.Render(context.Background(), c.Response().Writer)
    })
    e.GET("/add_member_10_rec", func(c echo.Context) error {
        for i := 0; i < 10; i++ {
            memberService.Create(generateRandomString(5))
        }
        members := memberService.GetAll()
        componet := components.MemberCards(members)
        return componet.Render(context.Background(), c.Response().Writer) 
    })

    e.Static("/static", "static")
    e.Logger.Fatal(e.Start(":1323"))
}

func generateRandomString(length int) string {
   b := make([]byte, length)
   _, err := rand.Read(b)
   if err != nil {
      panic(err)
   }
   return base64.StdEncoding.EncodeToString(b)
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
	CREATE TABLE IF NOT EXISTS Members(
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
