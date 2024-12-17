/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // remove field
  collection.fields.removeById("json1329388006")

  // add field
  collection.fields.addAt(8, new Field({
    "hidden": false,
    "id": "text1329388006",
    "maxSize": 0,
    "name": "birth_place",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "json"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // add field
  collection.fields.addAt(39, new Field({
    "hidden": false,
    "id": "json1329388006",
    "maxSize": 0,
    "name": "birth_place",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "json"
  }))

  // remove field
  collection.fields.removeById("text1329388006")

  return app.save(collection)
})
