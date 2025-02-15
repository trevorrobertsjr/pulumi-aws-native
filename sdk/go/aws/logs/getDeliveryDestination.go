// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This structure contains information about one delivery destination in your account.
//
// A delivery destination is an AWS resource that represents an AWS service that logs can be sent to CloudWatch Logs, Amazon S3, are supported as Kinesis Data Firehose delivery destinations.
func LookupDeliveryDestination(ctx *pulumi.Context, args *LookupDeliveryDestinationArgs, opts ...pulumi.InvokeOption) (*LookupDeliveryDestinationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeliveryDestinationResult
	err := ctx.Invoke("aws-native:logs:getDeliveryDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeliveryDestinationArgs struct {
	// The name of this delivery destination.
	Name string `pulumi:"name"`
}

type LookupDeliveryDestinationResult struct {
	// The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.
	Arn *string `pulumi:"arn"`
	// IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 51200
	DeliveryDestinationPolicy interface{} `pulumi:"deliveryDestinationPolicy"`
	// Displays whether this delivery destination is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.
	DeliveryDestinationType *string `pulumi:"deliveryDestinationType"`
	// The tags that have been assigned to this delivery destination.
	Tags []DeliveryDestinationTag `pulumi:"tags"`
}

func LookupDeliveryDestinationOutput(ctx *pulumi.Context, args LookupDeliveryDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupDeliveryDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeliveryDestinationResult, error) {
			args := v.(LookupDeliveryDestinationArgs)
			r, err := LookupDeliveryDestination(ctx, &args, opts...)
			var s LookupDeliveryDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeliveryDestinationResultOutput)
}

type LookupDeliveryDestinationOutputArgs struct {
	// The name of this delivery destination.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDeliveryDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeliveryDestinationArgs)(nil)).Elem()
}

type LookupDeliveryDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupDeliveryDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeliveryDestinationResult)(nil)).Elem()
}

func (o LookupDeliveryDestinationResultOutput) ToLookupDeliveryDestinationResultOutput() LookupDeliveryDestinationResultOutput {
	return o
}

func (o LookupDeliveryDestinationResultOutput) ToLookupDeliveryDestinationResultOutputWithContext(ctx context.Context) LookupDeliveryDestinationResultOutput {
	return o
}

func (o LookupDeliveryDestinationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDeliveryDestinationResult] {
	return pulumix.Output[LookupDeliveryDestinationResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.
func (o LookupDeliveryDestinationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeliveryDestinationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
//
// The policy must be in JSON string format.
//
// Length Constraints: Maximum length of 51200
func (o LookupDeliveryDestinationResultOutput) DeliveryDestinationPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDeliveryDestinationResult) interface{} { return v.DeliveryDestinationPolicy }).(pulumi.AnyOutput)
}

// Displays whether this delivery destination is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.
func (o LookupDeliveryDestinationResultOutput) DeliveryDestinationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeliveryDestinationResult) *string { return v.DeliveryDestinationType }).(pulumi.StringPtrOutput)
}

// The tags that have been assigned to this delivery destination.
func (o LookupDeliveryDestinationResultOutput) Tags() DeliveryDestinationTagArrayOutput {
	return o.ApplyT(func(v LookupDeliveryDestinationResult) []DeliveryDestinationTag { return v.Tags }).(DeliveryDestinationTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeliveryDestinationResultOutput{})
}
