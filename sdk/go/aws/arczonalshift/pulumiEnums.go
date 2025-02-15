// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package arczonalshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type ZonalAutoshiftConfigurationControlConditionType string

const (
	ZonalAutoshiftConfigurationControlConditionTypeCloudwatch = ZonalAutoshiftConfigurationControlConditionType("CLOUDWATCH")
)

func (ZonalAutoshiftConfigurationControlConditionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonalAutoshiftConfigurationControlConditionType)(nil)).Elem()
}

func (e ZonalAutoshiftConfigurationControlConditionType) ToZonalAutoshiftConfigurationControlConditionTypeOutput() ZonalAutoshiftConfigurationControlConditionTypeOutput {
	return pulumi.ToOutput(e).(ZonalAutoshiftConfigurationControlConditionTypeOutput)
}

func (e ZonalAutoshiftConfigurationControlConditionType) ToZonalAutoshiftConfigurationControlConditionTypeOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationControlConditionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ZonalAutoshiftConfigurationControlConditionTypeOutput)
}

func (e ZonalAutoshiftConfigurationControlConditionType) ToZonalAutoshiftConfigurationControlConditionTypePtrOutput() ZonalAutoshiftConfigurationControlConditionTypePtrOutput {
	return e.ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(context.Background())
}

func (e ZonalAutoshiftConfigurationControlConditionType) ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationControlConditionTypePtrOutput {
	return ZonalAutoshiftConfigurationControlConditionType(e).ToZonalAutoshiftConfigurationControlConditionTypeOutputWithContext(ctx).ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(ctx)
}

func (e ZonalAutoshiftConfigurationControlConditionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZonalAutoshiftConfigurationControlConditionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZonalAutoshiftConfigurationControlConditionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ZonalAutoshiftConfigurationControlConditionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ZonalAutoshiftConfigurationControlConditionTypeOutput struct{ *pulumi.OutputState }

func (ZonalAutoshiftConfigurationControlConditionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonalAutoshiftConfigurationControlConditionType)(nil)).Elem()
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToZonalAutoshiftConfigurationControlConditionTypeOutput() ZonalAutoshiftConfigurationControlConditionTypeOutput {
	return o
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToZonalAutoshiftConfigurationControlConditionTypeOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationControlConditionTypeOutput {
	return o
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToZonalAutoshiftConfigurationControlConditionTypePtrOutput() ZonalAutoshiftConfigurationControlConditionTypePtrOutput {
	return o.ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(context.Background())
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationControlConditionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ZonalAutoshiftConfigurationControlConditionType) *ZonalAutoshiftConfigurationControlConditionType {
		return &v
	}).(ZonalAutoshiftConfigurationControlConditionTypePtrOutput)
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToOutput(ctx context.Context) pulumix.Output[ZonalAutoshiftConfigurationControlConditionType] {
	return pulumix.Output[ZonalAutoshiftConfigurationControlConditionType]{
		OutputState: o.OutputState,
	}
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZonalAutoshiftConfigurationControlConditionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZonalAutoshiftConfigurationControlConditionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZonalAutoshiftConfigurationControlConditionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ZonalAutoshiftConfigurationControlConditionTypePtrOutput struct{ *pulumi.OutputState }

func (ZonalAutoshiftConfigurationControlConditionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZonalAutoshiftConfigurationControlConditionType)(nil)).Elem()
}

func (o ZonalAutoshiftConfigurationControlConditionTypePtrOutput) ToZonalAutoshiftConfigurationControlConditionTypePtrOutput() ZonalAutoshiftConfigurationControlConditionTypePtrOutput {
	return o
}

func (o ZonalAutoshiftConfigurationControlConditionTypePtrOutput) ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationControlConditionTypePtrOutput {
	return o
}

func (o ZonalAutoshiftConfigurationControlConditionTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ZonalAutoshiftConfigurationControlConditionType] {
	return pulumix.Output[*ZonalAutoshiftConfigurationControlConditionType]{
		OutputState: o.OutputState,
	}
}

func (o ZonalAutoshiftConfigurationControlConditionTypePtrOutput) Elem() ZonalAutoshiftConfigurationControlConditionTypeOutput {
	return o.ApplyT(func(v *ZonalAutoshiftConfigurationControlConditionType) ZonalAutoshiftConfigurationControlConditionType {
		if v != nil {
			return *v
		}
		var ret ZonalAutoshiftConfigurationControlConditionType
		return ret
	}).(ZonalAutoshiftConfigurationControlConditionTypeOutput)
}

