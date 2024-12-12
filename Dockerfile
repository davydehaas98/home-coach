FROM --platform=$BUILDPLATFORM golang:1.22.10@sha256:9855006ddcf40a79e9a2d90df11870331d24bcf2354232482ae132a7ba7b624f AS build

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
