// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::DirectoryService::MicrosoftAD
//
// Deprecated: MicrosoftAd is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type MicrosoftAd struct {
	pulumi.CustomResourceState

	Alias          pulumi.StringOutput          `pulumi:"alias"`
	CreateAlias    pulumi.BoolPtrOutput         `pulumi:"createAlias"`
	DnsIpAddresses pulumi.StringArrayOutput     `pulumi:"dnsIpAddresses"`
	Edition        pulumi.StringPtrOutput       `pulumi:"edition"`
	EnableSso      pulumi.BoolPtrOutput         `pulumi:"enableSso"`
	Name           pulumi.StringOutput          `pulumi:"name"`
	Password       pulumi.StringOutput          `pulumi:"password"`
	ShortName      pulumi.StringPtrOutput       `pulumi:"shortName"`
	VpcSettings    MicrosoftAdVpcSettingsOutput `pulumi:"vpcSettings"`
}

// NewMicrosoftAd registers a new resource with the given unique name, arguments, and options.
func NewMicrosoftAd(ctx *pulumi.Context,
	name string, args *MicrosoftAdArgs, opts ...pulumi.ResourceOption) (*MicrosoftAd, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.VpcSettings == nil {
		return nil, errors.New("invalid value for required argument 'VpcSettings'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"createAlias",
		"edition",
		"name",
		"password",
		"shortName",
		"vpcSettings",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MicrosoftAd
	err := ctx.RegisterResource("aws-native:directoryservice:MicrosoftAd", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMicrosoftAd gets an existing MicrosoftAd resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMicrosoftAd(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MicrosoftAdState, opts ...pulumi.ResourceOption) (*MicrosoftAd, error) {
	var resource MicrosoftAd
	err := ctx.ReadResource("aws-native:directoryservice:MicrosoftAd", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MicrosoftAd resources.
type microsoftAdState struct {
}

type MicrosoftAdState struct {
}

func (MicrosoftAdState) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftAdState)(nil)).Elem()
}

type microsoftAdArgs struct {
	CreateAlias *bool                  `pulumi:"createAlias"`
	Edition     *string                `pulumi:"edition"`
	EnableSso   *bool                  `pulumi:"enableSso"`
	Name        *string                `pulumi:"name"`
	Password    string                 `pulumi:"password"`
	ShortName   *string                `pulumi:"shortName"`
	VpcSettings MicrosoftAdVpcSettings `pulumi:"vpcSettings"`
}

// The set of arguments for constructing a MicrosoftAd resource.
type MicrosoftAdArgs struct {
	CreateAlias pulumi.BoolPtrInput
	Edition     pulumi.StringPtrInput
	EnableSso   pulumi.BoolPtrInput
	Name        pulumi.StringPtrInput
	Password    pulumi.StringInput
	ShortName   pulumi.StringPtrInput
	VpcSettings MicrosoftAdVpcSettingsInput
}

func (MicrosoftAdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftAdArgs)(nil)).Elem()
}

type MicrosoftAdInput interface {
	pulumi.Input

	ToMicrosoftAdOutput() MicrosoftAdOutput
	ToMicrosoftAdOutputWithContext(ctx context.Context) MicrosoftAdOutput
}

func (*MicrosoftAd) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftAd)(nil)).Elem()
}

func (i *MicrosoftAd) ToMicrosoftAdOutput() MicrosoftAdOutput {
	return i.ToMicrosoftAdOutputWithContext(context.Background())
}

func (i *MicrosoftAd) ToMicrosoftAdOutputWithContext(ctx context.Context) MicrosoftAdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicrosoftAdOutput)
}

func (i *MicrosoftAd) ToOutput(ctx context.Context) pulumix.Output[*MicrosoftAd] {
	return pulumix.Output[*MicrosoftAd]{
		OutputState: i.ToMicrosoftAdOutputWithContext(ctx).OutputState,
	}
}

type MicrosoftAdOutput struct{ *pulumi.OutputState }

func (MicrosoftAdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftAd)(nil)).Elem()
}

func (o MicrosoftAdOutput) ToMicrosoftAdOutput() MicrosoftAdOutput {
	return o
}

func (o MicrosoftAdOutput) ToMicrosoftAdOutputWithContext(ctx context.Context) MicrosoftAdOutput {
	return o
}

func (o MicrosoftAdOutput) ToOutput(ctx context.Context) pulumix.Output[*MicrosoftAd] {
	return pulumix.Output[*MicrosoftAd]{
		OutputState: o.OutputState,
	}
}

func (o MicrosoftAdOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftAd) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

func (o MicrosoftAdOutput) CreateAlias() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MicrosoftAd) pulumi.BoolPtrOutput { return v.CreateAlias }).(pulumi.BoolPtrOutput)
}

func (o MicrosoftAdOutput) DnsIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftAd) pulumi.StringArrayOutput { return v.DnsIpAddresses }).(pulumi.StringArrayOutput)
}

func (o MicrosoftAdOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftAd) pulumi.StringPtrOutput { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o MicrosoftAdOutput) EnableSso() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MicrosoftAd) pulumi.BoolPtrOutput { return v.EnableSso }).(pulumi.BoolPtrOutput)
}

func (o MicrosoftAdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftAd) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MicrosoftAdOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftAd) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func (o MicrosoftAdOutput) ShortName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftAd) pulumi.StringPtrOutput { return v.ShortName }).(pulumi.StringPtrOutput)
}

func (o MicrosoftAdOutput) VpcSettings() MicrosoftAdVpcSettingsOutput {
	return o.ApplyT(func(v *MicrosoftAd) MicrosoftAdVpcSettingsOutput { return v.VpcSettings }).(MicrosoftAdVpcSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MicrosoftAdInput)(nil)).Elem(), &MicrosoftAd{})
	pulumi.RegisterOutputType(MicrosoftAdOutput{})
}
