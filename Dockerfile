FROM golang:alpine AS builder

# install gcc, required by sqlite3 driver
# RUN apk add build-base

WORKDIR /app/

# Download dependencies first to leverage Docker cache
COPY go.* .
RUN go mod download

COPY . .

ENV ENV=production

# Test with cache
RUN --mount=type=cache,target=/root/.cache/go-build go test ./...

# Build with cache
RUN --mount=type=cache,target=/root/.cache/go-build go build -o /bin/server ./cmd/crud

FROM alpine

COPY --from=builder /bin/ /bin/

WORKDIR /app/

CMD server --port 8080

EXPOSE 8080