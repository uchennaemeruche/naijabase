FROM golang:latest

WORKDIR /app

COPY . .

CMD [ "go", "run", "main.go" ]