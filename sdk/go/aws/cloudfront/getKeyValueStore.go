// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::CloudFront::KeyValueStore
func LookupKeyValueStore(ctx *pulumi.Context, args *LookupKeyValueStoreArgs, opts ...pulumi.InvokeOption) (*LookupKeyValueStoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKeyValueStoreResult
	err := ctx.Invoke("aws-native:cloudfront:getKeyValueStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyValueStoreArgs struct {
	Name string `pulumi:"name"`
}

type LookupKeyValueStoreResult struct {
	Arn     *string `pulumi:"arn"`
	Comment *string `pulumi:"comment"`
	Id      *string `pulumi:"id"`
	Status  *string `pulumi:"status"`
}

func LookupKeyValueStoreOutput(ctx *pulumi.Context, args LookupKeyValueStoreOutputArgs, opts ...pulumi.InvokeOption) LookupKeyValueStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKeyValueStoreResult, error) {
			args := v.(LookupKeyValueStoreArgs)
			r, err := LookupKeyValueStore(ctx, &args, opts...)
			var s LookupKeyValueStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKeyValueStoreResultOutput)
}

type LookupKeyValueStoreOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupKeyValueStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyValueStoreArgs)(nil)).Elem()
}

type LookupKeyValueStoreResultOutput struct{ *pulumi.OutputState }

func (LookupKeyValueStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyValueStoreResult)(nil)).Elem()
}

func (o LookupKeyValueStoreResultOutput) ToLookupKeyValueStoreResultOutput() LookupKeyValueStoreResultOutput {
	return o
}

func (o LookupKeyValueStoreResultOutput) ToLookupKeyValueStoreResultOutputWithContext(ctx context.Context) LookupKeyValueStoreResultOutput {
	return o
}

func (o LookupKeyValueStoreResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupKeyValueStoreResult] {
	return pulumix.Output[LookupKeyValueStoreResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupKeyValueStoreResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyValueStoreResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupKeyValueStoreResultOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyValueStoreResult) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o LookupKeyValueStoreResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyValueStoreResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupKeyValueStoreResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyValueStoreResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeyValueStoreResultOutput{})
}
