# Reserve

This is a simple reservation system. The system supports user management, resource management, and reservation management, among other features.

## Getting Started

Clone the project:

```shell
git clone https://github.com/csguojin/reserve
```

Configure the MySQL connection and other parameters:

```shell
vim ./config/config.yaml
```

Start the application:

```shell
go run main.go
```

If you are using Docker, you can build the Docker image for the project:

```shell
docker build -t reserve:latest .
docker run -p 8080:8080 reserve:latest
```

## API

View the RESTful APIs of this project in [Swagger Editor](https://editor.swagger.io/?url=https://raw.githubusercontent.com/csguojin/reserve/main/docs/openapi.yaml).
