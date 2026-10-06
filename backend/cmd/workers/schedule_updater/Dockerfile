FROM golang:1.23.1-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/scheduleUpdater ./cmd/workers/schedule_updater

FROM alpine:3.21
RUN addgroup -S app && adduser -S -G app app
WORKDIR /app
COPY --from=build --chown=app:app /out/scheduleUpdater /app/scheduleUpdater
USER app:app
ENTRYPOINT ["/app/scheduleUpdater"]


