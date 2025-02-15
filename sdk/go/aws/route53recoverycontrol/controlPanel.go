// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53recoverycontrol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// AWS Route53 Recovery Control Control Panel resource schema .
type ControlPanel struct {
	pulumi.CustomResourceState

	// Cluster to associate with the Control Panel
	ClusterArn pulumi.StringPtrOutput `pulumi:"clusterArn"`
	// The Amazon Resource Name (ARN) of the cluster.
	ControlPanelArn pulumi.StringOutput `pulumi:"controlPanelArn"`
	// A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.
	DefaultControlPanel pulumi.BoolOutput `pulumi:"defaultControlPanel"`
	// The name of the control panel. You can use any non-white space character in the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Count of associated routing controls
	RoutingControlCount pulumi.IntOutput `pulumi:"routingControlCount"`
	// The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
	Status ControlPanelStatusOutput `pulumi:"status"`
	// A collection of tags associated with a resource
	Tags ControlPanelTagArrayOutput `pulumi:"tags"`
}

// NewControlPanel registers a new resource with the given unique name, arguments, and options.
func NewControlPanel(ctx *pulumi.Context,
	name string, args *ControlPanelArgs, opts ...pulumi.ResourceOption) (*ControlPanel, error) {
	if args == nil {
		args = &ControlPanelArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clusterArn",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ControlPanel
	err := ctx.RegisterResource("aws-native:route53recoverycontrol:ControlPanel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetControlPanel gets an existing ControlPanel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetControlPanel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControlPanelState, opts ...pulumi.ResourceOption) (*ControlPanel, error) {
	var resource ControlPanel
	err := ctx.ReadResource("aws-native:route53recoverycontrol:ControlPanel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ControlPanel resources.
type controlPanelState struct {
}

type ControlPanelState struct {
}

func (ControlPanelState) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPanelState)(nil)).Elem()
}

type controlPanelArgs struct {
	// Cluster to associate with the Control Panel
	ClusterArn *string `pulumi:"clusterArn"`
	// The name of the control panel. You can use any non-white space character in the name.
	Name *string `pulumi:"name"`
	// A collection of tags associated with a resource
	Tags []ControlPanelTag `pulumi:"tags"`
}

// The set of arguments for constructing a ControlPanel resource.
type ControlPanelArgs struct {
	// Cluster to associate with the Control Panel
	ClusterArn pulumi.StringPtrInput
	// The name of the control panel. You can use any non-white space character in the name.
	Name pulumi.StringPtrInput
	// A collection of tags associated with a resource
	Tags ControlPanelTagArrayInput
}

func (ControlPanelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPanelArgs)(nil)).Elem()
}

type ControlPanelInput interface {
	pulumi.Input

	ToControlPanelOutput() ControlPanelOutput
	ToControlPanelOutputWithContext(ctx context.Context) ControlPanelOutput
}

func (*ControlPanel) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPanel)(nil)).Elem()
}

func (i *ControlPanel) ToControlPanelOutput() ControlPanelOutput {
	return i.ToControlPanelOutputWithContext(context.Background())
}

func (i *ControlPanel) ToControlPanelOutputWithContext(ctx context.Context) ControlPanelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPanelOutput)
}

func (i *ControlPanel) ToOutput(ctx context.Context) pulumix.Output[*ControlPanel] {
	return pulumix.Output[*ControlPanel]{
		OutputState: i.ToControlPanelOutputWithContext(ctx).OutputState,
	}
}

type ControlPanelOutput struct{ *pulumi.OutputState }

func (ControlPanelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPanel)(nil)).Elem()
}

func (o ControlPanelOutput) ToControlPanelOutput() ControlPanelOutput {
	return o
}

func (o ControlPanelOutput) ToControlPanelOutputWithContext(ctx context.Context) ControlPanelOutput {
	return o
}

func (o ControlPanelOutput) ToOutput(ctx context.Context) pulumix.Output[*ControlPanel] {
	return pulumix.Output[*ControlPanel]{
		OutputState: o.OutputState,
	}
}

// Cluster to associate with the Control Panel
func (o ControlPanelOutput) ClusterArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPanel) pulumi.StringPtrOutput { return v.ClusterArn }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the cluster.
func (o ControlPanelOutput) ControlPanelArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPanel) pulumi.StringOutput { return v.ControlPanelArn }).(pulumi.StringOutput)
}

// A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.
func (o ControlPanelOutput) DefaultControlPanel() pulumi.BoolOutput {
	return o.ApplyT(func(v *ControlPanel) pulumi.BoolOutput { return v.DefaultControlPanel }).(pulumi.BoolOutput)
}

// The name of the control panel. You can use any non-white space character in the name.
func (o ControlPanelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPanel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Count of associated routing controls
func (o ControlPanelOutput) RoutingControlCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ControlPanel) pulumi.IntOutput { return v.RoutingControlCount }).(pulumi.IntOutput)
}

// The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
func (o ControlPanelOutput) Status() ControlPanelStatusOutput {
	return o.ApplyT(func(v *ControlPanel) ControlPanelStatusOutput { return v.Status }).(ControlPanelStatusOutput)
}

// A collection of tags associated with a resource
func (o ControlPanelOutput) Tags() ControlPanelTagArrayOutput {
	return o.ApplyT(func(v *ControlPanel) ControlPanelTagArrayOutput { return v.Tags }).(ControlPanelTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPanelInput)(nil)).Elem(), &ControlPanel{})
	pulumi.RegisterOutputType(ControlPanelOutput{})
}
