//go:build generate

package generate

//go:generate mkdir -p .bin

//go:generate env GOBIN=${PWD}${CI_GOBIN_PREFIX}/.bin go install github.com/bufbuild/buf/cmd/buf
//go:generate env GOBIN=${PWD}${CI_GOBIN_PREFIX}/.bin go install google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate env GOBIN=${PWD}${CI_GOBIN_PREFIX}/.bin go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go

//go:generate env PATH=${PWD}${CI_GOBIN_PREFIX}/.bin:$PATH buf lint
//go:generate env PATH=${PWD}${CI_GOBIN_PREFIX}/.bin:$PATH buf generate rpc/matchmaking/v1/api.proto -v
