// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegroups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type GroupConfigurationItem struct {
	Parameters []GroupConfigurationParameter `pulumi:"parameters"`
	Type       *string                       `pulumi:"type"`
}

// GroupConfigurationItemInput is an input type that accepts GroupConfigurationItemArgs and GroupConfigurationItemOutput values.
// You can construct a concrete instance of `GroupConfigurationItemInput` via:
//
//	GroupConfigurationItemArgs{...}
type GroupConfigurationItemInput interface {
	pulumi.Input

	ToGroupConfigurationItemOutput() GroupConfigurationItemOutput
	ToGroupConfigurationItemOutputWithContext(context.Context) GroupConfigurationItemOutput
}

type GroupConfigurationItemArgs struct {
	Parameters GroupConfigurationParameterArrayInput `pulumi:"parameters"`
	Type       pulumi.StringPtrInput                 `pulumi:"type"`
}

func (GroupConfigurationItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConfigurationItem)(nil)).Elem()
}

func (i GroupConfigurationItemArgs) ToGroupConfigurationItemOutput() GroupConfigurationItemOutput {
	return i.ToGroupConfigurationItemOutputWithContext(context.Background())
}

func (i GroupConfigurationItemArgs) ToGroupConfigurationItemOutputWithContext(ctx context.Context) GroupConfigurationItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupConfigurationItemOutput)
}

func (i GroupConfigurationItemArgs) ToOutput(ctx context.Context) pulumix.Output[GroupConfigurationItem] {
	return pulumix.Output[GroupConfigurationItem]{
		OutputState: i.ToGroupConfigurationItemOutputWithContext(ctx).OutputState,
	}
}

// GroupConfigurationItemArrayInput is an input type that accepts GroupConfigurationItemArray and GroupConfigurationItemArrayOutput values.
// You can construct a concrete instance of `GroupConfigurationItemArrayInput` via:
//
//	GroupConfigurationItemArray{ GroupConfigurationItemArgs{...} }
type GroupConfigurationItemArrayInput interface {
	pulumi.Input

	ToGroupConfigurationItemArrayOutput() GroupConfigurationItemArrayOutput
	ToGroupConfigurationItemArrayOutputWithContext(context.Context) GroupConfigurationItemArrayOutput
}

type GroupConfigurationItemArray []GroupConfigurationItemInput

func (GroupConfigurationItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupConfigurationItem)(nil)).Elem()
}

func (i GroupConfigurationItemArray) ToGroupConfigurationItemArrayOutput() GroupConfigurationItemArrayOutput {
	return i.ToGroupConfigurationItemArrayOutputWithContext(context.Background())
}

func (i GroupConfigurationItemArray) ToGroupConfigurationItemArrayOutputWithContext(ctx context.Context) GroupConfigurationItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupConfigurationItemArrayOutput)
}

func (i GroupConfigurationItemArray) ToOutput(ctx context.Context) pulumix.Output[[]GroupConfigurationItem] {
	return pulumix.Output[[]GroupConfigurationItem]{
		OutputState: i.ToGroupConfigurationItemArrayOutputWithContext(ctx).OutputState,
	}
}

type GroupConfigurationItemOutput struct{ *pulumi.OutputState }

func (GroupConfigurationItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConfigurationItem)(nil)).Elem()
}

func (o GroupConfigurationItemOutput) ToGroupConfigurationItemOutput() GroupConfigurationItemOutput {
	return o
}

func (o GroupConfigurationItemOutput) ToGroupConfigurationItemOutputWithContext(ctx context.Context) GroupConfigurationItemOutput {
	return o
}

func (o GroupConfigurationItemOutput) ToOutput(ctx context.Context) pulumix.Output[GroupConfigurationItem] {
	return pulumix.Output[GroupConfigurationItem]{
		OutputState: o.OutputState,
	}
}

func (o GroupConfigurationItemOutput) Parameters() GroupConfigurationParameterArrayOutput {
	return o.ApplyT(func(v GroupConfigurationItem) []GroupConfigurationParameter { return v.Parameters }).(GroupConfigurationParameterArrayOutput)
}

