POC - under development

Notes:

* To support opt-in while leaving pure-Go client around, probably would:
  * Move this package to `internal/sdkcore`
  * In `client/client.go`, make `var clientFactory func(Options) (Client, error)` and an init that sets it to
    `internal.NewClient` only if nil
  * Make a `client/client_sdkcore.go` with a build tag for `temporal-sdk-core` and an init that sets `clientFactory` to
    `sdkcore.NewClient` regardless of wither nil
  * Do the same for `worker`
  * For package-level functions with side effects, abstract into interface w/ default pure-go impl that can be
    overridden with sdk-core equivalents

Compatibility issues:

* TODO