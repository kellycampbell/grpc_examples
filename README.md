# grpc_examples

To run the cats_n_pups example:

Everything for this section is run from within the `cats_n_pups` directory.

```
cd cats_n_pups
```

Build & run the JSON -> gRPC gateway:
```
bazel build protos:gateway
./bazel-bin/protos/linux_amd64_stripped/gateway -logtostderr  # This will be different depending on your OS
```

This will take a while the first time as bazel builds all the tools and dependencies (could be 10+ mins).

While you could use `bazel run` for the above, that would tie up the bazel server and the below commands would
wait until the bazel server becomes available.


In another terminal build and start the catpics server:
```
bazel build services/catpics:catpic_server
./bazel-bin/services/catpics/catpic_server
```

At this point you should have a grpc-gateway running on port 8080, and a catpic_server listening on port 50051 for gRPC calls.

Send a REST API request to the catpic_server via the grpc_gateway:

```
http -v http://localhost:8080/api/v1/catpics/asdf
```

