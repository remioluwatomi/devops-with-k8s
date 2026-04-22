FROM golang:1.25.7-alpine as build

WORKDIR /log-output
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o log-output main.go



FROM alpine:3.19

WORKDIR /log-output

COPY --from=build log-output .

ENTRYPOINT [ "./log-output" ]

