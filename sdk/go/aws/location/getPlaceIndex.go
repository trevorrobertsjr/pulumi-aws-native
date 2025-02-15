// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::Location::PlaceIndex Resource Type
func LookupPlaceIndex(ctx *pulumi.Context, args *LookupPlaceIndexArgs, opts ...pulumi.InvokeOption) (*LookupPlaceIndexResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPlaceIndexResult
	err := ctx.Invoke("aws-native:location:getPlaceIndex", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlaceIndexArgs struct {
	IndexName string `pulumi:"indexName"`
}

type LookupPlaceIndexResult struct {
	Arn                     *string                            `pulumi:"arn"`
	CreateTime              *string                            `pulumi:"createTime"`
	DataSourceConfiguration *PlaceIndexDataSourceConfiguration `pulumi:"dataSourceConfiguration"`
	Description             *string                            `pulumi:"description"`
	IndexArn                *string                            `pulumi:"indexArn"`
	PricingPlan             *PlaceIndexPricingPlan             `pulumi:"pricingPlan"`
	// An array of key-value pairs to apply to this resource.
	Tags       []PlaceIndexTag `pulumi:"tags"`
	UpdateTime *string         `pulumi:"updateTime"`
}

func LookupPlaceIndexOutput(ctx *pulumi.Context, args LookupPlaceIndexOutputArgs, opts ...pulumi.InvokeOption) LookupPlaceIndexResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlaceIndexResult, error) {
			args := v.(LookupPlaceIndexArgs)
			r, err := LookupPlaceIndex(ctx, &args, opts...)
			var s LookupPlaceIndexResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPlaceIndexResultOutput)
}

type LookupPlaceIndexOutputArgs struct {
	IndexName pulumi.StringInput `pulumi:"indexName"`
}

func (LookupPlaceIndexOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlaceIndexArgs)(nil)).Elem()
}

type LookupPlaceIndexResultOutput struct{ *pulumi.OutputState }

func (LookupPlaceIndexResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlaceIndexResult)(nil)).Elem()
}

func (o LookupPlaceIndexResultOutput) ToLookupPlaceIndexResultOutput() LookupPlaceIndexResultOutput {
	return o
}

func (o LookupPlaceIndexResultOutput) ToLookupPlaceIndexResultOutputWithContext(ctx context.Context) LookupPlaceIndexResultOutput {
	return o
}

func (o LookupPlaceIndexResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPlaceIndexResult] {
	return pulumix.Output[LookupPlaceIndexResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupPlaceIndexResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupPlaceIndexResultOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) *string { return v.CreateTime }).(pulumi.StringPtrOutput)
}

func (o LookupPlaceIndexResultOutput) DataSourceConfiguration() PlaceIndexDataSourceConfigurationPtrOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) *PlaceIndexDataSourceConfiguration { return v.DataSourceConfiguration }).(PlaceIndexDataSourceConfigurationPtrOutput)
}

func (o LookupPlaceIndexResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPlaceIndexResultOutput) IndexArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) *string { return v.IndexArn }).(pulumi.StringPtrOutput)
}

func (o LookupPlaceIndexResultOutput) PricingPlan() PlaceIndexPricingPlanPtrOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) *PlaceIndexPricingPlan { return v.PricingPlan }).(PlaceIndexPricingPlanPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupPlaceIndexResultOutput) Tags() PlaceIndexTagArrayOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) []PlaceIndexTag { return v.Tags }).(PlaceIndexTagArrayOutput)
}

func (o LookupPlaceIndexResultOutput) UpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) *string { return v.UpdateTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlaceIndexResultOutput{})
}
