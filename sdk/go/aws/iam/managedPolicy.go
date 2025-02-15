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

// Resource Type definition for AWS::IAM::ManagedPolicy
type ManagedPolicy struct {
	pulumi.CustomResourceState

	// The number of entities (users, groups, and roles) that the policy is attached to.
	AttachmentCount pulumi.IntOutput `pulumi:"attachmentCount"`
	// The date and time, in ISO 8601 date-time format, when the policy was created.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// The identifier for the version of the policy that is set as the default version.
	DefaultVersionId pulumi.StringOutput `pulumi:"defaultVersionId"`
	// A friendly description of the policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name (friendly name, not ARN) of the group to attach the policy to.
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// Specifies whether the policy can be attached to an IAM user, group, or role.
	IsAttachable pulumi.BoolOutput `pulumi:"isAttachable"`
	// The friendly name of the policy.
	ManagedPolicyName pulumi.StringPtrOutput `pulumi:"managedPolicyName"`
	// The path for the policy.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The number of entities (users and roles) for which the policy is used to set the permissions boundary.
	PermissionsBoundaryUsageCount pulumi.IntOutput `pulumi:"permissionsBoundaryUsageCount"`
	// Amazon Resource Name (ARN) of the managed policy
	PolicyArn pulumi.StringOutput `pulumi:"policyArn"`
	// The JSON policy document that you want to use as the content for the new policy.
	PolicyDocument pulumi.AnyOutput `pulumi:"policyDocument"`
	// The stable and unique string identifying the policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The name (friendly name, not ARN) of the role to attach the policy to.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The date and time, in ISO 8601 date-time format, when the policy was last updated.
	UpdateDate pulumi.StringOutput `pulumi:"updateDate"`
	// The name (friendly name, not ARN) of the IAM user to attach the policy to.
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewManagedPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagedPolicy(ctx *pulumi.Context,
	name string, args *ManagedPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"description",
		"managedPolicyName",
		"path",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ManagedPolicy
	err := ctx.RegisterResource("aws-native:iam:ManagedPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPolicy gets an existing ManagedPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPolicyState, opts ...pulumi.ResourceOption) (*ManagedPolicy, error) {
	var resource ManagedPolicy
	err := ctx.ReadResource("aws-native:iam:ManagedPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPolicy resources.
type managedPolicyState struct {
}

type ManagedPolicyState struct {
}

func (ManagedPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPolicyState)(nil)).Elem()
}

type managedPolicyArgs struct {
	// A friendly description of the policy.
	Description *string `pulumi:"description"`
	// The name (friendly name, not ARN) of the group to attach the policy to.
	Groups []string `pulumi:"groups"`
	// The friendly name of the policy.
	ManagedPolicyName *string `pulumi:"managedPolicyName"`
	// The path for the policy.
	Path *string `pulumi:"path"`
	// The JSON policy document that you want to use as the content for the new policy.
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// The name (friendly name, not ARN) of the role to attach the policy to.
	Roles []string `pulumi:"roles"`
	// The name (friendly name, not ARN) of the IAM user to attach the policy to.
	Users []string `pulumi:"users"`
}

// The set of arguments for constructing a ManagedPolicy resource.
type ManagedPolicyArgs struct {
	// A friendly description of the policy.
	Description pulumi.StringPtrInput
	// The name (friendly name, not ARN) of the group to attach the policy to.
	Groups pulumi.StringArrayInput
	// The friendly name of the policy.
	ManagedPolicyName pulumi.StringPtrInput
	// The path for the policy.
	Path pulumi.StringPtrInput
	// The JSON policy document that you want to use as the content for the new policy.
	PolicyDocument pulumi.Input
	// The name (friendly name, not ARN) of the role to attach the policy to.
	Roles pulumi.StringArrayInput
	// The name (friendly name, not ARN) of the IAM user to attach the policy to.
	Users pulumi.StringArrayInput
}

func (ManagedPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPolicyArgs)(nil)).Elem()
}

type ManagedPolicyInput interface {
	pulumi.Input

	ToManagedPolicyOutput() ManagedPolicyOutput
	ToManagedPolicyOutputWithContext(ctx context.Context) ManagedPolicyOutput
}

func (*ManagedPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPolicy)(nil)).Elem()
}

func (i *ManagedPolicy) ToManagedPolicyOutput() ManagedPolicyOutput {
	return i.ToManagedPolicyOutputWithContext(context.Background())
}

func (i *ManagedPolicy) ToManagedPolicyOutputWithContext(ctx context.Context) ManagedPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPolicyOutput)
}

func (i *ManagedPolicy) ToOutput(ctx context.Context) pulumix.Output[*ManagedPolicy] {
	return pulumix.Output[*ManagedPolicy]{
		OutputState: i.ToManagedPolicyOutputWithContext(ctx).OutputState,
	}
}

type ManagedPolicyOutput struct{ *pulumi.OutputState }

func (ManagedPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPolicy)(nil)).Elem()
}

func (o ManagedPolicyOutput) ToManagedPolicyOutput() ManagedPolicyOutput {
	return o
}

func (o ManagedPolicyOutput) ToManagedPolicyOutputWithContext(ctx context.Context) ManagedPolicyOutput {
	return o
}

func (o ManagedPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedPolicy] {
	return pulumix.Output[*ManagedPolicy]{
		OutputState: o.OutputState,
	}
}

// The number of entities (users, groups, and roles) that the policy is attached to.
func (o ManagedPolicyOutput) AttachmentCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.IntOutput { return v.AttachmentCount }).(pulumi.IntOutput)
}

// The date and time, in ISO 8601 date-time format, when the policy was created.
func (o ManagedPolicyOutput) CreateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringOutput { return v.CreateDate }).(pulumi.StringOutput)
}

// The identifier for the version of the policy that is set as the default version.
func (o ManagedPolicyOutput) DefaultVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringOutput { return v.DefaultVersionId }).(pulumi.StringOutput)
}

// A friendly description of the policy.
func (o ManagedPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name (friendly name, not ARN) of the group to attach the policy to.
func (o ManagedPolicyOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

// Specifies whether the policy can be attached to an IAM user, group, or role.
func (o ManagedPolicyOutput) IsAttachable() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.BoolOutput { return v.IsAttachable }).(pulumi.BoolOutput)
}

// The friendly name of the policy.
func (o ManagedPolicyOutput) ManagedPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringPtrOutput { return v.ManagedPolicyName }).(pulumi.StringPtrOutput)
}

// The path for the policy.
func (o ManagedPolicyOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The number of entities (users and roles) for which the policy is used to set the permissions boundary.
func (o ManagedPolicyOutput) PermissionsBoundaryUsageCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.IntOutput { return v.PermissionsBoundaryUsageCount }).(pulumi.IntOutput)
}

// Amazon Resource Name (ARN) of the managed policy
func (o ManagedPolicyOutput) PolicyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringOutput { return v.PolicyArn }).(pulumi.StringOutput)
}

// The JSON policy document that you want to use as the content for the new policy.
func (o ManagedPolicyOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

// The stable and unique string identifying the policy.
func (o ManagedPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The name (friendly name, not ARN) of the role to attach the policy to.
func (o ManagedPolicyOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The date and time, in ISO 8601 date-time format, when the policy was last updated.
func (o ManagedPolicyOutput) UpdateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringOutput { return v.UpdateDate }).(pulumi.StringOutput)
}

// The name (friendly name, not ARN) of the IAM user to attach the policy to.
func (o ManagedPolicyOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPolicyInput)(nil)).Elem(), &ManagedPolicy{})
	pulumi.RegisterOutputType(ManagedPolicyOutput{})
}
