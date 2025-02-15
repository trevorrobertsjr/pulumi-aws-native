// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package panorama

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type ApplicationInstanceManifestOverridesPayload struct {
	PayloadData *string `pulumi:"payloadData"`
}

// ApplicationInstanceManifestOverridesPayloadInput is an input type that accepts ApplicationInstanceManifestOverridesPayloadArgs and ApplicationInstanceManifestOverridesPayloadOutput values.
// You can construct a concrete instance of `ApplicationInstanceManifestOverridesPayloadInput` via:
//
//	ApplicationInstanceManifestOverridesPayloadArgs{...}
type ApplicationInstanceManifestOverridesPayloadInput interface {
	pulumi.Input

	ToApplicationInstanceManifestOverridesPayloadOutput() ApplicationInstanceManifestOverridesPayloadOutput
	ToApplicationInstanceManifestOverridesPayloadOutputWithContext(context.Context) ApplicationInstanceManifestOverridesPayloadOutput
}

type ApplicationInstanceManifestOverridesPayloadArgs struct {
	PayloadData pulumi.StringPtrInput `pulumi:"payloadData"`
}

func (ApplicationInstanceManifestOverridesPayloadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceManifestOverridesPayload)(nil)).Elem()
}

func (i ApplicationInstanceManifestOverridesPayloadArgs) ToApplicationInstanceManifestOverridesPayloadOutput() ApplicationInstanceManifestOverridesPayloadOutput {
	return i.ToApplicationInstanceManifestOverridesPayloadOutputWithContext(context.Background())
}

func (i ApplicationInstanceManifestOverridesPayloadArgs) ToApplicationInstanceManifestOverridesPayloadOutputWithContext(ctx context.Context) ApplicationInstanceManifestOverridesPayloadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInstanceManifestOverridesPayloadOutput)
}

func (i ApplicationInstanceManifestOverridesPayloadArgs) ToOutput(ctx context.Context) pulumix.Output[ApplicationInstanceManifestOverridesPayload] {
	return pulumix.Output[ApplicationInstanceManifestOverridesPayload]{
		OutputState: i.ToApplicationInstanceManifestOverridesPayloadOutputWithContext(ctx).OutputState,
	}
}

func (i ApplicationInstanceManifestOverridesPayloadArgs) ToApplicationInstanceManifestOverridesPayloadPtrOutput() ApplicationInstanceManifestOverridesPayloadPtrOutput {
	return i.ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(context.Background())
}

func (i ApplicationInstanceManifestOverridesPayloadArgs) ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(ctx context.Context) ApplicationInstanceManifestOverridesPayloadPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInstanceManifestOverridesPayloadOutput).ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(ctx)
}

// ApplicationInstanceManifestOverridesPayloadPtrInput is an input type that accepts ApplicationInstanceManifestOverridesPayloadArgs, ApplicationInstanceManifestOverridesPayloadPtr and ApplicationInstanceManifestOverridesPayloadPtrOutput values.
// You can construct a concrete instance of `ApplicationInstanceManifestOverridesPayloadPtrInput` via:
//
//	        ApplicationInstanceManifestOverridesPayloadArgs{...}
//
//	or:
//
//	        nil
type ApplicationInstanceManifestOverridesPayloadPtrInput interface {
	pulumi.Input

	ToApplicationInstanceManifestOverridesPayloadPtrOutput() ApplicationInstanceManifestOverridesPayloadPtrOutput
	ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(context.Context) ApplicationInstanceManifestOverridesPayloadPtrOutput
}

type applicationInstanceManifestOverridesPayloadPtrType ApplicationInstanceManifestOverridesPayloadArgs

func ApplicationInstanceManifestOverridesPayloadPtr(v *ApplicationInstanceManifestOverridesPayloadArgs) ApplicationInstanceManifestOverridesPayloadPtrInput {
	return (*applicationInstanceManifestOverridesPayloadPtrType)(v)
}

func (*applicationInstanceManifestOverridesPayloadPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInstanceManifestOverridesPayload)(nil)).Elem()
}

