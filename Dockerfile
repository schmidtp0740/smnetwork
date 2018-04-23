FROM golang:1.8
WORKDIR /go/src/app
COPY . .
VOLUME [ "/go/src/app" ]
ENV CGO_ENABLED=0
RUN go get && go build