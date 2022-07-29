FROM golang:1.18 AS build
RUN go version

RUN mkdir /src
ADD . /src
WORKDIR /src

COPY go.mod /src
COPY go.sum /src

RUN go mod download
RUN go build -o /tmp/st-booker .


FROM alpine:edge AS runtime

RUN mkdir /bin
WORKDIR /bin

COPY --from=build /tmp/* /app

CMD pm-context