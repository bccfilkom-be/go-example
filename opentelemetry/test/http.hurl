GET http://localhost:8080/api/v1/pets?page=1
HTTP 200
[Asserts]
jsonpath "$" != null

GET http://localhost:8080/api/v1/pets/1
HTTP 200
[Asserts]
jsonpath "$" != null

GET http://localhost:8080/api/v1/pets/1001
HTTP 404

POST http://localhost:8080/api/v1/pets
{
	"name": "string",
	"photoURL": "url"
}
HTTP 201

PATCH http://localhost:8080/api/v1/pets/2
{
	"name": "string",
	"photoURL": "url"
}
HTTP 200

DELETE http://localhost:8080/api/v1/pets/1001
HTTP 200
