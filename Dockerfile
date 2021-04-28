FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./api/main.go

EXPOSE 4513
EXPOSE 4514

ENTRYPOINT [ "/app/main" ]