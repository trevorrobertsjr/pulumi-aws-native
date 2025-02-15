// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The AWS::S3::AccessGrantsInstance resource is an Amazon S3 resource type that hosts Access Grants and their associated locations
func LookupAccessGrantsInstance(ctx *pulumi.Context, args *LookupAccessGrantsInstanceArgs, opts ...pulumi.InvokeOption) (*LookupAccessGrantsInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessGrantsInstanceResult
	err := ctx.Invoke("aws-native:s3:getAccessGrantsInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessGrantsInstanceArgs struct {
	// The Amazon Resource Name (ARN) of the specified Access Grants instance.
	AccessGrantsInstanceArn string `pulumi:"accessGrantsInstanceArn"`
}

type LookupAccessGrantsInstanceResult struct {
	// The Amazon Resource Name (ARN) of the specified Access Grants instance.
	AccessGrantsInstanceArn *string `pulumi:"accessGrantsInstanceArn"`
	// A unique identifier for the specified access grants instance.
	AccessGrantsInstanceId *string `pulumi:"accessGrantsInstanceId"`
	// The Amazon Resource Name (ARN) of the specified AWS Identity Center.
	IdentityCenterArn *string `pulumi:"identityCenterArn"`
}

func LookupAccessGrantsInstanceOutput(ctx *pulumi.Context, args LookupAccessGrantsInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupAccessGrantsInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessGrantsInstanceResult, error) {
			args := v.(LookupAccessGrantsInstanceArgs)
			r, err := LookupAccessGrantsInstance(ctx, &args, opts...)
			var s LookupAccessGrantsInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessGrantsInstanceResultOutput)
}

type LookupAccessGrantsInstanceOutputArgs struct {
	// The Amazon Resource Name (ARN) of the specified Access Grants instance.
	AccessGrantsInstanceArn pulumi.StringInput `pulumi:"accessGrantsInstanceArn"`
}

func (LookupAccessGrantsInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessGrantsInstanceArgs)(nil)).Elem()
}

type LookupAccessGrantsInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupAccessGrantsInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessGrantsInstanceResult)(nil)).Elem()
}

func (o LookupAccessGrantsInstanceResultOutput) ToLookupAccessGrantsInstanceResultOutput() LookupAccessGrantsInstanceResultOutput {
	return o
}

func (o LookupAccessGrantsInstanceResultOutput) ToLookupAccessGrantsInstanceResultOutputWithContext(ctx context.Context) LookupAccessGrantsInstanceResultOutput {
	return o
}

func (o LookupAccessGrantsInstanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAccessGrantsInstanceResult] {
	return pulumix.Output[LookupAccessGrantsInstanceResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the specified Access Grants instance.
func (o LookupAccessGrantsInstanceResultOutput) AccessGrantsInstanceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessGrantsInstanceResult) *string { return v.AccessGrantsInstanceArn }).(pulumi.StringPtrOutput)
}

// A unique identifier for the specified access grants instance.
func (o LookupAccessGrantsInstanceResultOutput) AccessGrantsInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessGrantsInstanceResult) *string { return v.AccessGrantsInstanceId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the specified AWS Identity Center.
func (o LookupAccessGrantsInstanceResultOutput) IdentityCenterArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessGrantsInstanceResult) *string { return v.IdentityCenterArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessGrantsInstanceResultOutput{})
}
