package db_models

type Entity struct {
  Id  string  `db:"id"`
  Caption string `db:"caption"`
  Schema  string  `db:"schema"`
  Properties  interface{} `db:"properties"`
  Image string `db:"image"`
}
