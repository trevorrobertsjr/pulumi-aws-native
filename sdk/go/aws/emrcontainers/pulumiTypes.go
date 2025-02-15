// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emrcontainers

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type VirtualClusterContainerInfo struct {
	EksInfo VirtualClusterEksInfo `pulumi:"eksInfo"`
}

// VirtualClusterContainerInfoInput is an input type that accepts VirtualClusterContainerInfoArgs and VirtualClusterContainerInfoOutput values.
// You can construct a concrete instance of `VirtualClusterContainerInfoInput` via:
//
//	VirtualClusterContainerInfoArgs{...}
type VirtualClusterContainerInfoInput interface {
	pulumi.Input

	ToVirtualClusterContainerInfoOutput() VirtualClusterContainerInfoOutput
	ToVirtualClusterContainerInfoOutputWithContext(context.Context) VirtualClusterContainerInfoOutput
}

type VirtualClusterContainerInfoArgs struct {
	EksInfo VirtualClusterEksInfoInput `pulumi:"eksInfo"`
}

func (VirtualClusterContainerInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualClusterContainerInfo)(nil)).Elem()
}

func (i VirtualClusterContainerInfoArgs) ToVirtualClusterContainerInfoOutput() VirtualClusterContainerInfoOutput {
	return i.ToVirtualClusterContainerInfoOutputWithContext(context.Background())
}

func (i VirtualClusterContainerInfoArgs) ToVirtualClusterContainerInfoOutputWithContext(ctx context.Context) VirtualClusterContainerInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterContainerInfoOutput)
}

func (i VirtualClusterContainerInfoArgs) ToOutput(ctx context.Context) pulumix.Output[VirtualClusterContainerInfo] {
	return pulumix.Output[VirtualClusterContainerInfo]{
		OutputState: i.ToVirtualClusterContainerInfoOutputWithContext(ctx).OutputState,
	}
}

type VirtualClusterContainerInfoOutput struct{ *pulumi.OutputState }

func (VirtualClusterContainerInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualClusterContainerInfo)(nil)).Elem()
}

func (o VirtualClusterContainerInfoOutput) ToVirtualClusterContainerInfoOutput() VirtualClusterContainerInfoOutput {
	return o
}

func (o VirtualClusterContainerInfoOutput) ToVirtualClusterContainerInfoOutputWithContext(ctx context.Context) VirtualClusterContainerInfoOutput {
	return o
}

func (o VirtualClusterContainerInfoOutput) ToOutput(ctx context.Context) pulumix.Output[VirtualClusterContainerInfo] {
	return pulumix.Output[VirtualClusterContainerInfo]{
		OutputState: o.OutputState,
	}
}

func (o VirtualClusterContainerInfoOutput) EksInfo() VirtualClusterEksInfoOutput {
	return o.ApplyT(func(v VirtualClusterContainerInfo) VirtualClusterEksInfo { return v.EksInfo }).(VirtualClusterEksInfoOutput)
}

type VirtualClusterContainerProvider struct {
	// The ID of the container cluster
	Id   string                      `pulumi:"id"`
	Info VirtualClusterContainerInfo `pulumi:"info"`
	// The type of the container provider
	Type string `pulumi:"type"`
}

// VirtualClusterContainerProviderInput is an input type that accepts VirtualClusterContainerProviderArgs and VirtualClusterContainerProviderOutput values.
// You can construct a concrete instance of `VirtualClusterContainerProviderInput` via:
//
//	VirtualClusterContainerProviderArgs{...}
type VirtualClusterContainerProviderInput interface {
	pulumi.Input

	ToVirtualClusterContainerProviderOutput() VirtualClusterContainerProviderOutput
	ToVirtualClusterContainerProviderOutputWithContext(context.Context) VirtualClusterContainerProviderOutput
}

type VirtualClusterContainerProviderArgs struct {
	// The ID of the container cluster
	Id   pulumi.StringInput               `pulumi:"id"`
	Info VirtualClusterContainerInfoInput `pulumi:"info"`
	// The type of the container provider
	Type pulumi.StringInput `pulumi:"type"`
}

func (VirtualClusterContainerProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualClusterContainerProvider)(nil)).Elem()
}

func (i VirtualClusterContainerProviderArgs) ToVirtualClusterContainerProviderOutput() VirtualClusterContainerProviderOutput {
	return i.ToVirtualClusterContainerProviderOutputWithContext(context.Background())
}

func (i VirtualClusterContainerProviderArgs) ToVirtualClusterContainerProviderOutputWithContext(ctx context.Context) VirtualClusterContainerProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterContainerProviderOutput)
}

