package(default_visibility=['//visibility:public'])

load("@com_google_protobuf//:protobuf.bzl", "py_proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
go_prefix("google/api")


filegroup(
    name = "api_protos",
    srcs = [
        "google/api/annotations.proto",
        "google/api/http.proto",
    ]
)

proto_library(
    name = "api_proto",
    srcs = [
        "google/api/annotations.proto",
        "google/api/http.proto",
    ],
    deps = [
        "@com_google_protobuf//:descriptor_proto",
    ],
)

go_proto_library(
    name = "api_proto_go",
    proto = ":api_proto",
    deps = [
        "@com_github_golang_protobuf//protoc-gen-go/descriptor:go_default_library",
        "@org_golang_google_genproto//googleapis/api/annotations:go_default_library",
    ],
    importpath = "google.golang.org/genproto/googleapis/api"
)
