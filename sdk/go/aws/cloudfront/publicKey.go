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

// Resource Type definition for AWS::CloudFront::PublicKey
type PublicKey struct {
	pulumi.CustomResourceState

	CreatedTime     pulumi.StringOutput   `pulumi:"createdTime"`
	PublicKeyConfig PublicKeyConfigOutput `pulumi:"publicKeyConfig"`
}

// NewPublicKey registers a new resource with the given unique name, arguments, and options.
func NewPublicKey(ctx *pulumi.Context,
	name string, args *PublicKeyArgs, opts ...pulumi.ResourceOption) (*PublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicKeyConfig == nil {
		return nil, errors.New("invalid value for required argument 'PublicKeyConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicKey
	err := ctx.RegisterResource("aws-native:cloudfront:PublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicKey gets an existing PublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicKeyState, opts ...pulumi.ResourceOption) (*PublicKey, error) {
	var resource PublicKey
	err := ctx.ReadResource("aws-native:cloudfront:PublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicKey resources.
type publicKeyState struct {
}

type PublicKeyState struct {
}

func (PublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicKeyState)(nil)).Elem()
}

type publicKeyArgs struct {
	PublicKeyConfig PublicKeyConfig `pulumi:"publicKeyConfig"`
}

// The set of arguments for constructing a PublicKey resource.
type PublicKeyArgs struct {
	PublicKeyConfig PublicKeyConfigInput
}

func (PublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicKeyArgs)(nil)).Elem()
}

type PublicKeyInput interface {
	pulumi.Input

	ToPublicKeyOutput() PublicKeyOutput
	ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput
}

func (*PublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicKey)(nil)).Elem()
}

func (i *PublicKey) ToPublicKeyOutput() PublicKeyOutput {
	return i.ToPublicKeyOutputWithContext(context.Background())
}

func (i *PublicKey) ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicKeyOutput)
}

func (i *PublicKey) ToOutput(ctx context.Context) pulumix.Output[*PublicKey] {
	return pulumix.Output[*PublicKey]{
		OutputState: i.ToPublicKeyOutputWithContext(ctx).OutputState,
	}
}

type PublicKeyOutput struct{ *pulumi.OutputState }

func (PublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicKey)(nil)).Elem()
}

func (o PublicKeyOutput) ToPublicKeyOutput() PublicKeyOutput {
	return o
}

func (o PublicKeyOutput) ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput {
	return o
}

func (o PublicKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*PublicKey] {
	return pulumix.Output[*PublicKey]{
		OutputState: o.OutputState,
	}
}

func (o PublicKeyOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicKey) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o PublicKeyOutput) PublicKeyConfig() PublicKeyConfigOutput {
	return o.ApplyT(func(v *PublicKey) PublicKeyConfigOutput { return v.PublicKeyConfig }).(PublicKeyConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicKeyInput)(nil)).Elem(), &PublicKey{})
	pulumi.RegisterOutputType(PublicKeyOutput{})
}
