# ─────────────────────────────────────────────
# Stage 1: Build Vue/Vite frontend
# ─────────────────────────────────────────────
FROM node:22-alpine AS frontend-builder

WORKDIR /app/frontend

# Cache npm install separately from source
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci --prefer-offline

COPY frontend/ ./
RUN npm run build

# ─────────────────────────────────────────────
# Stage 2: Run Go tests + build binary
# ─────────────────────────────────────────────
FROM golang:1.26-alpine AS go-builder

# gcc + musl required for mattn/go-sqlite3 (CGO)
RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

# Cache Go module downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Copy frontend build output into the expected location
COPY --from=frontend-builder /app/frontend/dist ./web/html

# ── Run tests ──────────────────────────────────
RUN go test \
      -v \
      -count=1 \
      ./util/... \
      ./util/common/... \
      2>&1 | tee /tmp/test-results.txt \
    && echo "✅ Tests passed"

# ── Build binary ───────────────────────────────
RUN go build \
      -ldflags "-w -s" \
      -tags "with_quic,with_grpc,with_utls,with_acme,with_gvisor" \
      -o /app/zpanel \
      ./main.go \
    && echo "✅ Build succeeded"

# ─────────────────────────────────────────────
# Stage 3: Minimal runtime image
# ─────────────────────────────────────────────
FROM alpine:3.21

# ca-certificates for TLS outbound connections
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy binary and test results from builder
COPY --from=go-builder /app/zpanel          ./zpanel
COPY --from=go-builder /tmp/test-results.txt ./test-results.txt

# Default data directory (mount a volume here in production)
RUN mkdir -p /app/data

# Expose web UI port (adjust to match your config)
EXPOSE 2096

ENTRYPOINT ["/app/zpanel"]
