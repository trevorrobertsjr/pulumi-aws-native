// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::EC2::TransitGateway
func LookupTransitGateway(ctx *pulumi.Context, args *LookupTransitGatewayArgs, opts ...pulumi.InvokeOption) (*LookupTransitGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTransitGatewayResult
	err := ctx.Invoke("aws-native:ec2:getTransitGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransitGatewayArgs struct {
	Id string `pulumi:"id"`
}

type LookupTransitGatewayResult struct {
	AssociationDefaultRouteTableId *string             `pulumi:"associationDefaultRouteTableId"`
	AutoAcceptSharedAttachments    *string             `pulumi:"autoAcceptSharedAttachments"`
	DefaultRouteTableAssociation   *string             `pulumi:"defaultRouteTableAssociation"`
	DefaultRouteTablePropagation   *string             `pulumi:"defaultRouteTablePropagation"`
	Description                    *string             `pulumi:"description"`
	DnsSupport                     *string             `pulumi:"dnsSupport"`
	Id                             *string             `pulumi:"id"`
	PropagationDefaultRouteTableId *string             `pulumi:"propagationDefaultRouteTableId"`
	Tags                           []TransitGatewayTag `pulumi:"tags"`
	TransitGatewayArn              *string             `pulumi:"transitGatewayArn"`
	TransitGatewayCidrBlocks       []string            `pulumi:"transitGatewayCidrBlocks"`
	VpnEcmpSupport                 *string             `pulumi:"vpnEcmpSupport"`
}

func LookupTransitGatewayOutput(ctx *pulumi.Context, args LookupTransitGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupTransitGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransitGatewayResult, error) {
			args := v.(LookupTransitGatewayArgs)
			r, err := LookupTransitGateway(ctx, &args, opts...)
			var s LookupTransitGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTransitGatewayResultOutput)
}

type LookupTransitGatewayOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTransitGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransitGatewayArgs)(nil)).Elem()
}

type LookupTransitGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupTransitGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransitGatewayResult)(nil)).Elem()
}

func (o LookupTransitGatewayResultOutput) ToLookupTransitGatewayResultOutput() LookupTransitGatewayResultOutput {
	return o
}

func (o LookupTransitGatewayResultOutput) ToLookupTransitGatewayResultOutputWithContext(ctx context.Context) LookupTransitGatewayResultOutput {
	return o
}

func (o LookupTransitGatewayResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTransitGatewayResult] {
	return pulumix.Output[LookupTransitGatewayResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupTransitGatewayResultOutput) AssociationDefaultRouteTableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.AssociationDefaultRouteTableId }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) AutoAcceptSharedAttachments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.AutoAcceptSharedAttachments }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) DefaultRouteTableAssociation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.DefaultRouteTableAssociation }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) DefaultRouteTablePropagation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.DefaultRouteTablePropagation }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) DnsSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.DnsSupport }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) PropagationDefaultRouteTableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.PropagationDefaultRouteTableId }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) Tags() TransitGatewayTagArrayOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) []TransitGatewayTag { return v.Tags }).(TransitGatewayTagArrayOutput)
}

func (o LookupTransitGatewayResultOutput) TransitGatewayArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.TransitGatewayArn }).(pulumi.StringPtrOutput)
}

func (o LookupTransitGatewayResultOutput) TransitGatewayCidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) []string { return v.TransitGatewayCidrBlocks }).(pulumi.StringArrayOutput)
}

func (o LookupTransitGatewayResultOutput) VpnEcmpSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayResult) *string { return v.VpnEcmpSupport }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransitGatewayResultOutput{})
}
