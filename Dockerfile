ARG GO_VERSION=1.25

# ---------- dev: Go tools + Air, code is mounted from your PC ----------
FROM golang:${GO_VERSION}-alpine AS dev
ARG GOPROXY=https://goproxy.io,direct
ENV GOPROXY=$GOPROXY CGO_ENABLED=0
RUN go install github.com/air-verse/air@latest
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
CMD ["air", "-c", ".air.toml"]

# ---------- build: compile the binaries ----------
FROM golang:${GO_VERSION}-alpine AS build
ARG GOPROXY=https://goproxy.io,direct
ENV GOPROXY=$GOPROXY CGO_ENABLED=0
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api \
 && go build -trimpath -ldflags="-s -w" -o /out/migrate ./cmd/migrate

# ---------- prod: only the binaries ----------
FROM alpine:3.20 AS prod
RUN adduser -D -u 10001 app
WORKDIR /app
COPY --from=build /out/api /out/migrate /usr/local/bin/
COPY migrations ./migrations
USER app
EXPOSE 8080
ENTRYPOINT ["api"]