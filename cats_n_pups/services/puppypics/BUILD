package(default_visibility=["//visibility:public"])
# load("@build_bazel_rules_typescript//:defs.bzl", "ts_library")

# ts_library(
#     name = "my_code",
#     srcs = glob(["*.ts"]),
#     # deps = ["//path/to/other:library"],
#     # tsconfig = "//:tsconfig.json",
# )

load("@build_bazel_rules_nodejs//:defs.bzl", "nodejs_binary")


nodejs_binary(
    name = "hello",
    # # Ordinarily this defaults to //:node_modules
    # node_modules = "@foo//:node_modules",
    entry_point = "cats_n_pups/services/puppypics/server.js",
    data = [
        "server.js",
    ],
)