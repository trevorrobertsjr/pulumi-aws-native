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

var _ = internal.GetEnvOrDefault

// A key-value pair to associate with a resource.
type RuleGroupsNamespaceTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// RuleGroupsNamespaceTagInput is an input type that accepts RuleGroupsNamespaceTagArgs and RuleGroupsNamespaceTagOutput values.
// You can construct a concrete instance of `RuleGroupsNamespaceTagInput` via:
//
//	RuleGroupsNamespaceTagArgs{...}
type RuleGroupsNamespaceTagInput interface {
	pulumi.Input

	ToRuleGroupsNamespaceTagOutput() RuleGroupsNamespaceTagOutput
	ToRuleGroupsNamespaceTagOutputWithContext(context.Context) RuleGroupsNamespaceTagOutput
}

// A key-value pair to associate with a resource.
type RuleGroupsNamespaceTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (RuleGroupsNamespaceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroupsNamespaceTag)(nil)).Elem()
}

func (i RuleGroupsNamespaceTagArgs) ToRuleGroupsNamespaceTagOutput() RuleGroupsNamespaceTagOutput {
	return i.ToRuleGroupsNamespaceTagOutputWithContext(context.Background())
}

func (i RuleGroupsNamespaceTagArgs) ToRuleGroupsNamespaceTagOutputWithContext(ctx context.Context) RuleGroupsNamespaceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupsNamespaceTagOutput)
}

func (i RuleGroupsNamespaceTagArgs) ToOutput(ctx context.Context) pulumix.Output[RuleGroupsNamespaceTag] {
	return pulumix.Output[RuleGroupsNamespaceTag]{
		OutputState: i.ToRuleGroupsNamespaceTagOutputWithContext(ctx).OutputState,
	}
}

// RuleGroupsNamespaceTagArrayInput is an input type that accepts RuleGroupsNamespaceTagArray and RuleGroupsNamespaceTagArrayOutput values.
// You can construct a concrete instance of `RuleGroupsNamespaceTagArrayInput` via:
//
//	RuleGroupsNamespaceTagArray{ RuleGroupsNamespaceTagArgs{...} }
type RuleGroupsNamespaceTagArrayInput interface {
	pulumi.Input

	ToRuleGroupsNamespaceTagArrayOutput() RuleGroupsNamespaceTagArrayOutput
	ToRuleGroupsNamespaceTagArrayOutputWithContext(context.Context) RuleGroupsNamespaceTagArrayOutput
}

type RuleGroupsNamespaceTagArray []RuleGroupsNamespaceTagInput

func (RuleGroupsNamespaceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleGroupsNamespaceTag)(nil)).Elem()
}

func (i RuleGroupsNamespaceTagArray) ToRuleGroupsNamespaceTagArrayOutput() RuleGroupsNamespaceTagArrayOutput {
	return i.ToRuleGroupsNamespaceTagArrayOutputWithContext(context.Background())
}

func (i RuleGroupsNamespaceTagArray) ToRuleGroupsNamespaceTagArrayOutputWithContext(ctx context.Context) RuleGroupsNamespaceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupsNamespaceTagArrayOutput)
}

func (i RuleGroupsNamespaceTagArray) ToOutput(ctx context.Context) pulumix.Output[[]RuleGroupsNamespaceTag] {
	return pulumix.Output[[]RuleGroupsNamespaceTag]{
		OutputState: i.ToRuleGroupsNamespaceTagArrayOutputWithContext(ctx).OutputState,
	}
}

// A key-value pair to associate with a resource.
type RuleGroupsNamespaceTagOutput struct{ *pulumi.OutputState }

func (RuleGroupsNamespaceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroupsNamespaceTag)(nil)).Elem()
}

func (o RuleGroupsNamespaceTagOutput) ToRuleGroupsNamespaceTagOutput() RuleGroupsNamespaceTagOutput {
	return o
}

func (o RuleGroupsNamespaceTagOutput) ToRuleGroupsNamespaceTagOutputWithContext(ctx context.Context) RuleGroupsNamespaceTagOutput {
	return o
}

