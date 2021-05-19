// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lxd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerDevice struct {
	Name       string                 `pulumi:"name"`
	Properties map[string]interface{} `pulumi:"properties"`
	Type       string                 `pulumi:"type"`
}

// ContainerDeviceInput is an input type that accepts ContainerDeviceArgs and ContainerDeviceOutput values.
// You can construct a concrete instance of `ContainerDeviceInput` via:
//
//          ContainerDeviceArgs{...}
type ContainerDeviceInput interface {
	pulumi.Input

	ToContainerDeviceOutput() ContainerDeviceOutput
	ToContainerDeviceOutputWithContext(context.Context) ContainerDeviceOutput
}

type ContainerDeviceArgs struct {
	Name       pulumi.StringInput `pulumi:"name"`
	Properties pulumi.MapInput    `pulumi:"properties"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ContainerDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerDevice)(nil)).Elem()
}

func (i ContainerDeviceArgs) ToContainerDeviceOutput() ContainerDeviceOutput {
	return i.ToContainerDeviceOutputWithContext(context.Background())
}

func (i ContainerDeviceArgs) ToContainerDeviceOutputWithContext(ctx context.Context) ContainerDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerDeviceOutput)
}

// ContainerDeviceArrayInput is an input type that accepts ContainerDeviceArray and ContainerDeviceArrayOutput values.
// You can construct a concrete instance of `ContainerDeviceArrayInput` via:
//
//          ContainerDeviceArray{ ContainerDeviceArgs{...} }
type ContainerDeviceArrayInput interface {
	pulumi.Input

	ToContainerDeviceArrayOutput() ContainerDeviceArrayOutput
	ToContainerDeviceArrayOutputWithContext(context.Context) ContainerDeviceArrayOutput
}

type ContainerDeviceArray []ContainerDeviceInput

func (ContainerDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerDevice)(nil)).Elem()
}

func (i ContainerDeviceArray) ToContainerDeviceArrayOutput() ContainerDeviceArrayOutput {
	return i.ToContainerDeviceArrayOutputWithContext(context.Background())
}

func (i ContainerDeviceArray) ToContainerDeviceArrayOutputWithContext(ctx context.Context) ContainerDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerDeviceArrayOutput)
}

type ContainerDeviceOutput struct{ *pulumi.OutputState }

func (ContainerDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerDevice)(nil)).Elem()
}

func (o ContainerDeviceOutput) ToContainerDeviceOutput() ContainerDeviceOutput {
	return o
}

func (o ContainerDeviceOutput) ToContainerDeviceOutputWithContext(ctx context.Context) ContainerDeviceOutput {
	return o
}

func (o ContainerDeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerDevice) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerDeviceOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v ContainerDevice) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

func (o ContainerDeviceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerDevice) string { return v.Type }).(pulumi.StringOutput)
}

type ContainerDeviceArrayOutput struct{ *pulumi.OutputState }

func (ContainerDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerDevice)(nil)).Elem()
}

func (o ContainerDeviceArrayOutput) ToContainerDeviceArrayOutput() ContainerDeviceArrayOutput {
	return o
}

func (o ContainerDeviceArrayOutput) ToContainerDeviceArrayOutputWithContext(ctx context.Context) ContainerDeviceArrayOutput {
	return o
}

func (o ContainerDeviceArrayOutput) Index(i pulumi.IntInput) ContainerDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerDevice {
		return vs[0].([]ContainerDevice)[vs[1].(int)]
	}).(ContainerDeviceOutput)
}

type ContainerFileType struct {
	Content           *string `pulumi:"content"`
	CreateDirectories *bool   `pulumi:"createDirectories"`
	Gid               *int    `pulumi:"gid"`
	Mode              *string `pulumi:"mode"`
	Source            *string `pulumi:"source"`
	TargetFile        string  `pulumi:"targetFile"`
	Uid               *int    `pulumi:"uid"`
}

// ContainerFileTypeInput is an input type that accepts ContainerFileTypeArgs and ContainerFileTypeOutput values.
// You can construct a concrete instance of `ContainerFileTypeInput` via:
//
//          ContainerFileTypeArgs{...}
type ContainerFileTypeInput interface {
	pulumi.Input

	ToContainerFileTypeOutput() ContainerFileTypeOutput
	ToContainerFileTypeOutputWithContext(context.Context) ContainerFileTypeOutput
}

type ContainerFileTypeArgs struct {
	Content           pulumi.StringPtrInput `pulumi:"content"`
	CreateDirectories pulumi.BoolPtrInput   `pulumi:"createDirectories"`
	Gid               pulumi.IntPtrInput    `pulumi:"gid"`
	Mode              pulumi.StringPtrInput `pulumi:"mode"`
	Source            pulumi.StringPtrInput `pulumi:"source"`
	TargetFile        pulumi.StringInput    `pulumi:"targetFile"`
	Uid               pulumi.IntPtrInput    `pulumi:"uid"`
}

func (ContainerFileTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerFileType)(nil)).Elem()
}

func (i ContainerFileTypeArgs) ToContainerFileTypeOutput() ContainerFileTypeOutput {
	return i.ToContainerFileTypeOutputWithContext(context.Background())
}

func (i ContainerFileTypeArgs) ToContainerFileTypeOutputWithContext(ctx context.Context) ContainerFileTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerFileTypeOutput)
}

// ContainerFileTypeArrayInput is an input type that accepts ContainerFileTypeArray and ContainerFileTypeArrayOutput values.
// You can construct a concrete instance of `ContainerFileTypeArrayInput` via:
//
//          ContainerFileTypeArray{ ContainerFileTypeArgs{...} }
type ContainerFileTypeArrayInput interface {
	pulumi.Input

	ToContainerFileTypeArrayOutput() ContainerFileTypeArrayOutput
	ToContainerFileTypeArrayOutputWithContext(context.Context) ContainerFileTypeArrayOutput
}

type ContainerFileTypeArray []ContainerFileTypeInput

func (ContainerFileTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerFileType)(nil)).Elem()
}

func (i ContainerFileTypeArray) ToContainerFileTypeArrayOutput() ContainerFileTypeArrayOutput {
	return i.ToContainerFileTypeArrayOutputWithContext(context.Background())
}

func (i ContainerFileTypeArray) ToContainerFileTypeArrayOutputWithContext(ctx context.Context) ContainerFileTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerFileTypeArrayOutput)
}

type ContainerFileTypeOutput struct{ *pulumi.OutputState }

func (ContainerFileTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerFileType)(nil)).Elem()
}

func (o ContainerFileTypeOutput) ToContainerFileTypeOutput() ContainerFileTypeOutput {
	return o
}

func (o ContainerFileTypeOutput) ToContainerFileTypeOutputWithContext(ctx context.Context) ContainerFileTypeOutput {
	return o
}

func (o ContainerFileTypeOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerFileType) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o ContainerFileTypeOutput) CreateDirectories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ContainerFileType) *bool { return v.CreateDirectories }).(pulumi.BoolPtrOutput)
}

func (o ContainerFileTypeOutput) Gid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerFileType) *int { return v.Gid }).(pulumi.IntPtrOutput)
}

func (o ContainerFileTypeOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerFileType) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o ContainerFileTypeOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerFileType) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o ContainerFileTypeOutput) TargetFile() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerFileType) string { return v.TargetFile }).(pulumi.StringOutput)
}

func (o ContainerFileTypeOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerFileType) *int { return v.Uid }).(pulumi.IntPtrOutput)
}

type ContainerFileTypeArrayOutput struct{ *pulumi.OutputState }

func (ContainerFileTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerFileType)(nil)).Elem()
}

func (o ContainerFileTypeArrayOutput) ToContainerFileTypeArrayOutput() ContainerFileTypeArrayOutput {
	return o
}

func (o ContainerFileTypeArrayOutput) ToContainerFileTypeArrayOutputWithContext(ctx context.Context) ContainerFileTypeArrayOutput {
	return o
}

func (o ContainerFileTypeArrayOutput) Index(i pulumi.IntInput) ContainerFileTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerFileType {
		return vs[0].([]ContainerFileType)[vs[1].(int)]
	}).(ContainerFileTypeOutput)
}

type ProfileDevice struct {
	Name       string                 `pulumi:"name"`
	Properties map[string]interface{} `pulumi:"properties"`
	Type       string                 `pulumi:"type"`
}

// ProfileDeviceInput is an input type that accepts ProfileDeviceArgs and ProfileDeviceOutput values.
// You can construct a concrete instance of `ProfileDeviceInput` via:
//
//          ProfileDeviceArgs{...}
type ProfileDeviceInput interface {
	pulumi.Input

	ToProfileDeviceOutput() ProfileDeviceOutput
	ToProfileDeviceOutputWithContext(context.Context) ProfileDeviceOutput
}

type ProfileDeviceArgs struct {
	Name       pulumi.StringInput `pulumi:"name"`
	Properties pulumi.MapInput    `pulumi:"properties"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ProfileDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileDevice)(nil)).Elem()
}

