module github.com/GoogleCloudPlatform/terraform-validator

require (
	cloud.google.com/go/bigtable v1.13.0
	github.com/GoogleCloudPlatform/config-validator v0.0.0-20211122204404-f3fd77c5c355
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/errwrap v1.0.0
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/terraform-json v0.13.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1
	github.com/hashicorp/terraform-provider-google v1.20.1-0.20220502223715-1ef3676c447e
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/onsi/gomega v1.17.0 // indirect
	github.com/open-policy-agent/opa v0.36.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	google.golang.org/api v0.70.0
	google.golang.org/genproto v0.0.0-20220218161850-94dd64e39d7c
	google.golang.org/grpc v1.44.0
)

go 1.16

replace github.com/hashicorp/terraform-provider-google => github.com/modular-magician/terraform-provider-google v1.20.1-0.20220502222358-8f00f7dcf9c6
