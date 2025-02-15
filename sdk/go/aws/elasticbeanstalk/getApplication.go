// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationResult
	err := ctx.Invoke("aws-native:elasticbeanstalk:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
	ApplicationName string `pulumi:"applicationName"`
}

type LookupApplicationResult struct {
	// Your description of the application.
	Description *string `pulumi:"description"`
	// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
	ResourceLifecycleConfig *ApplicationResourceLifecycleConfig `pulumi:"resourceLifecycleConfig"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApplicationResult] {
	return pulumix.Output[LookupApplicationResult]{
		OutputState: o.OutputState,
	}
}

// Your description of the application.
func (o LookupApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
func (o LookupApplicationResultOutput) ResourceLifecycleConfig() ApplicationResourceLifecycleConfigPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *ApplicationResourceLifecycleConfig { return v.ResourceLifecycleConfig }).(ApplicationResourceLifecycleConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
