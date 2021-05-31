FROM golang:1.16-alpine

WORKDIR /opt
ENV GO111MODULE on
ENV CGO_ENABLED 0

COPY . .
RUN go mod download
RUN go build -o goplug .
RUN go build -o ./plugin-sql ./plugins/sql/main.go

CMD ["goplug", "--plugin", "./plugin-sql"]