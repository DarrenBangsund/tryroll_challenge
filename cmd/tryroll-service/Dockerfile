FROM golang as builder

WORKDIR /
COPY . ./app
WORKDIR /app/cmd/tryroll-service
RUN CGO_ENABLED=0 go build -o build/service .

FROM alpine:3.17.0

WORKDIR /app

RUN apk add ca-certificates

COPY --from=builder /app/cmd/tryroll-service/build/service tryroll-service

ENV PATH="${PATH}:/app"

EXPOSE 8002

CMD [ "/app/tryroll-service" ]