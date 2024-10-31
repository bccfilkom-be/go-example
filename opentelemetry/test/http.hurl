GET http://localhost:8080/api/v1/pets?page=1
HTTP 200
[Asserts]
jsonpath "$" != null
