// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The “AWS::ApiGatewayV2::Route“ resource creates a route for an API.
type Route struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Specifies whether an API key is required for the route. Supported only for WebSocket APIs.
	ApiKeyRequired pulumi.BoolPtrOutput `pulumi:"apiKeyRequired"`
	// The authorization scopes supported by this route.
	AuthorizationScopes pulumi.StringArrayOutput `pulumi:"authorizationScopes"`
	// The authorization type for the route. For WebSocket APIs, valid values are ``NONE`` for open access, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer. For HTTP APIs, valid values are ``NONE`` for open access, ``JWT`` for using JSON Web Tokens, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer.
	AuthorizationType pulumi.StringPtrOutput `pulumi:"authorizationType"`
	// The identifier of the ``Authorizer`` resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.
	AuthorizerId pulumi.StringPtrOutput `pulumi:"authorizerId"`
	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression pulumi.StringPtrOutput `pulumi:"modelSelectionExpression"`
	// The operation name for the route.
	OperationName pulumi.StringPtrOutput `pulumi:"operationName"`
	// The request models for the route. Supported only for WebSocket APIs.
	RequestModels pulumi.AnyOutput `pulumi:"requestModels"`
	// The request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters pulumi.AnyOutput    `pulumi:"requestParameters"`
	RouteId           pulumi.StringOutput `pulumi:"routeId"`
	// The route key for the route. For HTTP APIs, the route key can be either ``$default``, or a combination of an HTTP method and resource path, for example, ``GET /pets``.
	RouteKey pulumi.StringOutput `pulumi:"routeKey"`
	// The route response selection expression for the route. Supported only for WebSocket APIs.
	RouteResponseSelectionExpression pulumi.StringPtrOutput `pulumi:"routeResponseSelectionExpression"`
	// The target for the route.
	Target pulumi.StringPtrOutput `pulumi:"target"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.RouteKey == nil {
		return nil, errors.New("invalid value for required argument 'RouteKey'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"apiId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("aws-native:apigatewayv2:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws-native:apigatewayv2:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// Specifies whether an API key is required for the route. Supported only for WebSocket APIs.
	ApiKeyRequired *bool `pulumi:"apiKeyRequired"`
	// The authorization scopes supported by this route.
	AuthorizationScopes []string `pulumi:"authorizationScopes"`
	// The authorization type for the route. For WebSocket APIs, valid values are ``NONE`` for open access, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer. For HTTP APIs, valid values are ``NONE`` for open access, ``JWT`` for using JSON Web Tokens, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer.
	AuthorizationType *string `pulumi:"authorizationType"`
	// The identifier of the ``Authorizer`` resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.
	AuthorizerId *string `pulumi:"authorizerId"`
	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression *string `pulumi:"modelSelectionExpression"`
	// The operation name for the route.
	OperationName *string `pulumi:"operationName"`
	// The request models for the route. Supported only for WebSocket APIs.
	RequestModels interface{} `pulumi:"requestModels"`
	// The request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters interface{} `pulumi:"requestParameters"`
	// The route key for the route. For HTTP APIs, the route key can be either ``$default``, or a combination of an HTTP method and resource path, for example, ``GET /pets``.
	RouteKey string `pulumi:"routeKey"`
	// The route response selection expression for the route. Supported only for WebSocket APIs.
	RouteResponseSelectionExpression *string `pulumi:"routeResponseSelectionExpression"`
	// The target for the route.
	Target *string `pulumi:"target"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// Specifies whether an API key is required for the route. Supported only for WebSocket APIs.
	ApiKeyRequired pulumi.BoolPtrInput
	// The authorization scopes supported by this route.
	AuthorizationScopes pulumi.StringArrayInput
	// The authorization type for the route. For WebSocket APIs, valid values are ``NONE`` for open access, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer. For HTTP APIs, valid values are ``NONE`` for open access, ``JWT`` for using JSON Web Tokens, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer.
	AuthorizationType pulumi.StringPtrInput
	// The identifier of the ``Authorizer`` resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.
	AuthorizerId pulumi.StringPtrInput
	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression pulumi.StringPtrInput
	// The operation name for the route.
	OperationName pulumi.StringPtrInput
	// The request models for the route. Supported only for WebSocket APIs.
	RequestModels pulumi.Input
	// The request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters pulumi.Input
	// The route key for the route. For HTTP APIs, the route key can be either ``$default``, or a combination of an HTTP method and resource path, for example, ``GET /pets``.
	RouteKey pulumi.StringInput
	// The route response selection expression for the route. Supported only for WebSocket APIs.
	RouteResponseSelectionExpression pulumi.StringPtrInput
	// The target for the route.
	Target pulumi.StringPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i *Route) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: i.ToRouteOutputWithContext(ctx).OutputState,
	}
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: o.OutputState,
	}
}

// The API identifier.
func (o RouteOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// Specifies whether an API key is required for the route. Supported only for WebSocket APIs.
func (o RouteOutput) ApiKeyRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.BoolPtrOutput { return v.ApiKeyRequired }).(pulumi.BoolPtrOutput)
}

// The authorization scopes supported by this route.
func (o RouteOutput) AuthorizationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Route) pulumi.StringArrayOutput { return v.AuthorizationScopes }).(pulumi.StringArrayOutput)
}

// The authorization type for the route. For WebSocket APIs, valid values are “NONE“ for open access, “AWS_IAM“ for using AWS IAM permissions, and “CUSTOM“ for using a Lambda authorizer. For HTTP APIs, valid values are “NONE“ for open access, “JWT“ for using JSON Web Tokens, “AWS_IAM“ for using AWS IAM permissions, and “CUSTOM“ for using a Lambda authorizer.
func (o RouteOutput) AuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.AuthorizationType }).(pulumi.StringPtrOutput)
}

// The identifier of the “Authorizer“ resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.
func (o RouteOutput) AuthorizerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.AuthorizerId }).(pulumi.StringPtrOutput)
}

// The model selection expression for the route. Supported only for WebSocket APIs.
func (o RouteOutput) ModelSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.ModelSelectionExpression }).(pulumi.StringPtrOutput)
}

// The operation name for the route.
func (o RouteOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.OperationName }).(pulumi.StringPtrOutput)
}

// The request models for the route. Supported only for WebSocket APIs.
func (o RouteOutput) RequestModels() pulumi.AnyOutput {
	return o.ApplyT(func(v *Route) pulumi.AnyOutput { return v.RequestModels }).(pulumi.AnyOutput)
}

// The request parameters for the route. Supported only for WebSocket APIs.
func (o RouteOutput) RequestParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *Route) pulumi.AnyOutput { return v.RequestParameters }).(pulumi.AnyOutput)
}

func (o RouteOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.RouteId }).(pulumi.StringOutput)
}

// The route key for the route. For HTTP APIs, the route key can be either “$default“, or a combination of an HTTP method and resource path, for example, “GET /pets“.
func (o RouteOutput) RouteKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.RouteKey }).(pulumi.StringOutput)
}

// The route response selection expression for the route. Supported only for WebSocket APIs.
func (o RouteOutput) RouteResponseSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.RouteResponseSelectionExpression }).(pulumi.StringPtrOutput)
}

// The target for the route.
func (o RouteOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.Target }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterOutputType(RouteOutput{})
}
