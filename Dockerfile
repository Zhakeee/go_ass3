FROM golang:1.17-alpine

WORKDIR /assignment3

COPY ./ ./
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git && \
    go get -u gorm.io/driver/postgres && \    
    go get -u gorm.io/gorm && \
    go mod tidy

RUN go build -o assignment3 ./main.go

CMD ["./assignment3"] 