func (o GroupConfigurationItemOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupConfigurationItem) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type GroupConfigurationItemArrayOutput struct{ *pulumi.OutputState }

func (GroupConfigurationItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupConfigurationItem)(nil)).Elem()
}

func (o GroupConfigurationItemArrayOutput) ToGroupConfigurationItemArrayOutput() GroupConfigurationItemArrayOutput {
	return o
}

func (o GroupConfigurationItemArrayOutput) ToGroupConfigurationItemArrayOutputWithContext(ctx context.Context) GroupConfigurationItemArrayOutput {
	return o
}

func (o GroupConfigurationItemArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]GroupConfigurationItem] {
	return pulumix.Output[[]GroupConfigurationItem]{
		OutputState: o.OutputState,
	}
}

func (o GroupConfigurationItemArrayOutput) Index(i pulumi.IntInput) GroupConfigurationItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupConfigurationItem {
		return vs[0].([]GroupConfigurationItem)[vs[1].(int)]
	}).(GroupConfigurationItemOutput)
}

type GroupConfigurationParameter struct {
	Name   *string  `pulumi:"name"`
	Values []string `pulumi:"values"`
}

// GroupConfigurationParameterInput is an input type that accepts GroupConfigurationParameterArgs and GroupConfigurationParameterOutput values.
// You can construct a concrete instance of `GroupConfigurationParameterInput` via:
//
//	GroupConfigurationParameterArgs{...}
type GroupConfigurationParameterInput interface {
	pulumi.Input

	ToGroupConfigurationParameterOutput() GroupConfigurationParameterOutput
	ToGroupConfigurationParameterOutputWithContext(context.Context) GroupConfigurationParameterOutput
}

type GroupConfigurationParameterArgs struct {
	Name   pulumi.StringPtrInput   `pulumi:"name"`
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (GroupConfigurationParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConfigurationParameter)(nil)).Elem()
}

func (i GroupConfigurationParameterArgs) ToGroupConfigurationParameterOutput() GroupConfigurationParameterOutput {
	return i.ToGroupConfigurationParameterOutputWithContext(context.Background())
}

func (i GroupConfigurationParameterArgs) ToGroupConfigurationParameterOutputWithContext(ctx context.Context) GroupConfigurationParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupConfigurationParameterOutput)
}

func (i GroupConfigurationParameterArgs) ToOutput(ctx context.Context) pulumix.Output[GroupConfigurationParameter] {
	return pulumix.Output[GroupConfigurationParameter]{
		OutputState: i.ToGroupConfigurationParameterOutputWithContext(ctx).OutputState,
	}
}

// GroupConfigurationParameterArrayInput is an input type that accepts GroupConfigurationParameterArray and GroupConfigurationParameterArrayOutput values.
// You can construct a concrete instance of `GroupConfigurationParameterArrayInput` via:
//
//	GroupConfigurationParameterArray{ GroupConfigurationParameterArgs{...} }
type GroupConfigurationParameterArrayInput interface {
	pulumi.Input

	ToGroupConfigurationParameterArrayOutput() GroupConfigurationParameterArrayOutput
	ToGroupConfigurationParameterArrayOutputWithContext(context.Context) GroupConfigurationParameterArrayOutput
}

type GroupConfigurationParameterArray []GroupConfigurationParameterInput

func (GroupConfigurationParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupConfigurationParameter)(nil)).Elem()
}

func (i GroupConfigurationParameterArray) ToGroupConfigurationParameterArrayOutput() GroupConfigurationParameterArrayOutput {
	return i.ToGroupConfigurationParameterArrayOutputWithContext(context.Background())
}

func (i GroupConfigurationParameterArray) ToGroupConfigurationParameterArrayOutputWithContext(ctx context.Context) GroupConfigurationParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupConfigurationParameterArrayOutput)
}

func (i GroupConfigurationParameterArray) ToOutput(ctx context.Context) pulumix.Output[[]GroupConfigurationParameter] {
	return pulumix.Output[[]GroupConfigurationParameter]{
		OutputState: i.ToGroupConfigurationParameterArrayOutputWithContext(ctx).OutputState,
	}
}