func (i ProfileDeviceArgs) ToProfileDeviceOutput() ProfileDeviceOutput {
	return i.ToProfileDeviceOutputWithContext(context.Background())
}

func (i ProfileDeviceArgs) ToProfileDeviceOutputWithContext(ctx context.Context) ProfileDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileDeviceOutput)
}

// ProfileDeviceArrayInput is an input type that accepts ProfileDeviceArray and ProfileDeviceArrayOutput values.
// You can construct a concrete instance of `ProfileDeviceArrayInput` via:
//
//          ProfileDeviceArray{ ProfileDeviceArgs{...} }
type ProfileDeviceArrayInput interface {
	pulumi.Input

	ToProfileDeviceArrayOutput() ProfileDeviceArrayOutput
	ToProfileDeviceArrayOutputWithContext(context.Context) ProfileDeviceArrayOutput
}

type ProfileDeviceArray []ProfileDeviceInput

func (ProfileDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileDevice)(nil)).Elem()
}

func (i ProfileDeviceArray) ToProfileDeviceArrayOutput() ProfileDeviceArrayOutput {
	return i.ToProfileDeviceArrayOutputWithContext(context.Background())
}

func (i ProfileDeviceArray) ToProfileDeviceArrayOutputWithContext(ctx context.Context) ProfileDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileDeviceArrayOutput)
}

