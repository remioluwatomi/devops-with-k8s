FROM golang:1.25.7-alpine

WORKDIR /log-output
COPY go.mod ./ 
RUN go mod download

COPY . .

RUN go build -o log-output main.go

ENTRYPOINT [ "./log-output" ]

