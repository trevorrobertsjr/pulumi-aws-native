// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Pinpoint::App
func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAppResult
	err := ctx.Invoke("aws-native:pinpoint:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	Id string `pulumi:"id"`
}

type LookupAppResult struct {
	Arn  *string     `pulumi:"arn"`
	Id   *string     `pulumi:"id"`
	Tags interface{} `pulumi:"tags"`
}

func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppResult, error) {
			args := v.(LookupAppArgs)
			r, err := LookupApp(ctx, &args, opts...)
			var s LookupAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}

type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAppResult] {
	return pulumix.Output[LookupAppResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupAppResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAppResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
