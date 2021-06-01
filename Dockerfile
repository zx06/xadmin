FROM golang:1.16 as goBuilder

WORKDIR /go_build
COPY . /go_build

RUN go build -ldflags "-X main.VERSION=$VERSION_TAG" -o /go/bin/app /go_build/main.go


FROM alpine
COPY --from=goBuilder /go/bin/app /project/app
ENTRYPOINT ["/project/app"]