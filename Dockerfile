FROM golang:1.13-alpine
WORKDIR /go/src/github.com/pedidosya/web-asp
COPY . .
RUN go build -o service *.go
#
# Run stage
#
FROM alpine
COPY --from=0 /go/src /go/src
WORKDIR /go/src/github.com/pedidosya/web-asp
CMD ./service
EXPOSE 8080
