FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./



# build go app
RUN go mod download
RUN go build -o go-getir-challenge ./main.go

CMD ["./go-getir-challenge"]