func (i *applicationInstanceManifestOverridesPayloadPtrType) ToApplicationInstanceManifestOverridesPayloadPtrOutput() ApplicationInstanceManifestOverridesPayloadPtrOutput {
	return i.ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(context.Background())
}

func (i *applicationInstanceManifestOverridesPayloadPtrType) ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(ctx context.Context) ApplicationInstanceManifestOverridesPayloadPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInstanceManifestOverridesPayloadPtrOutput)
}

func (i *applicationInstanceManifestOverridesPayloadPtrType) ToOutput(ctx context.Context) pulumix.Output[*ApplicationInstanceManifestOverridesPayload] {
	return pulumix.Output[*ApplicationInstanceManifestOverridesPayload]{
		OutputState: i.ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(ctx).OutputState,
	}
}

type ApplicationInstanceManifestOverridesPayloadOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceManifestOverridesPayloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceManifestOverridesPayload)(nil)).Elem()
}

func (o ApplicationInstanceManifestOverridesPayloadOutput) ToApplicationInstanceManifestOverridesPayloadOutput() ApplicationInstanceManifestOverridesPayloadOutput {
	return o
}

func (o ApplicationInstanceManifestOverridesPayloadOutput) ToApplicationInstanceManifestOverridesPayloadOutputWithContext(ctx context.Context) ApplicationInstanceManifestOverridesPayloadOutput {
	return o
}

func (o ApplicationInstanceManifestOverridesPayloadOutput) ToApplicationInstanceManifestOverridesPayloadPtrOutput() ApplicationInstanceManifestOverridesPayloadPtrOutput {
	return o.ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceManifestOverridesPayloadOutput) ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(ctx context.Context) ApplicationInstanceManifestOverridesPayloadPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInstanceManifestOverridesPayload) *ApplicationInstanceManifestOverridesPayload {
		return &v
	}).(ApplicationInstanceManifestOverridesPayloadPtrOutput)
}

func (o ApplicationInstanceManifestOverridesPayloadOutput) ToOutput(ctx context.Context) pulumix.Output[ApplicationInstanceManifestOverridesPayload] {
	return pulumix.Output[ApplicationInstanceManifestOverridesPayload]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationInstanceManifestOverridesPayloadOutput) PayloadData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInstanceManifestOverridesPayload) *string { return v.PayloadData }).(pulumi.StringPtrOutput)
}

type ApplicationInstanceManifestOverridesPayloadPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceManifestOverridesPayloadPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInstanceManifestOverridesPayload)(nil)).Elem()
}

func (o ApplicationInstanceManifestOverridesPayloadPtrOutput) ToApplicationInstanceManifestOverridesPayloadPtrOutput() ApplicationInstanceManifestOverridesPayloadPtrOutput {
	return o
}

func (o ApplicationInstanceManifestOverridesPayloadPtrOutput) ToApplicationInstanceManifestOverridesPayloadPtrOutputWithContext(ctx context.Context) ApplicationInstanceManifestOverridesPayloadPtrOutput {
	return o
}

func (o ApplicationInstanceManifestOverridesPayloadPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationInstanceManifestOverridesPayload] {
	return pulumix.Output[*ApplicationInstanceManifestOverridesPayload]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationInstanceManifestOverridesPayloadPtrOutput) Elem() ApplicationInstanceManifestOverridesPayloadOutput {
	return o.ApplyT(func(v *ApplicationInstanceManifestOverridesPayload) ApplicationInstanceManifestOverridesPayload {
		if v != nil {
			return *v
		}
		var ret ApplicationInstanceManifestOverridesPayload
		return ret
	}).(ApplicationInstanceManifestOverridesPayloadOutput)
}

func (o ApplicationInstanceManifestOverridesPayloadPtrOutput) PayloadData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInstanceManifestOverridesPayload) *string {
		if v == nil {
			return nil
		}
		return v.PayloadData
	}).(pulumi.StringPtrOutput)
}

type ApplicationInstanceManifestPayload struct {
	PayloadData *string `pulumi:"payloadData"`
}

