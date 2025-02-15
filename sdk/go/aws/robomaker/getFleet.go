// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// AWS::RoboMaker::Fleet resource creates an AWS RoboMaker fleet. Fleets contain robots and can receive deployments.
func LookupFleet(ctx *pulumi.Context, args *LookupFleetArgs, opts ...pulumi.InvokeOption) (*LookupFleetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFleetResult
	err := ctx.Invoke("aws-native:robomaker:getFleet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupFleetResult struct {
	Arn  *string    `pulumi:"arn"`
	Tags *FleetTags `pulumi:"tags"`
}

func LookupFleetOutput(ctx *pulumi.Context, args LookupFleetOutputArgs, opts ...pulumi.InvokeOption) LookupFleetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFleetResult, error) {
			args := v.(LookupFleetArgs)
			r, err := LookupFleet(ctx, &args, opts...)
			var s LookupFleetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFleetResultOutput)
}

type LookupFleetOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupFleetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetArgs)(nil)).Elem()
}

type LookupFleetResultOutput struct{ *pulumi.OutputState }

func (LookupFleetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetResult)(nil)).Elem()
}

func (o LookupFleetResultOutput) ToLookupFleetResultOutput() LookupFleetResultOutput {
	return o
}

func (o LookupFleetResultOutput) ToLookupFleetResultOutputWithContext(ctx context.Context) LookupFleetResultOutput {
	return o
}

func (o LookupFleetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFleetResult] {
	return pulumix.Output[LookupFleetResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFleetResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) Tags() FleetTagsPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *FleetTags { return v.Tags }).(FleetTagsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetResultOutput{})
}
