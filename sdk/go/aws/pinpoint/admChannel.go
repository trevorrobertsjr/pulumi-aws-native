// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Pinpoint::ADMChannel
//
// Deprecated: AdmChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type AdmChannel struct {
	pulumi.CustomResourceState

	ApplicationId pulumi.StringOutput  `pulumi:"applicationId"`
	ClientId      pulumi.StringOutput  `pulumi:"clientId"`
	ClientSecret  pulumi.StringOutput  `pulumi:"clientSecret"`
	Enabled       pulumi.BoolPtrOutput `pulumi:"enabled"`
}

// NewAdmChannel registers a new resource with the given unique name, arguments, and options.
func NewAdmChannel(ctx *pulumi.Context,
	name string, args *AdmChannelArgs, opts ...pulumi.ResourceOption) (*AdmChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"applicationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AdmChannel
	err := ctx.RegisterResource("aws-native:pinpoint:AdmChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdmChannel gets an existing AdmChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdmChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdmChannelState, opts ...pulumi.ResourceOption) (*AdmChannel, error) {
	var resource AdmChannel
	err := ctx.ReadResource("aws-native:pinpoint:AdmChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdmChannel resources.
type admChannelState struct {
}

type AdmChannelState struct {
}

func (AdmChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*admChannelState)(nil)).Elem()
}

type admChannelArgs struct {
	ApplicationId string `pulumi:"applicationId"`
	ClientId      string `pulumi:"clientId"`
	ClientSecret  string `pulumi:"clientSecret"`
	Enabled       *bool  `pulumi:"enabled"`
}

// The set of arguments for constructing a AdmChannel resource.
type AdmChannelArgs struct {
	ApplicationId pulumi.StringInput
	ClientId      pulumi.StringInput
	ClientSecret  pulumi.StringInput
	Enabled       pulumi.BoolPtrInput
}

func (AdmChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*admChannelArgs)(nil)).Elem()
}

type AdmChannelInput interface {
	pulumi.Input

	ToAdmChannelOutput() AdmChannelOutput
	ToAdmChannelOutputWithContext(ctx context.Context) AdmChannelOutput
}

func (*AdmChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmChannel)(nil)).Elem()
}

func (i *AdmChannel) ToAdmChannelOutput() AdmChannelOutput {
	return i.ToAdmChannelOutputWithContext(context.Background())
}

func (i *AdmChannel) ToAdmChannelOutputWithContext(ctx context.Context) AdmChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmChannelOutput)
}

func (i *AdmChannel) ToOutput(ctx context.Context) pulumix.Output[*AdmChannel] {
	return pulumix.Output[*AdmChannel]{
		OutputState: i.ToAdmChannelOutputWithContext(ctx).OutputState,
	}
}

type AdmChannelOutput struct{ *pulumi.OutputState }

func (AdmChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmChannel)(nil)).Elem()
}

func (o AdmChannelOutput) ToAdmChannelOutput() AdmChannelOutput {
	return o
}

func (o AdmChannelOutput) ToAdmChannelOutputWithContext(ctx context.Context) AdmChannelOutput {
	return o
}

func (o AdmChannelOutput) ToOutput(ctx context.Context) pulumix.Output[*AdmChannel] {
	return pulumix.Output[*AdmChannel]{
		OutputState: o.OutputState,
	}
}

func (o AdmChannelOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdmChannel) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o AdmChannelOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdmChannel) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o AdmChannelOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *AdmChannel) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o AdmChannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdmChannel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdmChannelInput)(nil)).Elem(), &AdmChannel{})
	pulumi.RegisterOutputType(AdmChannelOutput{})
}
