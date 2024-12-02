/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // add field
  collection.fields.addAt(39, new Field({
    "hidden": false,
    "id": "file854170509",
    "maxSelect": 1,
    "maxSize": 0,
    "mimeTypes": [],
    "name": "profile_image",
    "presentable": false,
    "protected": false,
    "required": false,
    "system": false,
    "thumbs": [],
    "type": "file"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // remove field
  collection.fields.removeById("file854170509")

  return app.save(collection)
})
