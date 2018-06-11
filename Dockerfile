FROM golang:alpine AS builder

RUN apk update && apk add git && apk add ca-certificates

WORKDIR $GOPATH/src/github.com/williammartin/trumpeter
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /trumpeter .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /trumpeter ./
ENTRYPOINT ["./trumpeter"]
