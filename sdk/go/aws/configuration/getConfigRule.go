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

// Schema for AWS Config ConfigRule
func LookupConfigRule(ctx *pulumi.Context, args *LookupConfigRuleArgs, opts ...pulumi.InvokeOption) (*LookupConfigRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigRuleResult
	err := ctx.Invoke("aws-native:configuration:getConfigRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigRuleArgs struct {
	// Name for the AWS Config rule
	ConfigRuleName string `pulumi:"configRuleName"`
}

type LookupConfigRuleResult struct {
	// ARN generated for the AWS Config rule
	Arn *string `pulumi:"arn"`
	// Compliance details of the Config rule
	Compliance *ComplianceProperties `pulumi:"compliance"`
	// ID of the config rule
	ConfigRuleId *string `pulumi:"configRuleId"`
	// Description provided for the AWS Config rule
	Description *string `pulumi:"description"`
	// List of EvaluationModeConfiguration objects
	EvaluationModes []ConfigRuleEvaluationModeConfiguration `pulumi:"evaluationModes"`
	// JSON string passed the Lambda function
	InputParameters *string `pulumi:"inputParameters"`
	// Maximum frequency at which the rule has to be evaluated
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// Scope to constrain which resources can trigger the AWS Config rule
	Scope *ConfigRuleScope `pulumi:"scope"`
	// Source of events for the AWS Config rule
	Source *ConfigRuleSource `pulumi:"source"`
}

func LookupConfigRuleOutput(ctx *pulumi.Context, args LookupConfigRuleOutputArgs, opts ...pulumi.InvokeOption) LookupConfigRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigRuleResult, error) {
			args := v.(LookupConfigRuleArgs)
			r, err := LookupConfigRule(ctx, &args, opts...)
			var s LookupConfigRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigRuleResultOutput)
}

type LookupConfigRuleOutputArgs struct {
	// Name for the AWS Config rule
	ConfigRuleName pulumi.StringInput `pulumi:"configRuleName"`
}

func (LookupConfigRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigRuleArgs)(nil)).Elem()
}

type LookupConfigRuleResultOutput struct{ *pulumi.OutputState }

func (LookupConfigRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigRuleResult)(nil)).Elem()
}

func (o LookupConfigRuleResultOutput) ToLookupConfigRuleResultOutput() LookupConfigRuleResultOutput {
	return o
}

func (o LookupConfigRuleResultOutput) ToLookupConfigRuleResultOutputWithContext(ctx context.Context) LookupConfigRuleResultOutput {
	return o
}

func (o LookupConfigRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConfigRuleResult] {
	return pulumix.Output[LookupConfigRuleResult]{
		OutputState: o.OutputState,
	}
}

// ARN generated for the AWS Config rule
func (o LookupConfigRuleResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Compliance details of the Config rule
func (o LookupConfigRuleResultOutput) Compliance() CompliancePropertiesPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *ComplianceProperties { return v.Compliance }).(CompliancePropertiesPtrOutput)
}

// ID of the config rule
func (o LookupConfigRuleResultOutput) ConfigRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.ConfigRuleId }).(pulumi.StringPtrOutput)
}

// Description provided for the AWS Config rule
func (o LookupConfigRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// List of EvaluationModeConfiguration objects
func (o LookupConfigRuleResultOutput) EvaluationModes() ConfigRuleEvaluationModeConfigurationArrayOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) []ConfigRuleEvaluationModeConfiguration { return v.EvaluationModes }).(ConfigRuleEvaluationModeConfigurationArrayOutput)
}

// JSON string passed the Lambda function
func (o LookupConfigRuleResultOutput) InputParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.InputParameters }).(pulumi.StringPtrOutput)
}

// Maximum frequency at which the rule has to be evaluated
func (o LookupConfigRuleResultOutput) MaximumExecutionFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.MaximumExecutionFrequency }).(pulumi.StringPtrOutput)
}

// Scope to constrain which resources can trigger the AWS Config rule
func (o LookupConfigRuleResultOutput) Scope() ConfigRuleScopePtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *ConfigRuleScope { return v.Scope }).(ConfigRuleScopePtrOutput)
}

// Source of events for the AWS Config rule
func (o LookupConfigRuleResultOutput) Source() ConfigRuleSourcePtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *ConfigRuleSource { return v.Source }).(ConfigRuleSourcePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigRuleResultOutput{})
}
