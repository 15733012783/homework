FROM golang:1.20

RUN mkdir /grpc

COPY ./ /grpc

WORKDIR /grpc

EXPOSE 8882

RUN go mod tidy

RUN go build main.go

ENTRYPOINT ["go","run","main.go"]

#FROM scratch
#
#COPY --from=build /grpc/ /
#
#CMD ["/main"]