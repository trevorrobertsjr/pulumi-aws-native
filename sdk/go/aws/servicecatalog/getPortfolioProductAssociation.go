// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::ServiceCatalog::PortfolioProductAssociation
func LookupPortfolioProductAssociation(ctx *pulumi.Context, args *LookupPortfolioProductAssociationArgs, opts ...pulumi.InvokeOption) (*LookupPortfolioProductAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPortfolioProductAssociationResult
	err := ctx.Invoke("aws-native:servicecatalog:getPortfolioProductAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPortfolioProductAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupPortfolioProductAssociationResult struct {
	Id *string `pulumi:"id"`
}

func LookupPortfolioProductAssociationOutput(ctx *pulumi.Context, args LookupPortfolioProductAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupPortfolioProductAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPortfolioProductAssociationResult, error) {
			args := v.(LookupPortfolioProductAssociationArgs)
			r, err := LookupPortfolioProductAssociation(ctx, &args, opts...)
			var s LookupPortfolioProductAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPortfolioProductAssociationResultOutput)
}

type LookupPortfolioProductAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPortfolioProductAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortfolioProductAssociationArgs)(nil)).Elem()
}

type LookupPortfolioProductAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupPortfolioProductAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortfolioProductAssociationResult)(nil)).Elem()
}

func (o LookupPortfolioProductAssociationResultOutput) ToLookupPortfolioProductAssociationResultOutput() LookupPortfolioProductAssociationResultOutput {
	return o
}

func (o LookupPortfolioProductAssociationResultOutput) ToLookupPortfolioProductAssociationResultOutputWithContext(ctx context.Context) LookupPortfolioProductAssociationResultOutput {
	return o
}

func (o LookupPortfolioProductAssociationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPortfolioProductAssociationResult] {
	return pulumix.Output[LookupPortfolioProductAssociationResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupPortfolioProductAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortfolioProductAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPortfolioProductAssociationResultOutput{})
}
