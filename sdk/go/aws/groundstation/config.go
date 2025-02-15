// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package groundstation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// AWS Ground Station config resource type for CloudFormation.
type Config struct {
	pulumi.CustomResourceState

	Arn        pulumi.StringOutput  `pulumi:"arn"`
	ConfigData ConfigDataOutput     `pulumi:"configData"`
	Name       pulumi.StringOutput  `pulumi:"name"`
	Tags       ConfigTagArrayOutput `pulumi:"tags"`
	Type       pulumi.StringOutput  `pulumi:"type"`
}

// NewConfig registers a new resource with the given unique name, arguments, and options.
func NewConfig(ctx *pulumi.Context,
	name string, args *ConfigArgs, opts ...pulumi.ResourceOption) (*Config, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigData == nil {
		return nil, errors.New("invalid value for required argument 'ConfigData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Config
	err := ctx.RegisterResource("aws-native:groundstation:Config", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfig gets an existing Config resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigState, opts ...pulumi.ResourceOption) (*Config, error) {
	var resource Config
	err := ctx.ReadResource("aws-native:groundstation:Config", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Config resources.
type configState struct {
}

type ConfigState struct {
}

func (ConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*configState)(nil)).Elem()
}

type configArgs struct {
	ConfigData ConfigData  `pulumi:"configData"`
	Name       *string     `pulumi:"name"`
	Tags       []ConfigTag `pulumi:"tags"`
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	ConfigData ConfigDataInput
	Name       pulumi.StringPtrInput
	Tags       ConfigTagArrayInput
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configArgs)(nil)).Elem()
}

type ConfigInput interface {
	pulumi.Input

	ToConfigOutput() ConfigOutput
	ToConfigOutputWithContext(ctx context.Context) ConfigOutput
}

func (*Config) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil)).Elem()
}

func (i *Config) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i *Config) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

func (i *Config) ToOutput(ctx context.Context) pulumix.Output[*Config] {
	return pulumix.Output[*Config]{
		OutputState: i.ToConfigOutputWithContext(ctx).OutputState,
	}
}

type ConfigOutput struct{ *pulumi.OutputState }

func (ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil)).Elem()
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

func (o ConfigOutput) ToOutput(ctx context.Context) pulumix.Output[*Config] {
	return pulumix.Output[*Config]{
		OutputState: o.OutputState,
	}
}

func (o ConfigOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Config) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ConfigOutput) ConfigData() ConfigDataOutput {
	return o.ApplyT(func(v *Config) ConfigDataOutput { return v.ConfigData }).(ConfigDataOutput)
}

func (o ConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Config) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigOutput) Tags() ConfigTagArrayOutput {
	return o.ApplyT(func(v *Config) ConfigTagArrayOutput { return v.Tags }).(ConfigTagArrayOutput)
}

func (o ConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Config) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigInput)(nil)).Elem(), &Config{})
	pulumi.RegisterOutputType(ConfigOutput{})
}