type GroupConfigurationParameterOutput struct{ *pulumi.OutputState }

func (GroupConfigurationParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConfigurationParameter)(nil)).Elem()
}

func (o GroupConfigurationParameterOutput) ToGroupConfigurationParameterOutput() GroupConfigurationParameterOutput {
	return o
}

func (o GroupConfigurationParameterOutput) ToGroupConfigurationParameterOutputWithContext(ctx context.Context) GroupConfigurationParameterOutput {
	return o
}

func (o GroupConfigurationParameterOutput) ToOutput(ctx context.Context) pulumix.Output[GroupConfigurationParameter] {
	return pulumix.Output[GroupConfigurationParameter]{
		OutputState: o.OutputState,
	}
}

func (o GroupConfigurationParameterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupConfigurationParameter) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GroupConfigurationParameterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupConfigurationParameter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type GroupConfigurationParameterArrayOutput struct{ *pulumi.OutputState }

func (GroupConfigurationParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupConfigurationParameter)(nil)).Elem()
}

func (o GroupConfigurationParameterArrayOutput) ToGroupConfigurationParameterArrayOutput() GroupConfigurationParameterArrayOutput {
	return o
}

func (o GroupConfigurationParameterArrayOutput) ToGroupConfigurationParameterArrayOutputWithContext(ctx context.Context) GroupConfigurationParameterArrayOutput {
	return o
}

func (o GroupConfigurationParameterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]GroupConfigurationParameter] {
	return pulumix.Output[[]GroupConfigurationParameter]{
		OutputState: o.OutputState,
	}
}

func (o GroupConfigurationParameterArrayOutput) Index(i pulumi.IntInput) GroupConfigurationParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupConfigurationParameter {
		return vs[0].([]GroupConfigurationParameter)[vs[1].(int)]
	}).(GroupConfigurationParameterOutput)
}

type GroupQuery struct {
	ResourceTypeFilters []string         `pulumi:"resourceTypeFilters"`
	StackIdentifier     *string          `pulumi:"stackIdentifier"`
	TagFilters          []GroupTagFilter `pulumi:"tagFilters"`
}

// GroupQueryInput is an input type that accepts GroupQueryArgs and GroupQueryOutput values.
// You can construct a concrete instance of `GroupQueryInput` via:
//
//	GroupQueryArgs{...}
type GroupQueryInput interface {
	pulumi.Input

	ToGroupQueryOutput() GroupQueryOutput
	ToGroupQueryOutputWithContext(context.Context) GroupQueryOutput
}

type GroupQueryArgs struct {
	ResourceTypeFilters pulumi.StringArrayInput  `pulumi:"resourceTypeFilters"`
	StackIdentifier     pulumi.StringPtrInput    `pulumi:"stackIdentifier"`
	TagFilters          GroupTagFilterArrayInput `pulumi:"tagFilters"`
}

func (GroupQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupQuery)(nil)).Elem()
}

func (i GroupQueryArgs) ToGroupQueryOutput() GroupQueryOutput {
	return i.ToGroupQueryOutputWithContext(context.Background())
}

func (i GroupQueryArgs) ToGroupQueryOutputWithContext(ctx context.Context) GroupQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupQueryOutput)
}

func (i GroupQueryArgs) ToOutput(ctx context.Context) pulumix.Output[GroupQuery] {
	return pulumix.Output[GroupQuery]{
		OutputState: i.ToGroupQueryOutputWithContext(ctx).OutputState,
	}
}

func (i GroupQueryArgs) ToGroupQueryPtrOutput() GroupQueryPtrOutput {
	return i.ToGroupQueryPtrOutputWithContext(context.Background())
}

func (i GroupQueryArgs) ToGroupQueryPtrOutputWithContext(ctx context.Context) GroupQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupQueryOutput).ToGroupQueryPtrOutputWithContext(ctx)
}

