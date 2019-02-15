FROM golang:1.11.2 AS bs
MAINTAINER AutumnAoi

ENV GO111MODULE=on

RUN git clone https://github.com/brionac626/dockerHello.git
RUN cd ./dockerHello && go build -o hello

FROM alpine
WORKDIR /app
COPY --from=bs /go/dockerHello /app/
ENTRYPOINT [ "./hello" ]