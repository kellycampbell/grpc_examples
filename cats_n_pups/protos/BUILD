package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library", "go_grpc_library")

filegroup(
    name = "metadata_proto_file",
    srcs = ["metadata.proto"]
)

proto_library(
    name = "metadata_proto",
    srcs = [
        "metadata.proto",
    ],
    deps =  [
        "@com_github_googleapis//:api_proto",
        "@com_google_protobuf//:timestamp_proto",
    ]
)

go_proto_library(
    name = "metadata_proto_go",
    proto = ":metadata_proto",
    deps = [
        "@com_github_googleapis//:api_proto_go",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",
    ],
    importpath = "github.com/kellycampbell/grpc_examples/cats_n_pups/protos/metadata",

)

proto_library(
    name = "catpic_proto",
    srcs = [
        "catpic.proto",
    ],
    deps =  [
        ":metadata_proto",
        "@com_github_googleapis//:api_proto",
        "@com_google_protobuf//:empty_proto",
    ]
)

go_proto_library(
    name = "catpic_proto_go",
    proto = ":catpic_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    deps = [
        ":metadata_proto_go",
        "@com_github_googleapis//:api_proto_go",
        "@com_github_golang_protobuf//ptypes/empty:go_default_library",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",

    ],
    importpath = "github.com/kellycampbell/grpc_examples/cats_n_pups/protos/catpic",
)

go_proto_library(
    name = "catpic_proto_gw",
    proto = ":catpic_proto",
    compilers = ["//tools:go_grpc_gateway"],
    embed = [
        ":catpic_proto_go",
    ],
    deps = [
        ":metadata_proto_go",
    ],
    importpath = "github.com/kellycampbell/grpc_examples/cats_n_pups/protos/catpic",
)


proto_library(
    name = "puppypic_proto",
    srcs = [
        "puppypic.proto",
    ],
    deps =  [
        ":metadata_proto",
        "@com_github_googleapis//:api_proto",
        "@com_google_protobuf//:empty_proto",
    ]
)

go_proto_library(
    name = "puppypic_proto_go",
    proto = ":puppypic_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    deps = [
        ":metadata_proto_go",
        "@com_github_googleapis//:api_proto_go",
        "@com_github_golang_protobuf//ptypes/empty:go_default_library",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",

    ],
    importpath = "github.com/kellycampbell/grpc_examples/cats_n_pups/protos/puppypic",
)

go_proto_library(
    name = "puppypic_proto_gw",
    proto = ":puppypic_proto",
    compilers = ["//tools:go_grpc_gateway"],
    embed = [
        ":puppypic_proto_go",
    ],
    deps = [
        ":metadata_proto_go",
    ],
    importpath = "github.com/kellycampbell/grpc_examples/cats_n_pups/protos/puppypic",
)

GRPC_GATEWAY_DEPS = [
  "@com_github_grpc_ecosystem_grpc_gateway//runtime:go_default_library",
  "@com_github_grpc_ecosystem_grpc_gateway//utilities:go_default_library",
  "@org_golang_google_genproto//googleapis/api/annotations:go_default_library",
  "@org_golang_google_grpc//codes:go_default_library",
  "@org_golang_google_grpc//grpclog:go_default_library",
  "@org_golang_google_grpc//:go_default_library",
  "@org_golang_google_grpc//status:go_default_library",
  "@org_golang_x_net//context:go_default_library",
  "@com_github_golang_glog//:go_default_library",
]


go_binary(
    name = "gateway",
    srcs = ["gateway.go"],
    args = [
        "-logtostderr",
    ],
    deps = [
        ":catpic_proto_gw",
        ":puppypic_proto_gw",
    ] + GRPC_GATEWAY_DEPS

)


load("@org_pubref_rules_protobuf//python:rules.bzl", "py_proto_library")


py_proto_library(
    name = "py",
    protos = [
        "metadata.proto",
        "catpic.proto",
        "puppypic.proto"
    ],
    deps = [
        "@com_github_googleapis//:api_proto_py"
    ],
    imports = [
        "external/com_github_googleapis/",
        "external/com_google_protobuf/src/"
    ],
    inputs = [
        "@com_github_googleapis//:api_protos",
        "@com_google_protobuf//:well_known_protos"
    ],
    with_grpc = True,
    # verbose = 1,
)