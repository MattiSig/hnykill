# --- build stage ---
FROM golang:1.25-alpine AS build
WORKDIR /src

# Cache dependencies first.
COPY go.mod go.sum ./
RUN go mod download

# Build the statically-linked binary (templates are embedded via go:embed).
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/hnykill .

# --- run stage ---
FROM alpine:3.20
RUN adduser -D -u 10001 app
WORKDIR /app
COPY --from=build /app/hnykill /app/hnykill
USER app

# Railway provides $PORT; default to 8080 locally.
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/hnykill"]
