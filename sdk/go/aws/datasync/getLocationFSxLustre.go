// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::DataSync::LocationFSxLustre.
func LookupLocationFSxLustre(ctx *pulumi.Context, args *LookupLocationFSxLustreArgs, opts ...pulumi.InvokeOption) (*LookupLocationFSxLustreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocationFSxLustreResult
	err := ctx.Invoke("aws-native:datasync:getLocationFSxLustre", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationFSxLustreArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.
	LocationArn string `pulumi:"locationArn"`
}

type LookupLocationFSxLustreResult struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.
	LocationArn *string `pulumi:"locationArn"`
	// The URL of the FSx for Lustre location that was described.
	LocationUri *string `pulumi:"locationUri"`
	// An array of key-value pairs to apply to this resource.
	Tags []LocationFSxLustreTag `pulumi:"tags"`
}

func LookupLocationFSxLustreOutput(ctx *pulumi.Context, args LookupLocationFSxLustreOutputArgs, opts ...pulumi.InvokeOption) LookupLocationFSxLustreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocationFSxLustreResult, error) {
			args := v.(LookupLocationFSxLustreArgs)
			r, err := LookupLocationFSxLustre(ctx, &args, opts...)
			var s LookupLocationFSxLustreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocationFSxLustreResultOutput)
}

type LookupLocationFSxLustreOutputArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.
	LocationArn pulumi.StringInput `pulumi:"locationArn"`
}

func (LookupLocationFSxLustreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationFSxLustreArgs)(nil)).Elem()
}

type LookupLocationFSxLustreResultOutput struct{ *pulumi.OutputState }

func (LookupLocationFSxLustreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationFSxLustreResult)(nil)).Elem()
}

func (o LookupLocationFSxLustreResultOutput) ToLookupLocationFSxLustreResultOutput() LookupLocationFSxLustreResultOutput {
	return o
}

func (o LookupLocationFSxLustreResultOutput) ToLookupLocationFSxLustreResultOutputWithContext(ctx context.Context) LookupLocationFSxLustreResultOutput {
	return o
}

func (o LookupLocationFSxLustreResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLocationFSxLustreResult] {
	return pulumix.Output[LookupLocationFSxLustreResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.
func (o LookupLocationFSxLustreResultOutput) LocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationFSxLustreResult) *string { return v.LocationArn }).(pulumi.StringPtrOutput)
}

// The URL of the FSx for Lustre location that was described.
func (o LookupLocationFSxLustreResultOutput) LocationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationFSxLustreResult) *string { return v.LocationUri }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLocationFSxLustreResultOutput) Tags() LocationFSxLustreTagArrayOutput {
	return o.ApplyT(func(v LookupLocationFSxLustreResult) []LocationFSxLustreTag { return v.Tags }).(LocationFSxLustreTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationFSxLustreResultOutput{})
}
