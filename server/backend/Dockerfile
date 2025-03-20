FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN make

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/server .

COPY --from=builder /app/internal/ /root/internal

RUN chmod +x server

EXPOSE 8000

CMD [ "/root/server" ]

