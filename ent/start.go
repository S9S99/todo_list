package main

import (
  "context"
  "log"

  "entdemo/ent"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  client, err := ent.Open("mysql", "root@tcp(localhost:3306)/todo?parseTime=True")
  if err != nil {
    log.Fatalf("failed opening connection to mysql: %v", err)
  }
  defer client.Close()
  // Run the auto migration tool.
  if err := client.Schema.Create(context.Background()); err != nil {
    log.Fatalf("failed creating schema resources: %v", err)
  }
}