// ApplicationInstanceManifestPayloadInput is an input type that accepts ApplicationInstanceManifestPayloadArgs and ApplicationInstanceManifestPayloadOutput values.
// You can construct a concrete instance of `ApplicationInstanceManifestPayloadInput` via:
//
//	ApplicationInstanceManifestPayloadArgs{...}
type ApplicationInstanceManifestPayloadInput interface {
	pulumi.Input

	ToApplicationInstanceManifestPayloadOutput() ApplicationInstanceManifestPayloadOutput
	ToApplicationInstanceManifestPayloadOutputWithContext(context.Context) ApplicationInstanceManifestPayloadOutput
}

type ApplicationInstanceManifestPayloadArgs struct {
	PayloadData pulumi.StringPtrInput `pulumi:"payloadData"`
}

func (ApplicationInstanceManifestPayloadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceManifestPayload)(nil)).Elem()
}

func (i ApplicationInstanceManifestPayloadArgs) ToApplicationInstanceManifestPayloadOutput() ApplicationInstanceManifestPayloadOutput {
	return i.ToApplicationInstanceManifestPayloadOutputWithContext(context.Background())
}

func (i ApplicationInstanceManifestPayloadArgs) ToApplicationInstanceManifestPayloadOutputWithContext(ctx context.Context) ApplicationInstanceManifestPayloadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInstanceManifestPayloadOutput)
}

func (i ApplicationInstanceManifestPayloadArgs) ToOutput(ctx context.Context) pulumix.Output[ApplicationInstanceManifestPayload] {
	return pulumix.Output[ApplicationInstanceManifestPayload]{
		OutputState: i.ToApplicationInstanceManifestPayloadOutputWithContext(ctx).OutputState,
	}
}

type ApplicationInstanceManifestPayloadOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceManifestPayloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceManifestPayload)(nil)).Elem()
}

func (o ApplicationInstanceManifestPayloadOutput) ToApplicationInstanceManifestPayloadOutput() ApplicationInstanceManifestPayloadOutput {
	return o
}

func (o ApplicationInstanceManifestPayloadOutput) ToApplicationInstanceManifestPayloadOutputWithContext(ctx context.Context) ApplicationInstanceManifestPayloadOutput {
	return o
}

func (o ApplicationInstanceManifestPayloadOutput) ToOutput(ctx context.Context) pulumix.Output[ApplicationInstanceManifestPayload] {
	return pulumix.Output[ApplicationInstanceManifestPayload]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationInstanceManifestPayloadOutput) PayloadData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInstanceManifestPayload) *string { return v.PayloadData }).(pulumi.StringPtrOutput)
}

type ApplicationInstanceTag struct {
	// A string used to identify this tag
	Key string `pulumi:"key"`
	// A string containing the value for the tag
	Value string `pulumi:"value"`
}

// ApplicationInstanceTagInput is an input type that accepts ApplicationInstanceTagArgs and ApplicationInstanceTagOutput values.
// You can construct a concrete instance of `ApplicationInstanceTagInput` via:
//
//	ApplicationInstanceTagArgs{...}
type ApplicationInstanceTagInput interface {
	pulumi.Input

	ToApplicationInstanceTagOutput() ApplicationInstanceTagOutput
	ToApplicationInstanceTagOutputWithContext(context.Context) ApplicationInstanceTagOutput
}

type ApplicationInstanceTagArgs struct {
	// A string used to identify this tag
	Key pulumi.StringInput `pulumi:"key"`
	// A string containing the value for the tag
	Value pulumi.StringInput `pulumi:"value"`
}

func (ApplicationInstanceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceTag)(nil)).Elem()
}

func (i ApplicationInstanceTagArgs) ToApplicationInstanceTagOutput() ApplicationInstanceTagOutput {
	return i.ToApplicationInstanceTagOutputWithContext(context.Background())
}

func (i ApplicationInstanceTagArgs) ToApplicationInstanceTagOutputWithContext(ctx context.Context) ApplicationInstanceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInstanceTagOutput)
}

func (i ApplicationInstanceTagArgs) ToOutput(ctx context.Context) pulumix.Output[ApplicationInstanceTag] {
	return pulumix.Output[ApplicationInstanceTag]{
		OutputState: i.ToApplicationInstanceTagOutputWithContext(ctx).OutputState,
	}
}

