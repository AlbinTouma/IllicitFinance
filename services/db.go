package db

import (
	"database/sql"
	//"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
  "kys/models"
)

 
//http://go-database-sql.org/retrieving.html
func QueryEntity() ([]db_models.Person, error) {
  db, err := sql.Open("sqlite3", "./db/pb_data/data.db")
  if err != nil {
    log.Fatal(err)
  }
  rows, err := db.Query("SELECT * FROM Person;")
  if err != nil {
    log.Printf("Err %s", err)
    return nil, err
  }
  //fmt.Printf("%T\n", rows)
  var results []db_models.Person

  for rows.Next(){
    queryResult := db_models.Person{}
    rows.Scan(
      &queryResult.Id,
      &queryResult.Name, 
      &queryResult.FirstName,
      &queryResult.MiddleName,
      &queryResult.LastName,
      &queryResult.NameSuffix,
      &queryResult.FatherName,
      &queryResult.MotherName,
      &queryResult.BirthDate, 
      &queryResult.DeathDate, 
      &queryResult.Gender, 
      &queryResult.Ethnicity, 
      &queryResult.Height, 
      &queryResult.Weight, 
      &queryResult.EyeColor, 
      &queryResult.HairColor, 
      &queryResult.Appearance, 
      &queryResult.Religion,
      &queryResult.SpokenLanguage, 
      &queryResult.Education,
      &queryResult.Address,
      &queryResult.Nationality, 
      &queryResult.Citizenship, 
      &queryResult.BirthPlace, 
      &queryResult.BirthCountry, 
      &queryResult.Position,
      &queryResult.Email, 
      &queryResult.Phone, 
      &queryResult.PassportNumber, 
      &queryResult.SocialSecurityNumber, 
      &queryResult.SourceName,
      &queryResult.SourceUrl,
      &queryResult.Note,
      &queryResult.Image,
      &queryResult.Created,
      &queryResult.LastUpdated,

      )
    results = append(results, queryResult)
  }
  return results, nil

}