// GroupQueryPtrInput is an input type that accepts GroupQueryArgs, GroupQueryPtr and GroupQueryPtrOutput values.
// You can construct a concrete instance of `GroupQueryPtrInput` via:
//
//	        GroupQueryArgs{...}
//
//	or:
//
//	        nil
type GroupQueryPtrInput interface {
	pulumi.Input

	ToGroupQueryPtrOutput() GroupQueryPtrOutput
	ToGroupQueryPtrOutputWithContext(context.Context) GroupQueryPtrOutput
}

type groupQueryPtrType GroupQueryArgs

func GroupQueryPtr(v *GroupQueryArgs) GroupQueryPtrInput {
	return (*groupQueryPtrType)(v)
}

func (*groupQueryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupQuery)(nil)).Elem()
}

func (i *groupQueryPtrType) ToGroupQueryPtrOutput() GroupQueryPtrOutput {
	return i.ToGroupQueryPtrOutputWithContext(context.Background())
}

func (i *groupQueryPtrType) ToGroupQueryPtrOutputWithContext(ctx context.Context) GroupQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupQueryPtrOutput)
}

func (i *groupQueryPtrType) ToOutput(ctx context.Context) pulumix.Output[*GroupQuery] {
	return pulumix.Output[*GroupQuery]{
		OutputState: i.ToGroupQueryPtrOutputWithContext(ctx).OutputState,
	}
}

type GroupQueryOutput struct{ *pulumi.OutputState }

func (GroupQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupQuery)(nil)).Elem()
}

func (o GroupQueryOutput) ToGroupQueryOutput() GroupQueryOutput {
	return o
}

func (o GroupQueryOutput) ToGroupQueryOutputWithContext(ctx context.Context) GroupQueryOutput {
	return o
}

func (o GroupQueryOutput) ToGroupQueryPtrOutput() GroupQueryPtrOutput {
	return o.ToGroupQueryPtrOutputWithContext(context.Background())
}

func (o GroupQueryOutput) ToGroupQueryPtrOutputWithContext(ctx context.Context) GroupQueryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupQuery) *GroupQuery {
		return &v
	}).(GroupQueryPtrOutput)
}

func (o GroupQueryOutput) ToOutput(ctx context.Context) pulumix.Output[GroupQuery] {
	return pulumix.Output[GroupQuery]{
		OutputState: o.OutputState,
	}
}

func (o GroupQueryOutput) ResourceTypeFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupQuery) []string { return v.ResourceTypeFilters }).(pulumi.StringArrayOutput)
}

func (o GroupQueryOutput) StackIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupQuery) *string { return v.StackIdentifier }).(pulumi.StringPtrOutput)
}

func (o GroupQueryOutput) TagFilters() GroupTagFilterArrayOutput {
	return o.ApplyT(func(v GroupQuery) []GroupTagFilter { return v.TagFilters }).(GroupTagFilterArrayOutput)
}

type GroupQueryPtrOutput struct{ *pulumi.OutputState }

func (GroupQueryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupQuery)(nil)).Elem()
}

func (o GroupQueryPtrOutput) ToGroupQueryPtrOutput() GroupQueryPtrOutput {
	return o
}

func (o GroupQueryPtrOutput) ToGroupQueryPtrOutputWithContext(ctx context.Context) GroupQueryPtrOutput {
	return o
}

func (o GroupQueryPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*GroupQuery] {
	return pulumix.Output[*GroupQuery]{
		OutputState: o.OutputState,
	}
}

func (o GroupQueryPtrOutput) Elem() GroupQueryOutput {
	return o.ApplyT(func(v *GroupQuery) GroupQuery {
		if v != nil {
			return *v
		}
		var ret GroupQuery
		return ret
	}).(GroupQueryOutput)
}

func (o GroupQueryPtrOutput) ResourceTypeFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupQuery) []string {
		if v == nil {
			return nil
		}
		return v.ResourceTypeFilters
	}).(pulumi.StringArrayOutput)
}

func (o GroupQueryPtrOutput) StackIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupQuery) *string {
		if v == nil {
			return nil
		}
		return v.StackIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o GroupQueryPtrOutput) TagFilters() GroupTagFilterArrayOutput {
	return o.ApplyT(func(v *GroupQuery) []GroupTagFilter {
		if v == nil {
			return nil
		}
		return v.TagFilters
	}).(GroupTagFilterArrayOutput)
}

