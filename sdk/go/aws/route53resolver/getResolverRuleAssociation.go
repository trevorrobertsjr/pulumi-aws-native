// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Route53Resolver::ResolverRuleAssociation
func LookupResolverRuleAssociation(ctx *pulumi.Context, args *LookupResolverRuleAssociationArgs, opts ...pulumi.InvokeOption) (*LookupResolverRuleAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResolverRuleAssociationResult
	err := ctx.Invoke("aws-native:route53resolver:getResolverRuleAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResolverRuleAssociationArgs struct {
	// Primary Identifier for Resolver Rule Association
	ResolverRuleAssociationId string `pulumi:"resolverRuleAssociationId"`
}

type LookupResolverRuleAssociationResult struct {
	// Primary Identifier for Resolver Rule Association
	ResolverRuleAssociationId *string `pulumi:"resolverRuleAssociationId"`
}

func LookupResolverRuleAssociationOutput(ctx *pulumi.Context, args LookupResolverRuleAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupResolverRuleAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResolverRuleAssociationResult, error) {
			args := v.(LookupResolverRuleAssociationArgs)
			r, err := LookupResolverRuleAssociation(ctx, &args, opts...)
			var s LookupResolverRuleAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResolverRuleAssociationResultOutput)
}

type LookupResolverRuleAssociationOutputArgs struct {
	// Primary Identifier for Resolver Rule Association
	ResolverRuleAssociationId pulumi.StringInput `pulumi:"resolverRuleAssociationId"`
}

func (LookupResolverRuleAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverRuleAssociationArgs)(nil)).Elem()
}

type LookupResolverRuleAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupResolverRuleAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverRuleAssociationResult)(nil)).Elem()
}

func (o LookupResolverRuleAssociationResultOutput) ToLookupResolverRuleAssociationResultOutput() LookupResolverRuleAssociationResultOutput {
	return o
}

func (o LookupResolverRuleAssociationResultOutput) ToLookupResolverRuleAssociationResultOutputWithContext(ctx context.Context) LookupResolverRuleAssociationResultOutput {
	return o
}

func (o LookupResolverRuleAssociationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupResolverRuleAssociationResult] {
	return pulumix.Output[LookupResolverRuleAssociationResult]{
		OutputState: o.OutputState,
	}
}

// Primary Identifier for Resolver Rule Association
func (o LookupResolverRuleAssociationResultOutput) ResolverRuleAssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverRuleAssociationResult) *string { return v.ResolverRuleAssociationId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResolverRuleAssociationResultOutput{})
}