type ProfileDeviceOutput struct{ *pulumi.OutputState }

func (ProfileDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileDevice)(nil)).Elem()
}

func (o ProfileDeviceOutput) ToProfileDeviceOutput() ProfileDeviceOutput {
	return o
}

func (o ProfileDeviceOutput) ToProfileDeviceOutputWithContext(ctx context.Context) ProfileDeviceOutput {
	return o
}

func (o ProfileDeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProfileDevice) string { return v.Name }).(pulumi.StringOutput)
}

func (o ProfileDeviceOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v ProfileDevice) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

func (o ProfileDeviceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ProfileDevice) string { return v.Type }).(pulumi.StringOutput)
}

type ProfileDeviceArrayOutput struct{ *pulumi.OutputState }

func (ProfileDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileDevice)(nil)).Elem()
}

func (o ProfileDeviceArrayOutput) ToProfileDeviceArrayOutput() ProfileDeviceArrayOutput {
	return o
}

func (o ProfileDeviceArrayOutput) ToProfileDeviceArrayOutputWithContext(ctx context.Context) ProfileDeviceArrayOutput {
	return o
}

func (o ProfileDeviceArrayOutput) Index(i pulumi.IntInput) ProfileDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileDevice {
		return vs[0].([]ProfileDevice)[vs[1].(int)]
	}).(ProfileDeviceOutput)
}

type ProviderLXDRemote struct {
	Address  *string `pulumi:"address"`
	Default  *bool   `pulumi:"default"`
	Name     string  `pulumi:"name"`
	Password *string `pulumi:"password"`
	Port     *string `pulumi:"port"`
	Scheme   *string `pulumi:"scheme"`
}

// ProviderLXDRemoteInput is an input type that accepts ProviderLXDRemoteArgs and ProviderLXDRemoteOutput values.
// You can construct a concrete instance of `ProviderLXDRemoteInput` via:
//
//          ProviderLXDRemoteArgs{...}
type ProviderLXDRemoteInput interface {
	pulumi.Input

	ToProviderLXDRemoteOutput() ProviderLXDRemoteOutput
	ToProviderLXDRemoteOutputWithContext(context.Context) ProviderLXDRemoteOutput
}

