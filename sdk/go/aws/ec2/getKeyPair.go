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

// The AWS::EC2::KeyPair creates an SSH key pair
func LookupKeyPair(ctx *pulumi.Context, args *LookupKeyPairArgs, opts ...pulumi.InvokeOption) (*LookupKeyPairResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKeyPairResult
	err := ctx.Invoke("aws-native:ec2:getKeyPair", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyPairArgs struct {
	// The name of the SSH key pair
	KeyName string `pulumi:"keyName"`
}

type LookupKeyPairResult struct {
	// A short sequence of bytes used for public key verification
	KeyFingerprint *string `pulumi:"keyFingerprint"`
	// An AWS generated ID for the key pair
	KeyPairId *string `pulumi:"keyPairId"`
}

func LookupKeyPairOutput(ctx *pulumi.Context, args LookupKeyPairOutputArgs, opts ...pulumi.InvokeOption) LookupKeyPairResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKeyPairResult, error) {
			args := v.(LookupKeyPairArgs)
			r, err := LookupKeyPair(ctx, &args, opts...)
			var s LookupKeyPairResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKeyPairResultOutput)
}

type LookupKeyPairOutputArgs struct {
	// The name of the SSH key pair
	KeyName pulumi.StringInput `pulumi:"keyName"`
}

func (LookupKeyPairOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyPairArgs)(nil)).Elem()
}

type LookupKeyPairResultOutput struct{ *pulumi.OutputState }

func (LookupKeyPairResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyPairResult)(nil)).Elem()
}

func (o LookupKeyPairResultOutput) ToLookupKeyPairResultOutput() LookupKeyPairResultOutput {
	return o
}

func (o LookupKeyPairResultOutput) ToLookupKeyPairResultOutputWithContext(ctx context.Context) LookupKeyPairResultOutput {
	return o
}

func (o LookupKeyPairResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupKeyPairResult] {
	return pulumix.Output[LookupKeyPairResult]{
		OutputState: o.OutputState,
	}
}

// A short sequence of bytes used for public key verification
func (o LookupKeyPairResultOutput) KeyFingerprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyPairResult) *string { return v.KeyFingerprint }).(pulumi.StringPtrOutput)
}

// An AWS generated ID for the key pair
func (o LookupKeyPairResultOutput) KeyPairId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyPairResult) *string { return v.KeyPairId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeyPairResultOutput{})
}
