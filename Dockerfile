FROM golang:1.18 AS build
RUN go version

RUN mkdir /src
COPY . /src
WORKDIR /src

RUN go mod download
RUN export CGO_ENABLED=0; go build -o /tmp/st-booker .

FROM alpine:edge AS runtime

RUN mkdir /app
COPY --from=build /tmp/st-booker /app/st-booker
COPY ./internal/config/config.yml /app/config.yml
COPY ./migrations/01_bookings.up.sql /app/migrations/01_bookings.up.sql

ENTRYPOINT ["/app/st-booker"]
