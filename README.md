Running in root directory
$ go run ./cmd/JWT/main.go

```
GET http://localhost:5000/api/v1/auth/token?guid=<valid_guid>
```

Успешный запрос вернет:

```json
201 Created
```

Если GUID будет не соответствовать формату, то вернется ошибка:

```json
400 BadRequest
```

Если GUID уже задействован, то вернется ошибка:

```json
401 Unauthorized
```

Если в процессе возникнут другие неприятности, то вернется ошибка:

```json
500 InternalServerError
```

### Refresh

Маршрут выполняет Refresh операцию на пару Access, Refresh токенов.

```
POST http://localhost:5000/api/v1/auth/refresh

Authorization: Bearer <valid_JWT_Token>
Content-Type: application/json
```

```json
{
  "refresh_token": <valid_refresh_token>
}
```

Успешный запрос вернет:

```json
200 OK
```

Если запрос будет не соответствовать формату, то вернется ошибка:

```json
400 BadRequest
```

Если access_token ещё действителен или refresh_token не действителен, то вернется ошибка:

```json
401 Unauthorized
```

Если в процессе возникнут другие неприятности, то вернется ошибка:

```json
500 InternalServerError
```


## Тестовое задание на позицию Junior Backend Developer

**Используемые технологии:**

- Go
- JWT
- MongoDB

**Задание:**

Написать часть сервиса аутентификации.

Два REST маршрута:

- Первый маршрут выдает пару Access, Refresh токенов для пользователя с идентификатором (GUID) указанным в параметре запроса
- Второй маршрут выполняет Refresh операцию на пару Access, Refresh токенов

**Требования:**

Access токен тип JWT, алгоритм SHA512, хранить в базе строго запрещено.

Refresh токен тип произвольный, формат передачи base64, хранится в базе исключительно в виде bcrypt хеша, должен быть защищен от изменения на стороне клиента и попыток повторного использования.

Access, Refresh токены обоюдно связаны, Refresh операцию для Access токена можно выполнить только тем Refresh токеном который был выдан вместе с ним.