func (o RuleGroupsNamespaceTagOutput) ToOutput(ctx context.Context) pulumix.Output[RuleGroupsNamespaceTag] {
	return pulumix.Output[RuleGroupsNamespaceTag]{
		OutputState: o.OutputState,
	}
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o RuleGroupsNamespaceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v RuleGroupsNamespaceTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o RuleGroupsNamespaceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v RuleGroupsNamespaceTag) string { return v.Value }).(pulumi.StringOutput)
}

type RuleGroupsNamespaceTagArrayOutput struct{ *pulumi.OutputState }

func (RuleGroupsNamespaceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleGroupsNamespaceTag)(nil)).Elem()
}

func (o RuleGroupsNamespaceTagArrayOutput) ToRuleGroupsNamespaceTagArrayOutput() RuleGroupsNamespaceTagArrayOutput {
	return o
}

func (o RuleGroupsNamespaceTagArrayOutput) ToRuleGroupsNamespaceTagArrayOutputWithContext(ctx context.Context) RuleGroupsNamespaceTagArrayOutput {
	return o
}

func (o RuleGroupsNamespaceTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]RuleGroupsNamespaceTag] {
	return pulumix.Output[[]RuleGroupsNamespaceTag]{
		OutputState: o.OutputState,
	}
}

func (o RuleGroupsNamespaceTagArrayOutput) Index(i pulumi.IntInput) RuleGroupsNamespaceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleGroupsNamespaceTag {
		return vs[0].([]RuleGroupsNamespaceTag)[vs[1].(int)]
	}).(RuleGroupsNamespaceTagOutput)
}

// Logging configuration
type WorkspaceLoggingConfiguration struct {
	// CloudWatch log group ARN
	LogGroupArn *string `pulumi:"logGroupArn"`
}

// WorkspaceLoggingConfigurationInput is an input type that accepts WorkspaceLoggingConfigurationArgs and WorkspaceLoggingConfigurationOutput values.
// You can construct a concrete instance of `WorkspaceLoggingConfigurationInput` via:
//
//	WorkspaceLoggingConfigurationArgs{...}
type WorkspaceLoggingConfigurationInput interface {
	pulumi.Input

	ToWorkspaceLoggingConfigurationOutput() WorkspaceLoggingConfigurationOutput
	ToWorkspaceLoggingConfigurationOutputWithContext(context.Context) WorkspaceLoggingConfigurationOutput
}

// Logging configuration
type WorkspaceLoggingConfigurationArgs struct {
	// CloudWatch log group ARN
	LogGroupArn pulumi.StringPtrInput `pulumi:"logGroupArn"`
}

func (WorkspaceLoggingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceLoggingConfiguration)(nil)).Elem()
}

func (i WorkspaceLoggingConfigurationArgs) ToWorkspaceLoggingConfigurationOutput() WorkspaceLoggingConfigurationOutput {
	return i.ToWorkspaceLoggingConfigurationOutputWithContext(context.Background())
}

func (i WorkspaceLoggingConfigurationArgs) ToWorkspaceLoggingConfigurationOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceLoggingConfigurationOutput)
}

func (i WorkspaceLoggingConfigurationArgs) ToOutput(ctx context.Context) pulumix.Output[WorkspaceLoggingConfiguration] {
	return pulumix.Output[WorkspaceLoggingConfiguration]{
		OutputState: i.ToWorkspaceLoggingConfigurationOutputWithContext(ctx).OutputState,
	}
}

func (i WorkspaceLoggingConfigurationArgs) ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput {
	return i.ToWorkspaceLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (i WorkspaceLoggingConfigurationArgs) ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceLoggingConfigurationOutput).ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx)
}

// WorkspaceLoggingConfigurationPtrInput is an input type that accepts WorkspaceLoggingConfigurationArgs, WorkspaceLoggingConfigurationPtr and WorkspaceLoggingConfigurationPtrOutput values.
// You can construct a concrete instance of `WorkspaceLoggingConfigurationPtrInput` via:
//
//	        WorkspaceLoggingConfigurationArgs{...}
//
//	or:
//
//	        nil
type WorkspaceLoggingConfigurationPtrInput interface {
	pulumi.Input

	ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput
	ToWorkspaceLoggingConfigurationPtrOutputWithContext(context.Context) WorkspaceLoggingConfigurationPtrOutput
}

