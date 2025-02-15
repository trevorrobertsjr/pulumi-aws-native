// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::IAM::Group
type Group struct {
	pulumi.CustomResourceState

	// The Arn of the group to create
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the group to create
	GroupName pulumi.StringPtrOutput `pulumi:"groupName"`
	// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
	ManagedPolicyArns pulumi.StringArrayOutput `pulumi:"managedPolicyArns"`
	// The path to the group
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Adds or updates an inline policy document that is embedded in the specified IAM group
	Policies GroupPolicyTypeArrayOutput `pulumi:"policies"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"groupName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("aws-native:iam:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("aws-native:iam:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
}

type GroupState struct {
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The name of the group to create
	GroupName *string `pulumi:"groupName"`
	// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
	ManagedPolicyArns []string `pulumi:"managedPolicyArns"`
	// The path to the group
	Path *string `pulumi:"path"`
	// Adds or updates an inline policy document that is embedded in the specified IAM group
	Policies []GroupPolicyType `pulumi:"policies"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The name of the group to create
	GroupName pulumi.StringPtrInput
	// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
	ManagedPolicyArns pulumi.StringArrayInput
	// The path to the group
	Path pulumi.StringPtrInput
	// Adds or updates an inline policy document that is embedded in the specified IAM group
	Policies GroupPolicyTypeArrayInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

func (i *Group) ToOutput(ctx context.Context) pulumix.Output[*Group] {
	return pulumix.Output[*Group]{
		OutputState: i.ToGroupOutputWithContext(ctx).OutputState,
	}
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) ToOutput(ctx context.Context) pulumix.Output[*Group] {
	return pulumix.Output[*Group]{
		OutputState: o.OutputState,
	}
}

// The Arn of the group to create
func (o GroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the group to create
func (o GroupOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.GroupName }).(pulumi.StringPtrOutput)
}

// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
func (o GroupOutput) ManagedPolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.ManagedPolicyArns }).(pulumi.StringArrayOutput)
}

// The path to the group
func (o GroupOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// Adds or updates an inline policy document that is embedded in the specified IAM group
func (o GroupOutput) Policies() GroupPolicyTypeArrayOutput {
	return o.ApplyT(func(v *Group) GroupPolicyTypeArrayOutput { return v.Policies }).(GroupPolicyTypeArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterOutputType(GroupOutput{})
}
