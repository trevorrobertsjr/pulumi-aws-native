// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::ServiceDiscovery::Instance
//
// Deprecated: Instance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Instance struct {
	pulumi.CustomResourceState

	InstanceAttributes pulumi.AnyOutput    `pulumi:"instanceAttributes"`
	InstanceId         pulumi.StringOutput `pulumi:"instanceId"`
	ServiceId          pulumi.StringOutput `pulumi:"serviceId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceAttributes == nil {
		return nil, errors.New("invalid value for required argument 'InstanceAttributes'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"instanceId",
		"serviceId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("aws-native:servicediscovery:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("aws-native:servicediscovery:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	InstanceAttributes interface{} `pulumi:"instanceAttributes"`
	InstanceId         *string     `pulumi:"instanceId"`
	ServiceId          string      `pulumi:"serviceId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	InstanceAttributes pulumi.Input
	InstanceId         pulumi.StringPtrInput
	ServiceId          pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToOutput(ctx context.Context) pulumix.Output[*Instance] {
	return pulumix.Output[*Instance]{
		OutputState: i.ToInstanceOutputWithContext(ctx).OutputState,
	}
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToOutput(ctx context.Context) pulumix.Output[*Instance] {
	return pulumix.Output[*Instance]{
		OutputState: o.OutputState,
	}
}

func (o InstanceOutput) InstanceAttributes() pulumi.AnyOutput {
	return o.ApplyT(func(v *Instance) pulumi.AnyOutput { return v.InstanceAttributes }).(pulumi.AnyOutput)
}

func (o InstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o InstanceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterOutputType(InstanceOutput{})
}
