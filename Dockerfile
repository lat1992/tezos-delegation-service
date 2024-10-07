FROM alpine:3

RUN apk add --no-cache make go

ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
ENV PORT 8080
EXPOSE 8080

ADD . /app/
WORKDIR /app

RUN make

CMD ["/app/build/tezos-delegation-service"]
