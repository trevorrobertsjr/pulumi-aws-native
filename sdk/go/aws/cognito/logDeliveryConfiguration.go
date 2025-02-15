// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Cognito::LogDeliveryConfiguration
type LogDeliveryConfiguration struct {
	pulumi.CustomResourceState

	LogConfigurations LogDeliveryConfigurationLogConfigurationArrayOutput `pulumi:"logConfigurations"`
	UserPoolId        pulumi.StringOutput                                 `pulumi:"userPoolId"`
}

// NewLogDeliveryConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLogDeliveryConfiguration(ctx *pulumi.Context,
	name string, args *LogDeliveryConfigurationArgs, opts ...pulumi.ResourceOption) (*LogDeliveryConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserPoolId == nil {
		return nil, errors.New("invalid value for required argument 'UserPoolId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"userPoolId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogDeliveryConfiguration
	err := ctx.RegisterResource("aws-native:cognito:LogDeliveryConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogDeliveryConfiguration gets an existing LogDeliveryConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogDeliveryConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogDeliveryConfigurationState, opts ...pulumi.ResourceOption) (*LogDeliveryConfiguration, error) {
	var resource LogDeliveryConfiguration
	err := ctx.ReadResource("aws-native:cognito:LogDeliveryConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogDeliveryConfiguration resources.
type logDeliveryConfigurationState struct {
}

type LogDeliveryConfigurationState struct {
}

func (LogDeliveryConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*logDeliveryConfigurationState)(nil)).Elem()
}

type logDeliveryConfigurationArgs struct {
	LogConfigurations []LogDeliveryConfigurationLogConfiguration `pulumi:"logConfigurations"`
	UserPoolId        string                                     `pulumi:"userPoolId"`
}

// The set of arguments for constructing a LogDeliveryConfiguration resource.
type LogDeliveryConfigurationArgs struct {
	LogConfigurations LogDeliveryConfigurationLogConfigurationArrayInput
	UserPoolId        pulumi.StringInput
}

func (LogDeliveryConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logDeliveryConfigurationArgs)(nil)).Elem()
}

type LogDeliveryConfigurationInput interface {
	pulumi.Input

	ToLogDeliveryConfigurationOutput() LogDeliveryConfigurationOutput
	ToLogDeliveryConfigurationOutputWithContext(ctx context.Context) LogDeliveryConfigurationOutput
}

func (*LogDeliveryConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDeliveryConfiguration)(nil)).Elem()
}

func (i *LogDeliveryConfiguration) ToLogDeliveryConfigurationOutput() LogDeliveryConfigurationOutput {
	return i.ToLogDeliveryConfigurationOutputWithContext(context.Background())
}

func (i *LogDeliveryConfiguration) ToLogDeliveryConfigurationOutputWithContext(ctx context.Context) LogDeliveryConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDeliveryConfigurationOutput)
}

func (i *LogDeliveryConfiguration) ToOutput(ctx context.Context) pulumix.Output[*LogDeliveryConfiguration] {
	return pulumix.Output[*LogDeliveryConfiguration]{
		OutputState: i.ToLogDeliveryConfigurationOutputWithContext(ctx).OutputState,
	}
}

type LogDeliveryConfigurationOutput struct{ *pulumi.OutputState }

func (LogDeliveryConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDeliveryConfiguration)(nil)).Elem()
}

func (o LogDeliveryConfigurationOutput) ToLogDeliveryConfigurationOutput() LogDeliveryConfigurationOutput {
	return o
}

func (o LogDeliveryConfigurationOutput) ToLogDeliveryConfigurationOutputWithContext(ctx context.Context) LogDeliveryConfigurationOutput {
	return o
}

func (o LogDeliveryConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[*LogDeliveryConfiguration] {
	return pulumix.Output[*LogDeliveryConfiguration]{
		OutputState: o.OutputState,
	}
}

func (o LogDeliveryConfigurationOutput) LogConfigurations() LogDeliveryConfigurationLogConfigurationArrayOutput {
	return o.ApplyT(func(v *LogDeliveryConfiguration) LogDeliveryConfigurationLogConfigurationArrayOutput {
		return v.LogConfigurations
	}).(LogDeliveryConfigurationLogConfigurationArrayOutput)
}

func (o LogDeliveryConfigurationOutput) UserPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDeliveryConfiguration) pulumi.StringOutput { return v.UserPoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogDeliveryConfigurationInput)(nil)).Elem(), &LogDeliveryConfiguration{})
	pulumi.RegisterOutputType(LogDeliveryConfigurationOutput{})
}
