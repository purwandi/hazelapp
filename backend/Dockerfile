FROM golang:latest as build
ADD . /github.com/purwandi/hazelapp
WORKDIR /github.com/purwandi/hazelapp
RUN go get -u github.com/a-urth/go-bindata/...
RUN rm -rf gateway/schema/bindata.go
RUN make build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./gateway/cmd/http/

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /github.com/purwandi/hazelapp/main /github.com/purwandi/hazelapp/
WORKDIR /github.com/purwandi/hazelapp/
CMD ["/github.com/purwandi/hazelapp/main"]
