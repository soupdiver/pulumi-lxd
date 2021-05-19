// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LXDRemote struct {
	Address  *string `pulumi:"address"`
	Default  *bool   `pulumi:"default"`
	Name     string  `pulumi:"name"`
	Password *string `pulumi:"password"`
	Port     *string `pulumi:"port"`
	Scheme   *string `pulumi:"scheme"`
}

// LXDRemoteInput is an input type that accepts LXDRemoteArgs and LXDRemoteOutput values.
// You can construct a concrete instance of `LXDRemoteInput` via:
//
//          LXDRemoteArgs{...}
type LXDRemoteInput interface {
	pulumi.Input

	ToLXDRemoteOutput() LXDRemoteOutput
	ToLXDRemoteOutputWithContext(context.Context) LXDRemoteOutput
}

type LXDRemoteArgs struct {
	Address  pulumi.StringPtrInput `pulumi:"address"`
	Default  pulumi.BoolPtrInput   `pulumi:"default"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Port     pulumi.StringPtrInput `pulumi:"port"`
	Scheme   pulumi.StringPtrInput `pulumi:"scheme"`
}

func (LXDRemoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LXDRemote)(nil)).Elem()
}

func (i LXDRemoteArgs) ToLXDRemoteOutput() LXDRemoteOutput {
	return i.ToLXDRemoteOutputWithContext(context.Background())
}

func (i LXDRemoteArgs) ToLXDRemoteOutputWithContext(ctx context.Context) LXDRemoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXDRemoteOutput)
}

// LXDRemoteArrayInput is an input type that accepts LXDRemoteArray and LXDRemoteArrayOutput values.
// You can construct a concrete instance of `LXDRemoteArrayInput` via:
//
//          LXDRemoteArray{ LXDRemoteArgs{...} }
type LXDRemoteArrayInput interface {
	pulumi.Input

	ToLXDRemoteArrayOutput() LXDRemoteArrayOutput
	ToLXDRemoteArrayOutputWithContext(context.Context) LXDRemoteArrayOutput
}

type LXDRemoteArray []LXDRemoteInput

func (LXDRemoteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LXDRemote)(nil)).Elem()
}

func (i LXDRemoteArray) ToLXDRemoteArrayOutput() LXDRemoteArrayOutput {
	return i.ToLXDRemoteArrayOutputWithContext(context.Background())
}

func (i LXDRemoteArray) ToLXDRemoteArrayOutputWithContext(ctx context.Context) LXDRemoteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXDRemoteArrayOutput)
}

type LXDRemoteOutput struct{ *pulumi.OutputState }

func (LXDRemoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LXDRemote)(nil)).Elem()
}

func (o LXDRemoteOutput) ToLXDRemoteOutput() LXDRemoteOutput {
	return o
}

func (o LXDRemoteOutput) ToLXDRemoteOutputWithContext(ctx context.Context) LXDRemoteOutput {
	return o
}

func (o LXDRemoteOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LXDRemote) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o LXDRemoteOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LXDRemote) *bool { return v.Default }).(pulumi.BoolPtrOutput)
}

func (o LXDRemoteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LXDRemote) string { return v.Name }).(pulumi.StringOutput)
}

func (o LXDRemoteOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LXDRemote) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LXDRemoteOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LXDRemote) *string { return v.Port }).(pulumi.StringPtrOutput)
}

func (o LXDRemoteOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LXDRemote) *string { return v.Scheme }).(pulumi.StringPtrOutput)
}

type LXDRemoteArrayOutput struct{ *pulumi.OutputState }

func (LXDRemoteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LXDRemote)(nil)).Elem()
}

func (o LXDRemoteArrayOutput) ToLXDRemoteArrayOutput() LXDRemoteArrayOutput {
	return o
}

func (o LXDRemoteArrayOutput) ToLXDRemoteArrayOutputWithContext(ctx context.Context) LXDRemoteArrayOutput {
	return o
}

func (o LXDRemoteArrayOutput) Index(i pulumi.IntInput) LXDRemoteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LXDRemote {
		return vs[0].([]LXDRemote)[vs[1].(int)]
	}).(LXDRemoteOutput)
}

func init() {
	pulumi.RegisterOutputType(LXDRemoteOutput{})
	pulumi.RegisterOutputType(LXDRemoteArrayOutput{})
}
