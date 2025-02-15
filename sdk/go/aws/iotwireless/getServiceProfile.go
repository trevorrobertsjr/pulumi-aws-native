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

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupServiceProfile(ctx *pulumi.Context, args *LookupServiceProfileArgs, opts ...pulumi.InvokeOption) (*LookupServiceProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceProfileResult
	err := ctx.Invoke("aws-native:iotwireless:getServiceProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceProfileArgs struct {
	// Service profile Id. Returned after successful create.
	Id string `pulumi:"id"`
}

type LookupServiceProfileResult struct {
	// Service profile Arn. Returned after successful create.
	Arn *string `pulumi:"arn"`
	// Service profile Id. Returned after successful create.
	Id *string `pulumi:"id"`
	// LoRaWAN supports all LoRa specific attributes for service profile for CreateServiceProfile operation
	LoRaWan *ServiceProfileLoRaWanServiceProfile `pulumi:"loRaWan"`
	// Name of service profile
	Name *string `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the service profile.
	Tags []ServiceProfileTag `pulumi:"tags"`
}

func LookupServiceProfileOutput(ctx *pulumi.Context, args LookupServiceProfileOutputArgs, opts ...pulumi.InvokeOption) LookupServiceProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceProfileResult, error) {
			args := v.(LookupServiceProfileArgs)
			r, err := LookupServiceProfile(ctx, &args, opts...)
			var s LookupServiceProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceProfileResultOutput)
}

type LookupServiceProfileOutputArgs struct {
	// Service profile Id. Returned after successful create.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupServiceProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceProfileArgs)(nil)).Elem()
}

type LookupServiceProfileResultOutput struct{ *pulumi.OutputState }

func (LookupServiceProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceProfileResult)(nil)).Elem()
}

func (o LookupServiceProfileResultOutput) ToLookupServiceProfileResultOutput() LookupServiceProfileResultOutput {
	return o
}

func (o LookupServiceProfileResultOutput) ToLookupServiceProfileResultOutputWithContext(ctx context.Context) LookupServiceProfileResultOutput {
	return o
}

func (o LookupServiceProfileResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServiceProfileResult] {
	return pulumix.Output[LookupServiceProfileResult]{
		OutputState: o.OutputState,
	}
}

// Service profile Arn. Returned after successful create.
func (o LookupServiceProfileResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceProfileResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Service profile Id. Returned after successful create.
func (o LookupServiceProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// LoRaWAN supports all LoRa specific attributes for service profile for CreateServiceProfile operation
func (o LookupServiceProfileResultOutput) LoRaWan() ServiceProfileLoRaWanServiceProfilePtrOutput {
	return o.ApplyT(func(v LookupServiceProfileResult) *ServiceProfileLoRaWanServiceProfile { return v.LoRaWan }).(ServiceProfileLoRaWanServiceProfilePtrOutput)
}

// Name of service profile
func (o LookupServiceProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the service profile.
func (o LookupServiceProfileResultOutput) Tags() ServiceProfileTagArrayOutput {
	return o.ApplyT(func(v LookupServiceProfileResult) []ServiceProfileTag { return v.Tags }).(ServiceProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceProfileResultOutput{})
}
