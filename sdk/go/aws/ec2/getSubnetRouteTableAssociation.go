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

// Resource Type definition for AWS::EC2::SubnetRouteTableAssociation
func LookupSubnetRouteTableAssociation(ctx *pulumi.Context, args *LookupSubnetRouteTableAssociationArgs, opts ...pulumi.InvokeOption) (*LookupSubnetRouteTableAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSubnetRouteTableAssociationResult
	err := ctx.Invoke("aws-native:ec2:getSubnetRouteTableAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubnetRouteTableAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupSubnetRouteTableAssociationResult struct {
	Id *string `pulumi:"id"`
}

func LookupSubnetRouteTableAssociationOutput(ctx *pulumi.Context, args LookupSubnetRouteTableAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetRouteTableAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetRouteTableAssociationResult, error) {
			args := v.(LookupSubnetRouteTableAssociationArgs)
			r, err := LookupSubnetRouteTableAssociation(ctx, &args, opts...)
			var s LookupSubnetRouteTableAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetRouteTableAssociationResultOutput)
}

type LookupSubnetRouteTableAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSubnetRouteTableAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetRouteTableAssociationArgs)(nil)).Elem()
}

type LookupSubnetRouteTableAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetRouteTableAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetRouteTableAssociationResult)(nil)).Elem()
}

func (o LookupSubnetRouteTableAssociationResultOutput) ToLookupSubnetRouteTableAssociationResultOutput() LookupSubnetRouteTableAssociationResultOutput {
	return o
}

func (o LookupSubnetRouteTableAssociationResultOutput) ToLookupSubnetRouteTableAssociationResultOutputWithContext(ctx context.Context) LookupSubnetRouteTableAssociationResultOutput {
	return o
}

func (o LookupSubnetRouteTableAssociationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSubnetRouteTableAssociationResult] {
	return pulumix.Output[LookupSubnetRouteTableAssociationResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupSubnetRouteTableAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetRouteTableAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetRouteTableAssociationResultOutput{})
}