type GroupResourceQuery struct {
	Query *GroupQuery             `pulumi:"query"`
	Type  *GroupResourceQueryType `pulumi:"type"`
}

// GroupResourceQueryInput is an input type that accepts GroupResourceQueryArgs and GroupResourceQueryOutput values.
// You can construct a concrete instance of `GroupResourceQueryInput` via:
//
//	GroupResourceQueryArgs{...}
type GroupResourceQueryInput interface {
	pulumi.Input

	ToGroupResourceQueryOutput() GroupResourceQueryOutput
	ToGroupResourceQueryOutputWithContext(context.Context) GroupResourceQueryOutput
}

type GroupResourceQueryArgs struct {
	Query GroupQueryPtrInput             `pulumi:"query"`
	Type  GroupResourceQueryTypePtrInput `pulumi:"type"`
}

func (GroupResourceQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupResourceQuery)(nil)).Elem()
}

func (i GroupResourceQueryArgs) ToGroupResourceQueryOutput() GroupResourceQueryOutput {
	return i.ToGroupResourceQueryOutputWithContext(context.Background())
}

func (i GroupResourceQueryArgs) ToGroupResourceQueryOutputWithContext(ctx context.Context) GroupResourceQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupResourceQueryOutput)
}

func (i GroupResourceQueryArgs) ToOutput(ctx context.Context) pulumix.Output[GroupResourceQuery] {
	return pulumix.Output[GroupResourceQuery]{
		OutputState: i.ToGroupResourceQueryOutputWithContext(ctx).OutputState,
	}
}

func (i GroupResourceQueryArgs) ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput {
	return i.ToGroupResourceQueryPtrOutputWithContext(context.Background())
}

func (i GroupResourceQueryArgs) ToGroupResourceQueryPtrOutputWithContext(ctx context.Context) GroupResourceQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupResourceQueryOutput).ToGroupResourceQueryPtrOutputWithContext(ctx)
}

// GroupResourceQueryPtrInput is an input type that accepts GroupResourceQueryArgs, GroupResourceQueryPtr and GroupResourceQueryPtrOutput values.
// You can construct a concrete instance of `GroupResourceQueryPtrInput` via:
//
//	        GroupResourceQueryArgs{...}
//
//	or:
//
//	        nil
type GroupResourceQueryPtrInput interface {
	pulumi.Input

	ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput
	ToGroupResourceQueryPtrOutputWithContext(context.Context) GroupResourceQueryPtrOutput
}

type groupResourceQueryPtrType GroupResourceQueryArgs

func GroupResourceQueryPtr(v *GroupResourceQueryArgs) GroupResourceQueryPtrInput {
	return (*groupResourceQueryPtrType)(v)
}

func (*groupResourceQueryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupResourceQuery)(nil)).Elem()
}

func (i *groupResourceQueryPtrType) ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput {
	return i.ToGroupResourceQueryPtrOutputWithContext(context.Background())
}

func (i *groupResourceQueryPtrType) ToGroupResourceQueryPtrOutputWithContext(ctx context.Context) GroupResourceQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupResourceQueryPtrOutput)
}

func (i *groupResourceQueryPtrType) ToOutput(ctx context.Context) pulumix.Output[*GroupResourceQuery] {
	return pulumix.Output[*GroupResourceQuery]{
		OutputState: i.ToGroupResourceQueryPtrOutputWithContext(ctx).OutputState,
	}
}

type GroupResourceQueryOutput struct{ *pulumi.OutputState }

func (GroupResourceQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupResourceQuery)(nil)).Elem()
}

func (o GroupResourceQueryOutput) ToGroupResourceQueryOutput() GroupResourceQueryOutput {
	return o
}

func (o GroupResourceQueryOutput) ToGroupResourceQueryOutputWithContext(ctx context.Context) GroupResourceQueryOutput {
	return o
}

