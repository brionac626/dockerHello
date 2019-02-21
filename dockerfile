FROM golang:1.11.5 AS bs
MAINTAINER AutumnAoi

ENV GO111MODULE=on

RUN git clone https://github.com/brionac626/dockerHello.git
RUN cd ./dockerHello && go build -o main

FROM alpine

WORKDIR /app
ENV PIC_DIR /picture
VOLUME [ "/picture" ]

COPY --from=bs /go/dockerHello /app/
ENTRYPOINT [ "./main" ]