func (o ZonalAutoshiftConfigurationControlConditionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZonalAutoshiftConfigurationControlConditionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ZonalAutoshiftConfigurationControlConditionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ZonalAutoshiftConfigurationControlConditionTypeInput is an input type that accepts ZonalAutoshiftConfigurationControlConditionTypeArgs and ZonalAutoshiftConfigurationControlConditionTypeOutput values.
// You can construct a concrete instance of `ZonalAutoshiftConfigurationControlConditionTypeInput` via:
//
//	ZonalAutoshiftConfigurationControlConditionTypeArgs{...}
type ZonalAutoshiftConfigurationControlConditionTypeInput interface {
	pulumi.Input

	ToZonalAutoshiftConfigurationControlConditionTypeOutput() ZonalAutoshiftConfigurationControlConditionTypeOutput
	ToZonalAutoshiftConfigurationControlConditionTypeOutputWithContext(context.Context) ZonalAutoshiftConfigurationControlConditionTypeOutput
}

var zonalAutoshiftConfigurationControlConditionTypePtrType = reflect.TypeOf((**ZonalAutoshiftConfigurationControlConditionType)(nil)).Elem()

type ZonalAutoshiftConfigurationControlConditionTypePtrInput interface {
	pulumi.Input

	ToZonalAutoshiftConfigurationControlConditionTypePtrOutput() ZonalAutoshiftConfigurationControlConditionTypePtrOutput
	ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(context.Context) ZonalAutoshiftConfigurationControlConditionTypePtrOutput
}

type zonalAutoshiftConfigurationControlConditionTypePtr string

func ZonalAutoshiftConfigurationControlConditionTypePtr(v string) ZonalAutoshiftConfigurationControlConditionTypePtrInput {
	return (*zonalAutoshiftConfigurationControlConditionTypePtr)(&v)
}

func (*zonalAutoshiftConfigurationControlConditionTypePtr) ElementType() reflect.Type {
	return zonalAutoshiftConfigurationControlConditionTypePtrType
}

func (in *zonalAutoshiftConfigurationControlConditionTypePtr) ToZonalAutoshiftConfigurationControlConditionTypePtrOutput() ZonalAutoshiftConfigurationControlConditionTypePtrOutput {
	return pulumi.ToOutput(in).(ZonalAutoshiftConfigurationControlConditionTypePtrOutput)
}

func (in *zonalAutoshiftConfigurationControlConditionTypePtr) ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationControlConditionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ZonalAutoshiftConfigurationControlConditionTypePtrOutput)
}

func (in *zonalAutoshiftConfigurationControlConditionTypePtr) ToOutput(ctx context.Context) pulumix.Output[*ZonalAutoshiftConfigurationControlConditionType] {
	return pulumix.Output[*ZonalAutoshiftConfigurationControlConditionType]{
		OutputState: in.ToZonalAutoshiftConfigurationControlConditionTypePtrOutputWithContext(ctx).OutputState,
	}
}

type ZonalAutoshiftConfigurationZonalAutoshiftStatus string

const (
	ZonalAutoshiftConfigurationZonalAutoshiftStatusEnabled  = ZonalAutoshiftConfigurationZonalAutoshiftStatus("ENABLED")
	ZonalAutoshiftConfigurationZonalAutoshiftStatusDisabled = ZonalAutoshiftConfigurationZonalAutoshiftStatus("DISABLED")
)

func (ZonalAutoshiftConfigurationZonalAutoshiftStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonalAutoshiftConfigurationZonalAutoshiftStatus)(nil)).Elem()
}

func (e ZonalAutoshiftConfigurationZonalAutoshiftStatus) ToZonalAutoshiftConfigurationZonalAutoshiftStatusOutput() ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput {
	return pulumi.ToOutput(e).(ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput)
}

func (e ZonalAutoshiftConfigurationZonalAutoshiftStatus) ToZonalAutoshiftConfigurationZonalAutoshiftStatusOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput)
}

func (e ZonalAutoshiftConfigurationZonalAutoshiftStatus) ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput() ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return e.ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(context.Background())
}

func (e ZonalAutoshiftConfigurationZonalAutoshiftStatus) ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return ZonalAutoshiftConfigurationZonalAutoshiftStatus(e).ToZonalAutoshiftConfigurationZonalAutoshiftStatusOutputWithContext(ctx).ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(ctx)
}

