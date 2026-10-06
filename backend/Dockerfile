FROM golang:1.23.1-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/homeworkManager ./cmd/service

FROM alpine:3.21
RUN addgroup -S app && adduser -S -G app app
WORKDIR /app
COPY --from=build --chown=app:app /out/homeworkManager /app/homeworkManager
USER app:app
EXPOSE 8888 8889
ENTRYPOINT ["/app/homeworkManager"]


