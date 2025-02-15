// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::RDS::DBProxyTargetGroup
func LookupDbProxyTargetGroup(ctx *pulumi.Context, args *LookupDbProxyTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupDbProxyTargetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDbProxyTargetGroupResult
	err := ctx.Invoke("aws-native:rds:getDbProxyTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDbProxyTargetGroupArgs struct {
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn string `pulumi:"targetGroupArn"`
}

type LookupDbProxyTargetGroupResult struct {
	ConnectionPoolConfigurationInfo *DbProxyTargetGroupConnectionPoolConfigurationInfoFormat `pulumi:"connectionPoolConfigurationInfo"`
	DbClusterIdentifiers            []string                                                 `pulumi:"dbClusterIdentifiers"`
	DbInstanceIdentifiers           []string                                                 `pulumi:"dbInstanceIdentifiers"`
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn *string `pulumi:"targetGroupArn"`
}

func LookupDbProxyTargetGroupOutput(ctx *pulumi.Context, args LookupDbProxyTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDbProxyTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDbProxyTargetGroupResult, error) {
			args := v.(LookupDbProxyTargetGroupArgs)
			r, err := LookupDbProxyTargetGroup(ctx, &args, opts...)
			var s LookupDbProxyTargetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDbProxyTargetGroupResultOutput)
}

type LookupDbProxyTargetGroupOutputArgs struct {
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn pulumi.StringInput `pulumi:"targetGroupArn"`
}

func (LookupDbProxyTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbProxyTargetGroupArgs)(nil)).Elem()
}

type LookupDbProxyTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDbProxyTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbProxyTargetGroupResult)(nil)).Elem()
}

func (o LookupDbProxyTargetGroupResultOutput) ToLookupDbProxyTargetGroupResultOutput() LookupDbProxyTargetGroupResultOutput {
	return o
}

func (o LookupDbProxyTargetGroupResultOutput) ToLookupDbProxyTargetGroupResultOutputWithContext(ctx context.Context) LookupDbProxyTargetGroupResultOutput {
	return o
}

func (o LookupDbProxyTargetGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDbProxyTargetGroupResult] {
	return pulumix.Output[LookupDbProxyTargetGroupResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupDbProxyTargetGroupResultOutput) ConnectionPoolConfigurationInfo() DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput {
	return o.ApplyT(func(v LookupDbProxyTargetGroupResult) *DbProxyTargetGroupConnectionPoolConfigurationInfoFormat {
		return v.ConnectionPoolConfigurationInfo
	}).(DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput)
}

func (o LookupDbProxyTargetGroupResultOutput) DbClusterIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDbProxyTargetGroupResult) []string { return v.DbClusterIdentifiers }).(pulumi.StringArrayOutput)
}

func (o LookupDbProxyTargetGroupResultOutput) DbInstanceIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDbProxyTargetGroupResult) []string { return v.DbInstanceIdentifiers }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) representing the target group.
func (o LookupDbProxyTargetGroupResultOutput) TargetGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDbProxyTargetGroupResult) *string { return v.TargetGroupArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDbProxyTargetGroupResultOutput{})
}