func (e ZonalAutoshiftConfigurationZonalAutoshiftStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZonalAutoshiftConfigurationZonalAutoshiftStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZonalAutoshiftConfigurationZonalAutoshiftStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ZonalAutoshiftConfigurationZonalAutoshiftStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput struct{ *pulumi.OutputState }

func (ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonalAutoshiftConfigurationZonalAutoshiftStatus)(nil)).Elem()
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToZonalAutoshiftConfigurationZonalAutoshiftStatusOutput() ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput {
	return o
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToZonalAutoshiftConfigurationZonalAutoshiftStatusOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput {
	return o
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput() ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return o.ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(context.Background())
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ZonalAutoshiftConfigurationZonalAutoshiftStatus) *ZonalAutoshiftConfigurationZonalAutoshiftStatus {
		return &v
	}).(ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput)
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToOutput(ctx context.Context) pulumix.Output[ZonalAutoshiftConfigurationZonalAutoshiftStatus] {
	return pulumix.Output[ZonalAutoshiftConfigurationZonalAutoshiftStatus]{
		OutputState: o.OutputState,
	}
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZonalAutoshiftConfigurationZonalAutoshiftStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZonalAutoshiftConfigurationZonalAutoshiftStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput struct{ *pulumi.OutputState }

func (ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZonalAutoshiftConfigurationZonalAutoshiftStatus)(nil)).Elem()
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput) ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput() ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return o
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput) ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return o
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ZonalAutoshiftConfigurationZonalAutoshiftStatus] {
	return pulumix.Output[*ZonalAutoshiftConfigurationZonalAutoshiftStatus]{
		OutputState: o.OutputState,
	}
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput) Elem() ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput {
	return o.ApplyT(func(v *ZonalAutoshiftConfigurationZonalAutoshiftStatus) ZonalAutoshiftConfigurationZonalAutoshiftStatus {
		if v != nil {
			return *v
		}
		var ret ZonalAutoshiftConfigurationZonalAutoshiftStatus
		return ret
	}).(ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput)
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ZonalAutoshiftConfigurationZonalAutoshiftStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ZonalAutoshiftConfigurationZonalAutoshiftStatusInput is an input type that accepts ZonalAutoshiftConfigurationZonalAutoshiftStatusArgs and ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput values.
// You can construct a concrete instance of `ZonalAutoshiftConfigurationZonalAutoshiftStatusInput` via:
//
//	ZonalAutoshiftConfigurationZonalAutoshiftStatusArgs{...}
type ZonalAutoshiftConfigurationZonalAutoshiftStatusInput interface {
	pulumi.Input

	ToZonalAutoshiftConfigurationZonalAutoshiftStatusOutput() ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput
	ToZonalAutoshiftConfigurationZonalAutoshiftStatusOutputWithContext(context.Context) ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput
}

var zonalAutoshiftConfigurationZonalAutoshiftStatusPtrType = reflect.TypeOf((**ZonalAutoshiftConfigurationZonalAutoshiftStatus)(nil)).Elem()

type ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrInput interface {
	pulumi.Input

	ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput() ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput
	ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(context.Context) ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput
}

type zonalAutoshiftConfigurationZonalAutoshiftStatusPtr string

func ZonalAutoshiftConfigurationZonalAutoshiftStatusPtr(v string) ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrInput {
	return (*zonalAutoshiftConfigurationZonalAutoshiftStatusPtr)(&v)
}

func (*zonalAutoshiftConfigurationZonalAutoshiftStatusPtr) ElementType() reflect.Type {
	return zonalAutoshiftConfigurationZonalAutoshiftStatusPtrType
}

func (in *zonalAutoshiftConfigurationZonalAutoshiftStatusPtr) ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput() ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return pulumi.ToOutput(in).(ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput)
}

func (in *zonalAutoshiftConfigurationZonalAutoshiftStatusPtr) ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(ctx context.Context) ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput)
}

func (in *zonalAutoshiftConfigurationZonalAutoshiftStatusPtr) ToOutput(ctx context.Context) pulumix.Output[*ZonalAutoshiftConfigurationZonalAutoshiftStatus] {
	return pulumix.Output[*ZonalAutoshiftConfigurationZonalAutoshiftStatus]{
		OutputState: in.ToZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZonalAutoshiftConfigurationControlConditionTypeInput)(nil)).Elem(), ZonalAutoshiftConfigurationControlConditionType("CLOUDWATCH"))
	pulumi.RegisterInputType(reflect.TypeOf((*ZonalAutoshiftConfigurationControlConditionTypePtrInput)(nil)).Elem(), ZonalAutoshiftConfigurationControlConditionType("CLOUDWATCH"))
	pulumi.RegisterInputType(reflect.TypeOf((*ZonalAutoshiftConfigurationZonalAutoshiftStatusInput)(nil)).Elem(), ZonalAutoshiftConfigurationZonalAutoshiftStatus("ENABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrInput)(nil)).Elem(), ZonalAutoshiftConfigurationZonalAutoshiftStatus("ENABLED"))
	pulumi.RegisterOutputType(ZonalAutoshiftConfigurationControlConditionTypeOutput{})
	pulumi.RegisterOutputType(ZonalAutoshiftConfigurationControlConditionTypePtrOutput{})
	pulumi.RegisterOutputType(ZonalAutoshiftConfigurationZonalAutoshiftStatusOutput{})
	pulumi.RegisterOutputType(ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput{})
}
