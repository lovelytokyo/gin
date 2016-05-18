package main

import (
    "fmt"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "github.com/coopernurse/gorp"
)

func main() {
    const (
        dbUser = "root"
        dbPassword = "rootpasswd"
        dbHost = "192.168.99.100:3306"
        dbName = "saigon_facebook"
    )

    type Tags struct {
      Id             int32  `db:"id"`
      Name           string `db:"name"`
      TaggingType    string `db:"tagging_type"`
      TagsCategoryId int32  `db:"tags_category_id"`
      CreatedAt      int64  `db:"created_at"`
      UpdatedAt      int64  `db:"updated_at"`
    }

    fmt.Println("ORM test start.")

    db, err := sql.Open("mysql", dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName)
    if err != nil {
        fmt.Printf("db connection error. [%s]\n", err)
        return
    }
    defer db.Close()

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
    defer dbmap.Db.Close()

    sQuery := " SELECT " +
              "     * " +
              " FROM " +
              "     tags " +
              " WHERE " +
              "     tagging_type = :tagging_type AND " +
              "     tags_category_id = :tags_category_id "

    rows, err := dbmap.Select(&Tags{}, sQuery, map[string]interface{}{
      "tagging_type": "AUTO",
      "tags_category_id": 2,
    })
    if err != nil {
        fmt.Printf("db select error. [%s]\n", err)
        return
    }

    for _, row := range rows {
      t := row.(*Tags)
      fmt.Println(t.Id)
      fmt.Println(t.Name)
    }

    fmt.Println("ORM test end.")
}
