// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::SageMaker::Device
type Device struct {
	pulumi.CustomResourceState

	// The Edge Device you want to register against a device fleet
	Device DeviceTypePtrOutput `pulumi:"device"`
	// The name of the edge device fleet
	DeviceFleetName pulumi.StringOutput `pulumi:"deviceFleetName"`
	// Associate tags with the resource
	Tags DeviceTagArrayOutput `pulumi:"tags"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceFleetName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceFleetName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"device.deviceName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Device
	err := ctx.RegisterResource("aws-native:sagemaker:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("aws-native:sagemaker:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
}

type DeviceState struct {
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// The Edge Device you want to register against a device fleet
	Device *DeviceType `pulumi:"device"`
	// The name of the edge device fleet
	DeviceFleetName string `pulumi:"deviceFleetName"`
	// Associate tags with the resource
	Tags []DeviceTag `pulumi:"tags"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// The Edge Device you want to register against a device fleet
	Device DeviceTypePtrInput
	// The name of the edge device fleet
	DeviceFleetName pulumi.StringInput
	// Associate tags with the resource
	Tags DeviceTagArrayInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

func (i *Device) ToOutput(ctx context.Context) pulumix.Output[*Device] {
	return pulumix.Output[*Device]{
		OutputState: i.ToDeviceOutputWithContext(ctx).OutputState,
	}
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

func (o DeviceOutput) ToOutput(ctx context.Context) pulumix.Output[*Device] {
	return pulumix.Output[*Device]{
		OutputState: o.OutputState,
	}
}

// The Edge Device you want to register against a device fleet
func (o DeviceOutput) Device() DeviceTypePtrOutput {
	return o.ApplyT(func(v *Device) DeviceTypePtrOutput { return v.Device }).(DeviceTypePtrOutput)
}

// The name of the edge device fleet
func (o DeviceOutput) DeviceFleetName() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceFleetName }).(pulumi.StringOutput)
}

// Associate tags with the resource
func (o DeviceOutput) Tags() DeviceTagArrayOutput {
	return o.ApplyT(func(v *Device) DeviceTagArrayOutput { return v.Tags }).(DeviceTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceInput)(nil)).Elem(), &Device{})
	pulumi.RegisterOutputType(DeviceOutput{})
}
