// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Schema for IAM Group Policy
func LookupGroupPolicy(ctx *pulumi.Context, args *LookupGroupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupGroupPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupPolicyResult
	err := ctx.Invoke("aws-native:iam:getGroupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupPolicyArgs struct {
	// The name of the group to associate the policy with.
	GroupName string `pulumi:"groupName"`
	// The name of the policy document.
	PolicyName string `pulumi:"policyName"`
}

type LookupGroupPolicyResult struct {
	// The policy document.
	PolicyDocument interface{} `pulumi:"policyDocument"`
}

func LookupGroupPolicyOutput(ctx *pulumi.Context, args LookupGroupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupGroupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupPolicyResult, error) {
			args := v.(LookupGroupPolicyArgs)
			r, err := LookupGroupPolicy(ctx, &args, opts...)
			var s LookupGroupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupPolicyResultOutput)
}

type LookupGroupPolicyOutputArgs struct {
	// The name of the group to associate the policy with.
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// The name of the policy document.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
}

func (LookupGroupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupPolicyArgs)(nil)).Elem()
}

type LookupGroupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupGroupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupPolicyResult)(nil)).Elem()
}

func (o LookupGroupPolicyResultOutput) ToLookupGroupPolicyResultOutput() LookupGroupPolicyResultOutput {
	return o
}

func (o LookupGroupPolicyResultOutput) ToLookupGroupPolicyResultOutputWithContext(ctx context.Context) LookupGroupPolicyResultOutput {
	return o
}

func (o LookupGroupPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGroupPolicyResult] {
	return pulumix.Output[LookupGroupPolicyResult]{
		OutputState: o.OutputState,
	}
}

// The policy document.
func (o LookupGroupPolicyResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupGroupPolicyResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupPolicyResultOutput{})
}
