// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Associates a service with a service network.
type ServiceNetworkServiceAssociation struct {
	pulumi.CustomResourceState

	Arn                      pulumi.StringOutput                               `pulumi:"arn"`
	CreatedAt                pulumi.StringOutput                               `pulumi:"createdAt"`
	DnsEntry                 ServiceNetworkServiceAssociationDnsEntryPtrOutput `pulumi:"dnsEntry"`
	ServiceArn               pulumi.StringOutput                               `pulumi:"serviceArn"`
	ServiceId                pulumi.StringOutput                               `pulumi:"serviceId"`
	ServiceIdentifier        pulumi.StringPtrOutput                            `pulumi:"serviceIdentifier"`
	ServiceName              pulumi.StringOutput                               `pulumi:"serviceName"`
	ServiceNetworkArn        pulumi.StringOutput                               `pulumi:"serviceNetworkArn"`
	ServiceNetworkId         pulumi.StringOutput                               `pulumi:"serviceNetworkId"`
	ServiceNetworkIdentifier pulumi.StringPtrOutput                            `pulumi:"serviceNetworkIdentifier"`
	ServiceNetworkName       pulumi.StringOutput                               `pulumi:"serviceNetworkName"`
	Status                   ServiceNetworkServiceAssociationStatusOutput      `pulumi:"status"`
	Tags                     ServiceNetworkServiceAssociationTagArrayOutput    `pulumi:"tags"`
}