// ApplicationInstanceTagArrayInput is an input type that accepts ApplicationInstanceTagArray and ApplicationInstanceTagArrayOutput values.
// You can construct a concrete instance of `ApplicationInstanceTagArrayInput` via:
//
//	ApplicationInstanceTagArray{ ApplicationInstanceTagArgs{...} }
type ApplicationInstanceTagArrayInput interface {
	pulumi.Input

	ToApplicationInstanceTagArrayOutput() ApplicationInstanceTagArrayOutput
	ToApplicationInstanceTagArrayOutputWithContext(context.Context) ApplicationInstanceTagArrayOutput
}

type ApplicationInstanceTagArray []ApplicationInstanceTagInput

func (ApplicationInstanceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationInstanceTag)(nil)).Elem()
}

func (i ApplicationInstanceTagArray) ToApplicationInstanceTagArrayOutput() ApplicationInstanceTagArrayOutput {
	return i.ToApplicationInstanceTagArrayOutputWithContext(context.Background())
}

func (i ApplicationInstanceTagArray) ToApplicationInstanceTagArrayOutputWithContext(ctx context.Context) ApplicationInstanceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInstanceTagArrayOutput)
}

func (i ApplicationInstanceTagArray) ToOutput(ctx context.Context) pulumix.Output[[]ApplicationInstanceTag] {
	return pulumix.Output[[]ApplicationInstanceTag]{
		OutputState: i.ToApplicationInstanceTagArrayOutputWithContext(ctx).OutputState,
	}
}

type ApplicationInstanceTagOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceTag)(nil)).Elem()
}

func (o ApplicationInstanceTagOutput) ToApplicationInstanceTagOutput() ApplicationInstanceTagOutput {
	return o
}

func (o ApplicationInstanceTagOutput) ToApplicationInstanceTagOutputWithContext(ctx context.Context) ApplicationInstanceTagOutput {
	return o
}

func (o ApplicationInstanceTagOutput) ToOutput(ctx context.Context) pulumix.Output[ApplicationInstanceTag] {
	return pulumix.Output[ApplicationInstanceTag]{
		OutputState: o.OutputState,
	}
}

// A string used to identify this tag
func (o ApplicationInstanceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationInstanceTag) string { return v.Key }).(pulumi.StringOutput)
}

// A string containing the value for the tag
func (o ApplicationInstanceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationInstanceTag) string { return v.Value }).(pulumi.StringOutput)
}

type ApplicationInstanceTagArrayOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationInstanceTag)(nil)).Elem()
}

func (o ApplicationInstanceTagArrayOutput) ToApplicationInstanceTagArrayOutput() ApplicationInstanceTagArrayOutput {
	return o
}

func (o ApplicationInstanceTagArrayOutput) ToApplicationInstanceTagArrayOutputWithContext(ctx context.Context) ApplicationInstanceTagArrayOutput {
	return o
}

func (o ApplicationInstanceTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ApplicationInstanceTag] {
	return pulumix.Output[[]ApplicationInstanceTag]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationInstanceTagArrayOutput) Index(i pulumi.IntInput) ApplicationInstanceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationInstanceTag {
		return vs[0].([]ApplicationInstanceTag)[vs[1].(int)]
	}).(ApplicationInstanceTagOutput)
}

type PackageStorageLocation struct {
	BinaryPrefixLocation    *string `pulumi:"binaryPrefixLocation"`
	Bucket                  *string `pulumi:"bucket"`
	GeneratedPrefixLocation *string `pulumi:"generatedPrefixLocation"`
	ManifestPrefixLocation  *string `pulumi:"manifestPrefixLocation"`
	RepoPrefixLocation      *string `pulumi:"repoPrefixLocation"`
}

// PackageStorageLocationInput is an input type that accepts PackageStorageLocationArgs and PackageStorageLocationOutput values.
// You can construct a concrete instance of `PackageStorageLocationInput` via:
//
//	PackageStorageLocationArgs{...}
type PackageStorageLocationInput interface {
	pulumi.Input

	ToPackageStorageLocationOutput() PackageStorageLocationOutput
	ToPackageStorageLocationOutputWithContext(context.Context) PackageStorageLocationOutput
}

