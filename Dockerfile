FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o .

FROM ubuntu:22.04
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .
ENTRYPOINT ["./mobile_films_backend"]