// NewServiceNetworkServiceAssociation registers a new resource with the given unique name, arguments, and options.
func NewServiceNetworkServiceAssociation(ctx *pulumi.Context,
	name string, args *ServiceNetworkServiceAssociationArgs, opts ...pulumi.ResourceOption) (*ServiceNetworkServiceAssociation, error) {
	if args == nil {
		args = &ServiceNetworkServiceAssociationArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"serviceIdentifier",
		"serviceNetworkIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceNetworkServiceAssociation
	err := ctx.RegisterResource("aws-native:vpclattice:ServiceNetworkServiceAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceNetworkServiceAssociation gets an existing ServiceNetworkServiceAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceNetworkServiceAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceNetworkServiceAssociationState, opts ...pulumi.ResourceOption) (*ServiceNetworkServiceAssociation, error) {
	var resource ServiceNetworkServiceAssociation
	err := ctx.ReadResource("aws-native:vpclattice:ServiceNetworkServiceAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceNetworkServiceAssociation resources.
type serviceNetworkServiceAssociationState struct {
}

type ServiceNetworkServiceAssociationState struct {
}

func (ServiceNetworkServiceAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceNetworkServiceAssociationState)(nil)).Elem()
}

type serviceNetworkServiceAssociationArgs struct {
	DnsEntry                 *ServiceNetworkServiceAssociationDnsEntry `pulumi:"dnsEntry"`
	ServiceIdentifier        *string                                   `pulumi:"serviceIdentifier"`
	ServiceNetworkIdentifier *string                                   `pulumi:"serviceNetworkIdentifier"`
	Tags                     []ServiceNetworkServiceAssociationTag     `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceNetworkServiceAssociation resource.
type ServiceNetworkServiceAssociationArgs struct {
	DnsEntry                 ServiceNetworkServiceAssociationDnsEntryPtrInput
	ServiceIdentifier        pulumi.StringPtrInput
	ServiceNetworkIdentifier pulumi.StringPtrInput
	Tags                     ServiceNetworkServiceAssociationTagArrayInput
}

func (ServiceNetworkServiceAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceNetworkServiceAssociationArgs)(nil)).Elem()
}

type ServiceNetworkServiceAssociationInput interface {
	pulumi.Input

	ToServiceNetworkServiceAssociationOutput() ServiceNetworkServiceAssociationOutput
	ToServiceNetworkServiceAssociationOutputWithContext(ctx context.Context) ServiceNetworkServiceAssociationOutput
}

func (*ServiceNetworkServiceAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceNetworkServiceAssociation)(nil)).Elem()
}

func (i *ServiceNetworkServiceAssociation) ToServiceNetworkServiceAssociationOutput() ServiceNetworkServiceAssociationOutput {
	return i.ToServiceNetworkServiceAssociationOutputWithContext(context.Background())
}

func (i *ServiceNetworkServiceAssociation) ToServiceNetworkServiceAssociationOutputWithContext(ctx context.Context) ServiceNetworkServiceAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceNetworkServiceAssociationOutput)
}

func (i *ServiceNetworkServiceAssociation) ToOutput(ctx context.Context) pulumix.Output[*ServiceNetworkServiceAssociation] {
	return pulumix.Output[*ServiceNetworkServiceAssociation]{
		OutputState: i.ToServiceNetworkServiceAssociationOutputWithContext(ctx).OutputState,
	}
}

type ServiceNetworkServiceAssociationOutput struct{ *pulumi.OutputState }

func (ServiceNetworkServiceAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceNetworkServiceAssociation)(nil)).Elem()
}

func (o ServiceNetworkServiceAssociationOutput) ToServiceNetworkServiceAssociationOutput() ServiceNetworkServiceAssociationOutput {
	return o
}

func (o ServiceNetworkServiceAssociationOutput) ToServiceNetworkServiceAssociationOutputWithContext(ctx context.Context) ServiceNetworkServiceAssociationOutput {
	return o
}

func (o ServiceNetworkServiceAssociationOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceNetworkServiceAssociation] {
	return pulumix.Output[*ServiceNetworkServiceAssociation]{
		OutputState: o.OutputState,
	}
}

func (o ServiceNetworkServiceAssociationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ServiceNetworkServiceAssociationOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ServiceNetworkServiceAssociationOutput) DnsEntry() ServiceNetworkServiceAssociationDnsEntryPtrOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) ServiceNetworkServiceAssociationDnsEntryPtrOutput {
		return v.DnsEntry
	}).(ServiceNetworkServiceAssociationDnsEntryPtrOutput)
}

func (o ServiceNetworkServiceAssociationOutput) ServiceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringOutput { return v.ServiceArn }).(pulumi.StringOutput)
}

func (o ServiceNetworkServiceAssociationOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

func (o ServiceNetworkServiceAssociationOutput) ServiceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringPtrOutput { return v.ServiceIdentifier }).(pulumi.StringPtrOutput)
}

func (o ServiceNetworkServiceAssociationOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func (o ServiceNetworkServiceAssociationOutput) ServiceNetworkArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringOutput { return v.ServiceNetworkArn }).(pulumi.StringOutput)
}

func (o ServiceNetworkServiceAssociationOutput) ServiceNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringOutput { return v.ServiceNetworkId }).(pulumi.StringOutput)
}

func (o ServiceNetworkServiceAssociationOutput) ServiceNetworkIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringPtrOutput { return v.ServiceNetworkIdentifier }).(pulumi.StringPtrOutput)
}

func (o ServiceNetworkServiceAssociationOutput) ServiceNetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) pulumi.StringOutput { return v.ServiceNetworkName }).(pulumi.StringOutput)
}

func (o ServiceNetworkServiceAssociationOutput) Status() ServiceNetworkServiceAssociationStatusOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) ServiceNetworkServiceAssociationStatusOutput {
		return v.Status
	}).(ServiceNetworkServiceAssociationStatusOutput)
}

func (o ServiceNetworkServiceAssociationOutput) Tags() ServiceNetworkServiceAssociationTagArrayOutput {
	return o.ApplyT(func(v *ServiceNetworkServiceAssociation) ServiceNetworkServiceAssociationTagArrayOutput {
		return v.Tags
	}).(ServiceNetworkServiceAssociationTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceNetworkServiceAssociationInput)(nil)).Elem(), &ServiceNetworkServiceAssociation{})
	pulumi.RegisterOutputType(ServiceNetworkServiceAssociationOutput{})
}
