// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::CloudFront::OriginRequestPolicy
type OriginRequestPolicy struct {
	pulumi.CustomResourceState

	LastModifiedTime          pulumi.StringOutput             `pulumi:"lastModifiedTime"`
	OriginRequestPolicyConfig OriginRequestPolicyConfigOutput `pulumi:"originRequestPolicyConfig"`
}

// NewOriginRequestPolicy registers a new resource with the given unique name, arguments, and options.
func NewOriginRequestPolicy(ctx *pulumi.Context,
	name string, args *OriginRequestPolicyArgs, opts ...pulumi.ResourceOption) (*OriginRequestPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OriginRequestPolicyConfig == nil {
		return nil, errors.New("invalid value for required argument 'OriginRequestPolicyConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OriginRequestPolicy
	err := ctx.RegisterResource("aws-native:cloudfront:OriginRequestPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOriginRequestPolicy gets an existing OriginRequestPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOriginRequestPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginRequestPolicyState, opts ...pulumi.ResourceOption) (*OriginRequestPolicy, error) {
	var resource OriginRequestPolicy
	err := ctx.ReadResource("aws-native:cloudfront:OriginRequestPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OriginRequestPolicy resources.
type originRequestPolicyState struct {
}

type OriginRequestPolicyState struct {
}

func (OriginRequestPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*originRequestPolicyState)(nil)).Elem()
}

type originRequestPolicyArgs struct {
	OriginRequestPolicyConfig OriginRequestPolicyConfig `pulumi:"originRequestPolicyConfig"`
}

// The set of arguments for constructing a OriginRequestPolicy resource.
type OriginRequestPolicyArgs struct {
	OriginRequestPolicyConfig OriginRequestPolicyConfigInput
}

func (OriginRequestPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originRequestPolicyArgs)(nil)).Elem()
}

type OriginRequestPolicyInput interface {
	pulumi.Input

	ToOriginRequestPolicyOutput() OriginRequestPolicyOutput
	ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput
}

func (*OriginRequestPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginRequestPolicy)(nil)).Elem()
}

func (i *OriginRequestPolicy) ToOriginRequestPolicyOutput() OriginRequestPolicyOutput {
	return i.ToOriginRequestPolicyOutputWithContext(context.Background())
}

func (i *OriginRequestPolicy) ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginRequestPolicyOutput)
}

func (i *OriginRequestPolicy) ToOutput(ctx context.Context) pulumix.Output[*OriginRequestPolicy] {
	return pulumix.Output[*OriginRequestPolicy]{
		OutputState: i.ToOriginRequestPolicyOutputWithContext(ctx).OutputState,
	}
}

type OriginRequestPolicyOutput struct{ *pulumi.OutputState }

func (OriginRequestPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginRequestPolicy)(nil)).Elem()
}

func (o OriginRequestPolicyOutput) ToOriginRequestPolicyOutput() OriginRequestPolicyOutput {
	return o
}

func (o OriginRequestPolicyOutput) ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput {
	return o
}

func (o OriginRequestPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*OriginRequestPolicy] {
	return pulumix.Output[*OriginRequestPolicy]{
		OutputState: o.OutputState,
	}
}

func (o OriginRequestPolicyOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginRequestPolicy) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o OriginRequestPolicyOutput) OriginRequestPolicyConfig() OriginRequestPolicyConfigOutput {
	return o.ApplyT(func(v *OriginRequestPolicy) OriginRequestPolicyConfigOutput { return v.OriginRequestPolicyConfig }).(OriginRequestPolicyConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OriginRequestPolicyInput)(nil)).Elem(), &OriginRequestPolicy{})
	pulumi.RegisterOutputType(OriginRequestPolicyOutput{})
}
