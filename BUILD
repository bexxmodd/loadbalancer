load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "main",
    srcs = ["main.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//server:server",
    ],
)