// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The “AWS::ApiGatewayV2::Model“ resource updates data model for a WebSocket API. For more information, see [Model Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) in the *API Gateway Developer Guide*.
func LookupModel(ctx *pulumi.Context, args *LookupModelArgs, opts ...pulumi.InvokeOption) (*LookupModelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupModelResult
	err := ctx.Invoke("aws-native:apigatewayv2:getModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelArgs struct {
	// The API identifier.
	ApiId   string `pulumi:"apiId"`
	ModelId string `pulumi:"modelId"`
}

type LookupModelResult struct {
	// The content-type for the model, for example, "application/json".
	ContentType *string `pulumi:"contentType"`
	// The description of the model.
	Description *string `pulumi:"description"`
	ModelId     *string `pulumi:"modelId"`
	// The name of the model.
	Name *string `pulumi:"name"`
	// The schema for the model. For application/json models, this should be JSON schema draft 4 model.
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
	// The API identifier.
	ApiId   pulumi.StringInput `pulumi:"apiId"`
	ModelId pulumi.StringInput `pulumi:"modelId"`
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

// The content-type for the model, for example, "application/json".
func (o LookupModelResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The description of the model.
func (o LookupModelResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupModelResultOutput) ModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelResult) *string { return v.ModelId }).(pulumi.StringPtrOutput)
}

// The name of the model.
func (o LookupModelResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The schema for the model. For application/json models, this should be JSON schema draft 4 model.
func (o LookupModelResultOutput) Schema() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupModelResult) interface{} { return v.Schema }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelResultOutput{})
}
