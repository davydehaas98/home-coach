FROM golang:1.22.5 AS build

ARG TARGET_OS
ARG TARGET_ARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGET_OS GOARCH=$TARGET_ARCH go build -o /app

FROM scratch AS runtime
COPY --from=build /app /app
USER app

ENTRYPOINT ["/app"]
