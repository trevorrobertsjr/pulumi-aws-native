// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xray

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This schema provides construct and validation rules for AWS-XRay Group resource parameters.
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("aws-native:xray:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	// The ARN of the group that was generated on creation.
	GroupArn string `pulumi:"groupArn"`
}

type LookupGroupResult struct {
	// The filter expression defining criteria by which to group traces.
	FilterExpression *string `pulumi:"filterExpression"`
	// The ARN of the group that was generated on creation.
	GroupArn *string `pulumi:"groupArn"`
	// The case-sensitive name of the new group. Names must be unique.
	GroupName             *string                     `pulumi:"groupName"`
	InsightsConfiguration *GroupInsightsConfiguration `pulumi:"insightsConfiguration"`
	Tags                  []GroupTag                  `pulumi:"tags"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

type LookupGroupOutputArgs struct {
	// The ARN of the group that was generated on creation.
	GroupArn pulumi.StringInput `pulumi:"groupArn"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGroupResult] {
	return pulumix.Output[LookupGroupResult]{
		OutputState: o.OutputState,
	}
}

// The filter expression defining criteria by which to group traces.
func (o LookupGroupResultOutput) FilterExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.FilterExpression }).(pulumi.StringPtrOutput)
}

// The ARN of the group that was generated on creation.
func (o LookupGroupResultOutput) GroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.GroupArn }).(pulumi.StringPtrOutput)
}

// The case-sensitive name of the new group. Names must be unique.
func (o LookupGroupResultOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) InsightsConfiguration() GroupInsightsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *GroupInsightsConfiguration { return v.InsightsConfiguration }).(GroupInsightsConfigurationPtrOutput)
}

func (o LookupGroupResultOutput) Tags() GroupTagArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GroupTag { return v.Tags }).(GroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
