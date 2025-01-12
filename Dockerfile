# syntax=docker/dockerfile:1

# builder stage
FROM docker.io/golang:latest AS builder

WORKDIR /src

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o seized main.go

# static error pages from https://github.com/tarampampam/error-pages
# rather than worrying about building it, let's just go ahead and pick the
# pre-made version
FROM ghcr.io/tarampampam/error-pages:latest AS error-pages

# actual image
FROM scratch

# copy the binary from the builder
COPY --from=builder /src/seized /bin/seized
# copy all the error pages from the tarampampam's image
COPY --from=error-pages /opt/html/connection/ /srv/static/
# add our own
COPY static/* /srv/static/

WORKDIR /srv/

EXPOSE 4000

CMD ["/bin/seized"]
