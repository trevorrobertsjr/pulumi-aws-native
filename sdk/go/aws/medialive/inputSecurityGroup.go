// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::MediaLive::InputSecurityGroup
//
// Deprecated: InputSecurityGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type InputSecurityGroup struct {
	pulumi.CustomResourceState

	Arn            pulumi.StringOutput                                 `pulumi:"arn"`
	Tags           pulumi.AnyOutput                                    `pulumi:"tags"`
	WhitelistRules InputSecurityGroupInputWhitelistRuleCidrArrayOutput `pulumi:"whitelistRules"`
}

// NewInputSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewInputSecurityGroup(ctx *pulumi.Context,
	name string, args *InputSecurityGroupArgs, opts ...pulumi.ResourceOption) (*InputSecurityGroup, error) {
	if args == nil {
		args = &InputSecurityGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InputSecurityGroup
	err := ctx.RegisterResource("aws-native:medialive:InputSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInputSecurityGroup gets an existing InputSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInputSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputSecurityGroupState, opts ...pulumi.ResourceOption) (*InputSecurityGroup, error) {
	var resource InputSecurityGroup
	err := ctx.ReadResource("aws-native:medialive:InputSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InputSecurityGroup resources.
type inputSecurityGroupState struct {
}

type InputSecurityGroupState struct {
}

func (InputSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputSecurityGroupState)(nil)).Elem()
}

type inputSecurityGroupArgs struct {
	Tags           interface{}                                `pulumi:"tags"`
	WhitelistRules []InputSecurityGroupInputWhitelistRuleCidr `pulumi:"whitelistRules"`
}

// The set of arguments for constructing a InputSecurityGroup resource.
type InputSecurityGroupArgs struct {
	Tags           pulumi.Input
	WhitelistRules InputSecurityGroupInputWhitelistRuleCidrArrayInput
}

func (InputSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputSecurityGroupArgs)(nil)).Elem()
}

type InputSecurityGroupInput interface {
	pulumi.Input

	ToInputSecurityGroupOutput() InputSecurityGroupOutput
	ToInputSecurityGroupOutputWithContext(ctx context.Context) InputSecurityGroupOutput
}

func (*InputSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**InputSecurityGroup)(nil)).Elem()
}

func (i *InputSecurityGroup) ToInputSecurityGroupOutput() InputSecurityGroupOutput {
	return i.ToInputSecurityGroupOutputWithContext(context.Background())
}

func (i *InputSecurityGroup) ToInputSecurityGroupOutputWithContext(ctx context.Context) InputSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputSecurityGroupOutput)
}

func (i *InputSecurityGroup) ToOutput(ctx context.Context) pulumix.Output[*InputSecurityGroup] {
	return pulumix.Output[*InputSecurityGroup]{
		OutputState: i.ToInputSecurityGroupOutputWithContext(ctx).OutputState,
	}
}

type InputSecurityGroupOutput struct{ *pulumi.OutputState }

func (InputSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputSecurityGroup)(nil)).Elem()
}

func (o InputSecurityGroupOutput) ToInputSecurityGroupOutput() InputSecurityGroupOutput {
	return o
}

func (o InputSecurityGroupOutput) ToInputSecurityGroupOutputWithContext(ctx context.Context) InputSecurityGroupOutput {
	return o
}

func (o InputSecurityGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*InputSecurityGroup] {
	return pulumix.Output[*InputSecurityGroup]{
		OutputState: o.OutputState,
	}
}

func (o InputSecurityGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *InputSecurityGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o InputSecurityGroupOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *InputSecurityGroup) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o InputSecurityGroupOutput) WhitelistRules() InputSecurityGroupInputWhitelistRuleCidrArrayOutput {
	return o.ApplyT(func(v *InputSecurityGroup) InputSecurityGroupInputWhitelistRuleCidrArrayOutput {
		return v.WhitelistRules
	}).(InputSecurityGroupInputWhitelistRuleCidrArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputSecurityGroupInput)(nil)).Elem(), &InputSecurityGroup{})
	pulumi.RegisterOutputType(InputSecurityGroupOutput{})
}
