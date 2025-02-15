// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspacesweb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::WorkSpacesWeb::NetworkSettings Resource Type
type NetworkSettings struct {
	pulumi.CustomResourceState

	AssociatedPortalArns pulumi.StringArrayOutput      `pulumi:"associatedPortalArns"`
	NetworkSettingsArn   pulumi.StringOutput           `pulumi:"networkSettingsArn"`
	SecurityGroupIds     pulumi.StringArrayOutput      `pulumi:"securityGroupIds"`
	SubnetIds            pulumi.StringArrayOutput      `pulumi:"subnetIds"`
	Tags                 NetworkSettingsTagArrayOutput `pulumi:"tags"`
	VpcId                pulumi.StringOutput           `pulumi:"vpcId"`
}

// NewNetworkSettings registers a new resource with the given unique name, arguments, and options.
func NewNetworkSettings(ctx *pulumi.Context,
	name string, args *NetworkSettingsArgs, opts ...pulumi.ResourceOption) (*NetworkSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkSettings
	err := ctx.RegisterResource("aws-native:workspacesweb:NetworkSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSettings gets an existing NetworkSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSettingsState, opts ...pulumi.ResourceOption) (*NetworkSettings, error) {
	var resource NetworkSettings
	err := ctx.ReadResource("aws-native:workspacesweb:NetworkSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSettings resources.
type networkSettingsState struct {
}

type NetworkSettingsState struct {
}

func (NetworkSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSettingsState)(nil)).Elem()
}

type networkSettingsArgs struct {
	SecurityGroupIds []string             `pulumi:"securityGroupIds"`
	SubnetIds        []string             `pulumi:"subnetIds"`
	Tags             []NetworkSettingsTag `pulumi:"tags"`
	VpcId            string               `pulumi:"vpcId"`
}

// The set of arguments for constructing a NetworkSettings resource.
type NetworkSettingsArgs struct {
	SecurityGroupIds pulumi.StringArrayInput
	SubnetIds        pulumi.StringArrayInput
	Tags             NetworkSettingsTagArrayInput
	VpcId            pulumi.StringInput
}

func (NetworkSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSettingsArgs)(nil)).Elem()
}

type NetworkSettingsInput interface {
	pulumi.Input

	ToNetworkSettingsOutput() NetworkSettingsOutput
	ToNetworkSettingsOutputWithContext(ctx context.Context) NetworkSettingsOutput
}

func (*NetworkSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSettings)(nil)).Elem()
}

func (i *NetworkSettings) ToNetworkSettingsOutput() NetworkSettingsOutput {
	return i.ToNetworkSettingsOutputWithContext(context.Background())
}

func (i *NetworkSettings) ToNetworkSettingsOutputWithContext(ctx context.Context) NetworkSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSettingsOutput)
}

func (i *NetworkSettings) ToOutput(ctx context.Context) pulumix.Output[*NetworkSettings] {
	return pulumix.Output[*NetworkSettings]{
		OutputState: i.ToNetworkSettingsOutputWithContext(ctx).OutputState,
	}
}

type NetworkSettingsOutput struct{ *pulumi.OutputState }

func (NetworkSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSettings)(nil)).Elem()
}

func (o NetworkSettingsOutput) ToNetworkSettingsOutput() NetworkSettingsOutput {
	return o
}

func (o NetworkSettingsOutput) ToNetworkSettingsOutputWithContext(ctx context.Context) NetworkSettingsOutput {
	return o
}

func (o NetworkSettingsOutput) ToOutput(ctx context.Context) pulumix.Output[*NetworkSettings] {
	return pulumix.Output[*NetworkSettings]{
		OutputState: o.OutputState,
	}
}

func (o NetworkSettingsOutput) AssociatedPortalArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkSettings) pulumi.StringArrayOutput { return v.AssociatedPortalArns }).(pulumi.StringArrayOutput)
}

func (o NetworkSettingsOutput) NetworkSettingsArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSettings) pulumi.StringOutput { return v.NetworkSettingsArn }).(pulumi.StringOutput)
}

func (o NetworkSettingsOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkSettings) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o NetworkSettingsOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkSettings) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o NetworkSettingsOutput) Tags() NetworkSettingsTagArrayOutput {
	return o.ApplyT(func(v *NetworkSettings) NetworkSettingsTagArrayOutput { return v.Tags }).(NetworkSettingsTagArrayOutput)
}

func (o NetworkSettingsOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSettings) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSettingsInput)(nil)).Elem(), &NetworkSettings{})
	pulumi.RegisterOutputType(NetworkSettingsOutput{})
}
