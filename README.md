# golang-api-note
An example Golang RESTful API project

## Frameworks
- web framework : Go-fiber
- ORM : Gorm
- Database : MySQL
- Validation : Go playground validator
- config : GoDotEnv

## Architecture
Controller -> Service -> Repository

reference : [https://github.com/khannedy/golang-clean-architecture](https://github.com/khannedy/golang-clean-architecture)

## API Spec

### Create Note

Request :

- Method : POST
- Endpoint : `api/notes`
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

Response :

```json
{
  "code": 0,
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

Response :

```json
{
  "code": 0,
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

- Method : PUT/PATCH
- Endpoint : `api/notes/{noteId}`
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

## Delete Mahasiswa

Request :

- Method : DELETE
- Endpoint : `api/notes/{noteId}`

Response :

```json
{
  "code": "number",
  "status": "string",
  "message": "string"
}
```
