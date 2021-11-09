FROM golang:1.17 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /bank-api

COPY / . 

RUN go get -d -v ./...  && \
    go install -v ./... && \
    go build -v -o /bin/bank-apid

# Image
FROM alpine

ENV TZ=America/Fortaleza

WORKDIR /bank-api

RUN apk update --no-cache && \
    apk --no-cache add tzdata ca-certificates && \
    adduser -D --uid 1000 bank bank 

COPY --chown=bank:bank --from=builder /bin/bank-apid .

ENTRYPOINT /bank-api/bank-apid

USER bank