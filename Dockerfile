FROM --platform=$BUILDPLATFORM golang:1.22.11@sha256:d5b17d684180648e16ea974bea677498945e8b619f7b26325958d8d99e97f9ea AS build

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
