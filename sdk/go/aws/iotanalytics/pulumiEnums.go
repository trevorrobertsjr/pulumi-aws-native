// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotanalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type DatasetResourceConfigurationComputeType string

const (
	DatasetResourceConfigurationComputeTypeAcu1 = DatasetResourceConfigurationComputeType("ACU_1")
	DatasetResourceConfigurationComputeTypeAcu2 = DatasetResourceConfigurationComputeType("ACU_2")
)

func (DatasetResourceConfigurationComputeType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetResourceConfigurationComputeType)(nil)).Elem()
}

func (e DatasetResourceConfigurationComputeType) ToDatasetResourceConfigurationComputeTypeOutput() DatasetResourceConfigurationComputeTypeOutput {
	return pulumi.ToOutput(e).(DatasetResourceConfigurationComputeTypeOutput)
}

func (e DatasetResourceConfigurationComputeType) ToDatasetResourceConfigurationComputeTypeOutputWithContext(ctx context.Context) DatasetResourceConfigurationComputeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatasetResourceConfigurationComputeTypeOutput)
}

func (e DatasetResourceConfigurationComputeType) ToDatasetResourceConfigurationComputeTypePtrOutput() DatasetResourceConfigurationComputeTypePtrOutput {
	return e.ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(context.Background())
}

func (e DatasetResourceConfigurationComputeType) ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(ctx context.Context) DatasetResourceConfigurationComputeTypePtrOutput {
	return DatasetResourceConfigurationComputeType(e).ToDatasetResourceConfigurationComputeTypeOutputWithContext(ctx).ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(ctx)
}

func (e DatasetResourceConfigurationComputeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetResourceConfigurationComputeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetResourceConfigurationComputeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatasetResourceConfigurationComputeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatasetResourceConfigurationComputeTypeOutput struct{ *pulumi.OutputState }

func (DatasetResourceConfigurationComputeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetResourceConfigurationComputeType)(nil)).Elem()
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToDatasetResourceConfigurationComputeTypeOutput() DatasetResourceConfigurationComputeTypeOutput {
	return o
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToDatasetResourceConfigurationComputeTypeOutputWithContext(ctx context.Context) DatasetResourceConfigurationComputeTypeOutput {
	return o
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToDatasetResourceConfigurationComputeTypePtrOutput() DatasetResourceConfigurationComputeTypePtrOutput {
	return o.ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(context.Background())
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(ctx context.Context) DatasetResourceConfigurationComputeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetResourceConfigurationComputeType) *DatasetResourceConfigurationComputeType {
		return &v
	}).(DatasetResourceConfigurationComputeTypePtrOutput)
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToOutput(ctx context.Context) pulumix.Output[DatasetResourceConfigurationComputeType] {
	return pulumix.Output[DatasetResourceConfigurationComputeType]{
		OutputState: o.OutputState,
	}
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetResourceConfigurationComputeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetResourceConfigurationComputeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetResourceConfigurationComputeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatasetResourceConfigurationComputeTypePtrOutput struct{ *pulumi.OutputState }

func (DatasetResourceConfigurationComputeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetResourceConfigurationComputeType)(nil)).Elem()
}

func (o DatasetResourceConfigurationComputeTypePtrOutput) ToDatasetResourceConfigurationComputeTypePtrOutput() DatasetResourceConfigurationComputeTypePtrOutput {
	return o
}

func (o DatasetResourceConfigurationComputeTypePtrOutput) ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(ctx context.Context) DatasetResourceConfigurationComputeTypePtrOutput {
	return o
}

func (o DatasetResourceConfigurationComputeTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DatasetResourceConfigurationComputeType] {
	return pulumix.Output[*DatasetResourceConfigurationComputeType]{
		OutputState: o.OutputState,
	}
}

func (o DatasetResourceConfigurationComputeTypePtrOutput) Elem() DatasetResourceConfigurationComputeTypeOutput {
	return o.ApplyT(func(v *DatasetResourceConfigurationComputeType) DatasetResourceConfigurationComputeType {
		if v != nil {
			return *v
		}
		var ret DatasetResourceConfigurationComputeType
		return ret
	}).(DatasetResourceConfigurationComputeTypeOutput)
}

func (o DatasetResourceConfigurationComputeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetResourceConfigurationComputeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatasetResourceConfigurationComputeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DatasetResourceConfigurationComputeTypeInput is an input type that accepts DatasetResourceConfigurationComputeTypeArgs and DatasetResourceConfigurationComputeTypeOutput values.
// You can construct a concrete instance of `DatasetResourceConfigurationComputeTypeInput` via:
//
//	DatasetResourceConfigurationComputeTypeArgs{...}
type DatasetResourceConfigurationComputeTypeInput interface {
	pulumi.Input

	ToDatasetResourceConfigurationComputeTypeOutput() DatasetResourceConfigurationComputeTypeOutput
	ToDatasetResourceConfigurationComputeTypeOutputWithContext(context.Context) DatasetResourceConfigurationComputeTypeOutput
}

var datasetResourceConfigurationComputeTypePtrType = reflect.TypeOf((**DatasetResourceConfigurationComputeType)(nil)).Elem()

type DatasetResourceConfigurationComputeTypePtrInput interface {
	pulumi.Input

	ToDatasetResourceConfigurationComputeTypePtrOutput() DatasetResourceConfigurationComputeTypePtrOutput
	ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(context.Context) DatasetResourceConfigurationComputeTypePtrOutput
}

type datasetResourceConfigurationComputeTypePtr string

func DatasetResourceConfigurationComputeTypePtr(v string) DatasetResourceConfigurationComputeTypePtrInput {
	return (*datasetResourceConfigurationComputeTypePtr)(&v)
}

func (*datasetResourceConfigurationComputeTypePtr) ElementType() reflect.Type {
	return datasetResourceConfigurationComputeTypePtrType
}

func (in *datasetResourceConfigurationComputeTypePtr) ToDatasetResourceConfigurationComputeTypePtrOutput() DatasetResourceConfigurationComputeTypePtrOutput {
	return pulumi.ToOutput(in).(DatasetResourceConfigurationComputeTypePtrOutput)
}

func (in *datasetResourceConfigurationComputeTypePtr) ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(ctx context.Context) DatasetResourceConfigurationComputeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatasetResourceConfigurationComputeTypePtrOutput)
}

func (in *datasetResourceConfigurationComputeTypePtr) ToOutput(ctx context.Context) pulumix.Output[*DatasetResourceConfigurationComputeType] {
	return pulumix.Output[*DatasetResourceConfigurationComputeType]{
		OutputState: in.ToDatasetResourceConfigurationComputeTypePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetResourceConfigurationComputeTypeInput)(nil)).Elem(), DatasetResourceConfigurationComputeType("ACU_1"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetResourceConfigurationComputeTypePtrInput)(nil)).Elem(), DatasetResourceConfigurationComputeType("ACU_1"))
	pulumi.RegisterOutputType(DatasetResourceConfigurationComputeTypeOutput{})
	pulumi.RegisterOutputType(DatasetResourceConfigurationComputeTypePtrOutput{})
}
