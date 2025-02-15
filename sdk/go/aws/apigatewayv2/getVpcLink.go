// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::ApiGatewayV2::VpcLink
func LookupVpcLink(ctx *pulumi.Context, args *LookupVpcLinkArgs, opts ...pulumi.InvokeOption) (*LookupVpcLinkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcLinkResult
	err := ctx.Invoke("aws-native:apigatewayv2:getVpcLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcLinkArgs struct {
	VpcLinkId string `pulumi:"vpcLinkId"`
}

type LookupVpcLinkResult struct {
	Name *string `pulumi:"name"`
	// This resource type use map for Tags, suggest to use List of Tag
	Tags      interface{} `pulumi:"tags"`
	VpcLinkId *string     `pulumi:"vpcLinkId"`
}

func LookupVpcLinkOutput(ctx *pulumi.Context, args LookupVpcLinkOutputArgs, opts ...pulumi.InvokeOption) LookupVpcLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcLinkResult, error) {
			args := v.(LookupVpcLinkArgs)
			r, err := LookupVpcLink(ctx, &args, opts...)
			var s LookupVpcLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcLinkResultOutput)
}

type LookupVpcLinkOutputArgs struct {
	VpcLinkId pulumi.StringInput `pulumi:"vpcLinkId"`
}

func (LookupVpcLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcLinkArgs)(nil)).Elem()
}

type LookupVpcLinkResultOutput struct{ *pulumi.OutputState }

func (LookupVpcLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcLinkResult)(nil)).Elem()
}

func (o LookupVpcLinkResultOutput) ToLookupVpcLinkResultOutput() LookupVpcLinkResultOutput {
	return o
}

func (o LookupVpcLinkResultOutput) ToLookupVpcLinkResultOutputWithContext(ctx context.Context) LookupVpcLinkResultOutput {
	return o
}

func (o LookupVpcLinkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVpcLinkResult] {
	return pulumix.Output[LookupVpcLinkResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupVpcLinkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcLinkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// This resource type use map for Tags, suggest to use List of Tag
func (o LookupVpcLinkResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVpcLinkResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupVpcLinkResultOutput) VpcLinkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcLinkResult) *string { return v.VpcLinkId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcLinkResultOutput{})
}