type PackageStorageLocationArgs struct {
	BinaryPrefixLocation    pulumi.StringPtrInput `pulumi:"binaryPrefixLocation"`
	Bucket                  pulumi.StringPtrInput `pulumi:"bucket"`
	GeneratedPrefixLocation pulumi.StringPtrInput `pulumi:"generatedPrefixLocation"`
	ManifestPrefixLocation  pulumi.StringPtrInput `pulumi:"manifestPrefixLocation"`
	RepoPrefixLocation      pulumi.StringPtrInput `pulumi:"repoPrefixLocation"`
}

func (PackageStorageLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageStorageLocation)(nil)).Elem()
}

func (i PackageStorageLocationArgs) ToPackageStorageLocationOutput() PackageStorageLocationOutput {
	return i.ToPackageStorageLocationOutputWithContext(context.Background())
}

func (i PackageStorageLocationArgs) ToPackageStorageLocationOutputWithContext(ctx context.Context) PackageStorageLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageStorageLocationOutput)
}

func (i PackageStorageLocationArgs) ToOutput(ctx context.Context) pulumix.Output[PackageStorageLocation] {
	return pulumix.Output[PackageStorageLocation]{
		OutputState: i.ToPackageStorageLocationOutputWithContext(ctx).OutputState,
	}
}

func (i PackageStorageLocationArgs) ToPackageStorageLocationPtrOutput() PackageStorageLocationPtrOutput {
	return i.ToPackageStorageLocationPtrOutputWithContext(context.Background())
}

func (i PackageStorageLocationArgs) ToPackageStorageLocationPtrOutputWithContext(ctx context.Context) PackageStorageLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageStorageLocationOutput).ToPackageStorageLocationPtrOutputWithContext(ctx)
}

// PackageStorageLocationPtrInput is an input type that accepts PackageStorageLocationArgs, PackageStorageLocationPtr and PackageStorageLocationPtrOutput values.
// You can construct a concrete instance of `PackageStorageLocationPtrInput` via:
//
//	        PackageStorageLocationArgs{...}
//
//	or:
//
//	        nil
type PackageStorageLocationPtrInput interface {
	pulumi.Input

	ToPackageStorageLocationPtrOutput() PackageStorageLocationPtrOutput
	ToPackageStorageLocationPtrOutputWithContext(context.Context) PackageStorageLocationPtrOutput
}

type packageStorageLocationPtrType PackageStorageLocationArgs

func PackageStorageLocationPtr(v *PackageStorageLocationArgs) PackageStorageLocationPtrInput {
	return (*packageStorageLocationPtrType)(v)
}

func (*packageStorageLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageStorageLocation)(nil)).Elem()
}

func (i *packageStorageLocationPtrType) ToPackageStorageLocationPtrOutput() PackageStorageLocationPtrOutput {
	return i.ToPackageStorageLocationPtrOutputWithContext(context.Background())
}

func (i *packageStorageLocationPtrType) ToPackageStorageLocationPtrOutputWithContext(ctx context.Context) PackageStorageLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageStorageLocationPtrOutput)
}

func (i *packageStorageLocationPtrType) ToOutput(ctx context.Context) pulumix.Output[*PackageStorageLocation] {
	return pulumix.Output[*PackageStorageLocation]{
		OutputState: i.ToPackageStorageLocationPtrOutputWithContext(ctx).OutputState,
	}
}

type PackageStorageLocationOutput struct{ *pulumi.OutputState }

func (PackageStorageLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageStorageLocation)(nil)).Elem()
}

func (o PackageStorageLocationOutput) ToPackageStorageLocationOutput() PackageStorageLocationOutput {
	return o
}

func (o PackageStorageLocationOutput) ToPackageStorageLocationOutputWithContext(ctx context.Context) PackageStorageLocationOutput {
	return o
}

func (o PackageStorageLocationOutput) ToPackageStorageLocationPtrOutput() PackageStorageLocationPtrOutput {
	return o.ToPackageStorageLocationPtrOutputWithContext(context.Background())
}

