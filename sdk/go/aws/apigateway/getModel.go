// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The “AWS::ApiGateway::Model“ resource defines the structure of a request or response payload for an API method.
func LookupModel(ctx *pulumi.Context, args *LookupModelArgs, opts ...pulumi.InvokeOption) (*LookupModelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupModelResult
	err := ctx.Invoke("aws-native:apigateway:getModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelArgs struct {
	// A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name string `pulumi:"name"`
	// The string identifier of the associated RestApi.
	RestApiId string `pulumi:"restApiId"`
}

type LookupModelResult struct {
	// The description of the model.
	Description *string `pulumi:"description"`
	// The schema for the model. For ``application/json`` models, this should be JSON schema draft 4 model. Do not include "\*/" characters in the description of any properties because such "\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
	Schema interface{} `pulumi:"schema"`
}

func LookupModelOutput(ctx *pulumi.Context, args LookupModelOutputArgs, opts ...pulumi.InvokeOption) LookupModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelResult, error) {
			args := v.(LookupModelArgs)
			r, err := LookupModel(ctx, &args, opts...)
			var s LookupModelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelResultOutput)
}

type LookupModelOutputArgs struct {
	// A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name pulumi.StringInput `pulumi:"name"`
	// The string identifier of the associated RestApi.
	RestApiId pulumi.StringInput `pulumi:"restApiId"`
}

func (LookupModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelArgs)(nil)).Elem()
}

type LookupModelResultOutput struct{ *pulumi.OutputState }

func (LookupModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelResult)(nil)).Elem()
}

func (o LookupModelResultOutput) ToLookupModelResultOutput() LookupModelResultOutput {
	return o
}

func (o LookupModelResultOutput) ToLookupModelResultOutputWithContext(ctx context.Context) LookupModelResultOutput {
	return o
}

func (o LookupModelResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupModelResult] {
	return pulumix.Output[LookupModelResult]{
		OutputState: o.OutputState,
	}
}

// The description of the model.
func (o LookupModelResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The schema for the model. For “application/json“ models, this should be JSON schema draft 4 model. Do not include "\*/" characters in the description of any properties because such "\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
func (o LookupModelResultOutput) Schema() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupModelResult) interface{} { return v.Schema }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelResultOutput{})
}
