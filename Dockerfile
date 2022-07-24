FROM golang:1.16-alpine
RUN go version
#ENV GOPATH=/

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . ./
RUN go mod download

#RUN go env -w GO111MODULE=on
RUN go build -o helloworld cmd/server/main.go
#RUN go run cmd/server/main.go


EXPOSE 8000
CMD ["./helloworld"]
