FROM --platform=$BUILDPLATFORM golang:1.22.11@sha256:cd31706dd21bb47260286c0105f928b4d938ee7fa7b44ae398b8cc1f84d3150f AS build

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
