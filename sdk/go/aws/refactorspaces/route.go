// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package refactorspaces

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::RefactorSpaces::Route Resource Type
type Route struct {
	pulumi.CustomResourceState

	ApplicationIdentifier pulumi.StringOutput             `pulumi:"applicationIdentifier"`
	Arn                   pulumi.StringOutput             `pulumi:"arn"`
	DefaultRoute          RouteDefaultRouteInputPtrOutput `pulumi:"defaultRoute"`
	EnvironmentIdentifier pulumi.StringOutput             `pulumi:"environmentIdentifier"`
	PathResourceToId      pulumi.StringOutput             `pulumi:"pathResourceToId"`
	RouteIdentifier       pulumi.StringOutput             `pulumi:"routeIdentifier"`
	RouteType             RouteTypeOutput                 `pulumi:"routeType"`
	ServiceIdentifier     pulumi.StringOutput             `pulumi:"serviceIdentifier"`
	// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
	Tags         RouteTagArrayOutput             `pulumi:"tags"`
	UriPathRoute RouteUriPathRouteInputPtrOutput `pulumi:"uriPathRoute"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationIdentifier'")
	}
	if args.EnvironmentIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentIdentifier'")
	}
	if args.RouteType == nil {
		return nil, errors.New("invalid value for required argument 'RouteType'")
	}
	if args.ServiceIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ServiceIdentifier'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"applicationIdentifier",
		"environmentIdentifier",
		"routeType",
		"serviceIdentifier",
		"uriPathRoute.appendSourcePath",
		"uriPathRoute.includeChildPaths",
		"uriPathRoute.methods[*]",
		"uriPathRoute.sourcePath",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("aws-native:refactorspaces:Route", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:refactorspaces:Route", name, id, state, &resource, opts...)
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
	ApplicationIdentifier string                  `pulumi:"applicationIdentifier"`
	DefaultRoute          *RouteDefaultRouteInput `pulumi:"defaultRoute"`
	EnvironmentIdentifier string                  `pulumi:"environmentIdentifier"`
	RouteType             RouteType               `pulumi:"routeType"`
	ServiceIdentifier     string                  `pulumi:"serviceIdentifier"`
	// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
	Tags         []RouteTag              `pulumi:"tags"`
	UriPathRoute *RouteUriPathRouteInput `pulumi:"uriPathRoute"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	ApplicationIdentifier pulumi.StringInput
	DefaultRoute          RouteDefaultRouteInputPtrInput
	EnvironmentIdentifier pulumi.StringInput
	RouteType             RouteTypeInput
	ServiceIdentifier     pulumi.StringInput
	// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
	Tags         RouteTagArrayInput
	UriPathRoute RouteUriPathRouteInputPtrInput
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

func (o RouteOutput) ApplicationIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ApplicationIdentifier }).(pulumi.StringOutput)
}

func (o RouteOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o RouteOutput) DefaultRoute() RouteDefaultRouteInputPtrOutput {
	return o.ApplyT(func(v *Route) RouteDefaultRouteInputPtrOutput { return v.DefaultRoute }).(RouteDefaultRouteInputPtrOutput)
}

func (o RouteOutput) EnvironmentIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.EnvironmentIdentifier }).(pulumi.StringOutput)
}

func (o RouteOutput) PathResourceToId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.PathResourceToId }).(pulumi.StringOutput)
}

func (o RouteOutput) RouteIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.RouteIdentifier }).(pulumi.StringOutput)
}

func (o RouteOutput) RouteType() RouteTypeOutput {
	return o.ApplyT(func(v *Route) RouteTypeOutput { return v.RouteType }).(RouteTypeOutput)
}

func (o RouteOutput) ServiceIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ServiceIdentifier }).(pulumi.StringOutput)
}

// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
func (o RouteOutput) Tags() RouteTagArrayOutput {
	return o.ApplyT(func(v *Route) RouteTagArrayOutput { return v.Tags }).(RouteTagArrayOutput)
}

func (o RouteOutput) UriPathRoute() RouteUriPathRouteInputPtrOutput {
	return o.ApplyT(func(v *Route) RouteUriPathRouteInputPtrOutput { return v.UriPathRoute }).(RouteUriPathRouteInputPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterOutputType(RouteOutput{})
}
