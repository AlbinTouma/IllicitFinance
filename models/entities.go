package db_models

type Entity struct {
  Id  string  `db:"id"`
  Schema  string  `db:"Schema"`
  Properties  interface{} `db:"Properties"`
  Image string `db:"Image"`
}