type workspaceLoggingConfigurationPtrType WorkspaceLoggingConfigurationArgs

func WorkspaceLoggingConfigurationPtr(v *WorkspaceLoggingConfigurationArgs) WorkspaceLoggingConfigurationPtrInput {
	return (*workspaceLoggingConfigurationPtrType)(v)
}

func (*workspaceLoggingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceLoggingConfiguration)(nil)).Elem()
}

func (i *workspaceLoggingConfigurationPtrType) ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput {
	return i.ToWorkspaceLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (i *workspaceLoggingConfigurationPtrType) ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceLoggingConfigurationPtrOutput)
}

func (i *workspaceLoggingConfigurationPtrType) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceLoggingConfiguration] {
	return pulumix.Output[*WorkspaceLoggingConfiguration]{
		OutputState: i.ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx).OutputState,
	}
}

// Logging configuration
type WorkspaceLoggingConfigurationOutput struct{ *pulumi.OutputState }

func (WorkspaceLoggingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceLoggingConfiguration)(nil)).Elem()
}

func (o WorkspaceLoggingConfigurationOutput) ToWorkspaceLoggingConfigurationOutput() WorkspaceLoggingConfigurationOutput {
	return o
}

func (o WorkspaceLoggingConfigurationOutput) ToWorkspaceLoggingConfigurationOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationOutput {
	return o
}

func (o WorkspaceLoggingConfigurationOutput) ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput {
	return o.ToWorkspaceLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (o WorkspaceLoggingConfigurationOutput) ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceLoggingConfiguration) *WorkspaceLoggingConfiguration {
		return &v
	}).(WorkspaceLoggingConfigurationPtrOutput)
}

func (o WorkspaceLoggingConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[WorkspaceLoggingConfiguration] {
	return pulumix.Output[WorkspaceLoggingConfiguration]{
		OutputState: o.OutputState,
	}
}

// CloudWatch log group ARN
func (o WorkspaceLoggingConfigurationOutput) LogGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceLoggingConfiguration) *string { return v.LogGroupArn }).(pulumi.StringPtrOutput)
}

type WorkspaceLoggingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceLoggingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceLoggingConfiguration)(nil)).Elem()
}

func (o WorkspaceLoggingConfigurationPtrOutput) ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput {
	return o
}

func (o WorkspaceLoggingConfigurationPtrOutput) ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationPtrOutput {
	return o
}

func (o WorkspaceLoggingConfigurationPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceLoggingConfiguration] {
	return pulumix.Output[*WorkspaceLoggingConfiguration]{
		OutputState: o.OutputState,
	}
}

func (o WorkspaceLoggingConfigurationPtrOutput) Elem() WorkspaceLoggingConfigurationOutput {
	return o.ApplyT(func(v *WorkspaceLoggingConfiguration) WorkspaceLoggingConfiguration {
		if v != nil {
			return *v
		}
		var ret WorkspaceLoggingConfiguration
		return ret
	}).(WorkspaceLoggingConfigurationOutput)
}

// CloudWatch log group ARN
func (o WorkspaceLoggingConfigurationPtrOutput) LogGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceLoggingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.LogGroupArn
	}).(pulumi.StringPtrOutput)
}

// A key-value pair to associate with a resource.
type WorkspaceTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// WorkspaceTagInput is an input type that accepts WorkspaceTagArgs and WorkspaceTagOutput values.
// You can construct a concrete instance of `WorkspaceTagInput` via:
//
//	WorkspaceTagArgs{...}
type WorkspaceTagInput interface {
	pulumi.Input

	ToWorkspaceTagOutput() WorkspaceTagOutput
	ToWorkspaceTagOutputWithContext(context.Context) WorkspaceTagOutput
}

// A key-value pair to associate with a resource.
type WorkspaceTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (WorkspaceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceTag)(nil)).Elem()
}

func (i WorkspaceTagArgs) ToWorkspaceTagOutput() WorkspaceTagOutput {
	return i.ToWorkspaceTagOutputWithContext(context.Background())
}

