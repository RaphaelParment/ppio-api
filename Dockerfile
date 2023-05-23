FROM golang:1.20.4-alpine3.18 AS builder

WORKDIR /app

COPY . ./

RUN go build -o ppio-api

FROM scratch
COPY --from=builder /app/ppio-api /ppio-api

EXPOSE 9001

CMD ["/ppio-api"]