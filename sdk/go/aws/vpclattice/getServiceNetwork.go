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

// A service network is a logical boundary for a collection of services. You can associate services and VPCs with a service network.
func LookupServiceNetwork(ctx *pulumi.Context, args *LookupServiceNetworkArgs, opts ...pulumi.InvokeOption) (*LookupServiceNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceNetworkResult
	err := ctx.Invoke("aws-native:vpclattice:getServiceNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceNetworkArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupServiceNetworkResult struct {
	Arn           *string                 `pulumi:"arn"`
	AuthType      *ServiceNetworkAuthType `pulumi:"authType"`
	CreatedAt     *string                 `pulumi:"createdAt"`
	Id            *string                 `pulumi:"id"`
	LastUpdatedAt *string                 `pulumi:"lastUpdatedAt"`
	Tags          []ServiceNetworkTag     `pulumi:"tags"`
}

func LookupServiceNetworkOutput(ctx *pulumi.Context, args LookupServiceNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupServiceNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceNetworkResult, error) {
			args := v.(LookupServiceNetworkArgs)
			r, err := LookupServiceNetwork(ctx, &args, opts...)
			var s LookupServiceNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceNetworkResultOutput)
}

type LookupServiceNetworkOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupServiceNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceNetworkArgs)(nil)).Elem()
}

type LookupServiceNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupServiceNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceNetworkResult)(nil)).Elem()
}

func (o LookupServiceNetworkResultOutput) ToLookupServiceNetworkResultOutput() LookupServiceNetworkResultOutput {
	return o
}

func (o LookupServiceNetworkResultOutput) ToLookupServiceNetworkResultOutputWithContext(ctx context.Context) LookupServiceNetworkResultOutput {
	return o
}

func (o LookupServiceNetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServiceNetworkResult] {
	return pulumix.Output[LookupServiceNetworkResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupServiceNetworkResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceNetworkResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupServiceNetworkResultOutput) AuthType() ServiceNetworkAuthTypePtrOutput {
	return o.ApplyT(func(v LookupServiceNetworkResult) *ServiceNetworkAuthType { return v.AuthType }).(ServiceNetworkAuthTypePtrOutput)
}

func (o LookupServiceNetworkResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceNetworkResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupServiceNetworkResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceNetworkResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupServiceNetworkResultOutput) LastUpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceNetworkResult) *string { return v.LastUpdatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupServiceNetworkResultOutput) Tags() ServiceNetworkTagArrayOutput {
	return o.ApplyT(func(v LookupServiceNetworkResult) []ServiceNetworkTag { return v.Tags }).(ServiceNetworkTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceNetworkResultOutput{})
}
