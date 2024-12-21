package db

import (
	//	"database/sql"
	//"fmt"
	"encoding/json"
	"kys/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//http://go-database-sql.org/retrieving.html
func QueryEntity() ([]db_models.Person, error) {
  db, err := sqlx.Open("sqlite3", "./db/pb_data/data.db")
  if err != nil {
    log.Printf("Error in sqlx %s", err)
  }
  defer db.Close()
  
  var entities []db_models.Entity
  err = db.Select(&entities, "SELECT id, Schema, Properties, Image FROM test;")
  if err != nil {
    log.Printf("Error fetching entities from db %s", err)
    return nil, err
  }

  log.Printf("Ent %s", entities)
  
  var persons []db_models.Person

  for _, entity := range entities {
    var person db_models.Person
    
    err := json.Unmarshal([]byte(entity.Properties.(string)), &person)
    if err != nil {
      log.Printf("Error unmarshalling json", entity.Id, err)
      continue
    }
    log.Printf("Person marshal %s", person)
    persons = append(persons, person)
  }


  log.Printf("Persons %s", persons)
  return persons, nil
}
