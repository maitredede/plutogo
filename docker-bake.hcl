variable "ALPINE_MIRROR" {
    default = "https://dl-cdn.alpinelinux.org/alpine" # original
}

variable "GOPROXY" {
    default = "direct"
}

variable "GONOSUMDB" {
    default = ""
}

variable "PLUTOBOOK_VERSION" {
    default = "v0.9.0"
}

group "default" {
    targets = [
    "example-basic",
    "example-pagerendering",
    ]
}

target "_common" {
    context = "."
    dockerfile = "Dockerfile"
    args = {
        ALPINE_MIRROR = "${ALPINE_MIRROR}"
        GOPROXY = "${GOPROXY}"
        GONOSUMDB = "${GONOSUMDB}"
        PLUTOBOOK_VERSION = "${PLUTOBOOK_VERSION}"
    }
    platforms = ["linux/amd64", "linux/arm64"]
    pull = true
}

target "example-basic" {
    inherits = ["_common"]
    target = "example-basic"
}

target "example-pagerendering" {
    inherits = ["_common"]
    target = "example-pagerendering"
}
