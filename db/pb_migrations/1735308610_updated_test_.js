/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_4285667772")

  // add field
  collection.fields.addAt(7, new Field({
    "hidden": false,
    "id": "bool1181691900",
    "name": "target",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "bool"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_4285667772")

  // remove field
  collection.fields.removeById("bool1181691900")

  return app.save(collection)
})