func (i VirtualClusterContainerProviderArgs) ToOutput(ctx context.Context) pulumix.Output[VirtualClusterContainerProvider] {
	return pulumix.Output[VirtualClusterContainerProvider]{
		OutputState: i.ToVirtualClusterContainerProviderOutputWithContext(ctx).OutputState,
	}
}

type VirtualClusterContainerProviderOutput struct{ *pulumi.OutputState }

func (VirtualClusterContainerProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualClusterContainerProvider)(nil)).Elem()
}

func (o VirtualClusterContainerProviderOutput) ToVirtualClusterContainerProviderOutput() VirtualClusterContainerProviderOutput {
	return o
}

func (o VirtualClusterContainerProviderOutput) ToVirtualClusterContainerProviderOutputWithContext(ctx context.Context) VirtualClusterContainerProviderOutput {
	return o
}

func (o VirtualClusterContainerProviderOutput) ToOutput(ctx context.Context) pulumix.Output[VirtualClusterContainerProvider] {
	return pulumix.Output[VirtualClusterContainerProvider]{
		OutputState: o.OutputState,
	}
}

// The ID of the container cluster
func (o VirtualClusterContainerProviderOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualClusterContainerProvider) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualClusterContainerProviderOutput) Info() VirtualClusterContainerInfoOutput {
	return o.ApplyT(func(v VirtualClusterContainerProvider) VirtualClusterContainerInfo { return v.Info }).(VirtualClusterContainerInfoOutput)
}

// The type of the container provider
func (o VirtualClusterContainerProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualClusterContainerProvider) string { return v.Type }).(pulumi.StringOutput)
}

type VirtualClusterEksInfo struct {
	Namespace string `pulumi:"namespace"`
}

// VirtualClusterEksInfoInput is an input type that accepts VirtualClusterEksInfoArgs and VirtualClusterEksInfoOutput values.
// You can construct a concrete instance of `VirtualClusterEksInfoInput` via:
//
//	VirtualClusterEksInfoArgs{...}
type VirtualClusterEksInfoInput interface {
	pulumi.Input

	ToVirtualClusterEksInfoOutput() VirtualClusterEksInfoOutput
	ToVirtualClusterEksInfoOutputWithContext(context.Context) VirtualClusterEksInfoOutput
}

type VirtualClusterEksInfoArgs struct {
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (VirtualClusterEksInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualClusterEksInfo)(nil)).Elem()
}

func (i VirtualClusterEksInfoArgs) ToVirtualClusterEksInfoOutput() VirtualClusterEksInfoOutput {
	return i.ToVirtualClusterEksInfoOutputWithContext(context.Background())
}

func (i VirtualClusterEksInfoArgs) ToVirtualClusterEksInfoOutputWithContext(ctx context.Context) VirtualClusterEksInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterEksInfoOutput)
}

func (i VirtualClusterEksInfoArgs) ToOutput(ctx context.Context) pulumix.Output[VirtualClusterEksInfo] {
	return pulumix.Output[VirtualClusterEksInfo]{
		OutputState: i.ToVirtualClusterEksInfoOutputWithContext(ctx).OutputState,
	}
}

type VirtualClusterEksInfoOutput struct{ *pulumi.OutputState }

func (VirtualClusterEksInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualClusterEksInfo)(nil)).Elem()
}

func (o VirtualClusterEksInfoOutput) ToVirtualClusterEksInfoOutput() VirtualClusterEksInfoOutput {
	return o
}

func (o VirtualClusterEksInfoOutput) ToVirtualClusterEksInfoOutputWithContext(ctx context.Context) VirtualClusterEksInfoOutput {
	return o
}

func (o VirtualClusterEksInfoOutput) ToOutput(ctx context.Context) pulumix.Output[VirtualClusterEksInfo] {
	return pulumix.Output[VirtualClusterEksInfo]{
		OutputState: o.OutputState,
	}
}

func (o VirtualClusterEksInfoOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualClusterEksInfo) string { return v.Namespace }).(pulumi.StringOutput)
}

// An arbitrary set of tags (key-value pairs) for this virtual cluster.
type VirtualClusterTag struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// VirtualClusterTagInput is an input type that accepts VirtualClusterTagArgs and VirtualClusterTagOutput values.
// You can construct a concrete instance of `VirtualClusterTagInput` via:
//
//	VirtualClusterTagArgs{...}
type VirtualClusterTagInput interface {
	pulumi.Input

	ToVirtualClusterTagOutput() VirtualClusterTagOutput
	ToVirtualClusterTagOutputWithContext(context.Context) VirtualClusterTagOutput
}

// An arbitrary set of tags (key-value pairs) for this virtual cluster.
type VirtualClusterTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (VirtualClusterTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualClusterTag)(nil)).Elem()
}

