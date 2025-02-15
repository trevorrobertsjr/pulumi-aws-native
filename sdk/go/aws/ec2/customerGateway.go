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

// Resource Type definition for AWS::EC2::CustomerGateway
type CustomerGateway struct {
	pulumi.CustomResourceState

	// For devices that support BGP, the customer gateway's BGP ASN.
	BgpAsn pulumi.IntOutput `pulumi:"bgpAsn"`
	// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
	CustomerGatewayId pulumi.StringOutput `pulumi:"customerGatewayId"`
	// A name for the customer gateway device.
	DeviceName pulumi.StringPtrOutput `pulumi:"deviceName"`
	// The internet-routable IP address for the customer gateway's outside interface. The address must be static.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// One or more tags for the customer gateway.
	Tags CustomerGatewayTagArrayOutput `pulumi:"tags"`
	// The type of VPN connection that this customer gateway supports.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewCustomerGateway(ctx *pulumi.Context,
	name string, args *CustomerGatewayArgs, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BgpAsn == nil {
		return nil, errors.New("invalid value for required argument 'BgpAsn'")
	}
	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bgpAsn",
		"deviceName",
		"ipAddress",
		"type",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomerGateway
	err := ctx.RegisterResource("aws-native:ec2:CustomerGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerGateway gets an existing CustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerGatewayState, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	var resource CustomerGateway
	err := ctx.ReadResource("aws-native:ec2:CustomerGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerGateway resources.
type customerGatewayState struct {
}

type CustomerGatewayState struct {
}

func (CustomerGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayState)(nil)).Elem()
}

type customerGatewayArgs struct {
	// For devices that support BGP, the customer gateway's BGP ASN.
	BgpAsn int `pulumi:"bgpAsn"`
	// A name for the customer gateway device.
	DeviceName *string `pulumi:"deviceName"`
	// The internet-routable IP address for the customer gateway's outside interface. The address must be static.
	IpAddress string `pulumi:"ipAddress"`
	// One or more tags for the customer gateway.
	Tags []CustomerGatewayTag `pulumi:"tags"`
	// The type of VPN connection that this customer gateway supports.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a CustomerGateway resource.
type CustomerGatewayArgs struct {
	// For devices that support BGP, the customer gateway's BGP ASN.
	BgpAsn pulumi.IntInput
	// A name for the customer gateway device.
	DeviceName pulumi.StringPtrInput
	// The internet-routable IP address for the customer gateway's outside interface. The address must be static.
	IpAddress pulumi.StringInput
	// One or more tags for the customer gateway.
	Tags CustomerGatewayTagArrayInput
	// The type of VPN connection that this customer gateway supports.
	Type pulumi.StringInput
}

func (CustomerGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayArgs)(nil)).Elem()
}

type CustomerGatewayInput interface {
	pulumi.Input

	ToCustomerGatewayOutput() CustomerGatewayOutput
	ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput
}

func (*CustomerGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (i *CustomerGateway) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return i.ToCustomerGatewayOutputWithContext(context.Background())
}

func (i *CustomerGateway) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayOutput)
}

func (i *CustomerGateway) ToOutput(ctx context.Context) pulumix.Output[*CustomerGateway] {
	return pulumix.Output[*CustomerGateway]{
		OutputState: i.ToCustomerGatewayOutputWithContext(ctx).OutputState,
	}
}

type CustomerGatewayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomerGateway] {
	return pulumix.Output[*CustomerGateway]{
		OutputState: o.OutputState,
	}
}

// For devices that support BGP, the customer gateway's BGP ASN.
func (o CustomerGatewayOutput) BgpAsn() pulumi.IntOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.IntOutput { return v.BgpAsn }).(pulumi.IntOutput)
}

// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
func (o CustomerGatewayOutput) CustomerGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.CustomerGatewayId }).(pulumi.StringOutput)
}

// A name for the customer gateway device.
func (o CustomerGatewayOutput) DeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringPtrOutput { return v.DeviceName }).(pulumi.StringPtrOutput)
}

// The internet-routable IP address for the customer gateway's outside interface. The address must be static.
func (o CustomerGatewayOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// One or more tags for the customer gateway.
func (o CustomerGatewayOutput) Tags() CustomerGatewayTagArrayOutput {
	return o.ApplyT(func(v *CustomerGateway) CustomerGatewayTagArrayOutput { return v.Tags }).(CustomerGatewayTagArrayOutput)
}

// The type of VPN connection that this customer gateway supports.
func (o CustomerGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayInput)(nil)).Elem(), &CustomerGateway{})
	pulumi.RegisterOutputType(CustomerGatewayOutput{})
}
