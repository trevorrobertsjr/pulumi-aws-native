// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package detective

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::Detective::Graph
type Graph struct {
	pulumi.CustomResourceState

	// The Detective graph ARN
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.
	AutoEnableMembers pulumi.BoolPtrOutput `pulumi:"autoEnableMembers"`
	Tags              GraphTagArrayOutput  `pulumi:"tags"`
}

// NewGraph registers a new resource with the given unique name, arguments, and options.
func NewGraph(ctx *pulumi.Context,
	name string, args *GraphArgs, opts ...pulumi.ResourceOption) (*Graph, error) {
	if args == nil {
		args = &GraphArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Graph
	err := ctx.RegisterResource("aws-native:detective:Graph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraph gets an existing Graph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphState, opts ...pulumi.ResourceOption) (*Graph, error) {
	var resource Graph
	err := ctx.ReadResource("aws-native:detective:Graph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Graph resources.
type graphState struct {
}

type GraphState struct {
}

func (GraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphState)(nil)).Elem()
}

type graphArgs struct {
	// Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.
	AutoEnableMembers *bool      `pulumi:"autoEnableMembers"`
	Tags              []GraphTag `pulumi:"tags"`
}

// The set of arguments for constructing a Graph resource.
type GraphArgs struct {
	// Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.
	AutoEnableMembers pulumi.BoolPtrInput
	Tags              GraphTagArrayInput
}

func (GraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphArgs)(nil)).Elem()
}

type GraphInput interface {
	pulumi.Input

	ToGraphOutput() GraphOutput
	ToGraphOutputWithContext(ctx context.Context) GraphOutput
}

func (*Graph) ElementType() reflect.Type {
	return reflect.TypeOf((**Graph)(nil)).Elem()
}

func (i *Graph) ToGraphOutput() GraphOutput {
	return i.ToGraphOutputWithContext(context.Background())
}

func (i *Graph) ToGraphOutputWithContext(ctx context.Context) GraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphOutput)
}

func (i *Graph) ToOutput(ctx context.Context) pulumix.Output[*Graph] {
	return pulumix.Output[*Graph]{
		OutputState: i.ToGraphOutputWithContext(ctx).OutputState,
	}
}

type GraphOutput struct{ *pulumi.OutputState }

func (GraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Graph)(nil)).Elem()
}

func (o GraphOutput) ToGraphOutput() GraphOutput {
	return o
}

func (o GraphOutput) ToGraphOutputWithContext(ctx context.Context) GraphOutput {
	return o
}

func (o GraphOutput) ToOutput(ctx context.Context) pulumix.Output[*Graph] {
	return pulumix.Output[*Graph]{
		OutputState: o.OutputState,
	}
}

// The Detective graph ARN
func (o GraphOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Graph) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.
func (o GraphOutput) AutoEnableMembers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Graph) pulumi.BoolPtrOutput { return v.AutoEnableMembers }).(pulumi.BoolPtrOutput)
}

func (o GraphOutput) Tags() GraphTagArrayOutput {
	return o.ApplyT(func(v *Graph) GraphTagArrayOutput { return v.Tags }).(GraphTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GraphInput)(nil)).Elem(), &Graph{})
	pulumi.RegisterOutputType(GraphOutput{})
}
