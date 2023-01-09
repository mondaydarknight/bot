ARG GO_VERSION=1.19.4

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/api

FROM scratch AS dev

WORKDIR /dist

COPY --from=builder /build/app /build/.env ./

ENTRYPOINT ["./app"]