func (o PackageStorageLocationOutput) ToPackageStorageLocationPtrOutputWithContext(ctx context.Context) PackageStorageLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PackageStorageLocation) *PackageStorageLocation {
		return &v
	}).(PackageStorageLocationPtrOutput)
}

func (o PackageStorageLocationOutput) ToOutput(ctx context.Context) pulumix.Output[PackageStorageLocation] {
	return pulumix.Output[PackageStorageLocation]{
		OutputState: o.OutputState,
	}
}

func (o PackageStorageLocationOutput) BinaryPrefixLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageStorageLocation) *string { return v.BinaryPrefixLocation }).(pulumi.StringPtrOutput)
}

func (o PackageStorageLocationOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageStorageLocation) *string { return v.Bucket }).(pulumi.StringPtrOutput)
}

func (o PackageStorageLocationOutput) GeneratedPrefixLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageStorageLocation) *string { return v.GeneratedPrefixLocation }).(pulumi.StringPtrOutput)
}

func (o PackageStorageLocationOutput) ManifestPrefixLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageStorageLocation) *string { return v.ManifestPrefixLocation }).(pulumi.StringPtrOutput)
}

func (o PackageStorageLocationOutput) RepoPrefixLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageStorageLocation) *string { return v.RepoPrefixLocation }).(pulumi.StringPtrOutput)
}

type PackageStorageLocationPtrOutput struct{ *pulumi.OutputState }

func (PackageStorageLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageStorageLocation)(nil)).Elem()
}

func (o PackageStorageLocationPtrOutput) ToPackageStorageLocationPtrOutput() PackageStorageLocationPtrOutput {
	return o
}

func (o PackageStorageLocationPtrOutput) ToPackageStorageLocationPtrOutputWithContext(ctx context.Context) PackageStorageLocationPtrOutput {
	return o
}

func (o PackageStorageLocationPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*PackageStorageLocation] {
	return pulumix.Output[*PackageStorageLocation]{
		OutputState: o.OutputState,
	}
}

func (o PackageStorageLocationPtrOutput) Elem() PackageStorageLocationOutput {
	return o.ApplyT(func(v *PackageStorageLocation) PackageStorageLocation {
		if v != nil {
			return *v
		}
		var ret PackageStorageLocation
		return ret
	}).(PackageStorageLocationOutput)
}

func (o PackageStorageLocationPtrOutput) BinaryPrefixLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageStorageLocation) *string {
		if v == nil {
			return nil
		}
		return v.BinaryPrefixLocation
	}).(pulumi.StringPtrOutput)
}

func (o PackageStorageLocationPtrOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageStorageLocation) *string {
		if v == nil {
			return nil
		}
		return v.Bucket
	}).(pulumi.StringPtrOutput)
}

func (o PackageStorageLocationPtrOutput) GeneratedPrefixLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageStorageLocation) *string {
		if v == nil {
			return nil
		}
		return v.GeneratedPrefixLocation
	}).(pulumi.StringPtrOutput)
}

func (o PackageStorageLocationPtrOutput) ManifestPrefixLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageStorageLocation) *string {
		if v == nil {
			return nil
		}
		return v.ManifestPrefixLocation
	}).(pulumi.StringPtrOutput)
}

func (o PackageStorageLocationPtrOutput) RepoPrefixLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageStorageLocation) *string {
		if v == nil {
			return nil
		}
		return v.RepoPrefixLocation
	}).(pulumi.StringPtrOutput)
}

type PackageTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// PackageTagInput is an input type that accepts PackageTagArgs and PackageTagOutput values.
// You can construct a concrete instance of `PackageTagInput` via:
//
//	PackageTagArgs{...}
type PackageTagInput interface {
	pulumi.Input

	ToPackageTagOutput() PackageTagOutput
	ToPackageTagOutputWithContext(context.Context) PackageTagOutput
}

type PackageTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (PackageTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageTag)(nil)).Elem()
}

func (i PackageTagArgs) ToPackageTagOutput() PackageTagOutput {
	return i.ToPackageTagOutputWithContext(context.Background())
}

func (i PackageTagArgs) ToPackageTagOutputWithContext(ctx context.Context) PackageTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageTagOutput)
}

