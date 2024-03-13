FROM golang:1.22.1

RUN mkdir /app

COPY ./ /app

WORKDIR /app

EXPOSE 8883

RUN go mod tidy

RUN go build main.go

ENTRYPOINT ["go", "run","main.go"]