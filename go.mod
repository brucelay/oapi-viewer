module github.com/brucelay/oapi-viewer

go 1.24.0

require (
	github.com/pb33f/libopenapi v0.21.7
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/speakeasy-api/jsonpath v0.6.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.9-0.20240815153524-6ea36470d1bd // indirect
	golang.org/x/sys v0.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v1.0.1
	// On second thought this is just a CLI tool so no need to publish it to pkg.go.dev
	v1.0.0
)