type ProviderLXDRemoteArgs struct {
	Address  pulumi.StringPtrInput `pulumi:"address"`
	Default  pulumi.BoolPtrInput   `pulumi:"default"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Port     pulumi.StringPtrInput `pulumi:"port"`
	Scheme   pulumi.StringPtrInput `pulumi:"scheme"`
}

func (ProviderLXDRemoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderLXDRemote)(nil)).Elem()
}

func (i ProviderLXDRemoteArgs) ToProviderLXDRemoteOutput() ProviderLXDRemoteOutput {
	return i.ToProviderLXDRemoteOutputWithContext(context.Background())
}

func (i ProviderLXDRemoteArgs) ToProviderLXDRemoteOutputWithContext(ctx context.Context) ProviderLXDRemoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderLXDRemoteOutput)
}

// ProviderLXDRemoteArrayInput is an input type that accepts ProviderLXDRemoteArray and ProviderLXDRemoteArrayOutput values.
// You can construct a concrete instance of `ProviderLXDRemoteArrayInput` via:
//
//          ProviderLXDRemoteArray{ ProviderLXDRemoteArgs{...} }
type ProviderLXDRemoteArrayInput interface {
	pulumi.Input

	ToProviderLXDRemoteArrayOutput() ProviderLXDRemoteArrayOutput
	ToProviderLXDRemoteArrayOutputWithContext(context.Context) ProviderLXDRemoteArrayOutput
}

type ProviderLXDRemoteArray []ProviderLXDRemoteInput

func (ProviderLXDRemoteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderLXDRemote)(nil)).Elem()
}

func (i ProviderLXDRemoteArray) ToProviderLXDRemoteArrayOutput() ProviderLXDRemoteArrayOutput {
	return i.ToProviderLXDRemoteArrayOutputWithContext(context.Background())
}

func (i ProviderLXDRemoteArray) ToProviderLXDRemoteArrayOutputWithContext(ctx context.Context) ProviderLXDRemoteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderLXDRemoteArrayOutput)
}

type ProviderLXDRemoteOutput struct{ *pulumi.OutputState }

func (ProviderLXDRemoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderLXDRemote)(nil)).Elem()
}

func (o ProviderLXDRemoteOutput) ToProviderLXDRemoteOutput() ProviderLXDRemoteOutput {
	return o
}

func (o ProviderLXDRemoteOutput) ToProviderLXDRemoteOutputWithContext(ctx context.Context) ProviderLXDRemoteOutput {
	return o
}

func (o ProviderLXDRemoteOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderLXDRemote) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ProviderLXDRemoteOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderLXDRemote) *bool { return v.Default }).(pulumi.BoolPtrOutput)
}

func (o ProviderLXDRemoteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderLXDRemote) string { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderLXDRemoteOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderLXDRemote) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ProviderLXDRemoteOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderLXDRemote) *string { return v.Port }).(pulumi.StringPtrOutput)
}

func (o ProviderLXDRemoteOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderLXDRemote) *string { return v.Scheme }).(pulumi.StringPtrOutput)
}

type ProviderLXDRemoteArrayOutput struct{ *pulumi.OutputState }

func (ProviderLXDRemoteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderLXDRemote)(nil)).Elem()
}

func (o ProviderLXDRemoteArrayOutput) ToProviderLXDRemoteArrayOutput() ProviderLXDRemoteArrayOutput {
	return o
}

func (o ProviderLXDRemoteArrayOutput) ToProviderLXDRemoteArrayOutputWithContext(ctx context.Context) ProviderLXDRemoteArrayOutput {
	return o
}

func (o ProviderLXDRemoteArrayOutput) Index(i pulumi.IntInput) ProviderLXDRemoteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderLXDRemote {
		return vs[0].([]ProviderLXDRemote)[vs[1].(int)]
	}).(ProviderLXDRemoteOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerDeviceOutput{})
	pulumi.RegisterOutputType(ContainerDeviceArrayOutput{})
	pulumi.RegisterOutputType(ContainerFileTypeOutput{})
	pulumi.RegisterOutputType(ContainerFileTypeArrayOutput{})
	pulumi.RegisterOutputType(ProfileDeviceOutput{})
	pulumi.RegisterOutputType(ProfileDeviceArrayOutput{})
	pulumi.RegisterOutputType(ProviderLXDRemoteOutput{})
	pulumi.RegisterOutputType(ProviderLXDRemoteArrayOutput{})
}