FROM --platform=$BUILDPLATFORM golang:1.22.10@sha256:5110696075239f78fc471cc1ceaa75e809d12b39f0aefd9c819b81c66b04927c AS build

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
