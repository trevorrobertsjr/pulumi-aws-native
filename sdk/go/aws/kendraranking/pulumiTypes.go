// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kendraranking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type ExecutionPlanCapacityUnitsConfiguration struct {
	RescoreCapacityUnits int `pulumi:"rescoreCapacityUnits"`
}

// ExecutionPlanCapacityUnitsConfigurationInput is an input type that accepts ExecutionPlanCapacityUnitsConfigurationArgs and ExecutionPlanCapacityUnitsConfigurationOutput values.
// You can construct a concrete instance of `ExecutionPlanCapacityUnitsConfigurationInput` via:
//
//	ExecutionPlanCapacityUnitsConfigurationArgs{...}
type ExecutionPlanCapacityUnitsConfigurationInput interface {
	pulumi.Input

	ToExecutionPlanCapacityUnitsConfigurationOutput() ExecutionPlanCapacityUnitsConfigurationOutput
	ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(context.Context) ExecutionPlanCapacityUnitsConfigurationOutput
}

type ExecutionPlanCapacityUnitsConfigurationArgs struct {
	RescoreCapacityUnits pulumi.IntInput `pulumi:"rescoreCapacityUnits"`
}

func (ExecutionPlanCapacityUnitsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExecutionPlanCapacityUnitsConfiguration)(nil)).Elem()
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToExecutionPlanCapacityUnitsConfigurationOutput() ExecutionPlanCapacityUnitsConfigurationOutput {
	return i.ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(context.Background())
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionPlanCapacityUnitsConfigurationOutput)
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToOutput(ctx context.Context) pulumix.Output[ExecutionPlanCapacityUnitsConfiguration] {
	return pulumix.Output[ExecutionPlanCapacityUnitsConfiguration]{
		OutputState: i.ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(ctx).OutputState,
	}
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return i.ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(context.Background())
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionPlanCapacityUnitsConfigurationOutput).ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx)
}

// ExecutionPlanCapacityUnitsConfigurationPtrInput is an input type that accepts ExecutionPlanCapacityUnitsConfigurationArgs, ExecutionPlanCapacityUnitsConfigurationPtr and ExecutionPlanCapacityUnitsConfigurationPtrOutput values.
// You can construct a concrete instance of `ExecutionPlanCapacityUnitsConfigurationPtrInput` via:
//
//	        ExecutionPlanCapacityUnitsConfigurationArgs{...}
//
//	or:
//
//	        nil
type ExecutionPlanCapacityUnitsConfigurationPtrInput interface {
	pulumi.Input

	ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput
	ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput
}

type executionPlanCapacityUnitsConfigurationPtrType ExecutionPlanCapacityUnitsConfigurationArgs

func ExecutionPlanCapacityUnitsConfigurationPtr(v *ExecutionPlanCapacityUnitsConfigurationArgs) ExecutionPlanCapacityUnitsConfigurationPtrInput {
	return (*executionPlanCapacityUnitsConfigurationPtrType)(v)
}

func (*executionPlanCapacityUnitsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExecutionPlanCapacityUnitsConfiguration)(nil)).Elem()
}

func (i *executionPlanCapacityUnitsConfigurationPtrType) ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return i.ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(context.Background())
}

func (i *executionPlanCapacityUnitsConfigurationPtrType) ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionPlanCapacityUnitsConfigurationPtrOutput)
}

func (i *executionPlanCapacityUnitsConfigurationPtrType) ToOutput(ctx context.Context) pulumix.Output[*ExecutionPlanCapacityUnitsConfiguration] {
	return pulumix.Output[*ExecutionPlanCapacityUnitsConfiguration]{
		OutputState: i.ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx).OutputState,
	}
}

type ExecutionPlanCapacityUnitsConfigurationOutput struct{ *pulumi.OutputState }

func (ExecutionPlanCapacityUnitsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExecutionPlanCapacityUnitsConfiguration)(nil)).Elem()
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToExecutionPlanCapacityUnitsConfigurationOutput() ExecutionPlanCapacityUnitsConfigurationOutput {
	return o
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationOutput {
	return o
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return o.ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(context.Background())
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExecutionPlanCapacityUnitsConfiguration) *ExecutionPlanCapacityUnitsConfiguration {
		return &v
	}).(ExecutionPlanCapacityUnitsConfigurationPtrOutput)
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[ExecutionPlanCapacityUnitsConfiguration] {
	return pulumix.Output[ExecutionPlanCapacityUnitsConfiguration]{
		OutputState: o.OutputState,
	}
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) RescoreCapacityUnits() pulumi.IntOutput {
	return o.ApplyT(func(v ExecutionPlanCapacityUnitsConfiguration) int { return v.RescoreCapacityUnits }).(pulumi.IntOutput)
}

type ExecutionPlanCapacityUnitsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ExecutionPlanCapacityUnitsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExecutionPlanCapacityUnitsConfiguration)(nil)).Elem()
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return o
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return o
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ExecutionPlanCapacityUnitsConfiguration] {
	return pulumix.Output[*ExecutionPlanCapacityUnitsConfiguration]{
		OutputState: o.OutputState,
	}
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) Elem() ExecutionPlanCapacityUnitsConfigurationOutput {
	return o.ApplyT(func(v *ExecutionPlanCapacityUnitsConfiguration) ExecutionPlanCapacityUnitsConfiguration {
		if v != nil {
			return *v
		}
		var ret ExecutionPlanCapacityUnitsConfiguration
		return ret
	}).(ExecutionPlanCapacityUnitsConfigurationOutput)
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) RescoreCapacityUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExecutionPlanCapacityUnitsConfiguration) *int {
		if v == nil {
			return nil
		}
		return &v.RescoreCapacityUnits
	}).(pulumi.IntPtrOutput)
}

// A label for tagging KendraRanking resources
type ExecutionPlanTag struct {
	// A string used to identify this tag
	Key string `pulumi:"key"`
	// A string containing the value for the tag
	Value string `pulumi:"value"`
}

// ExecutionPlanTagInput is an input type that accepts ExecutionPlanTagArgs and ExecutionPlanTagOutput values.
// You can construct a concrete instance of `ExecutionPlanTagInput` via:
//
//	ExecutionPlanTagArgs{...}
type ExecutionPlanTagInput interface {
	pulumi.Input

	ToExecutionPlanTagOutput() ExecutionPlanTagOutput
	ToExecutionPlanTagOutputWithContext(context.Context) ExecutionPlanTagOutput
}

// A label for tagging KendraRanking resources
type ExecutionPlanTagArgs struct {
	// A string used to identify this tag
	Key pulumi.StringInput `pulumi:"key"`
	// A string containing the value for the tag
	Value pulumi.StringInput `pulumi:"value"`
}

func (ExecutionPlanTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExecutionPlanTag)(nil)).Elem()
}

func (i ExecutionPlanTagArgs) ToExecutionPlanTagOutput() ExecutionPlanTagOutput {
	return i.ToExecutionPlanTagOutputWithContext(context.Background())
}

func (i ExecutionPlanTagArgs) ToExecutionPlanTagOutputWithContext(ctx context.Context) ExecutionPlanTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionPlanTagOutput)
}

func (i ExecutionPlanTagArgs) ToOutput(ctx context.Context) pulumix.Output[ExecutionPlanTag] {
	return pulumix.Output[ExecutionPlanTag]{
		OutputState: i.ToExecutionPlanTagOutputWithContext(ctx).OutputState,
	}
}

