// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Glue::Trigger
//
// Deprecated: Trigger is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Trigger struct {
	pulumi.CustomResourceState

	Actions                TriggerActionArrayOutput               `pulumi:"actions"`
	Description            pulumi.StringPtrOutput                 `pulumi:"description"`
	EventBatchingCondition TriggerEventBatchingConditionPtrOutput `pulumi:"eventBatchingCondition"`
	Name                   pulumi.StringPtrOutput                 `pulumi:"name"`
	Predicate              TriggerPredicatePtrOutput              `pulumi:"predicate"`
	Schedule               pulumi.StringPtrOutput                 `pulumi:"schedule"`
	StartOnCreation        pulumi.BoolPtrOutput                   `pulumi:"startOnCreation"`
	Tags                   pulumi.AnyOutput                       `pulumi:"tags"`
	Type                   pulumi.StringOutput                    `pulumi:"type"`
	WorkflowName           pulumi.StringPtrOutput                 `pulumi:"workflowName"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"type",
		"workflowName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Trigger
	err := ctx.RegisterResource("aws-native:glue:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("aws-native:glue:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
}

type TriggerState struct {
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	Actions                []TriggerAction                `pulumi:"actions"`
	Description            *string                        `pulumi:"description"`
	EventBatchingCondition *TriggerEventBatchingCondition `pulumi:"eventBatchingCondition"`
	Name                   *string                        `pulumi:"name"`
	Predicate              *TriggerPredicate              `pulumi:"predicate"`
	Schedule               *string                        `pulumi:"schedule"`
	StartOnCreation        *bool                          `pulumi:"startOnCreation"`
	Tags                   interface{}                    `pulumi:"tags"`
	Type                   string                         `pulumi:"type"`
	WorkflowName           *string                        `pulumi:"workflowName"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	Actions                TriggerActionArrayInput
	Description            pulumi.StringPtrInput
	EventBatchingCondition TriggerEventBatchingConditionPtrInput
	Name                   pulumi.StringPtrInput
	Predicate              TriggerPredicatePtrInput
	Schedule               pulumi.StringPtrInput
	StartOnCreation        pulumi.BoolPtrInput
	Tags                   pulumi.Input
	Type                   pulumi.StringInput
	WorkflowName           pulumi.StringPtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

func (i *Trigger) ToOutput(ctx context.Context) pulumix.Output[*Trigger] {
	return pulumix.Output[*Trigger]{
		OutputState: i.ToTriggerOutputWithContext(ctx).OutputState,
	}
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

func (o TriggerOutput) ToOutput(ctx context.Context) pulumix.Output[*Trigger] {
	return pulumix.Output[*Trigger]{
		OutputState: o.OutputState,
	}
}

func (o TriggerOutput) Actions() TriggerActionArrayOutput {
	return o.ApplyT(func(v *Trigger) TriggerActionArrayOutput { return v.Actions }).(TriggerActionArrayOutput)
}

func (o TriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TriggerOutput) EventBatchingCondition() TriggerEventBatchingConditionPtrOutput {
	return o.ApplyT(func(v *Trigger) TriggerEventBatchingConditionPtrOutput { return v.EventBatchingCondition }).(TriggerEventBatchingConditionPtrOutput)
}

func (o TriggerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TriggerOutput) Predicate() TriggerPredicatePtrOutput {
	return o.ApplyT(func(v *Trigger) TriggerPredicatePtrOutput { return v.Predicate }).(TriggerPredicatePtrOutput)
}

func (o TriggerOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Schedule }).(pulumi.StringPtrOutput)
}

func (o TriggerOutput) StartOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.BoolPtrOutput { return v.StartOnCreation }).(pulumi.BoolPtrOutput)
}

func (o TriggerOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Trigger) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o TriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o TriggerOutput) WorkflowName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.WorkflowName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterOutputType(TriggerOutput{})
}
