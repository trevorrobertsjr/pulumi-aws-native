// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::SageMaker::Image
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupImageResult
	err := ctx.Invoke("aws-native:sagemaker:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImageArgs struct {
	ImageArn string `pulumi:"imageArn"`
}

type LookupImageResult struct {
	ImageArn         *string `pulumi:"imageArn"`
	ImageDescription *string `pulumi:"imageDescription"`
	ImageDisplayName *string `pulumi:"imageDisplayName"`
	ImageRoleArn     *string `pulumi:"imageRoleArn"`
	// An array of key-value pairs to apply to this resource.
	Tags []ImageTag `pulumi:"tags"`
}

func LookupImageOutput(ctx *pulumi.Context, args LookupImageOutputArgs, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageResult, error) {
			args := v.(LookupImageArgs)
			r, err := LookupImage(ctx, &args, opts...)
			var s LookupImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageResultOutput)
}

type LookupImageOutputArgs struct {
	ImageArn pulumi.StringInput `pulumi:"imageArn"`
}

func (LookupImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupImageResult] {
	return pulumix.Output[LookupImageResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupImageResultOutput) ImageArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.ImageArn }).(pulumi.StringPtrOutput)
}

func (o LookupImageResultOutput) ImageDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.ImageDescription }).(pulumi.StringPtrOutput)
}

func (o LookupImageResultOutput) ImageDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.ImageDisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupImageResultOutput) ImageRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.ImageRoleArn }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupImageResultOutput) Tags() ImageTagArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []ImageTag { return v.Tags }).(ImageTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
