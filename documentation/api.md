### Task API

> POST /api/v1/tasks

```bash
curl --location --request POST 'localhost:8080/api/v1/tasks' \
--header 'Content-Type: application/json' \
--data-raw '    {
        "title": "t2",
        "description": "t2"
    }
]'
```

> GET /api/v1/tasks/:id

```bash
curl --location --request GET 'localhost:8080/tasks/1'
```

> DELETE /api/v1/tasks/:id

```bash
curl --location --request DELETE 'localhost:8080/tasks/1'
```

> PUT /api/v1/tasks

```bash
curl --location --request PUT 'localhost:8080/tasks/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "1",
    "title": "Update 2",
    "body": "Update 2"
}'
```

> GET /api/v1/tasks/

```bash
curl --location --request GET 'localhost:8080/api/v1/tasks'
```
