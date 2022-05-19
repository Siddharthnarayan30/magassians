FROM golang:1.18.0-alpine3.14
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 9000
RUN go build -o ./server .
CMD ["/app/server"]