func (i WorkspaceTagArgs) ToWorkspaceTagOutputWithContext(ctx context.Context) WorkspaceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceTagOutput)
}

func (i WorkspaceTagArgs) ToOutput(ctx context.Context) pulumix.Output[WorkspaceTag] {
	return pulumix.Output[WorkspaceTag]{
		OutputState: i.ToWorkspaceTagOutputWithContext(ctx).OutputState,
	}
}

// WorkspaceTagArrayInput is an input type that accepts WorkspaceTagArray and WorkspaceTagArrayOutput values.
// You can construct a concrete instance of `WorkspaceTagArrayInput` via:
//
//	WorkspaceTagArray{ WorkspaceTagArgs{...} }
type WorkspaceTagArrayInput interface {
	pulumi.Input

	ToWorkspaceTagArrayOutput() WorkspaceTagArrayOutput
	ToWorkspaceTagArrayOutputWithContext(context.Context) WorkspaceTagArrayOutput
}

type WorkspaceTagArray []WorkspaceTagInput

func (WorkspaceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkspaceTag)(nil)).Elem()
}

func (i WorkspaceTagArray) ToWorkspaceTagArrayOutput() WorkspaceTagArrayOutput {
	return i.ToWorkspaceTagArrayOutputWithContext(context.Background())
}

func (i WorkspaceTagArray) ToWorkspaceTagArrayOutputWithContext(ctx context.Context) WorkspaceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceTagArrayOutput)
}

func (i WorkspaceTagArray) ToOutput(ctx context.Context) pulumix.Output[[]WorkspaceTag] {
	return pulumix.Output[[]WorkspaceTag]{
		OutputState: i.ToWorkspaceTagArrayOutputWithContext(ctx).OutputState,
	}
}

// A key-value pair to associate with a resource.
type WorkspaceTagOutput struct{ *pulumi.OutputState }

func (WorkspaceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceTag)(nil)).Elem()
}

func (o WorkspaceTagOutput) ToWorkspaceTagOutput() WorkspaceTagOutput {
	return o
}

func (o WorkspaceTagOutput) ToWorkspaceTagOutputWithContext(ctx context.Context) WorkspaceTagOutput {
	return o
}

func (o WorkspaceTagOutput) ToOutput(ctx context.Context) pulumix.Output[WorkspaceTag] {
	return pulumix.Output[WorkspaceTag]{
		OutputState: o.OutputState,
	}
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o WorkspaceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o WorkspaceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceTag) string { return v.Value }).(pulumi.StringOutput)
}

type WorkspaceTagArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkspaceTag)(nil)).Elem()
}

func (o WorkspaceTagArrayOutput) ToWorkspaceTagArrayOutput() WorkspaceTagArrayOutput {
	return o
}

func (o WorkspaceTagArrayOutput) ToWorkspaceTagArrayOutputWithContext(ctx context.Context) WorkspaceTagArrayOutput {
	return o
}

func (o WorkspaceTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]WorkspaceTag] {
	return pulumix.Output[[]WorkspaceTag]{
		OutputState: o.OutputState,
	}
}

func (o WorkspaceTagArrayOutput) Index(i pulumi.IntInput) WorkspaceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkspaceTag {
		return vs[0].([]WorkspaceTag)[vs[1].(int)]
	}).(WorkspaceTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleGroupsNamespaceTagInput)(nil)).Elem(), RuleGroupsNamespaceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleGroupsNamespaceTagArrayInput)(nil)).Elem(), RuleGroupsNamespaceTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceLoggingConfigurationInput)(nil)).Elem(), WorkspaceLoggingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceLoggingConfigurationPtrInput)(nil)).Elem(), WorkspaceLoggingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceTagInput)(nil)).Elem(), WorkspaceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceTagArrayInput)(nil)).Elem(), WorkspaceTagArray{})
	pulumi.RegisterOutputType(RuleGroupsNamespaceTagOutput{})
	pulumi.RegisterOutputType(RuleGroupsNamespaceTagArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceLoggingConfigurationOutput{})
	pulumi.RegisterOutputType(WorkspaceLoggingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceTagOutput{})
	pulumi.RegisterOutputType(WorkspaceTagArrayOutput{})
}
