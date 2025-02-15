// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackagev2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::MediaPackageV2::ChannelGroup Resource Type
func LookupChannelGroup(ctx *pulumi.Context, args *LookupChannelGroupArgs, opts ...pulumi.InvokeOption) (*LookupChannelGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupChannelGroupResult
	err := ctx.Invoke("aws-native:mediapackagev2:getChannelGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelGroupArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupChannelGroupResult struct {
	Arn          *string           `pulumi:"arn"`
	CreatedAt    *string           `pulumi:"createdAt"`
	Description  *string           `pulumi:"description"`
	EgressDomain *string           `pulumi:"egressDomain"`
	ModifiedAt   *string           `pulumi:"modifiedAt"`
	Tags         []ChannelGroupTag `pulumi:"tags"`
}

func LookupChannelGroupOutput(ctx *pulumi.Context, args LookupChannelGroupOutputArgs, opts ...pulumi.InvokeOption) LookupChannelGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelGroupResult, error) {
			args := v.(LookupChannelGroupArgs)
			r, err := LookupChannelGroup(ctx, &args, opts...)
			var s LookupChannelGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelGroupResultOutput)
}

type LookupChannelGroupOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupChannelGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelGroupArgs)(nil)).Elem()
}

type LookupChannelGroupResultOutput struct{ *pulumi.OutputState }

func (LookupChannelGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelGroupResult)(nil)).Elem()
}

func (o LookupChannelGroupResultOutput) ToLookupChannelGroupResultOutput() LookupChannelGroupResultOutput {
	return o
}

func (o LookupChannelGroupResultOutput) ToLookupChannelGroupResultOutputWithContext(ctx context.Context) LookupChannelGroupResultOutput {
	return o
}

func (o LookupChannelGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupChannelGroupResult] {
	return pulumix.Output[LookupChannelGroupResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupChannelGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupChannelGroupResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelGroupResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupChannelGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupChannelGroupResultOutput) EgressDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelGroupResult) *string { return v.EgressDomain }).(pulumi.StringPtrOutput)
}

func (o LookupChannelGroupResultOutput) ModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelGroupResult) *string { return v.ModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupChannelGroupResultOutput) Tags() ChannelGroupTagArrayOutput {
	return o.ApplyT(func(v LookupChannelGroupResult) []ChannelGroupTag { return v.Tags }).(ChannelGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelGroupResultOutput{})
}
