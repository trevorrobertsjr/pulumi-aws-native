// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Amazon OpenSearchServerless access policy resource
func LookupAccessPolicy(ctx *pulumi.Context, args *LookupAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessPolicyResult
	err := ctx.Invoke("aws-native:opensearchserverless:getAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPolicyArgs struct {
	// The name of the policy
	Name string           `pulumi:"name"`
	Type AccessPolicyType `pulumi:"type"`
}

type LookupAccessPolicyResult struct {
	// The description of the policy
	Description *string `pulumi:"description"`
	// The JSON policy document that is the content for the policy
	Policy *string `pulumi:"policy"`
}

func LookupAccessPolicyOutput(ctx *pulumi.Context, args LookupAccessPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessPolicyResult, error) {
			args := v.(LookupAccessPolicyArgs)
			r, err := LookupAccessPolicy(ctx, &args, opts...)
			var s LookupAccessPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessPolicyResultOutput)
}

type LookupAccessPolicyOutputArgs struct {
	// The name of the policy
	Name pulumi.StringInput    `pulumi:"name"`
	Type AccessPolicyTypeInput `pulumi:"type"`
}

func (LookupAccessPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyArgs)(nil)).Elem()
}

type LookupAccessPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyResult)(nil)).Elem()
}

func (o LookupAccessPolicyResultOutput) ToLookupAccessPolicyResultOutput() LookupAccessPolicyResultOutput {
	return o
}

func (o LookupAccessPolicyResultOutput) ToLookupAccessPolicyResultOutputWithContext(ctx context.Context) LookupAccessPolicyResultOutput {
	return o
}

func (o LookupAccessPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAccessPolicyResult] {
	return pulumix.Output[LookupAccessPolicyResult]{
		OutputState: o.OutputState,
	}
}

// The description of the policy
func (o LookupAccessPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The JSON policy document that is the content for the policy
func (o LookupAccessPolicyResultOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPolicyResultOutput{})
}
