package main

import (
  "database/sql"
  "fmt"
  "bufio"
  "os"
  "net/http?
  html/template"
  "regexp"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "hireme123"
  dbname   = "postgres"
) //postgresql server info

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!") //once connected to server
  
  fmt.Println("What is your street address?")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  address := scanner.Text()
  fmt.Println("What is your city?")
  scanner2 := bufio.NewScanner(os.Stdin)
  scanner2.Scan()
  city := scanner2.Text()
  fmt.Println("What state do you live in?")
  scanner3 := bufio.NewScanner(os.Stdin)
  scanner3.Scan()
  state := scanner3.Text()
  fmt.Println("What is your Zip code?")
  scanner4 := bufio.NewScanner(os.Stdin)
  scanner4.Scan()
  zip := scanner4.Text()
  
  sqlStatement := `
  INSERT INTO addresses (Address, City, State, Zip)
  VALUES ($1, $2, $3, $4)`
  id := 0
  err = db.QueryRow(sqlStatement, address, city, state, zip).Scan(&id)
   //post to postgresql server
  
  
  
}
