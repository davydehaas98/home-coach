FROM --platform=$BUILDPLATFORM golang:1.22.6@sha256:2bd56f00ff47baf33e64eae7996b65846c7cb5e0a46e0a882ef179fd89654afa AS build

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
