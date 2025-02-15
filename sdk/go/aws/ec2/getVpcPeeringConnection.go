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

// Resource Type definition for AWS::EC2::VPCPeeringConnection
func LookupVpcPeeringConnection(ctx *pulumi.Context, args *LookupVpcPeeringConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpcPeeringConnectionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcPeeringConnectionResult
	err := ctx.Invoke("aws-native:ec2:getVpcPeeringConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcPeeringConnectionArgs struct {
	Id string `pulumi:"id"`
}

type LookupVpcPeeringConnectionResult struct {
	Id   *string                   `pulumi:"id"`
	Tags []VpcPeeringConnectionTag `pulumi:"tags"`
}

func LookupVpcPeeringConnectionOutput(ctx *pulumi.Context, args LookupVpcPeeringConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPeeringConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcPeeringConnectionResult, error) {
			args := v.(LookupVpcPeeringConnectionArgs)
			r, err := LookupVpcPeeringConnection(ctx, &args, opts...)
			var s LookupVpcPeeringConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcPeeringConnectionResultOutput)
}

type LookupVpcPeeringConnectionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVpcPeeringConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPeeringConnectionArgs)(nil)).Elem()
}

type LookupVpcPeeringConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPeeringConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPeeringConnectionResult)(nil)).Elem()
}

func (o LookupVpcPeeringConnectionResultOutput) ToLookupVpcPeeringConnectionResultOutput() LookupVpcPeeringConnectionResultOutput {
	return o
}

func (o LookupVpcPeeringConnectionResultOutput) ToLookupVpcPeeringConnectionResultOutputWithContext(ctx context.Context) LookupVpcPeeringConnectionResultOutput {
	return o
}

func (o LookupVpcPeeringConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVpcPeeringConnectionResult] {
	return pulumix.Output[LookupVpcPeeringConnectionResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupVpcPeeringConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) Tags() VpcPeeringConnectionTagArrayOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) []VpcPeeringConnectionTag { return v.Tags }).(VpcPeeringConnectionTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPeeringConnectionResultOutput{})
}
