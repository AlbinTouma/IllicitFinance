/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // remove field
  collection.fields.removeById("date3845444698")

  // remove field
  collection.fields.removeById("date4230635969")

  // add field
  collection.fields.addAt(34, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text3845444698",
    "max": 0,
    "min": 0,
    "name": "birth_date",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(35, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text4230635969",
    "max": 0,
    "min": 0,
    "name": "death_date",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(36, new Field({
    "exceptDomains": null,
    "hidden": false,
    "id": "url2776776943",
    "name": "source_url",
    "onlyDomains": null,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "url"
  }))

  // add field
  collection.fields.addAt(37, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text1604975365",
    "max": 0,
    "min": 0,
    "name": "source_name",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // add field
  collection.fields.addAt(8, new Field({
    "hidden": false,
    "id": "date3845444698",
    "max": "",
    "min": "",
    "name": "birth_date",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "date"
  }))

  // add field
  collection.fields.addAt(9, new Field({
    "hidden": false,
    "id": "date4230635969",
    "max": "",
    "min": "",
    "name": "death_date",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "date"
  }))

  // remove field
  collection.fields.removeById("text3845444698")

  // remove field
  collection.fields.removeById("text4230635969")

  // remove field
  collection.fields.removeById("url2776776943")

  // remove field
  collection.fields.removeById("text1604975365")

  return app.save(collection)
})
