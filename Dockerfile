FROM golang:latest as builder 

LABEL maintainer = "Wesley Mochiemo <wesleymochiemo@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -o main .

# Starting a new stage from scratch 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 10000

CMD ["./main"]
