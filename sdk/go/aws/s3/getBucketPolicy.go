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

// Resource Type definition for AWS::S3::BucketPolicy
func LookupBucketPolicy(ctx *pulumi.Context, args *LookupBucketPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBucketPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBucketPolicyResult
	err := ctx.Invoke("aws-native:s3:getBucketPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBucketPolicyArgs struct {
	// The name of the Amazon S3 bucket to which the policy applies.
	Bucket string `pulumi:"bucket"`
}

type LookupBucketPolicyResult struct {
	// A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
	PolicyDocument interface{} `pulumi:"policyDocument"`
}

func LookupBucketPolicyOutput(ctx *pulumi.Context, args LookupBucketPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBucketPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketPolicyResult, error) {
			args := v.(LookupBucketPolicyArgs)
			r, err := LookupBucketPolicy(ctx, &args, opts...)
			var s LookupBucketPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBucketPolicyResultOutput)
}

type LookupBucketPolicyOutputArgs struct {
	// The name of the Amazon S3 bucket to which the policy applies.
	Bucket pulumi.StringInput `pulumi:"bucket"`
}

func (LookupBucketPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketPolicyArgs)(nil)).Elem()
}

type LookupBucketPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBucketPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketPolicyResult)(nil)).Elem()
}

func (o LookupBucketPolicyResultOutput) ToLookupBucketPolicyResultOutput() LookupBucketPolicyResultOutput {
	return o
}

func (o LookupBucketPolicyResultOutput) ToLookupBucketPolicyResultOutputWithContext(ctx context.Context) LookupBucketPolicyResultOutput {
	return o
}

func (o LookupBucketPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBucketPolicyResult] {
	return pulumix.Output[LookupBucketPolicyResult]{
		OutputState: o.OutputState,
	}
}

// A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
func (o LookupBucketPolicyResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBucketPolicyResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketPolicyResultOutput{})
}
