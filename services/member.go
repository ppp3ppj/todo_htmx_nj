package services

import (
	"database/sql"
	"log"
	"github.com/ppp3ppj/todo_htmx_nj/dto"
)

type MemberService struct {
    DB *sql.DB
}

func (ms *MemberService) GetAll() []*dto.MemberCardDto {
    members := []*dto.MemberCardDto{}
    rows, err := ms.DB.Query("SELECT * FROM Members")
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

func (ms *MemberService) Create(text string) *dto.MemberCardDto {
    member := &dto.MemberCardDto{
        Id: 0,
        Name: text,
    }
    sql_str := `INSERT INTO Members(Name) values(?)`
    _, err := ms.DB.Exec(sql_str, member.Name)

    if err != nil {
        log.Printf("%q: %s\n", err, sql_str)
        return nil
    }
    return member
}

func (ms *MemberService) Get(id int) *dto.MemberCardDto {
    sql_str, err := ms.DB.Prepare("SELECT Name FROM Members where Id = ?")
    if err != nil {
        log.Fatal(err)
        return nil
    }
    defer sql_str.Close()
    var name string
    err = sql_str.QueryRow(id).Scan(&name)
    if err != nil {
        log.Fatal(err)
        return nil
    }
    return &dto.MemberCardDto {
        Id: id,
        Name: name,
    }
}

func (ms *MemberService) Update(id int, text string) *dto.MemberCardDto {
	member := &dto.MemberCardDto{
		Id:      id,
		Name:    text,
	}
	sql_str := `UPDATE Members SET Name=? WHERE Id = ?`
	_, err := ms.DB.Exec(sql_str, member.Name, member.Id)
	if err != nil {
		log.Printf("%q: %s\n", err, sql_str)
		return nil
	}
	return member
}

func (ms *MemberService) Delete(id int) error {
    sql_str := `DELETE FROM Members WHERE Id = ?`
    _, err := ms.DB.Exec(sql_str, id)
    if err != nil {
        log.Printf("%q: %s\n", err, sql_str)
        return err
    }
    return nil
}
