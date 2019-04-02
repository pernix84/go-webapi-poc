FROM golang:1.12

WORKDIR /go/src/go-webapi-poc
COPY . . 
RUN go get -d -v ./...
RUN go install -v ./...

# ADD . /go-webapi-poc  

RUN go build -o go-webapi-poc .

EXPOSE 8000:8000

CMD ["./go-webapi-poc"]