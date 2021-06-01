FROM golang:1.16 as goBuilder

WORKDIR /go_build
COPY . /go_build

RUN go build -ldflags "-X main.VERSION=$VERSION_TAG" -o /go/bin/app /go_build/main.go


FROM node:15 as jsBuilder

WORKDIR /js_build
COPY fe /js_build
RUN yarn build



FROM ubuntu:20.04

COPY --from=goBuilder /go/bin/app /project/app
COPY --from=jsBuilder /js_build/build /project/public
ARG PORT
ARG DATABASE_URL
ARG REDIS_URL

ENTRYPOINT ["/project/app"]