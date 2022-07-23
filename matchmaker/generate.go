//go:build generate

package generate

//go:generate mkdir -p .bin

//go:generate env GOBIN=${PWD}/.bin go install github.com/bufbuild/buf/cmd/buf
//go:generate env GOBIN=${PWD}/.bin go install google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate env GOBIN=${PWD}/.bin go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go

//go:generate echo "Installed tools"

//go:generate env PATH=${PWD}/.bin:$PATH buf lint

//go:generate echo "Verified protobuf definitions"

//go:generate env PATH=${PWD}/.bin:$PATH buf generate rpc/v1/api.proto -v

//go:generate echo "Generated protobuf code"
