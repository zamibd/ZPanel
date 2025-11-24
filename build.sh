#!/bin/sh

echo "Running linters..."
echo "Linting Backend..."
go fmt ./api/... ./service/... ./cmd/... ./app/... ./config/... ./core/... ./cronjob/... ./database/... ./logger/... ./middleware/... ./network/... ./sub/... ./util/... ./web/...
echo "Backend lint complete ✓"

echo "Linting Frontend..."
cd frontend
npm run lint:check
echo "Frontend lint complete ✓"

echo "Installing dependencies..."
npm i

echo "Building Frontend..."
npm run build

cd ..
echo "Backend"

mkdir -p web/html
rm -fr web/html/*
cp -R frontend/dist/* web/html/

echo "Building Go binary..."
go build -ldflags "-w -s" -tags "with_quic,with_grpc,with_utls,with_acme,with_gvisor" -o sui main.go

echo "Build complete ✓"