func (o GroupResourceQueryOutput) ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput {
	return o.ToGroupResourceQueryPtrOutputWithContext(context.Background())
}

func (o GroupResourceQueryOutput) ToGroupResourceQueryPtrOutputWithContext(ctx context.Context) GroupResourceQueryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupResourceQuery) *GroupResourceQuery {
		return &v
	}).(GroupResourceQueryPtrOutput)
}

func (o GroupResourceQueryOutput) ToOutput(ctx context.Context) pulumix.Output[GroupResourceQuery] {
	return pulumix.Output[GroupResourceQuery]{
		OutputState: o.OutputState,
	}
}

func (o GroupResourceQueryOutput) Query() GroupQueryPtrOutput {
	return o.ApplyT(func(v GroupResourceQuery) *GroupQuery { return v.Query }).(GroupQueryPtrOutput)
}

func (o GroupResourceQueryOutput) Type() GroupResourceQueryTypePtrOutput {
	return o.ApplyT(func(v GroupResourceQuery) *GroupResourceQueryType { return v.Type }).(GroupResourceQueryTypePtrOutput)
}

type GroupResourceQueryPtrOutput struct{ *pulumi.OutputState }

func (GroupResourceQueryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupResourceQuery)(nil)).Elem()
}

func (o GroupResourceQueryPtrOutput) ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput {
	return o
}

func (o GroupResourceQueryPtrOutput) ToGroupResourceQueryPtrOutputWithContext(ctx context.Context) GroupResourceQueryPtrOutput {
	return o
}

func (o GroupResourceQueryPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*GroupResourceQuery] {
	return pulumix.Output[*GroupResourceQuery]{
		OutputState: o.OutputState,
	}
}

func (o GroupResourceQueryPtrOutput) Elem() GroupResourceQueryOutput {
	return o.ApplyT(func(v *GroupResourceQuery) GroupResourceQuery {
		if v != nil {
			return *v
		}
		var ret GroupResourceQuery
		return ret
	}).(GroupResourceQueryOutput)
}

func (o GroupResourceQueryPtrOutput) Query() GroupQueryPtrOutput {
	return o.ApplyT(func(v *GroupResourceQuery) *GroupQuery {
		if v == nil {
			return nil
		}
		return v.Query
	}).(GroupQueryPtrOutput)
}

