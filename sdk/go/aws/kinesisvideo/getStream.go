// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisvideo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type Definition for AWS::KinesisVideo::Stream
func LookupStream(ctx *pulumi.Context, args *LookupStreamArgs, opts ...pulumi.InvokeOption) (*LookupStreamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStreamResult
	err := ctx.Invoke("aws-native:kinesisvideo:getStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamArgs struct {
	// The name of the Kinesis Video stream.
	Name string `pulumi:"name"`
}

type LookupStreamResult struct {
	// The Amazon Resource Name (ARN) of the Kinesis Video stream.
	Arn *string `pulumi:"arn"`
	// The number of hours till which Kinesis Video will retain the data in the stream
	DataRetentionInHours *int `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream.
	DeviceName *string `pulumi:"deviceName"`
	// AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream.
	MediaType *string `pulumi:"mediaType"`
	// An array of key-value pairs associated with the Kinesis Video Stream.
	Tags []StreamTag `pulumi:"tags"`
}

func LookupStreamOutput(ctx *pulumi.Context, args LookupStreamOutputArgs, opts ...pulumi.InvokeOption) LookupStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamResult, error) {
			args := v.(LookupStreamArgs)
			r, err := LookupStream(ctx, &args, opts...)
			var s LookupStreamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamResultOutput)
}

type LookupStreamOutputArgs struct {
	// The name of the Kinesis Video stream.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamArgs)(nil)).Elem()
}

type LookupStreamResultOutput struct{ *pulumi.OutputState }

func (LookupStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamResult)(nil)).Elem()
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutput() LookupStreamResultOutput {
	return o
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutputWithContext(ctx context.Context) LookupStreamResultOutput {
	return o
}

func (o LookupStreamResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStreamResult] {
	return pulumix.Output[LookupStreamResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the Kinesis Video stream.
func (o LookupStreamResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The number of hours till which Kinesis Video will retain the data in the stream
func (o LookupStreamResultOutput) DataRetentionInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *int { return v.DataRetentionInHours }).(pulumi.IntPtrOutput)
}

// The name of the device that is writing to the stream.
func (o LookupStreamResultOutput) DeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *string { return v.DeviceName }).(pulumi.StringPtrOutput)
}

// AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.
func (o LookupStreamResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The media type of the stream. Consumers of the stream can use this information when processing the stream.
func (o LookupStreamResultOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *string { return v.MediaType }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs associated with the Kinesis Video Stream.
func (o LookupStreamResultOutput) Tags() StreamTagArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []StreamTag { return v.Tags }).(StreamTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamResultOutput{})
}
