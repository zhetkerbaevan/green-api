FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o out ./cmd/green-api

FROM gcr.io/distroless/base

COPY --from=build /app/out /app/out

CMD ["/app/out"]
