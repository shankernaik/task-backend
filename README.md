# task-backend

Task is a simple task management service deployed using Echo framework.

Application layout as below

```
├── LICENSE
├── Makefile
├── README.md
├── go.mod
├── go.sum
└── task-backend
    ├── cmd
    │   └── server
    │       └── main.go
    ├── controllers
    │   └── task_controller.go
    ├── models
    │   └── type.go
    ├── routes
    │   └── route.go
    └── services
        └── service.go
```

The task service api documentation can be found here [documentation](/documentation/api.md)

To build the application :

`$ make build`

Output (Executable located in target folder):

```
cd task-backend/cmd/server/ && go build -o ../../../target/task-backend
```

To build and run the application :

`$ make run`

Output (Application runs on port=:8080) :

```bash
cd task-backend/cmd/server/ && go build -o ../../../target/task-backend
./target/task-backend

        ################################## ENDPOINT(S) #####################################

GET /api/v1/tasks
POST /api/v1/tasks
GET /api/v1/tasks/:id
DELETE /api/v1/tasks/:id
POST /api/v1/tasks

---

/ **/\_**/ / **_
/ _// **/ _ \/ _ \
/**\_/\__/_//\_/\_**/ v3.3.10-dev
High performance, minimalist Go web framework
https://echo.labstack.com
****************\_\_\_\_****************O/**\_\_\_**
O\
⇨ http server started on [::]:8080
```
