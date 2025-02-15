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

// The “AWS::ApiGatewayV2::RouteResponse“ resource creates a route response for a WebSocket API. For more information, see [Set up Route Responses for a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-route-response.html) in the *API Gateway Developer Guide*.
func LookupRouteResponse(ctx *pulumi.Context, args *LookupRouteResponseArgs, opts ...pulumi.InvokeOption) (*LookupRouteResponseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteResponseResult
	err := ctx.Invoke("aws-native:apigatewayv2:getRouteResponse", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteResponseArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The route ID.
	RouteId         string `pulumi:"routeId"`
	RouteResponseId string `pulumi:"routeResponseId"`
}

type LookupRouteResponseResult struct {
	// The model selection expression for the route response. Supported only for WebSocket APIs.
	ModelSelectionExpression *string `pulumi:"modelSelectionExpression"`
	// The response models for the route response.
	ResponseModels interface{} `pulumi:"responseModels"`
	// The route response parameters.
	ResponseParameters *RouteResponseRouteParameters `pulumi:"responseParameters"`
	RouteResponseId    *string                       `pulumi:"routeResponseId"`
	// The route response key.
	RouteResponseKey *string `pulumi:"routeResponseKey"`
}

func LookupRouteResponseOutput(ctx *pulumi.Context, args LookupRouteResponseOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResponseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResponseResult, error) {
			args := v.(LookupRouteResponseArgs)
			r, err := LookupRouteResponse(ctx, &args, opts...)
			var s LookupRouteResponseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResponseResultOutput)
}

type LookupRouteResponseOutputArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// The route ID.
	RouteId         pulumi.StringInput `pulumi:"routeId"`
	RouteResponseId pulumi.StringInput `pulumi:"routeResponseId"`
}

func (LookupRouteResponseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResponseArgs)(nil)).Elem()
}

type LookupRouteResponseResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResponseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResponseResult)(nil)).Elem()
}

func (o LookupRouteResponseResultOutput) ToLookupRouteResponseResultOutput() LookupRouteResponseResultOutput {
	return o
}

func (o LookupRouteResponseResultOutput) ToLookupRouteResponseResultOutputWithContext(ctx context.Context) LookupRouteResponseResultOutput {
	return o
}

func (o LookupRouteResponseResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRouteResponseResult] {
	return pulumix.Output[LookupRouteResponseResult]{
		OutputState: o.OutputState,
	}
}

// The model selection expression for the route response. Supported only for WebSocket APIs.
func (o LookupRouteResponseResultOutput) ModelSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResponseResult) *string { return v.ModelSelectionExpression }).(pulumi.StringPtrOutput)
}

// The response models for the route response.
func (o LookupRouteResponseResultOutput) ResponseModels() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRouteResponseResult) interface{} { return v.ResponseModels }).(pulumi.AnyOutput)
}

// The route response parameters.
func (o LookupRouteResponseResultOutput) ResponseParameters() RouteResponseRouteParametersPtrOutput {
	return o.ApplyT(func(v LookupRouteResponseResult) *RouteResponseRouteParameters { return v.ResponseParameters }).(RouteResponseRouteParametersPtrOutput)
}

func (o LookupRouteResponseResultOutput) RouteResponseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResponseResult) *string { return v.RouteResponseId }).(pulumi.StringPtrOutput)
}

// The route response key.
func (o LookupRouteResponseResultOutput) RouteResponseKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResponseResult) *string { return v.RouteResponseKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResponseResultOutput{})
}
