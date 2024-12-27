/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_4285667772")

  // update field
  collection.fields.addAt(9, new Field({
    "hidden": false,
    "id": "autodate2685905599",
    "name": "last_change",
    "onCreate": true,
    "onUpdate": false,
    "presentable": false,
    "system": false,
    "type": "autodate"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_4285667772")

  // update field
  collection.fields.addAt(9, new Field({
    "hidden": false,
    "id": "autodate2685905599",
    "name": "last_updated",
    "onCreate": true,
    "onUpdate": false,
    "presentable": false,
    "system": false,
    "type": "autodate"
  }))

  return app.save(collection)
})
