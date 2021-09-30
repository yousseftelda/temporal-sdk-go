
(under development)

TODO: Right now expects static lib to be copied from `target` to `lib`. In the future, will have build script move it
based on triple. And will make sure to only commit the static libs when tagging. E.g.

    cp target/debug/libtemporal_sdk_core_bridge.a lib/linux-x86_64/