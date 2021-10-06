

## Build Library

From this dir with Rust installed (need gnu version not msvc version for Windows), run:

    go run ./sdk-core-bridge build

## Generate Protos

From this dir with [Proto prereqs](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers),
run:

    go run ./sdk-core-bridge genprotos