FROM golang as builder

WORKDIR /
COPY . ./app
WORKDIR /app/cmd/gql-api
RUN CGO_ENABLED=0 go build -o build/service .

FROM alpine:3.17.0

WORKDIR /app

RUN apk add ca-certificates

COPY --from=builder /app/cmd/gql-api/build/service gql-api

ENV PATH="${PATH}:/app"

CMD [ "/app/gql-api" ]