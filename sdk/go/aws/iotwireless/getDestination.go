// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotwireless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Destination's resource schema demonstrating some basic constructs and validation rules.
func LookupDestination(ctx *pulumi.Context, args *LookupDestinationArgs, opts ...pulumi.InvokeOption) (*LookupDestinationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationResult
	err := ctx.Invoke("aws-native:iotwireless:getDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDestinationArgs struct {
	// Unique name of destination
	Name string `pulumi:"name"`
}

type LookupDestinationResult struct {
	// Destination arn. Returned after successful create.
	Arn *string `pulumi:"arn"`
	// Destination description
	Description *string `pulumi:"description"`
	// Destination expression
	Expression *string `pulumi:"expression"`
	// Must be RuleName
	ExpressionType *DestinationExpressionType `pulumi:"expressionType"`
	// AWS role ARN that grants access
	RoleArn *string `pulumi:"roleArn"`
	// A list of key-value pairs that contain metadata for the destination.
	Tags []DestinationTag `pulumi:"tags"`
}

func LookupDestinationOutput(ctx *pulumi.Context, args LookupDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationResult, error) {
			args := v.(LookupDestinationArgs)
			r, err := LookupDestination(ctx, &args, opts...)
			var s LookupDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationResultOutput)
}

type LookupDestinationOutputArgs struct {
	// Unique name of destination
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationArgs)(nil)).Elem()
}

type LookupDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationResult)(nil)).Elem()
}

func (o LookupDestinationResultOutput) ToLookupDestinationResultOutput() LookupDestinationResultOutput {
	return o
}

func (o LookupDestinationResultOutput) ToLookupDestinationResultOutputWithContext(ctx context.Context) LookupDestinationResultOutput {
	return o
}

func (o LookupDestinationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDestinationResult] {
	return pulumix.Output[LookupDestinationResult]{
		OutputState: o.OutputState,
	}
}

// Destination arn. Returned after successful create.
func (o LookupDestinationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Destination description
func (o LookupDestinationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Destination expression
func (o LookupDestinationResultOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *string { return v.Expression }).(pulumi.StringPtrOutput)
}

// Must be RuleName
func (o LookupDestinationResultOutput) ExpressionType() DestinationExpressionTypePtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *DestinationExpressionType { return v.ExpressionType }).(DestinationExpressionTypePtrOutput)
}

// AWS role ARN that grants access
func (o LookupDestinationResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the destination.
func (o LookupDestinationResultOutput) Tags() DestinationTagArrayOutput {
	return o.ApplyT(func(v LookupDestinationResult) []DestinationTag { return v.Tags }).(DestinationTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationResultOutput{})
}
