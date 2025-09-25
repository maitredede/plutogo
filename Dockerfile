#syntax=docker/dockerfile:1
FROM alpine:edge AS alpine-base
ARG ALPINE_MIRROR=https://dl-cdn.alpinelinux.org/alpine
RUN set -ex ;\
    echo "${ALPINE_MIRROR}/edge/main" > /etc/apk/repositories ;\
    echo "${ALPINE_MIRROR}/edge/community" >> /etc/apk/repositories ;\
    echo "@testing ${ALPINE_MIRROR}/edge/testing" >> /etc/apk/repositories ;\
    apk upgrade --no-cache --available

FROM alpine-base AS build-plutobook-alpine
RUN set -ex ;\
    apk add --no-cache build-base cairo-dev curl-dev expat-dev fontconfig-dev freetype-dev \
    git harfbuzz-dev icu-dev libjpeg-turbo-dev libwebp-dev \
    meson ninja-build
WORKDIR /src/plutobook
ARG PLUTOBOOK_VERSION=v0.9.0
RUN git clone https://github.com/plutoprint/plutobook.git . --branch=${PLUTOBOOK_VERSION}
RUN meson setup build --prefix /usr-plutobook
RUN meson compile -C build
RUN meson install -C build 
RUN sed -i 's|/usr-plutobook|/usr|g' /usr-plutobook/lib/pkgconfig/plutobook.pc

FROM golang:1.25-alpine AS build-go
ARG ALPINE_MIRROR=https://dl-cdn.alpinelinux.org/alpine
RUN set -ex ;\
    echo "${ALPINE_MIRROR}/edge/main" > /etc/apk/repositories ;\
    echo "${ALPINE_MIRROR}/edge/community" >> /etc/apk/repositories ;\
    echo "@testing ${ALPINE_MIRROR}/edge/testing" >> /etc/apk/repositories ;\
    apk upgrade --no-cache --available ;\
    apk add --no-cache build-base cairo-dev curl-dev expat-dev fontconfig-dev freetype-dev \
    git harfbuzz-dev icu-dev libjpeg-turbo-dev libwebp-dev
COPY --from=build-plutobook-alpine /usr-plutobook/ /usr/
WORKDIR /src
COPY go.mod go.sum ./
ARG GOPROXY
ARG GONOSUMDB
ARG CGO_ENABLED=1
RUN go mod download -x
COPY . ./
WORKDIR /src
COPY go.mod go.sum ./
ARG GOPROXY
ARG GONOSUMDB
ARG CGO_ENABLED=1
RUN go mod download -x
COPY . ./
RUN set -x ;\
    go get -v tool ;\
    go generate ./... || true ;\
    set -ex ;\
    go vet ./... ;\
    go test -short ./... ;\
    go install ./...

FROM alpine-base AS application-base
RUN set -ex ;\
    apk add --no-cache font-awesome font-dejavu font-inconsolata font-misc-misc font-noto \
    font-noto-cjk font-noto-extra font-terminus ;\
    apk add --no-cache cairo fontconfig freetype harfbuzz icu-data-full icu-libs libcurl libexpat libturbojpeg libwebp
COPY --from=build-plutobook-alpine /usr-plutobook/ /usr/

FROM application-base AS example-basic
COPY --from=build-go /go/bin/basic /usr/local/bin/basic
ENTRYPOINT [ "/usr/local/bin/basic" ]

FROM application-base AS example-pagerendering
COPY --from=build-go /go/bin/pagerendering /usr/local/bin/pagerendering
ENTRYPOINT [ "/usr/local/bin/pagerendering" ]
