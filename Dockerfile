FROM golang:1.20.2-alpine3.17 AS builder
  
WORKDIR /build
RUN adduser -u 10001 -D handsome

ENV GOPROXY https://goproxy.cn

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o aws_ec2_status .

FROM alpine:3.17 AS final

RUN adduser -u 10001 -D handsome
RUN chmod -R 755 /usr/local/


WORKDIR /usr/local/
COPY --from=builder /build/aws_ec2_status /usr/local/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

USER handsome
ENTRYPOINT ["/usr/local/aws_ec2_status"]
