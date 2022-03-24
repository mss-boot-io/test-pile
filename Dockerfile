FROM golang:alpine as builder

MAINTAINER lwnmengjing

ENV GOPROXY https://goproxy.cn/

WORKDIR /go/release
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add tzdata git

COPY go.mod ./go.mod
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o test-pile .

FROM alpine

COPY --from=builder /go/release/test-pile /

EXPOSE 8060

ENTRYPOINT [ "/test-pile" ]