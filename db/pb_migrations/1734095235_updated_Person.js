/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // remove field
  collection.fields.removeById("text1329388006")

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // add field
  collection.fields.addAt(8, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text1329388006",
    "max": 0,
    "min": 0,
    "name": "birth_place",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  return app.save(collection)
})
