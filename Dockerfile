FROM golang:1.11-alpine3.8 AS build
# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git
COPY . .
WORKDIR /

RUN go get github.com/gorilla/mux
RUN go get gopkg.in/src-d/go-billy.v4
RUN go get gopkg.in/src-d/go-git.v4
RUN GOOS=linux go build -ldflags="-s -w" -o go/bin/test go/main.go

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build go/bin /go/bin
COPY --from=build go/dist /go/bin/dist
EXPOSE 80
ENTRYPOINT /go/bin/test --port 80
