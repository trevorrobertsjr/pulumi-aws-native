// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of the AWS::QuickSight::Template Resource Type.
func LookupTemplate(ctx *pulumi.Context, args *LookupTemplateArgs, opts ...pulumi.InvokeOption) (*LookupTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTemplateResult
	err := ctx.Invoke("aws-native:quicksight:getTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateArgs struct {
	AwsAccountId string `pulumi:"awsAccountId"`
	TemplateId   string `pulumi:"templateId"`
}

type LookupTemplateResult struct {
	Arn         *string                      `pulumi:"arn"`
	Name        *string                      `pulumi:"name"`
	Permissions []TemplateResourcePermission `pulumi:"permissions"`
	Tags        []TemplateTag                `pulumi:"tags"`
}

func LookupTemplateOutput(ctx *pulumi.Context, args LookupTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplateResult, error) {
			args := v.(LookupTemplateArgs)
			r, err := LookupTemplate(ctx, &args, opts...)
			var s LookupTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplateResultOutput)
}

type LookupTemplateOutputArgs struct {
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
	TemplateId   pulumi.StringInput `pulumi:"templateId"`
}

func (LookupTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateArgs)(nil)).Elem()
}

type LookupTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateResult)(nil)).Elem()
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutput() LookupTemplateResultOutput {
	return o
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutputWithContext(ctx context.Context) LookupTemplateResultOutput {
	return o
}

func (o LookupTemplateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTemplateResult] {
	return pulumix.Output[LookupTemplateResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupTemplateResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateResultOutput) Permissions() TemplateResourcePermissionArrayOutput {
	return o.ApplyT(func(v LookupTemplateResult) []TemplateResourcePermission { return v.Permissions }).(TemplateResourcePermissionArrayOutput)
}

func (o LookupTemplateResultOutput) Tags() TemplateTagArrayOutput {
	return o.ApplyT(func(v LookupTemplateResult) []TemplateTag { return v.Tags }).(TemplateTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateResultOutput{})
}