// ExecutionPlanTagArrayInput is an input type that accepts ExecutionPlanTagArray and ExecutionPlanTagArrayOutput values.
// You can construct a concrete instance of `ExecutionPlanTagArrayInput` via:
//
//	ExecutionPlanTagArray{ ExecutionPlanTagArgs{...} }
type ExecutionPlanTagArrayInput interface {
	pulumi.Input

	ToExecutionPlanTagArrayOutput() ExecutionPlanTagArrayOutput
	ToExecutionPlanTagArrayOutputWithContext(context.Context) ExecutionPlanTagArrayOutput
}

type ExecutionPlanTagArray []ExecutionPlanTagInput

func (ExecutionPlanTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExecutionPlanTag)(nil)).Elem()
}

func (i ExecutionPlanTagArray) ToExecutionPlanTagArrayOutput() ExecutionPlanTagArrayOutput {
	return i.ToExecutionPlanTagArrayOutputWithContext(context.Background())
}

func (i ExecutionPlanTagArray) ToExecutionPlanTagArrayOutputWithContext(ctx context.Context) ExecutionPlanTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionPlanTagArrayOutput)
}

func (i ExecutionPlanTagArray) ToOutput(ctx context.Context) pulumix.Output[[]ExecutionPlanTag] {
	return pulumix.Output[[]ExecutionPlanTag]{
		OutputState: i.ToExecutionPlanTagArrayOutputWithContext(ctx).OutputState,
	}
}

// A label for tagging KendraRanking resources
type ExecutionPlanTagOutput struct{ *pulumi.OutputState }

func (ExecutionPlanTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExecutionPlanTag)(nil)).Elem()
}

func (o ExecutionPlanTagOutput) ToExecutionPlanTagOutput() ExecutionPlanTagOutput {
	return o
}

func (o ExecutionPlanTagOutput) ToExecutionPlanTagOutputWithContext(ctx context.Context) ExecutionPlanTagOutput {
	return o
}

func (o ExecutionPlanTagOutput) ToOutput(ctx context.Context) pulumix.Output[ExecutionPlanTag] {
	return pulumix.Output[ExecutionPlanTag]{
		OutputState: o.OutputState,
	}
}

// A string used to identify this tag
func (o ExecutionPlanTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ExecutionPlanTag) string { return v.Key }).(pulumi.StringOutput)
}

// A string containing the value for the tag
func (o ExecutionPlanTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ExecutionPlanTag) string { return v.Value }).(pulumi.StringOutput)
}

type ExecutionPlanTagArrayOutput struct{ *pulumi.OutputState }

func (ExecutionPlanTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExecutionPlanTag)(nil)).Elem()
}

func (o ExecutionPlanTagArrayOutput) ToExecutionPlanTagArrayOutput() ExecutionPlanTagArrayOutput {
	return o
}

func (o ExecutionPlanTagArrayOutput) ToExecutionPlanTagArrayOutputWithContext(ctx context.Context) ExecutionPlanTagArrayOutput {
	return o
}

func (o ExecutionPlanTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ExecutionPlanTag] {
	return pulumix.Output[[]ExecutionPlanTag]{
		OutputState: o.OutputState,
	}
}

func (o ExecutionPlanTagArrayOutput) Index(i pulumi.IntInput) ExecutionPlanTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExecutionPlanTag {
		return vs[0].([]ExecutionPlanTag)[vs[1].(int)]
	}).(ExecutionPlanTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionPlanCapacityUnitsConfigurationInput)(nil)).Elem(), ExecutionPlanCapacityUnitsConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionPlanCapacityUnitsConfigurationPtrInput)(nil)).Elem(), ExecutionPlanCapacityUnitsConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionPlanTagInput)(nil)).Elem(), ExecutionPlanTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionPlanTagArrayInput)(nil)).Elem(), ExecutionPlanTagArray{})
	pulumi.RegisterOutputType(ExecutionPlanCapacityUnitsConfigurationOutput{})
	pulumi.RegisterOutputType(ExecutionPlanCapacityUnitsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ExecutionPlanTagOutput{})
	pulumi.RegisterOutputType(ExecutionPlanTagArrayOutput{})
}
