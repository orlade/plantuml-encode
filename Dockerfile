# The produced image is published to https://hub.docker.com/r/orlade/plantuml-encode

ARG go_ver=1.13.8
ARG alpine_ver=3.11

FROM golang:${go_ver}-alpine${alpine_ver} as builder

RUN apk --no-cache add git make

WORKDIR /base

COPY . .

RUN make build

FROM golang:${go_ver}-alpine${alpine_ver} as runner

COPY --from=builder /base/dist/plantuml-encode /

ENTRYPOINT ["/plantuml-encode"]
