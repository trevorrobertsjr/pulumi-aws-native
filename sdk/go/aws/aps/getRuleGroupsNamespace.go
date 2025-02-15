// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// RuleGroupsNamespace schema for cloudformation.
func LookupRuleGroupsNamespace(ctx *pulumi.Context, args *LookupRuleGroupsNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupRuleGroupsNamespaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRuleGroupsNamespaceResult
	err := ctx.Invoke("aws-native:aps:getRuleGroupsNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleGroupsNamespaceArgs struct {
	// The RuleGroupsNamespace ARN.
	Arn string `pulumi:"arn"`
}

type LookupRuleGroupsNamespaceResult struct {
	// The RuleGroupsNamespace ARN.
	Arn *string `pulumi:"arn"`
	// The RuleGroupsNamespace data.
	Data *string `pulumi:"data"`
	// An array of key-value pairs to apply to this resource.
	Tags []RuleGroupsNamespaceTag `pulumi:"tags"`
	// Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
	Workspace *string `pulumi:"workspace"`
}

func LookupRuleGroupsNamespaceOutput(ctx *pulumi.Context, args LookupRuleGroupsNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupRuleGroupsNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleGroupsNamespaceResult, error) {
			args := v.(LookupRuleGroupsNamespaceArgs)
			r, err := LookupRuleGroupsNamespace(ctx, &args, opts...)
			var s LookupRuleGroupsNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleGroupsNamespaceResultOutput)
}

type LookupRuleGroupsNamespaceOutputArgs struct {
	// The RuleGroupsNamespace ARN.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupRuleGroupsNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleGroupsNamespaceArgs)(nil)).Elem()
}

type LookupRuleGroupsNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupRuleGroupsNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleGroupsNamespaceResult)(nil)).Elem()
}

func (o LookupRuleGroupsNamespaceResultOutput) ToLookupRuleGroupsNamespaceResultOutput() LookupRuleGroupsNamespaceResultOutput {
	return o
}

func (o LookupRuleGroupsNamespaceResultOutput) ToLookupRuleGroupsNamespaceResultOutputWithContext(ctx context.Context) LookupRuleGroupsNamespaceResultOutput {
	return o
}

func (o LookupRuleGroupsNamespaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRuleGroupsNamespaceResult] {
	return pulumix.Output[LookupRuleGroupsNamespaceResult]{
		OutputState: o.OutputState,
	}
}

// The RuleGroupsNamespace ARN.
func (o LookupRuleGroupsNamespaceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupsNamespaceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The RuleGroupsNamespace data.
func (o LookupRuleGroupsNamespaceResultOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupsNamespaceResult) *string { return v.Data }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupRuleGroupsNamespaceResultOutput) Tags() RuleGroupsNamespaceTagArrayOutput {
	return o.ApplyT(func(v LookupRuleGroupsNamespaceResult) []RuleGroupsNamespaceTag { return v.Tags }).(RuleGroupsNamespaceTagArrayOutput)
}

// Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
func (o LookupRuleGroupsNamespaceResultOutput) Workspace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupsNamespaceResult) *string { return v.Workspace }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleGroupsNamespaceResultOutput{})
}
