package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct{
  Id int
  First_Name string
  Last_Name string
}

func main() {
        person := Person{}
	db, err := sql.Open("mysql", "root:system@tcp(127.0.0.1:3306)/vk_hrm")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

        row := db.QueryRow("select id, first_name, last_name from person where id = ?;", 123)
	err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)
	if err != nil {
	} else {
	  fmt.Println("result for query row",person)
	}

	stmt, err := db.Prepare("select * from person;")
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Person select : %#v",res)
	}

}
