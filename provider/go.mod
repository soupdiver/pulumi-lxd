module github.com/soupdiver/pulumi-lxd/provider

go 1.16

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0

replace github.com/hashicorp/terraform-plugin-test v1.2.0 => github.com/hashicorp/terraform-plugin-test v1.4.3

require (
	github.com/hashicorp/terraform-plugin-sdk v1.16.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/pkg/v3 v3.2.1 // indirect
	github.com/pulumi/pulumi/sdk/v3 v3.2.1
	github.com/terraform-lxd/terraform-provider-lxd v1.5.0
)
