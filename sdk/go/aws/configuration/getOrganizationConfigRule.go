// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Config::OrganizationConfigRule
func LookupOrganizationConfigRule(ctx *pulumi.Context, args *LookupOrganizationConfigRuleArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationConfigRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationConfigRuleResult
	err := ctx.Invoke("aws-native:configuration:getOrganizationConfigRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationConfigRuleArgs struct {
	Id string `pulumi:"id"`
}

type LookupOrganizationConfigRuleResult struct {
	ExcludedAccounts                     []string                                                    `pulumi:"excludedAccounts"`
	Id                                   *string                                                     `pulumi:"id"`
	OrganizationCustomPolicyRuleMetadata *OrganizationConfigRuleOrganizationCustomPolicyRuleMetadata `pulumi:"organizationCustomPolicyRuleMetadata"`
	OrganizationCustomRuleMetadata       *OrganizationConfigRuleOrganizationCustomRuleMetadata       `pulumi:"organizationCustomRuleMetadata"`
	OrganizationManagedRuleMetadata      *OrganizationConfigRuleOrganizationManagedRuleMetadata      `pulumi:"organizationManagedRuleMetadata"`
}

func LookupOrganizationConfigRuleOutput(ctx *pulumi.Context, args LookupOrganizationConfigRuleOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationConfigRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationConfigRuleResult, error) {
			args := v.(LookupOrganizationConfigRuleArgs)
			r, err := LookupOrganizationConfigRule(ctx, &args, opts...)
			var s LookupOrganizationConfigRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationConfigRuleResultOutput)
}

type LookupOrganizationConfigRuleOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupOrganizationConfigRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationConfigRuleArgs)(nil)).Elem()
}

type LookupOrganizationConfigRuleResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationConfigRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationConfigRuleResult)(nil)).Elem()
}

func (o LookupOrganizationConfigRuleResultOutput) ToLookupOrganizationConfigRuleResultOutput() LookupOrganizationConfigRuleResultOutput {
	return o
}

func (o LookupOrganizationConfigRuleResultOutput) ToLookupOrganizationConfigRuleResultOutputWithContext(ctx context.Context) LookupOrganizationConfigRuleResultOutput {
	return o
}

func (o LookupOrganizationConfigRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOrganizationConfigRuleResult] {
	return pulumix.Output[LookupOrganizationConfigRuleResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupOrganizationConfigRuleResultOutput) ExcludedAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrganizationConfigRuleResult) []string { return v.ExcludedAccounts }).(pulumi.StringArrayOutput)
}

func (o LookupOrganizationConfigRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupOrganizationConfigRuleResultOutput) OrganizationCustomPolicyRuleMetadata() OrganizationConfigRuleOrganizationCustomPolicyRuleMetadataPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigRuleResult) *OrganizationConfigRuleOrganizationCustomPolicyRuleMetadata {
		return v.OrganizationCustomPolicyRuleMetadata
	}).(OrganizationConfigRuleOrganizationCustomPolicyRuleMetadataPtrOutput)
}

func (o LookupOrganizationConfigRuleResultOutput) OrganizationCustomRuleMetadata() OrganizationConfigRuleOrganizationCustomRuleMetadataPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigRuleResult) *OrganizationConfigRuleOrganizationCustomRuleMetadata {
		return v.OrganizationCustomRuleMetadata
	}).(OrganizationConfigRuleOrganizationCustomRuleMetadataPtrOutput)
}

func (o LookupOrganizationConfigRuleResultOutput) OrganizationManagedRuleMetadata() OrganizationConfigRuleOrganizationManagedRuleMetadataPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigRuleResult) *OrganizationConfigRuleOrganizationManagedRuleMetadata {
		return v.OrganizationManagedRuleMetadata
	}).(OrganizationConfigRuleOrganizationManagedRuleMetadataPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationConfigRuleResultOutput{})
}
