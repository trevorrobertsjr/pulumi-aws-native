// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::GuardDuty::IPSet
func LookupIpSet(ctx *pulumi.Context, args *LookupIpSetArgs, opts ...pulumi.InvokeOption) (*LookupIpSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpSetResult
	err := ctx.Invoke("aws-native:guardduty:getIpSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpSetArgs struct {
	DetectorId string `pulumi:"detectorId"`
	Id         string `pulumi:"id"`
}

type LookupIpSetResult struct {
	Id       *string        `pulumi:"id"`
	Location *string        `pulumi:"location"`
	Name     *string        `pulumi:"name"`
	Tags     []IpSetTagItem `pulumi:"tags"`
}

func LookupIpSetOutput(ctx *pulumi.Context, args LookupIpSetOutputArgs, opts ...pulumi.InvokeOption) LookupIpSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpSetResult, error) {
			args := v.(LookupIpSetArgs)
			r, err := LookupIpSet(ctx, &args, opts...)
			var s LookupIpSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpSetResultOutput)
}

type LookupIpSetOutputArgs struct {
	DetectorId pulumi.StringInput `pulumi:"detectorId"`
	Id         pulumi.StringInput `pulumi:"id"`
}

func (LookupIpSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpSetArgs)(nil)).Elem()
}

type LookupIpSetResultOutput struct{ *pulumi.OutputState }

func (LookupIpSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpSetResult)(nil)).Elem()
}

func (o LookupIpSetResultOutput) ToLookupIpSetResultOutput() LookupIpSetResultOutput {
	return o
}

func (o LookupIpSetResultOutput) ToLookupIpSetResultOutputWithContext(ctx context.Context) LookupIpSetResultOutput {
	return o
}

func (o LookupIpSetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIpSetResult] {
	return pulumix.Output[LookupIpSetResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupIpSetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpSetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupIpSetResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpSetResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIpSetResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpSetResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupIpSetResultOutput) Tags() IpSetTagItemArrayOutput {
	return o.ApplyT(func(v LookupIpSetResult) []IpSetTagItem { return v.Tags }).(IpSetTagItemArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpSetResultOutput{})
}
