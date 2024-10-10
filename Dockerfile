FROM --platform=$BUILDPLATFORM golang:1.22.8@sha256:b274ff14d8eb9309b61b1a45333bf0559a554ebcf6732fa2012dbed9b01ea56f AS build

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
