# syntax=docker/dockerfile:1

# builder stage
FROM docker.io/golang:latest AS builder

WORKDIR /src

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux  go build -o seized main.go

# actual image
FROM scratch

COPY --from=builder /src/seized /bin/seized

WORKDIR /srv/
COPY static /srv/static

EXPOSE 4000

CMD ["/bin/seized"]
