package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_prefix")

go_prefix("github.com/kellycampbell/grpc_examples/cats_n_pups")

exports_files(["package.json", "package-lock.json"])

filegroup(name = "node_modules", srcs = glob([
    "node_modules/**/*.js",
    "node_modules/**/*.d.ts",
    "node_modules/**/*.json",
]))