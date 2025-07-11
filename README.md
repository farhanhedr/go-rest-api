

- Create new items:
curl -X POST http://localhost:8088/items -H "Content-Type: application/json" -d '{"name":"Apple G", "description":"This is a green apple"}'

- Get all items
curl -X GET http://localhost:8088/items

- Get item by ID
curl -X GET http://localhost:8088/items/1

- Update item by ID
curl -X POST http://localhost:8088/items/1 -H "Content-Type: application/json" -d '{"name":"Apple B", "description":"Now this is a black apple"}'

- Delete item by ID
curl -X DELETE http://localhost:8088/items/1
