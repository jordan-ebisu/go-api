# Go-API

This is a golang api application that connects to MongoDB on the backend.

## Running

Run the application:

```bash
go run main.go
```

Open up Postman

Creating a Person:
create a new window
click on POST
Type in http://localhost:12345/person
Body Tab
raw
JSON (in dropdown)

{
  "firstname": "test",
  "lastname": "test"
}

