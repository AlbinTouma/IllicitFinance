/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_4285667772")

  // add field
  collection.fields.addAt(6, new Field({
    "hidden": false,
    "id": "json1609857236",
    "maxSize": 0,
    "name": "referents",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "json"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_4285667772")

  // remove field
  collection.fields.removeById("json1609857236")

  return app.save(collection)
})