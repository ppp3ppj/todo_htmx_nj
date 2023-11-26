package services

import (
	"database/sql"
	"log"
	"todo_htmx_nj/dto"
)

type MemberService struct {
    DB *sql.DB
}

func (ts *MemberService) GetAll() []*dto.MemberCardDto {
    members := []*dto.MemberCardDto{}
    rows, err := ts.DB.Query("SELECT * FROM Members")
    if err != nil {
        log.Fatal(err)
        return members
    }

    defer rows.Close()
    for rows.Next() {
        var id int
        var name string
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        members = append(members, &dto.MemberCardDto{
            Id: id,
            Name: name,
        })
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
        return members
    }
    return members
}
