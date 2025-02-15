// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// You can use AWS::Organizations::ResourcePolicy to delegate policy management for AWS Organizations to specified member accounts to perform policy actions that are by default available only to the management account.
type ResourcePolicy struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the resource policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The policy document. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
	Content pulumi.AnyOutput `pulumi:"content"`
	// A list of tags that you want to attach to the resource policy
	Tags ResourcePolicyTagArrayOutput `pulumi:"tags"`
}

// NewResourcePolicy registers a new resource with the given unique name, arguments, and options.
func NewResourcePolicy(ctx *pulumi.Context,
	name string, args *ResourcePolicyArgs, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourcePolicy
	err := ctx.RegisterResource("aws-native:organizations:ResourcePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePolicy gets an existing ResourcePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePolicyState, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	var resource ResourcePolicy
	err := ctx.ReadResource("aws-native:organizations:ResourcePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePolicy resources.
type resourcePolicyState struct {
}

type ResourcePolicyState struct {
}

func (ResourcePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyState)(nil)).Elem()
}

type resourcePolicyArgs struct {
	// The policy document. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
	Content interface{} `pulumi:"content"`
	// A list of tags that you want to attach to the resource policy
	Tags []ResourcePolicyTag `pulumi:"tags"`
}

// The set of arguments for constructing a ResourcePolicy resource.
type ResourcePolicyArgs struct {
	// The policy document. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
	Content pulumi.Input
	// A list of tags that you want to attach to the resource policy
	Tags ResourcePolicyTagArrayInput
}

func (ResourcePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyArgs)(nil)).Elem()
}

type ResourcePolicyInput interface {
	pulumi.Input

	ToResourcePolicyOutput() ResourcePolicyOutput
	ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput
}

func (*ResourcePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicy)(nil)).Elem()
}

func (i *ResourcePolicy) ToResourcePolicyOutput() ResourcePolicyOutput {
	return i.ToResourcePolicyOutputWithContext(context.Background())
}

func (i *ResourcePolicy) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyOutput)
}

func (i *ResourcePolicy) ToOutput(ctx context.Context) pulumix.Output[*ResourcePolicy] {
	return pulumix.Output[*ResourcePolicy]{
		OutputState: i.ToResourcePolicyOutputWithContext(ctx).OutputState,
	}
}

type ResourcePolicyOutput struct{ *pulumi.OutputState }

func (ResourcePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicy)(nil)).Elem()
}

func (o ResourcePolicyOutput) ToResourcePolicyOutput() ResourcePolicyOutput {
	return o
}

func (o ResourcePolicyOutput) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return o
}

func (o ResourcePolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ResourcePolicy] {
	return pulumix.Output[*ResourcePolicy]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the resource policy.
func (o ResourcePolicyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePolicy) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The policy document. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
func (o ResourcePolicyOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v *ResourcePolicy) pulumi.AnyOutput { return v.Content }).(pulumi.AnyOutput)
}

// A list of tags that you want to attach to the resource policy
func (o ResourcePolicyOutput) Tags() ResourcePolicyTagArrayOutput {
	return o.ApplyT(func(v *ResourcePolicy) ResourcePolicyTagArrayOutput { return v.Tags }).(ResourcePolicyTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePolicyInput)(nil)).Elem(), &ResourcePolicy{})
	pulumi.RegisterOutputType(ResourcePolicyOutput{})
}
