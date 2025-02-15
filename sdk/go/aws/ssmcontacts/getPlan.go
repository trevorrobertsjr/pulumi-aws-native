// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmcontacts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Engagement Plan for a SSM Incident Manager Contact.
func LookupPlan(ctx *pulumi.Context, args *LookupPlanArgs, opts ...pulumi.InvokeOption) (*LookupPlanResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPlanResult
	err := ctx.Invoke("aws-native:ssmcontacts:getPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlanArgs struct {
	// The Amazon Resource Name (ARN) of the contact.
	Arn string `pulumi:"arn"`
}

type LookupPlanResult struct {
	// The Amazon Resource Name (ARN) of the contact.
	Arn *string `pulumi:"arn"`
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	Stages []PlanStage `pulumi:"stages"`
}

func LookupPlanOutput(ctx *pulumi.Context, args LookupPlanOutputArgs, opts ...pulumi.InvokeOption) LookupPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlanResult, error) {
			args := v.(LookupPlanArgs)
			r, err := LookupPlan(ctx, &args, opts...)
			var s LookupPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPlanResultOutput)
}

type LookupPlanOutputArgs struct {
	// The Amazon Resource Name (ARN) of the contact.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlanArgs)(nil)).Elem()
}

type LookupPlanResultOutput struct{ *pulumi.OutputState }

func (LookupPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlanResult)(nil)).Elem()
}

func (o LookupPlanResultOutput) ToLookupPlanResultOutput() LookupPlanResultOutput {
	return o
}

func (o LookupPlanResultOutput) ToLookupPlanResultOutputWithContext(ctx context.Context) LookupPlanResultOutput {
	return o
}

func (o LookupPlanResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPlanResult] {
	return pulumix.Output[LookupPlanResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the contact.
func (o LookupPlanResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlanResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
func (o LookupPlanResultOutput) Stages() PlanStageArrayOutput {
	return o.ApplyT(func(v LookupPlanResult) []PlanStage { return v.Stages }).(PlanStageArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlanResultOutput{})
}
