// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apprunner

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The AWS::AppRunner::ObservabilityConfiguration resource  is an AWS App Runner resource type that specifies an App Runner observability configuration
func LookupObservabilityConfiguration(ctx *pulumi.Context, args *LookupObservabilityConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupObservabilityConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupObservabilityConfigurationResult
	err := ctx.Invoke("aws-native:apprunner:getObservabilityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupObservabilityConfigurationArgs struct {
	// The Amazon Resource Name (ARN) of this ObservabilityConfiguration
	ObservabilityConfigurationArn string `pulumi:"observabilityConfigurationArn"`
}

type LookupObservabilityConfigurationResult struct {
	// It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.
	Latest *bool `pulumi:"latest"`
	// The Amazon Resource Name (ARN) of this ObservabilityConfiguration
	ObservabilityConfigurationArn *string `pulumi:"observabilityConfigurationArn"`
	// The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.
	ObservabilityConfigurationRevision *int `pulumi:"observabilityConfigurationRevision"`
}

func LookupObservabilityConfigurationOutput(ctx *pulumi.Context, args LookupObservabilityConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupObservabilityConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObservabilityConfigurationResult, error) {
			args := v.(LookupObservabilityConfigurationArgs)
			r, err := LookupObservabilityConfiguration(ctx, &args, opts...)
			var s LookupObservabilityConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupObservabilityConfigurationResultOutput)
}

type LookupObservabilityConfigurationOutputArgs struct {
	// The Amazon Resource Name (ARN) of this ObservabilityConfiguration
	ObservabilityConfigurationArn pulumi.StringInput `pulumi:"observabilityConfigurationArn"`
}

func (LookupObservabilityConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObservabilityConfigurationArgs)(nil)).Elem()
}

type LookupObservabilityConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupObservabilityConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObservabilityConfigurationResult)(nil)).Elem()
}

func (o LookupObservabilityConfigurationResultOutput) ToLookupObservabilityConfigurationResultOutput() LookupObservabilityConfigurationResultOutput {
	return o
}

func (o LookupObservabilityConfigurationResultOutput) ToLookupObservabilityConfigurationResultOutputWithContext(ctx context.Context) LookupObservabilityConfigurationResultOutput {
	return o
}

func (o LookupObservabilityConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupObservabilityConfigurationResult] {
	return pulumix.Output[LookupObservabilityConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.
func (o LookupObservabilityConfigurationResultOutput) Latest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupObservabilityConfigurationResult) *bool { return v.Latest }).(pulumi.BoolPtrOutput)
}

// The Amazon Resource Name (ARN) of this ObservabilityConfiguration
func (o LookupObservabilityConfigurationResultOutput) ObservabilityConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObservabilityConfigurationResult) *string { return v.ObservabilityConfigurationArn }).(pulumi.StringPtrOutput)
}

// The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.
func (o LookupObservabilityConfigurationResultOutput) ObservabilityConfigurationRevision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupObservabilityConfigurationResult) *int { return v.ObservabilityConfigurationRevision }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObservabilityConfigurationResultOutput{})
}
