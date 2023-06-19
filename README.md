# golang-api-note
An example Golang RESTful API project

## Frameworks
- Web framework : Gofiber
- ORM : Gorm
- Database : MySQL
- Validation : Go playground validator
- Config : GoDotEnv
- Testing : Testify

## Architecture
Controller -> Service -> Repository

reference : [https://github.com/khannedy/golang-clean-architecture](https://github.com/khannedy/golang-clean-architecture)

## API Spec

### Login
Request :
- Method : POST
- Endpoint : `/api/auth/login`
- Body : 
```json
{
    "username": "string",
    "password":"string"
}
```

Response : 
```json
{
  "code": "number",
  "status": "string",
  "message": "string",
  "data": {
    "token": "string"
  }
}
```

### Register :
request :
- Method : POST
- Endpoint : `/api/auth/register`
- Body : 
```json
{
    "name": "string",
    "username": "string",
    "password":"string"
}
```

Response : 
```json
{
  "code": "number",
  "status": "string",
  "message": "string",
  "data": {
    "name": "string",
    "username": "string"
  }
}
```

### Create Note

Request :

- Method : POST
- Endpoint : `api/notes`
- Header : 
  - Authorization : token
  - content-type : application/json
- Body :

```json
{
  "title": "Judul catatan",
  "tags": [
    "tags 1",
    "tags 2"
  ],
  "body": "Isi catatan"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "message": "string",
  "data": {
    "id": "string"
  }
}
```

## Get all Notes

Request :

- Method : GET
- Endpoint : `api/notes`
- Header : 
  - Authorization : token
  - content-type : application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "message": "string",
  "data": [
    {
      "id": "string",
      "title": "string",
      "createdAt": "string",
      "updatedAt": "string",
      "tags": [
        "string"
      ],
      "body": "string"
    }
  ]
}
```

## Get Notes by Id

Request :

- Method : GET
- Endpoint : `api/notes/{noteId}`
- Header : 
  - Authorization : token
  - content-type : application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "message": "string",
  "data": {
    "id": "string",
    "title": "string",
    "createdAt": "string",
    "updatedAt": "string",
    "tags": [
      "string"
    ],
    "body": "string"
  }
}
```

## Update Note

Request :

- Method : PUT
- Endpoint : `api/notes/{noteId}`
- Header : 
  - Authorization : token
  - content-type : application/json
- Body :

```json
{
  "title": "Judul catatan",
  "tags": [
    "tags 1",
    "tags 2"
  ],
  "body": "Isi catatan"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "message": "string",
  "data": {
    "title": "Judul catatan",
    "tags": [
      "tags 1",
      "tags 2"
    ],
    "body": "Isi catatan"
  }
}
```

## Delete Note

Request :

- Method : DELETE
- Endpoint : `api/notes/{noteId}`
- Header : 
  - Authorization : token
  - content-type : application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "message": "string"
}
```
