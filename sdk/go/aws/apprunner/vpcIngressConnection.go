// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apprunner

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The AWS::AppRunner::VpcIngressConnection resource is an App Runner resource that specifies an App Runner VpcIngressConnection.
type VpcIngressConnection struct {
	pulumi.CustomResourceState

	// The Domain name associated with the VPC Ingress Connection.
	DomainName              pulumi.StringOutput                               `pulumi:"domainName"`
	IngressVpcConfiguration VpcIngressConnectionIngressVpcConfigurationOutput `pulumi:"ingressVpcConfiguration"`
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn pulumi.StringOutput `pulumi:"serviceArn"`
	// The current status of the VpcIngressConnection.
	Status VpcIngressConnectionStatusOutput   `pulumi:"status"`
	Tags   VpcIngressConnectionTagArrayOutput `pulumi:"tags"`
	// The Amazon Resource Name (ARN) of the VpcIngressConnection.
	VpcIngressConnectionArn pulumi.StringOutput `pulumi:"vpcIngressConnectionArn"`
	// The customer-provided Vpc Ingress Connection name.
	VpcIngressConnectionName pulumi.StringPtrOutput `pulumi:"vpcIngressConnectionName"`
}

// NewVpcIngressConnection registers a new resource with the given unique name, arguments, and options.
func NewVpcIngressConnection(ctx *pulumi.Context,
	name string, args *VpcIngressConnectionArgs, opts ...pulumi.ResourceOption) (*VpcIngressConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IngressVpcConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'IngressVpcConfiguration'")
	}
	if args.ServiceArn == nil {
		return nil, errors.New("invalid value for required argument 'ServiceArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"serviceArn",
		"tags[*]",
		"vpcIngressConnectionName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcIngressConnection
	err := ctx.RegisterResource("aws-native:apprunner:VpcIngressConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIngressConnection gets an existing VpcIngressConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIngressConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIngressConnectionState, opts ...pulumi.ResourceOption) (*VpcIngressConnection, error) {
	var resource VpcIngressConnection
	err := ctx.ReadResource("aws-native:apprunner:VpcIngressConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIngressConnection resources.
type vpcIngressConnectionState struct {
}

type VpcIngressConnectionState struct {
}

func (VpcIngressConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIngressConnectionState)(nil)).Elem()
}

type vpcIngressConnectionArgs struct {
	IngressVpcConfiguration VpcIngressConnectionIngressVpcConfiguration `pulumi:"ingressVpcConfiguration"`
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn string                    `pulumi:"serviceArn"`
	Tags       []VpcIngressConnectionTag `pulumi:"tags"`
	// The customer-provided Vpc Ingress Connection name.
	VpcIngressConnectionName *string `pulumi:"vpcIngressConnectionName"`
}

// The set of arguments for constructing a VpcIngressConnection resource.
type VpcIngressConnectionArgs struct {
	IngressVpcConfiguration VpcIngressConnectionIngressVpcConfigurationInput
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn pulumi.StringInput
	Tags       VpcIngressConnectionTagArrayInput
	// The customer-provided Vpc Ingress Connection name.
	VpcIngressConnectionName pulumi.StringPtrInput
}

func (VpcIngressConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIngressConnectionArgs)(nil)).Elem()
}

type VpcIngressConnectionInput interface {
	pulumi.Input

	ToVpcIngressConnectionOutput() VpcIngressConnectionOutput
	ToVpcIngressConnectionOutputWithContext(ctx context.Context) VpcIngressConnectionOutput
}

func (*VpcIngressConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIngressConnection)(nil)).Elem()
}

func (i *VpcIngressConnection) ToVpcIngressConnectionOutput() VpcIngressConnectionOutput {
	return i.ToVpcIngressConnectionOutputWithContext(context.Background())
}

func (i *VpcIngressConnection) ToVpcIngressConnectionOutputWithContext(ctx context.Context) VpcIngressConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIngressConnectionOutput)
}

func (i *VpcIngressConnection) ToOutput(ctx context.Context) pulumix.Output[*VpcIngressConnection] {
	return pulumix.Output[*VpcIngressConnection]{
		OutputState: i.ToVpcIngressConnectionOutputWithContext(ctx).OutputState,
	}
}

type VpcIngressConnectionOutput struct{ *pulumi.OutputState }

func (VpcIngressConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIngressConnection)(nil)).Elem()
}

func (o VpcIngressConnectionOutput) ToVpcIngressConnectionOutput() VpcIngressConnectionOutput {
	return o
}

func (o VpcIngressConnectionOutput) ToVpcIngressConnectionOutputWithContext(ctx context.Context) VpcIngressConnectionOutput {
	return o
}

func (o VpcIngressConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*VpcIngressConnection] {
	return pulumix.Output[*VpcIngressConnection]{
		OutputState: o.OutputState,
	}
}

// The Domain name associated with the VPC Ingress Connection.
func (o VpcIngressConnectionOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIngressConnection) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o VpcIngressConnectionOutput) IngressVpcConfiguration() VpcIngressConnectionIngressVpcConfigurationOutput {
	return o.ApplyT(func(v *VpcIngressConnection) VpcIngressConnectionIngressVpcConfigurationOutput {
		return v.IngressVpcConfiguration
	}).(VpcIngressConnectionIngressVpcConfigurationOutput)
}

// The Amazon Resource Name (ARN) of the service.
func (o VpcIngressConnectionOutput) ServiceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIngressConnection) pulumi.StringOutput { return v.ServiceArn }).(pulumi.StringOutput)
}

// The current status of the VpcIngressConnection.
func (o VpcIngressConnectionOutput) Status() VpcIngressConnectionStatusOutput {
	return o.ApplyT(func(v *VpcIngressConnection) VpcIngressConnectionStatusOutput { return v.Status }).(VpcIngressConnectionStatusOutput)
}

func (o VpcIngressConnectionOutput) Tags() VpcIngressConnectionTagArrayOutput {
	return o.ApplyT(func(v *VpcIngressConnection) VpcIngressConnectionTagArrayOutput { return v.Tags }).(VpcIngressConnectionTagArrayOutput)
}

// The Amazon Resource Name (ARN) of the VpcIngressConnection.
func (o VpcIngressConnectionOutput) VpcIngressConnectionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIngressConnection) pulumi.StringOutput { return v.VpcIngressConnectionArn }).(pulumi.StringOutput)
}

// The customer-provided Vpc Ingress Connection name.
func (o VpcIngressConnectionOutput) VpcIngressConnectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIngressConnection) pulumi.StringPtrOutput { return v.VpcIngressConnectionName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIngressConnectionInput)(nil)).Elem(), &VpcIngressConnection{})
	pulumi.RegisterOutputType(VpcIngressConnectionOutput{})
}
