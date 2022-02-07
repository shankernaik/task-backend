# Download Base Image of golang
FROM golang

# Clone the repo
RUN git clone https://github.com/shankernaik/task-backend.git

# Move the project folder inside src/
RUN mv task-backend src/

# Enable downloading of packages from go.mod
RUN go env -w GO111MODULE=on

# Clean up the go.mod and go.sum files
RUN cd src/task-backend/ && go mod tidy

# Download dependencies from go.mod
RUN cd src/task-backend/ && go mod download

# Go to main.go folder and build the project
RUN cd src/task-backend/task-backend/cmd/server && go build -o ../../../target/task-backend

# Run the executable
CMD [ "./src/task-backend/target/task-backend" ]