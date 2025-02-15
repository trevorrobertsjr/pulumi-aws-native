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

// Resource Schema of AWS::EC2::IPAMPoolCidr Type
func LookupIpamPoolCidr(ctx *pulumi.Context, args *LookupIpamPoolCidrArgs, opts ...pulumi.InvokeOption) (*LookupIpamPoolCidrResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpamPoolCidrResult
	err := ctx.Invoke("aws-native:ec2:getIpamPoolCidr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpamPoolCidrArgs struct {
	// Id of the IPAM Pool Cidr.
	IpamPoolCidrId string `pulumi:"ipamPoolCidrId"`
	// Id of the IPAM Pool.
	IpamPoolId string `pulumi:"ipamPoolId"`
}

type LookupIpamPoolCidrResult struct {
	// Id of the IPAM Pool Cidr.
	IpamPoolCidrId *string `pulumi:"ipamPoolCidrId"`
	// Provisioned state of the cidr.
	State *string `pulumi:"state"`
}

func LookupIpamPoolCidrOutput(ctx *pulumi.Context, args LookupIpamPoolCidrOutputArgs, opts ...pulumi.InvokeOption) LookupIpamPoolCidrResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpamPoolCidrResult, error) {
			args := v.(LookupIpamPoolCidrArgs)
			r, err := LookupIpamPoolCidr(ctx, &args, opts...)
			var s LookupIpamPoolCidrResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpamPoolCidrResultOutput)
}

type LookupIpamPoolCidrOutputArgs struct {
	// Id of the IPAM Pool Cidr.
	IpamPoolCidrId pulumi.StringInput `pulumi:"ipamPoolCidrId"`
	// Id of the IPAM Pool.
	IpamPoolId pulumi.StringInput `pulumi:"ipamPoolId"`
}

func (LookupIpamPoolCidrOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpamPoolCidrArgs)(nil)).Elem()
}

type LookupIpamPoolCidrResultOutput struct{ *pulumi.OutputState }

func (LookupIpamPoolCidrResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpamPoolCidrResult)(nil)).Elem()
}

func (o LookupIpamPoolCidrResultOutput) ToLookupIpamPoolCidrResultOutput() LookupIpamPoolCidrResultOutput {
	return o
}

func (o LookupIpamPoolCidrResultOutput) ToLookupIpamPoolCidrResultOutputWithContext(ctx context.Context) LookupIpamPoolCidrResultOutput {
	return o
}

func (o LookupIpamPoolCidrResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIpamPoolCidrResult] {
	return pulumix.Output[LookupIpamPoolCidrResult]{
		OutputState: o.OutputState,
	}
}

// Id of the IPAM Pool Cidr.
func (o LookupIpamPoolCidrResultOutput) IpamPoolCidrId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolCidrResult) *string { return v.IpamPoolCidrId }).(pulumi.StringPtrOutput)
}

// Provisioned state of the cidr.
func (o LookupIpamPoolCidrResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolCidrResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpamPoolCidrResultOutput{})
}
