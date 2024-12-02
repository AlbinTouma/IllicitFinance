/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_73640206")

  // add field
  collection.fields.addAt(27, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text2637366347",
    "max": 0,
    "min": 0,
    "name": "address_full",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(28, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text3417790055",
    "max": 0,
    "min": 0,
    "name": "address_street",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(29, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text3385392544",
    "max": 0,
    "min": 0,
    "name": "address_street2",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(30, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text1343738591",
    "max": 0,
    "min": 0,
    "name": "address_city",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(31, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text4250386621",
    "max": 0,
    "min": 0,
    "name": "address_postal_code",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(32, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text876311753",
    "max": 0,
    "min": 0,
    "name": "address_region",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(33, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text2486609371",
    "max": 0,
    "min": 0,
    "name": "address_state",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(34, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text144064254",
    "max": 0,
    "min": 0,
    "name": "address_country",
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
    "id": "text672117379",
    "max": 0,
    "min": 0,
    "name": "address_post_office_box",
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

  // remove field
  collection.fields.removeById("text2637366347")

  // remove field
  collection.fields.removeById("text3417790055")

  // remove field
  collection.fields.removeById("text3385392544")

  // remove field
  collection.fields.removeById("text1343738591")

  // remove field
  collection.fields.removeById("text4250386621")

  // remove field
  collection.fields.removeById("text876311753")

  // remove field
  collection.fields.removeById("text2486609371")

  // remove field
  collection.fields.removeById("text144064254")

  // remove field
  collection.fields.removeById("text672117379")

  return app.save(collection)
})
