FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o fairshare

FROM alpine:3.19
COPY --from=builder /app/fairshare /bin/fairshare
MAINTAINER manojnakp
EXPOSE 8080
CMD ["/bin/fairshare"]
