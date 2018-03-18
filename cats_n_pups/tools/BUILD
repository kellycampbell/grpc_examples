load("@io_bazel_rules_go//proto:compiler.bzl", "go_proto_compiler")


go_proto_compiler(
    name = "go_grpc_gateway",
    plugin = "@com_github_grpc_ecosystem_grpc_gateway//protoc-gen-grpc-gateway",
    suffix = ".pb.gw.go",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_grpc_ecosystem_grpc_gateway//runtime:go_default_library",
        "@com_github_grpc_ecosystem_grpc_gateway//utilities:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//grpclog:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)
