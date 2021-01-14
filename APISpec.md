## API Spec
### BASE URL
```
http://localhost:8080/api
```

### Response Format:
```json
{
  "status": "number",
  "message": "string",
  "data": "dynamic"
}
```
### Add User
- Headers:
  - Content-Type: "application/json"
- Method: POST
- Endpoint: /user
- Request:
```json
{
  "user_name": "string, required, no space allowed",
  "password": "string, required",
  "birth_place": "string, required",
  "birth_date": "string, required, format_string:'YYYY-MM-DD'"
}
```
- Example Request
```json
{
  "user_name": "naufaldy.mahas",
  "password": "qwerty123",
  "birth_place": "Jakarta",
  "birth_date": "1996-06-22"
}
```
- Example Response:
```json
{
  "status": 200,
  "message": "",
  "data": {
    "user_id": 1,
    "user_name": "naufaldy.mahas",
    "employee_id": "210001",
    "birth_place": "Jakarta",
    "birh_date": "1996-06-22T00:00:00Z",
    "created_at": "1996-06-22T22:19:12.95916+07:00",
    "updated_at": "1996-06-22T22:19:12.95916+07:00"
  }
}
```

### Login
- Headers:
  - Content-Type: "application/json"
- Method: POST
- Endpoint: /user/login
- Request:
```json
{
  "user_name": "string, required",
  "password": "string, required"
}
```
- Example Request:
```json
{
  "user_name": "naufaldy.mahas",
  "password": "qwerty123",
}
```

- Example Response:
```json
{
  "status": 200,
  "message": "",
  "data": {
    "access_token": "randomtoken",
    "refresh_token": "randomtoken"
  }
}
```

### Search User
- Headers:
  - Content-Type: "application/json"
- Method: GET
- Endpoint: /user?search={search}
  - if {search} null then get all users

- Example Response:
```json
{
  "status": 200,
  "message": "",
  "data": [
    {
      "user_id": 1,
      "user_name": "naufaldy.mahas",
      "employee_id": "210001",
      "birth_place": "Jakarta",
      "birh_date": "1996-06-22T00:00:00Z",
      "created_at": "1996-06-22T22:19:12.95916+07:00",
      "updated_at": "1996-06-22T22:19:12.95916+07:00"
    },
    ...
  ]
}
```
