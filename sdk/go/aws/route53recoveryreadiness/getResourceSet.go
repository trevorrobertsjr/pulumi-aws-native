// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53recoveryreadiness

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.
func LookupResourceSet(ctx *pulumi.Context, args *LookupResourceSetArgs, opts ...pulumi.InvokeOption) (*LookupResourceSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceSetResult
	err := ctx.Invoke("aws-native:route53recoveryreadiness:getResourceSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceSetArgs struct {
	// The name of the resource set to create.
	ResourceSetName string `pulumi:"resourceSetName"`
}

type LookupResourceSetResult struct {
	// The Amazon Resource Name (ARN) of the resource set.
	ResourceSetArn *string `pulumi:"resourceSetArn"`
	// A list of resource objects in the resource set.
	Resources []ResourceSetResource `pulumi:"resources"`
	// A tag to associate with the parameters for a resource set.
	Tags []ResourceSetTag `pulumi:"tags"`
}

func LookupResourceSetOutput(ctx *pulumi.Context, args LookupResourceSetOutputArgs, opts ...pulumi.InvokeOption) LookupResourceSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceSetResult, error) {
			args := v.(LookupResourceSetArgs)
			r, err := LookupResourceSet(ctx, &args, opts...)
			var s LookupResourceSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceSetResultOutput)
}

type LookupResourceSetOutputArgs struct {
	// The name of the resource set to create.
	ResourceSetName pulumi.StringInput `pulumi:"resourceSetName"`
}

func (LookupResourceSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceSetArgs)(nil)).Elem()
}

type LookupResourceSetResultOutput struct{ *pulumi.OutputState }

func (LookupResourceSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceSetResult)(nil)).Elem()
}

func (o LookupResourceSetResultOutput) ToLookupResourceSetResultOutput() LookupResourceSetResultOutput {
	return o
}

func (o LookupResourceSetResultOutput) ToLookupResourceSetResultOutputWithContext(ctx context.Context) LookupResourceSetResultOutput {
	return o
}

func (o LookupResourceSetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupResourceSetResult] {
	return pulumix.Output[LookupResourceSetResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the resource set.
func (o LookupResourceSetResultOutput) ResourceSetArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceSetResult) *string { return v.ResourceSetArn }).(pulumi.StringPtrOutput)
}

// A list of resource objects in the resource set.
func (o LookupResourceSetResultOutput) Resources() ResourceSetResourceArrayOutput {
	return o.ApplyT(func(v LookupResourceSetResult) []ResourceSetResource { return v.Resources }).(ResourceSetResourceArrayOutput)
}

// A tag to associate with the parameters for a resource set.
func (o LookupResourceSetResultOutput) Tags() ResourceSetTagArrayOutput {
	return o.ApplyT(func(v LookupResourceSetResult) []ResourceSetTag { return v.Tags }).(ResourceSetTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceSetResultOutput{})
}
