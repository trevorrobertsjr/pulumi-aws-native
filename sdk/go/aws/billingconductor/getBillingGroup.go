// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package billingconductor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A billing group is a set of linked account which belong to the same end customer. It can be seen as a virtual consolidated billing family.
func LookupBillingGroup(ctx *pulumi.Context, args *LookupBillingGroupArgs, opts ...pulumi.InvokeOption) (*LookupBillingGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBillingGroupResult
	err := ctx.Invoke("aws-native:billingconductor:getBillingGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBillingGroupArgs struct {
	// Billing Group ARN
	Arn string `pulumi:"arn"`
}

type LookupBillingGroupResult struct {
	AccountGrouping *BillingGroupAccountGrouping `pulumi:"accountGrouping"`
	// Billing Group ARN
	Arn                   *string                            `pulumi:"arn"`
	ComputationPreference *BillingGroupComputationPreference `pulumi:"computationPreference"`
	// Creation timestamp in UNIX epoch time format
	CreationTime *int    `pulumi:"creationTime"`
	Description  *string `pulumi:"description"`
	// Latest modified timestamp in UNIX epoch time format
	LastModifiedTime *int    `pulumi:"lastModifiedTime"`
	Name             *string `pulumi:"name"`
	// Number of accounts in the billing group
	Size         *int                `pulumi:"size"`
	Status       *BillingGroupStatus `pulumi:"status"`
	StatusReason *string             `pulumi:"statusReason"`
	Tags         []BillingGroupTag   `pulumi:"tags"`
}

func LookupBillingGroupOutput(ctx *pulumi.Context, args LookupBillingGroupOutputArgs, opts ...pulumi.InvokeOption) LookupBillingGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBillingGroupResult, error) {
			args := v.(LookupBillingGroupArgs)
			r, err := LookupBillingGroup(ctx, &args, opts...)
			var s LookupBillingGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBillingGroupResultOutput)
}

type LookupBillingGroupOutputArgs struct {
	// Billing Group ARN
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupBillingGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingGroupArgs)(nil)).Elem()
}

type LookupBillingGroupResultOutput struct{ *pulumi.OutputState }

func (LookupBillingGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingGroupResult)(nil)).Elem()
}

func (o LookupBillingGroupResultOutput) ToLookupBillingGroupResultOutput() LookupBillingGroupResultOutput {
	return o
}

func (o LookupBillingGroupResultOutput) ToLookupBillingGroupResultOutputWithContext(ctx context.Context) LookupBillingGroupResultOutput {
	return o
}

func (o LookupBillingGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBillingGroupResult] {
	return pulumix.Output[LookupBillingGroupResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBillingGroupResultOutput) AccountGrouping() BillingGroupAccountGroupingPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *BillingGroupAccountGrouping { return v.AccountGrouping }).(BillingGroupAccountGroupingPtrOutput)
}

// Billing Group ARN
func (o LookupBillingGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupBillingGroupResultOutput) ComputationPreference() BillingGroupComputationPreferencePtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *BillingGroupComputationPreference { return v.ComputationPreference }).(BillingGroupComputationPreferencePtrOutput)
}

// Creation timestamp in UNIX epoch time format
func (o LookupBillingGroupResultOutput) CreationTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *int { return v.CreationTime }).(pulumi.IntPtrOutput)
}

func (o LookupBillingGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Latest modified timestamp in UNIX epoch time format
func (o LookupBillingGroupResultOutput) LastModifiedTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *int { return v.LastModifiedTime }).(pulumi.IntPtrOutput)
}

func (o LookupBillingGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Number of accounts in the billing group
func (o LookupBillingGroupResultOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *int { return v.Size }).(pulumi.IntPtrOutput)
}

func (o LookupBillingGroupResultOutput) Status() BillingGroupStatusPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *BillingGroupStatus { return v.Status }).(BillingGroupStatusPtrOutput)
}

func (o LookupBillingGroupResultOutput) StatusReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *string { return v.StatusReason }).(pulumi.StringPtrOutput)
}

func (o LookupBillingGroupResultOutput) Tags() BillingGroupTagArrayOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) []BillingGroupTag { return v.Tags }).(BillingGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBillingGroupResultOutput{})
}
