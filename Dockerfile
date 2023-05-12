FROM golang:1.20

WORKDIR /app/backend

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY backend/go.mod backend/go.sum ./
RUN go mod download && go mod verify

COPY backend .
COPY my-ts-lib/bundle_dist ../my-ts-lib/bundle_dist
RUN go build -o /app/backend/core-main cmd/main.go && pwd && ls

CMD ["/app/backend/core-main"]