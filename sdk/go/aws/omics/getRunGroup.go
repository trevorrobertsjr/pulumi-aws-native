// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package omics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::Omics::RunGroup Resource Type
func LookupRunGroup(ctx *pulumi.Context, args *LookupRunGroupArgs, opts ...pulumi.InvokeOption) (*LookupRunGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRunGroupResult
	err := ctx.Invoke("aws-native:omics:getRunGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRunGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupRunGroupResult struct {
	Arn          *string         `pulumi:"arn"`
	CreationTime *string         `pulumi:"creationTime"`
	Id           *string         `pulumi:"id"`
	MaxCpus      *float64        `pulumi:"maxCpus"`
	MaxDuration  *float64        `pulumi:"maxDuration"`
	MaxGpus      *float64        `pulumi:"maxGpus"`
	MaxRuns      *float64        `pulumi:"maxRuns"`
	Name         *string         `pulumi:"name"`
	Tags         *RunGroupTagMap `pulumi:"tags"`
}

func LookupRunGroupOutput(ctx *pulumi.Context, args LookupRunGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRunGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRunGroupResult, error) {
			args := v.(LookupRunGroupArgs)
			r, err := LookupRunGroup(ctx, &args, opts...)
			var s LookupRunGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRunGroupResultOutput)
}

type LookupRunGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupRunGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRunGroupArgs)(nil)).Elem()
}

type LookupRunGroupResultOutput struct{ *pulumi.OutputState }

func (LookupRunGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRunGroupResult)(nil)).Elem()
}

func (o LookupRunGroupResultOutput) ToLookupRunGroupResultOutput() LookupRunGroupResultOutput {
	return o
}

func (o LookupRunGroupResultOutput) ToLookupRunGroupResultOutputWithContext(ctx context.Context) LookupRunGroupResultOutput {
	return o
}

func (o LookupRunGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRunGroupResult] {
	return pulumix.Output[LookupRunGroupResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupRunGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupRunGroupResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupRunGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRunGroupResultOutput) MaxCpus() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *float64 { return v.MaxCpus }).(pulumi.Float64PtrOutput)
}

func (o LookupRunGroupResultOutput) MaxDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *float64 { return v.MaxDuration }).(pulumi.Float64PtrOutput)
}

func (o LookupRunGroupResultOutput) MaxGpus() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *float64 { return v.MaxGpus }).(pulumi.Float64PtrOutput)
}

func (o LookupRunGroupResultOutput) MaxRuns() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *float64 { return v.MaxRuns }).(pulumi.Float64PtrOutput)
}

func (o LookupRunGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRunGroupResultOutput) Tags() RunGroupTagMapPtrOutput {
	return o.ApplyT(func(v LookupRunGroupResult) *RunGroupTagMap { return v.Tags }).(RunGroupTagMapPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRunGroupResultOutput{})
}
