FROM golang:1.18-alpine as build

RUN apk update && apk add make

WORKDIR /app

COPY ./recommender/go.mod ./
COPY ./recommender/go.sum ./

RUN go mod download

COPY ./recommender/main.go ./main.go
COPY ./recommender/cmd ./cmd
COPY ./recommender/pkg ./pkg

RUN mkdir -p /app/bin
RUN go build -o ./bin/server main.go

FROM alpine:latest

WORKDIR /

COPY --from=build /app/bin/server /server
COPY ./.env /.env

EXPOSE 9090

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

ENTRYPOINT ["/server"]