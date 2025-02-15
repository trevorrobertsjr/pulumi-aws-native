// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose. The service network owner can use the access logs to audit the services in the network. The service network owner will only see access logs from clients and services that are associated with their service network. Access log entries represent traffic originated from VPCs associated with that network.
func LookupAccessLogSubscription(ctx *pulumi.Context, args *LookupAccessLogSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAccessLogSubscriptionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessLogSubscriptionResult
	err := ctx.Invoke("aws-native:vpclattice:getAccessLogSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessLogSubscriptionArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupAccessLogSubscriptionResult struct {
	Arn            *string                    `pulumi:"arn"`
	DestinationArn *string                    `pulumi:"destinationArn"`
	Id             *string                    `pulumi:"id"`
	ResourceArn    *string                    `pulumi:"resourceArn"`
	ResourceId     *string                    `pulumi:"resourceId"`
	Tags           []AccessLogSubscriptionTag `pulumi:"tags"`
}

func LookupAccessLogSubscriptionOutput(ctx *pulumi.Context, args LookupAccessLogSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupAccessLogSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessLogSubscriptionResult, error) {
			args := v.(LookupAccessLogSubscriptionArgs)
			r, err := LookupAccessLogSubscription(ctx, &args, opts...)
			var s LookupAccessLogSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessLogSubscriptionResultOutput)
}

type LookupAccessLogSubscriptionOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupAccessLogSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessLogSubscriptionArgs)(nil)).Elem()
}

type LookupAccessLogSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupAccessLogSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessLogSubscriptionResult)(nil)).Elem()
}

func (o LookupAccessLogSubscriptionResultOutput) ToLookupAccessLogSubscriptionResultOutput() LookupAccessLogSubscriptionResultOutput {
	return o
}

func (o LookupAccessLogSubscriptionResultOutput) ToLookupAccessLogSubscriptionResultOutputWithContext(ctx context.Context) LookupAccessLogSubscriptionResultOutput {
	return o
}

func (o LookupAccessLogSubscriptionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAccessLogSubscriptionResult] {
	return pulumix.Output[LookupAccessLogSubscriptionResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupAccessLogSubscriptionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessLogSubscriptionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupAccessLogSubscriptionResultOutput) DestinationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessLogSubscriptionResult) *string { return v.DestinationArn }).(pulumi.StringPtrOutput)
}

func (o LookupAccessLogSubscriptionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessLogSubscriptionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAccessLogSubscriptionResultOutput) ResourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessLogSubscriptionResult) *string { return v.ResourceArn }).(pulumi.StringPtrOutput)
}

func (o LookupAccessLogSubscriptionResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessLogSubscriptionResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupAccessLogSubscriptionResultOutput) Tags() AccessLogSubscriptionTagArrayOutput {
	return o.ApplyT(func(v LookupAccessLogSubscriptionResult) []AccessLogSubscriptionTag { return v.Tags }).(AccessLogSubscriptionTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessLogSubscriptionResultOutput{})
}
