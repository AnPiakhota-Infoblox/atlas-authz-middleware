/*
	Package wasm provides fast and light authorization by
    direct call to OPA ingrained via its Wasm SDK.

	See https://www.openpolicyagent.org/docs/latest/integration/#webassembly-wasm

	An example of integrating it in services:

	import authz_fl "github.com/infobloxopen/atlas-authz-middleware/pkg/eval/wasm"

	interceptors = append(interceptors, authz_fl.UnaryServerInterceptor(
		authz_fl.ForApplicaton(appID),
		authz_fl.WithLogger(logger),
		authz_fl.WithBundleResourcePath("/bundle/bundle.tar.gz"),
		authz_fl.WithDecisionPath("/authz/rbac/validate_v1"),
	))
*/
package wasm
