package main

import (
    "database/sql"
    "fmt"
    "time"
    _ "github.com/mattn/go-sqlite3"
)

const (
    DatabaseType = "sqlite3"
    PathToDatabase = "./test.db"
)

func databaseConnection(databaseType string, pathToDB string) (db *sql.DB){
    db, err := sql.Open(databaseType, pathToDB)
    checkErr(err)

    // Turning on Forgein key support
    stmt, err := db.Prepare("PRAGMA foreign_keys = ON;")
    checkErr(err)
    _, err = stmt.Exec()
    checkErr(err)

    return db
}

func main(){
    db := databaseConnection(DatabaseType, PathToDatabase)

    // insert
    stmt, err := db.Prepare("INSERT INTO person(name, country) values(?, ?)")
    checkErr(err)

    res, err := stmt.Exec("astaxie", 1)
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Printf("id: %d\n", id)

    //update
    stmt, err = db.Prepare("update person set name=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("astaxieupdate", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Print("affected: ")
    fmt.Println(affect)

    //query
    rows, err := db.Query("SELECT * FROM person")
    checkErr(err)

    var uid int
    var name string
    var country int
    var created time.Time

    for rows.Next() {
        err = rows.Scan(&uid, &name, &created, &country)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(name)
        fmt.Println(country)
        fmt.Println(created)
    }

    rows.Close() // good habit to close

    // delete
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}