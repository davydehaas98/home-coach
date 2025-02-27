FROM --platform=$BUILDPLATFORM golang:1.24.0@sha256:3f7444391c51a11a039bf0359ee81cc64e663c17d787ad0e637a4de1a3f62a71 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .

ARG TARGETOS ARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /app

FROM scratch AS runtime
COPY --from=build /app /app
USER app

ENTRYPOINT ["/app"]
