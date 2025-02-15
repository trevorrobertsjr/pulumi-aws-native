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

// Resource Type definition for AWS::EC2::VPNConnection
func LookupVpnConnection(ctx *pulumi.Context, args *LookupVpnConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpnConnectionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpnConnectionResult
	err := ctx.Invoke("aws-native:ec2:getVpnConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnConnectionArgs struct {
	// The provider-assigned unique ID for this managed resource
	VpnConnectionId string `pulumi:"vpnConnectionId"`
}

type LookupVpnConnectionResult struct {
	// Any tags assigned to the VPN connection.
	Tags []VpnConnectionTag `pulumi:"tags"`
	// The provider-assigned unique ID for this managed resource
	VpnConnectionId *string `pulumi:"vpnConnectionId"`
}

func LookupVpnConnectionOutput(ctx *pulumi.Context, args LookupVpnConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVpnConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnConnectionResult, error) {
			args := v.(LookupVpnConnectionArgs)
			r, err := LookupVpnConnection(ctx, &args, opts...)
			var s LookupVpnConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnConnectionResultOutput)
}

type LookupVpnConnectionOutputArgs struct {
	// The provider-assigned unique ID for this managed resource
	VpnConnectionId pulumi.StringInput `pulumi:"vpnConnectionId"`
}

func (LookupVpnConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnConnectionArgs)(nil)).Elem()
}

type LookupVpnConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVpnConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnConnectionResult)(nil)).Elem()
}

func (o LookupVpnConnectionResultOutput) ToLookupVpnConnectionResultOutput() LookupVpnConnectionResultOutput {
	return o
}

func (o LookupVpnConnectionResultOutput) ToLookupVpnConnectionResultOutputWithContext(ctx context.Context) LookupVpnConnectionResultOutput {
	return o
}

func (o LookupVpnConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVpnConnectionResult] {
	return pulumix.Output[LookupVpnConnectionResult]{
		OutputState: o.OutputState,
	}
}

// Any tags assigned to the VPN connection.
func (o LookupVpnConnectionResultOutput) Tags() VpnConnectionTagArrayOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) []VpnConnectionTag { return v.Tags }).(VpnConnectionTagArrayOutput)
}

// The provider-assigned unique ID for this managed resource
func (o LookupVpnConnectionResultOutput) VpnConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.VpnConnectionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnConnectionResultOutput{})
}
