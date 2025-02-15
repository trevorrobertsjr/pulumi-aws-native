// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Schema for IAM Group Policy
type GroupPolicy struct {
	pulumi.CustomResourceState

	// The name of the group to associate the policy with.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The policy document.
	PolicyDocument pulumi.AnyOutput `pulumi:"policyDocument"`
	// The name of the policy document.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
}

// NewGroupPolicy registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicy(ctx *pulumi.Context,
	name string, args *GroupPolicyArgs, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"groupName",
		"policyName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupPolicy
	err := ctx.RegisterResource("aws-native:iam:GroupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicy gets an existing GroupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPolicyState, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	var resource GroupPolicy
	err := ctx.ReadResource("aws-native:iam:GroupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicy resources.
type groupPolicyState struct {
}

type GroupPolicyState struct {
}

func (GroupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyState)(nil)).Elem()
}

type groupPolicyArgs struct {
	// The name of the group to associate the policy with.
	GroupName string `pulumi:"groupName"`
	// The policy document.
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// The name of the policy document.
	PolicyName string `pulumi:"policyName"`
}

// The set of arguments for constructing a GroupPolicy resource.
type GroupPolicyArgs struct {
	// The name of the group to associate the policy with.
	GroupName pulumi.StringInput
	// The policy document.
	PolicyDocument pulumi.Input
	// The name of the policy document.
	PolicyName pulumi.StringInput
}

func (GroupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyArgs)(nil)).Elem()
}

type GroupPolicyInput interface {
	pulumi.Input

	ToGroupPolicyOutput() GroupPolicyOutput
	ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput
}

func (*GroupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicy)(nil)).Elem()
}

func (i *GroupPolicy) ToGroupPolicyOutput() GroupPolicyOutput {
	return i.ToGroupPolicyOutputWithContext(context.Background())
}

func (i *GroupPolicy) ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyOutput)
}

func (i *GroupPolicy) ToOutput(ctx context.Context) pulumix.Output[*GroupPolicy] {
	return pulumix.Output[*GroupPolicy]{
		OutputState: i.ToGroupPolicyOutputWithContext(ctx).OutputState,
	}
}

type GroupPolicyOutput struct{ *pulumi.OutputState }

func (GroupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicy)(nil)).Elem()
}

func (o GroupPolicyOutput) ToGroupPolicyOutput() GroupPolicyOutput {
	return o
}

func (o GroupPolicyOutput) ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput {
	return o
}

func (o GroupPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*GroupPolicy] {
	return pulumix.Output[*GroupPolicy]{
		OutputState: o.OutputState,
	}
}

// The name of the group to associate the policy with.
func (o GroupPolicyOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicy) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// The policy document.
func (o GroupPolicyOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *GroupPolicy) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

// The name of the policy document.
func (o GroupPolicyOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicy) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPolicyInput)(nil)).Elem(), &GroupPolicy{})
	pulumi.RegisterOutputType(GroupPolicyOutput{})
}
