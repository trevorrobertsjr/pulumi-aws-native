// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::AppMesh::VirtualNode
func LookupVirtualNode(ctx *pulumi.Context, args *LookupVirtualNodeArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNodeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualNodeResult
	err := ctx.Invoke("aws-native:appmesh:getVirtualNode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNodeArgs struct {
	Id string `pulumi:"id"`
}

type LookupVirtualNodeResult struct {
	Arn           *string          `pulumi:"arn"`
	Id            *string          `pulumi:"id"`
	ResourceOwner *string          `pulumi:"resourceOwner"`
	Spec          *VirtualNodeSpec `pulumi:"spec"`
	Tags          []VirtualNodeTag `pulumi:"tags"`
	Uid           *string          `pulumi:"uid"`
}

func LookupVirtualNodeOutput(ctx *pulumi.Context, args LookupVirtualNodeOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNodeResult, error) {
			args := v.(LookupVirtualNodeArgs)
			r, err := LookupVirtualNode(ctx, &args, opts...)
			var s LookupVirtualNodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNodeResultOutput)
}

type LookupVirtualNodeOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVirtualNodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNodeArgs)(nil)).Elem()
}

type LookupVirtualNodeResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNodeResult)(nil)).Elem()
}

func (o LookupVirtualNodeResultOutput) ToLookupVirtualNodeResultOutput() LookupVirtualNodeResultOutput {
	return o
}

func (o LookupVirtualNodeResultOutput) ToLookupVirtualNodeResultOutputWithContext(ctx context.Context) LookupVirtualNodeResultOutput {
	return o
}

func (o LookupVirtualNodeResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualNodeResult] {
	return pulumix.Output[LookupVirtualNodeResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupVirtualNodeResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNodeResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNodeResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNodeResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNodeResultOutput) ResourceOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNodeResult) *string { return v.ResourceOwner }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNodeResultOutput) Spec() VirtualNodeSpecPtrOutput {
	return o.ApplyT(func(v LookupVirtualNodeResult) *VirtualNodeSpec { return v.Spec }).(VirtualNodeSpecPtrOutput)
}

func (o LookupVirtualNodeResultOutput) Tags() VirtualNodeTagArrayOutput {
	return o.ApplyT(func(v LookupVirtualNodeResult) []VirtualNodeTag { return v.Tags }).(VirtualNodeTagArrayOutput)
}

func (o LookupVirtualNodeResultOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNodeResult) *string { return v.Uid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNodeResultOutput{})
}