func (o GroupResourceQueryPtrOutput) Type() GroupResourceQueryTypePtrOutput {
	return o.ApplyT(func(v *GroupResourceQuery) *GroupResourceQueryType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(GroupResourceQueryTypePtrOutput)
}

type GroupTag struct {
	Key   *string `pulumi:"key"`
	Value *string `pulumi:"value"`
}

// GroupTagInput is an input type that accepts GroupTagArgs and GroupTagOutput values.
// You can construct a concrete instance of `GroupTagInput` via:
//
//	GroupTagArgs{...}
type GroupTagInput interface {
	pulumi.Input

	ToGroupTagOutput() GroupTagOutput
	ToGroupTagOutputWithContext(context.Context) GroupTagOutput
}

type GroupTagArgs struct {
	Key   pulumi.StringPtrInput `pulumi:"key"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (GroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupTag)(nil)).Elem()
}

func (i GroupTagArgs) ToGroupTagOutput() GroupTagOutput {
	return i.ToGroupTagOutputWithContext(context.Background())
}

func (i GroupTagArgs) ToGroupTagOutputWithContext(ctx context.Context) GroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTagOutput)
}

func (i GroupTagArgs) ToOutput(ctx context.Context) pulumix.Output[GroupTag] {
	return pulumix.Output[GroupTag]{
		OutputState: i.ToGroupTagOutputWithContext(ctx).OutputState,
	}
}

// GroupTagArrayInput is an input type that accepts GroupTagArray and GroupTagArrayOutput values.
// You can construct a concrete instance of `GroupTagArrayInput` via:
//
//	GroupTagArray{ GroupTagArgs{...} }
type GroupTagArrayInput interface {
	pulumi.Input

	ToGroupTagArrayOutput() GroupTagArrayOutput
	ToGroupTagArrayOutputWithContext(context.Context) GroupTagArrayOutput
}

type GroupTagArray []GroupTagInput

func (GroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupTag)(nil)).Elem()
}

func (i GroupTagArray) ToGroupTagArrayOutput() GroupTagArrayOutput {
	return i.ToGroupTagArrayOutputWithContext(context.Background())
}

func (i GroupTagArray) ToGroupTagArrayOutputWithContext(ctx context.Context) GroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTagArrayOutput)
}

func (i GroupTagArray) ToOutput(ctx context.Context) pulumix.Output[[]GroupTag] {
	return pulumix.Output[[]GroupTag]{
		OutputState: i.ToGroupTagArrayOutputWithContext(ctx).OutputState,
	}
}

type GroupTagOutput struct{ *pulumi.OutputState }

func (GroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupTag)(nil)).Elem()
}

func (o GroupTagOutput) ToGroupTagOutput() GroupTagOutput {
	return o
}

func (o GroupTagOutput) ToGroupTagOutputWithContext(ctx context.Context) GroupTagOutput {
	return o
}

func (o GroupTagOutput) ToOutput(ctx context.Context) pulumix.Output[GroupTag] {
	return pulumix.Output[GroupTag]{
		OutputState: o.OutputState,
	}
}

func (o GroupTagOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupTag) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o GroupTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type GroupTagArrayOutput struct{ *pulumi.OutputState }

func (GroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupTag)(nil)).Elem()
}

func (o GroupTagArrayOutput) ToGroupTagArrayOutput() GroupTagArrayOutput {
	return o
}

func (o GroupTagArrayOutput) ToGroupTagArrayOutputWithContext(ctx context.Context) GroupTagArrayOutput {
	return o
}

func (o GroupTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]GroupTag] {
	return pulumix.Output[[]GroupTag]{
		OutputState: o.OutputState,
	}
}

func (o GroupTagArrayOutput) Index(i pulumi.IntInput) GroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupTag {
		return vs[0].([]GroupTag)[vs[1].(int)]
	}).(GroupTagOutput)
}

type GroupTagFilter struct {
	Key    *string  `pulumi:"key"`
	Values []string `pulumi:"values"`
}

// GroupTagFilterInput is an input type that accepts GroupTagFilterArgs and GroupTagFilterOutput values.
// You can construct a concrete instance of `GroupTagFilterInput` via:
//
//	GroupTagFilterArgs{...}
type GroupTagFilterInput interface {
	pulumi.Input

	ToGroupTagFilterOutput() GroupTagFilterOutput
	ToGroupTagFilterOutputWithContext(context.Context) GroupTagFilterOutput
}

type GroupTagFilterArgs struct {
	Key    pulumi.StringPtrInput   `pulumi:"key"`
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (GroupTagFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupTagFilter)(nil)).Elem()
}

func (i GroupTagFilterArgs) ToGroupTagFilterOutput() GroupTagFilterOutput {
	return i.ToGroupTagFilterOutputWithContext(context.Background())
}

func (i GroupTagFilterArgs) ToGroupTagFilterOutputWithContext(ctx context.Context) GroupTagFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTagFilterOutput)
}

func (i GroupTagFilterArgs) ToOutput(ctx context.Context) pulumix.Output[GroupTagFilter] {
	return pulumix.Output[GroupTagFilter]{
		OutputState: i.ToGroupTagFilterOutputWithContext(ctx).OutputState,
	}
}

// GroupTagFilterArrayInput is an input type that accepts GroupTagFilterArray and GroupTagFilterArrayOutput values.
// You can construct a concrete instance of `GroupTagFilterArrayInput` via:
//
//	GroupTagFilterArray{ GroupTagFilterArgs{...} }
type GroupTagFilterArrayInput interface {
	pulumi.Input

	ToGroupTagFilterArrayOutput() GroupTagFilterArrayOutput
	ToGroupTagFilterArrayOutputWithContext(context.Context) GroupTagFilterArrayOutput
}

type GroupTagFilterArray []GroupTagFilterInput

func (GroupTagFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupTagFilter)(nil)).Elem()
}

func (i GroupTagFilterArray) ToGroupTagFilterArrayOutput() GroupTagFilterArrayOutput {
	return i.ToGroupTagFilterArrayOutputWithContext(context.Background())
}

func (i GroupTagFilterArray) ToGroupTagFilterArrayOutputWithContext(ctx context.Context) GroupTagFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTagFilterArrayOutput)
}

func (i GroupTagFilterArray) ToOutput(ctx context.Context) pulumix.Output[[]GroupTagFilter] {
	return pulumix.Output[[]GroupTagFilter]{
		OutputState: i.ToGroupTagFilterArrayOutputWithContext(ctx).OutputState,
	}
}

type GroupTagFilterOutput struct{ *pulumi.OutputState }

func (GroupTagFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupTagFilter)(nil)).Elem()
}

func (o GroupTagFilterOutput) ToGroupTagFilterOutput() GroupTagFilterOutput {
	return o
}

func (o GroupTagFilterOutput) ToGroupTagFilterOutputWithContext(ctx context.Context) GroupTagFilterOutput {
	return o
}

func (o GroupTagFilterOutput) ToOutput(ctx context.Context) pulumix.Output[GroupTagFilter] {
	return pulumix.Output[GroupTagFilter]{
		OutputState: o.OutputState,
	}
}

func (o GroupTagFilterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupTagFilter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o GroupTagFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupTagFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type GroupTagFilterArrayOutput struct{ *pulumi.OutputState }

func (GroupTagFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupTagFilter)(nil)).Elem()
}

func (o GroupTagFilterArrayOutput) ToGroupTagFilterArrayOutput() GroupTagFilterArrayOutput {
	return o
}

func (o GroupTagFilterArrayOutput) ToGroupTagFilterArrayOutputWithContext(ctx context.Context) GroupTagFilterArrayOutput {
	return o
}

func (o GroupTagFilterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]GroupTagFilter] {
	return pulumix.Output[[]GroupTagFilter]{
		OutputState: o.OutputState,
	}
}

func (o GroupTagFilterArrayOutput) Index(i pulumi.IntInput) GroupTagFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupTagFilter {
		return vs[0].([]GroupTagFilter)[vs[1].(int)]
	}).(GroupTagFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupConfigurationItemInput)(nil)).Elem(), GroupConfigurationItemArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupConfigurationItemArrayInput)(nil)).Elem(), GroupConfigurationItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupConfigurationParameterInput)(nil)).Elem(), GroupConfigurationParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupConfigurationParameterArrayInput)(nil)).Elem(), GroupConfigurationParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupQueryInput)(nil)).Elem(), GroupQueryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupQueryPtrInput)(nil)).Elem(), GroupQueryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupResourceQueryInput)(nil)).Elem(), GroupResourceQueryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupResourceQueryPtrInput)(nil)).Elem(), GroupResourceQueryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupTagInput)(nil)).Elem(), GroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupTagArrayInput)(nil)).Elem(), GroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupTagFilterInput)(nil)).Elem(), GroupTagFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupTagFilterArrayInput)(nil)).Elem(), GroupTagFilterArray{})
	pulumi.RegisterOutputType(GroupConfigurationItemOutput{})
	pulumi.RegisterOutputType(GroupConfigurationItemArrayOutput{})
	pulumi.RegisterOutputType(GroupConfigurationParameterOutput{})
	pulumi.RegisterOutputType(GroupConfigurationParameterArrayOutput{})
	pulumi.RegisterOutputType(GroupQueryOutput{})
	pulumi.RegisterOutputType(GroupQueryPtrOutput{})
	pulumi.RegisterOutputType(GroupResourceQueryOutput{})
	pulumi.RegisterOutputType(GroupResourceQueryPtrOutput{})
	pulumi.RegisterOutputType(GroupTagOutput{})
	pulumi.RegisterOutputType(GroupTagArrayOutput{})
	pulumi.RegisterOutputType(GroupTagFilterOutput{})
	pulumi.RegisterOutputType(GroupTagFilterArrayOutput{})
}
