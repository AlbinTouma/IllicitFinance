package db_models

type Entity struct {
  ID  string  `db:"id"`
  Schema  string  `db:"Schema"`
  Properties  interface{} `db:"Properties"`
}