func (i PackageTagArgs) ToOutput(ctx context.Context) pulumix.Output[PackageTag] {
	return pulumix.Output[PackageTag]{
		OutputState: i.ToPackageTagOutputWithContext(ctx).OutputState,
	}
}

// PackageTagArrayInput is an input type that accepts PackageTagArray and PackageTagArrayOutput values.
// You can construct a concrete instance of `PackageTagArrayInput` via:
//
//	PackageTagArray{ PackageTagArgs{...} }
type PackageTagArrayInput interface {
	pulumi.Input

	ToPackageTagArrayOutput() PackageTagArrayOutput
	ToPackageTagArrayOutputWithContext(context.Context) PackageTagArrayOutput
}

type PackageTagArray []PackageTagInput

func (PackageTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PackageTag)(nil)).Elem()
}

func (i PackageTagArray) ToPackageTagArrayOutput() PackageTagArrayOutput {
	return i.ToPackageTagArrayOutputWithContext(context.Background())
}

func (i PackageTagArray) ToPackageTagArrayOutputWithContext(ctx context.Context) PackageTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageTagArrayOutput)
}

func (i PackageTagArray) ToOutput(ctx context.Context) pulumix.Output[[]PackageTag] {
	return pulumix.Output[[]PackageTag]{
		OutputState: i.ToPackageTagArrayOutputWithContext(ctx).OutputState,
	}
}

type PackageTagOutput struct{ *pulumi.OutputState }

func (PackageTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageTag)(nil)).Elem()
}

func (o PackageTagOutput) ToPackageTagOutput() PackageTagOutput {
	return o
}

func (o PackageTagOutput) ToPackageTagOutputWithContext(ctx context.Context) PackageTagOutput {
	return o
}

func (o PackageTagOutput) ToOutput(ctx context.Context) pulumix.Output[PackageTag] {
	return pulumix.Output[PackageTag]{
		OutputState: o.OutputState,
	}
}

func (o PackageTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v PackageTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o PackageTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v PackageTag) string { return v.Value }).(pulumi.StringOutput)
}

type PackageTagArrayOutput struct{ *pulumi.OutputState }

func (PackageTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PackageTag)(nil)).Elem()
}

func (o PackageTagArrayOutput) ToPackageTagArrayOutput() PackageTagArrayOutput {
	return o
}

func (o PackageTagArrayOutput) ToPackageTagArrayOutputWithContext(ctx context.Context) PackageTagArrayOutput {
	return o
}

func (o PackageTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]PackageTag] {
	return pulumix.Output[[]PackageTag]{
		OutputState: o.OutputState,
	}
}

func (o PackageTagArrayOutput) Index(i pulumi.IntInput) PackageTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PackageTag {
		return vs[0].([]PackageTag)[vs[1].(int)]
	}).(PackageTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInstanceManifestOverridesPayloadInput)(nil)).Elem(), ApplicationInstanceManifestOverridesPayloadArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInstanceManifestOverridesPayloadPtrInput)(nil)).Elem(), ApplicationInstanceManifestOverridesPayloadArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInstanceManifestPayloadInput)(nil)).Elem(), ApplicationInstanceManifestPayloadArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInstanceTagInput)(nil)).Elem(), ApplicationInstanceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInstanceTagArrayInput)(nil)).Elem(), ApplicationInstanceTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PackageStorageLocationInput)(nil)).Elem(), PackageStorageLocationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PackageStorageLocationPtrInput)(nil)).Elem(), PackageStorageLocationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PackageTagInput)(nil)).Elem(), PackageTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PackageTagArrayInput)(nil)).Elem(), PackageTagArray{})
	pulumi.RegisterOutputType(ApplicationInstanceManifestOverridesPayloadOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceManifestOverridesPayloadPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceManifestPayloadOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceTagOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceTagArrayOutput{})
	pulumi.RegisterOutputType(PackageStorageLocationOutput{})
	pulumi.RegisterOutputType(PackageStorageLocationPtrOutput{})
	pulumi.RegisterOutputType(PackageTagOutput{})
	pulumi.RegisterOutputType(PackageTagArrayOutput{})
}
