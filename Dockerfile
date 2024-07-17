FROM golang:1.22.2 AS build

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o out ./cmd/green-api

FROM gcr.io/distroless/base

COPY --from=build /app/out /app/out

CMD ["/app/out"]
