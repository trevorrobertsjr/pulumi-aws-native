// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::IoT::TopicRule
type TopicRule struct {
	pulumi.CustomResourceState

	Arn              pulumi.StringOutput     `pulumi:"arn"`
	RuleName         pulumi.StringPtrOutput  `pulumi:"ruleName"`
	Tags             TopicRuleTagArrayOutput `pulumi:"tags"`
	TopicRulePayload TopicRulePayloadOutput  `pulumi:"topicRulePayload"`
}

// NewTopicRule registers a new resource with the given unique name, arguments, and options.
func NewTopicRule(ctx *pulumi.Context,
	name string, args *TopicRuleArgs, opts ...pulumi.ResourceOption) (*TopicRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TopicRulePayload == nil {
		return nil, errors.New("invalid value for required argument 'TopicRulePayload'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"ruleName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TopicRule
	err := ctx.RegisterResource("aws-native:iot:TopicRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicRule gets an existing TopicRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicRuleState, opts ...pulumi.ResourceOption) (*TopicRule, error) {
	var resource TopicRule
	err := ctx.ReadResource("aws-native:iot:TopicRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicRule resources.
type topicRuleState struct {
}

type TopicRuleState struct {
}

func (TopicRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicRuleState)(nil)).Elem()
}

type topicRuleArgs struct {
	RuleName         *string          `pulumi:"ruleName"`
	Tags             []TopicRuleTag   `pulumi:"tags"`
	TopicRulePayload TopicRulePayload `pulumi:"topicRulePayload"`
}

// The set of arguments for constructing a TopicRule resource.
type TopicRuleArgs struct {
	RuleName         pulumi.StringPtrInput
	Tags             TopicRuleTagArrayInput
	TopicRulePayload TopicRulePayloadInput
}

func (TopicRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicRuleArgs)(nil)).Elem()
}

type TopicRuleInput interface {
	pulumi.Input

	ToTopicRuleOutput() TopicRuleOutput
	ToTopicRuleOutputWithContext(ctx context.Context) TopicRuleOutput
}

func (*TopicRule) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicRule)(nil)).Elem()
}

func (i *TopicRule) ToTopicRuleOutput() TopicRuleOutput {
	return i.ToTopicRuleOutputWithContext(context.Background())
}

func (i *TopicRule) ToTopicRuleOutputWithContext(ctx context.Context) TopicRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicRuleOutput)
}

func (i *TopicRule) ToOutput(ctx context.Context) pulumix.Output[*TopicRule] {
	return pulumix.Output[*TopicRule]{
		OutputState: i.ToTopicRuleOutputWithContext(ctx).OutputState,
	}
}

type TopicRuleOutput struct{ *pulumi.OutputState }

func (TopicRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicRule)(nil)).Elem()
}

func (o TopicRuleOutput) ToTopicRuleOutput() TopicRuleOutput {
	return o
}

func (o TopicRuleOutput) ToTopicRuleOutputWithContext(ctx context.Context) TopicRuleOutput {
	return o
}

func (o TopicRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*TopicRule] {
	return pulumix.Output[*TopicRule]{
		OutputState: o.OutputState,
	}
}

func (o TopicRuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicRule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o TopicRuleOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TopicRule) pulumi.StringPtrOutput { return v.RuleName }).(pulumi.StringPtrOutput)
}

func (o TopicRuleOutput) Tags() TopicRuleTagArrayOutput {
	return o.ApplyT(func(v *TopicRule) TopicRuleTagArrayOutput { return v.Tags }).(TopicRuleTagArrayOutput)
}

func (o TopicRuleOutput) TopicRulePayload() TopicRulePayloadOutput {
	return o.ApplyT(func(v *TopicRule) TopicRulePayloadOutput { return v.TopicRulePayload }).(TopicRulePayloadOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicRuleInput)(nil)).Elem(), &TopicRule{})
	pulumi.RegisterOutputType(TopicRuleOutput{})
}
