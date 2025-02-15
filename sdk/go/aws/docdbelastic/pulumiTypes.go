// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docdbelastic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type ClusterTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// ClusterTagInput is an input type that accepts ClusterTagArgs and ClusterTagOutput values.
// You can construct a concrete instance of `ClusterTagInput` via:
//
//	ClusterTagArgs{...}
type ClusterTagInput interface {
	pulumi.Input

	ToClusterTagOutput() ClusterTagOutput
	ToClusterTagOutputWithContext(context.Context) ClusterTagOutput
}

type ClusterTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (ClusterTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterTag)(nil)).Elem()
}

func (i ClusterTagArgs) ToClusterTagOutput() ClusterTagOutput {
	return i.ToClusterTagOutputWithContext(context.Background())
}

func (i ClusterTagArgs) ToClusterTagOutputWithContext(ctx context.Context) ClusterTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTagOutput)
}

func (i ClusterTagArgs) ToOutput(ctx context.Context) pulumix.Output[ClusterTag] {
	return pulumix.Output[ClusterTag]{
		OutputState: i.ToClusterTagOutputWithContext(ctx).OutputState,
	}
}

// ClusterTagArrayInput is an input type that accepts ClusterTagArray and ClusterTagArrayOutput values.
// You can construct a concrete instance of `ClusterTagArrayInput` via:
//
//	ClusterTagArray{ ClusterTagArgs{...} }
type ClusterTagArrayInput interface {
	pulumi.Input

	ToClusterTagArrayOutput() ClusterTagArrayOutput
	ToClusterTagArrayOutputWithContext(context.Context) ClusterTagArrayOutput
}

type ClusterTagArray []ClusterTagInput

func (ClusterTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterTag)(nil)).Elem()
}

func (i ClusterTagArray) ToClusterTagArrayOutput() ClusterTagArrayOutput {
	return i.ToClusterTagArrayOutputWithContext(context.Background())
}

func (i ClusterTagArray) ToClusterTagArrayOutputWithContext(ctx context.Context) ClusterTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTagArrayOutput)
}

func (i ClusterTagArray) ToOutput(ctx context.Context) pulumix.Output[[]ClusterTag] {
	return pulumix.Output[[]ClusterTag]{
		OutputState: i.ToClusterTagArrayOutputWithContext(ctx).OutputState,
	}
}

type ClusterTagOutput struct{ *pulumi.OutputState }

func (ClusterTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterTag)(nil)).Elem()
}

func (o ClusterTagOutput) ToClusterTagOutput() ClusterTagOutput {
	return o
}

func (o ClusterTagOutput) ToClusterTagOutputWithContext(ctx context.Context) ClusterTagOutput {
	return o
}

func (o ClusterTagOutput) ToOutput(ctx context.Context) pulumix.Output[ClusterTag] {
	return pulumix.Output[ClusterTag]{
		OutputState: o.OutputState,
	}
}

func (o ClusterTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o ClusterTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterTag) string { return v.Value }).(pulumi.StringOutput)
}

type ClusterTagArrayOutput struct{ *pulumi.OutputState }

func (ClusterTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterTag)(nil)).Elem()
}

func (o ClusterTagArrayOutput) ToClusterTagArrayOutput() ClusterTagArrayOutput {
	return o
}

func (o ClusterTagArrayOutput) ToClusterTagArrayOutputWithContext(ctx context.Context) ClusterTagArrayOutput {
	return o
}

func (o ClusterTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ClusterTag] {
	return pulumix.Output[[]ClusterTag]{
		OutputState: o.OutputState,
	}
}

func (o ClusterTagArrayOutput) Index(i pulumi.IntInput) ClusterTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterTag {
		return vs[0].([]ClusterTag)[vs[1].(int)]
	}).(ClusterTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTagInput)(nil)).Elem(), ClusterTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTagArrayInput)(nil)).Elem(), ClusterTagArray{})
	pulumi.RegisterOutputType(ClusterTagOutput{})
	pulumi.RegisterOutputType(ClusterTagArrayOutput{})
}