func (i VirtualClusterTagArgs) ToVirtualClusterTagOutput() VirtualClusterTagOutput {
	return i.ToVirtualClusterTagOutputWithContext(context.Background())
}

func (i VirtualClusterTagArgs) ToVirtualClusterTagOutputWithContext(ctx context.Context) VirtualClusterTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterTagOutput)
}

func (i VirtualClusterTagArgs) ToOutput(ctx context.Context) pulumix.Output[VirtualClusterTag] {
	return pulumix.Output[VirtualClusterTag]{
		OutputState: i.ToVirtualClusterTagOutputWithContext(ctx).OutputState,
	}
}

// VirtualClusterTagArrayInput is an input type that accepts VirtualClusterTagArray and VirtualClusterTagArrayOutput values.
// You can construct a concrete instance of `VirtualClusterTagArrayInput` via:
//
//	VirtualClusterTagArray{ VirtualClusterTagArgs{...} }
type VirtualClusterTagArrayInput interface {
	pulumi.Input

	ToVirtualClusterTagArrayOutput() VirtualClusterTagArrayOutput
	ToVirtualClusterTagArrayOutputWithContext(context.Context) VirtualClusterTagArrayOutput
}

type VirtualClusterTagArray []VirtualClusterTagInput

func (VirtualClusterTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualClusterTag)(nil)).Elem()
}

func (i VirtualClusterTagArray) ToVirtualClusterTagArrayOutput() VirtualClusterTagArrayOutput {
	return i.ToVirtualClusterTagArrayOutputWithContext(context.Background())
}

func (i VirtualClusterTagArray) ToVirtualClusterTagArrayOutputWithContext(ctx context.Context) VirtualClusterTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterTagArrayOutput)
}

func (i VirtualClusterTagArray) ToOutput(ctx context.Context) pulumix.Output[[]VirtualClusterTag] {
	return pulumix.Output[[]VirtualClusterTag]{
		OutputState: i.ToVirtualClusterTagArrayOutputWithContext(ctx).OutputState,
	}
}

// An arbitrary set of tags (key-value pairs) for this virtual cluster.
type VirtualClusterTagOutput struct{ *pulumi.OutputState }

func (VirtualClusterTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualClusterTag)(nil)).Elem()
}

func (o VirtualClusterTagOutput) ToVirtualClusterTagOutput() VirtualClusterTagOutput {
	return o
}

func (o VirtualClusterTagOutput) ToVirtualClusterTagOutputWithContext(ctx context.Context) VirtualClusterTagOutput {
	return o
}

func (o VirtualClusterTagOutput) ToOutput(ctx context.Context) pulumix.Output[VirtualClusterTag] {
	return pulumix.Output[VirtualClusterTag]{
		OutputState: o.OutputState,
	}
}

// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o VirtualClusterTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualClusterTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o VirtualClusterTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualClusterTag) string { return v.Value }).(pulumi.StringOutput)
}

type VirtualClusterTagArrayOutput struct{ *pulumi.OutputState }

func (VirtualClusterTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualClusterTag)(nil)).Elem()
}

func (o VirtualClusterTagArrayOutput) ToVirtualClusterTagArrayOutput() VirtualClusterTagArrayOutput {
	return o
}

func (o VirtualClusterTagArrayOutput) ToVirtualClusterTagArrayOutputWithContext(ctx context.Context) VirtualClusterTagArrayOutput {
	return o
}

func (o VirtualClusterTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]VirtualClusterTag] {
	return pulumix.Output[[]VirtualClusterTag]{
		OutputState: o.OutputState,
	}
}

func (o VirtualClusterTagArrayOutput) Index(i pulumi.IntInput) VirtualClusterTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualClusterTag {
		return vs[0].([]VirtualClusterTag)[vs[1].(int)]
	}).(VirtualClusterTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterContainerInfoInput)(nil)).Elem(), VirtualClusterContainerInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterContainerProviderInput)(nil)).Elem(), VirtualClusterContainerProviderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterEksInfoInput)(nil)).Elem(), VirtualClusterEksInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterTagInput)(nil)).Elem(), VirtualClusterTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterTagArrayInput)(nil)).Elem(), VirtualClusterTagArray{})
	pulumi.RegisterOutputType(VirtualClusterContainerInfoOutput{})
	pulumi.RegisterOutputType(VirtualClusterContainerProviderOutput{})
	pulumi.RegisterOutputType(VirtualClusterEksInfoOutput{})
	pulumi.RegisterOutputType(VirtualClusterTagOutput{})
	pulumi.RegisterOutputType(VirtualClusterTagArrayOutput{})
}
