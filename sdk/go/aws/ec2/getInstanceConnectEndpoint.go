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

// Resource Type definition for AWS::EC2::InstanceConnectEndpoint
func LookupInstanceConnectEndpoint(ctx *pulumi.Context, args *LookupInstanceConnectEndpointArgs, opts ...pulumi.InvokeOption) (*LookupInstanceConnectEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceConnectEndpointResult
	err := ctx.Invoke("aws-native:ec2:getInstanceConnectEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceConnectEndpointArgs struct {
	// The id of the instance connect endpoint
	Id string `pulumi:"id"`
}

type LookupInstanceConnectEndpointResult struct {
	// The id of the instance connect endpoint
	Id *string `pulumi:"id"`
	// The tags of the instance connect endpoint.
	Tags []InstanceConnectEndpointTag `pulumi:"tags"`
}

func LookupInstanceConnectEndpointOutput(ctx *pulumi.Context, args LookupInstanceConnectEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceConnectEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceConnectEndpointResult, error) {
			args := v.(LookupInstanceConnectEndpointArgs)
			r, err := LookupInstanceConnectEndpoint(ctx, &args, opts...)
			var s LookupInstanceConnectEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceConnectEndpointResultOutput)
}

type LookupInstanceConnectEndpointOutputArgs struct {
	// The id of the instance connect endpoint
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupInstanceConnectEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceConnectEndpointArgs)(nil)).Elem()
}

type LookupInstanceConnectEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceConnectEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceConnectEndpointResult)(nil)).Elem()
}

func (o LookupInstanceConnectEndpointResultOutput) ToLookupInstanceConnectEndpointResultOutput() LookupInstanceConnectEndpointResultOutput {
	return o
}

func (o LookupInstanceConnectEndpointResultOutput) ToLookupInstanceConnectEndpointResultOutputWithContext(ctx context.Context) LookupInstanceConnectEndpointResultOutput {
	return o
}

func (o LookupInstanceConnectEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInstanceConnectEndpointResult] {
	return pulumix.Output[LookupInstanceConnectEndpointResult]{
		OutputState: o.OutputState,
	}
}

// The id of the instance connect endpoint
func (o LookupInstanceConnectEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceConnectEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The tags of the instance connect endpoint.
func (o LookupInstanceConnectEndpointResultOutput) Tags() InstanceConnectEndpointTagArrayOutput {
	return o.ApplyT(func(v LookupInstanceConnectEndpointResult) []InstanceConnectEndpointTag { return v.Tags }).(InstanceConnectEndpointTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceConnectEndpointResultOutput{})
}
