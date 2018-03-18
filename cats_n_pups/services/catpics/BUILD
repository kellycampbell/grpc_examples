package(default_visibility = ["//visibility:public"])

load("@pip_grpcio//:requirements.bzl", "requirement")


py_library(
    name = "grpc_server",
    srcs = [
        "catpic_server.py",
    ],
    deps = [
        "//protos:py",
        requirement("grpcio"),
    ],
)

py_binary(
    name = "catpic_server",
    srcs = [":grpc_server"],
)