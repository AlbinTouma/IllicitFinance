/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // add field
  collection.fields.addAt(38, new Field({
    "convertURLs": false,
    "hidden": false,
    "id": "editor3485334036",
    "maxSize": 0,
    "name": "note",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "editor"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // remove field
  collection.fields.removeById("editor3485334036")

  return app.save(collection)
})
