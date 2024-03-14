FROM golang:alpine AS build
RUN apk update && apk add git
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go mod tidy
RUN go build -o /go/src/app/main main.go

FROM alpine
COPY --from=build /go/src/app/main /go/src/app/
COPY --from=build /go/src/app/.env /go/src/app/

COPY swagger.json .
EXPOSE 8080
ENTRYPOINT ["/go/src/app/main"]