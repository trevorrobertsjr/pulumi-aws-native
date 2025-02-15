// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::EC2::VolumeAttachment
type VolumeAttachment struct {
	pulumi.CustomResourceState

	Device     pulumi.StringPtrOutput `pulumi:"device"`
	InstanceId pulumi.StringOutput    `pulumi:"instanceId"`
	VolumeId   pulumi.StringOutput    `pulumi:"volumeId"`
}

// NewVolumeAttachment registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttachment(ctx *pulumi.Context,
	name string, args *VolumeAttachmentArgs, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"device",
		"instanceId",
		"volumeId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VolumeAttachment
	err := ctx.RegisterResource("aws-native:ec2:VolumeAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttachment gets an existing VolumeAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeAttachmentState, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	var resource VolumeAttachment
	err := ctx.ReadResource("aws-native:ec2:VolumeAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeAttachment resources.
type volumeAttachmentState struct {
}

type VolumeAttachmentState struct {
}

func (VolumeAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentState)(nil)).Elem()
}

type volumeAttachmentArgs struct {
	Device     *string `pulumi:"device"`
	InstanceId string  `pulumi:"instanceId"`
	VolumeId   string  `pulumi:"volumeId"`
}

// The set of arguments for constructing a VolumeAttachment resource.
type VolumeAttachmentArgs struct {
	Device     pulumi.StringPtrInput
	InstanceId pulumi.StringInput
	VolumeId   pulumi.StringInput
}

func (VolumeAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentArgs)(nil)).Elem()
}

type VolumeAttachmentInput interface {
	pulumi.Input

	ToVolumeAttachmentOutput() VolumeAttachmentOutput
	ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput
}

func (*VolumeAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachment)(nil)).Elem()
}

func (i *VolumeAttachment) ToVolumeAttachmentOutput() VolumeAttachmentOutput {
	return i.ToVolumeAttachmentOutputWithContext(context.Background())
}

func (i *VolumeAttachment) ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentOutput)
}

func (i *VolumeAttachment) ToOutput(ctx context.Context) pulumix.Output[*VolumeAttachment] {
	return pulumix.Output[*VolumeAttachment]{
		OutputState: i.ToVolumeAttachmentOutputWithContext(ctx).OutputState,
	}
}

type VolumeAttachmentOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachment)(nil)).Elem()
}

func (o VolumeAttachmentOutput) ToVolumeAttachmentOutput() VolumeAttachmentOutput {
	return o
}

func (o VolumeAttachmentOutput) ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput {
	return o
}

func (o VolumeAttachmentOutput) ToOutput(ctx context.Context) pulumix.Output[*VolumeAttachment] {
	return pulumix.Output[*VolumeAttachment]{
		OutputState: o.OutputState,
	}
}

func (o VolumeAttachmentOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachment) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

func (o VolumeAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o VolumeAttachmentOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachment) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentInput)(nil)).Elem(), &VolumeAttachment{})
	pulumi.RegisterOutputType(VolumeAttachmentOutput{})
}
