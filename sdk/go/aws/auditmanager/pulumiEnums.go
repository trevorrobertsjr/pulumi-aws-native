// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package auditmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The status of the delegation.
type AssessmentDelegationStatus string

const (
	AssessmentDelegationStatusInProgress  = AssessmentDelegationStatus("IN_PROGRESS")
	AssessmentDelegationStatusUnderReview = AssessmentDelegationStatus("UNDER_REVIEW")
	AssessmentDelegationStatusComplete    = AssessmentDelegationStatus("COMPLETE")
)

func (AssessmentDelegationStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentDelegationStatus)(nil)).Elem()
}

func (e AssessmentDelegationStatus) ToAssessmentDelegationStatusOutput() AssessmentDelegationStatusOutput {
	return pulumi.ToOutput(e).(AssessmentDelegationStatusOutput)
}

func (e AssessmentDelegationStatus) ToAssessmentDelegationStatusOutputWithContext(ctx context.Context) AssessmentDelegationStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssessmentDelegationStatusOutput)
}

func (e AssessmentDelegationStatus) ToAssessmentDelegationStatusPtrOutput() AssessmentDelegationStatusPtrOutput {
	return e.ToAssessmentDelegationStatusPtrOutputWithContext(context.Background())
}

func (e AssessmentDelegationStatus) ToAssessmentDelegationStatusPtrOutputWithContext(ctx context.Context) AssessmentDelegationStatusPtrOutput {
	return AssessmentDelegationStatus(e).ToAssessmentDelegationStatusOutputWithContext(ctx).ToAssessmentDelegationStatusPtrOutputWithContext(ctx)
}

func (e AssessmentDelegationStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentDelegationStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentDelegationStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentDelegationStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssessmentDelegationStatusOutput struct{ *pulumi.OutputState }

func (AssessmentDelegationStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentDelegationStatus)(nil)).Elem()
}

func (o AssessmentDelegationStatusOutput) ToAssessmentDelegationStatusOutput() AssessmentDelegationStatusOutput {
	return o
}

func (o AssessmentDelegationStatusOutput) ToAssessmentDelegationStatusOutputWithContext(ctx context.Context) AssessmentDelegationStatusOutput {
	return o
}

func (o AssessmentDelegationStatusOutput) ToAssessmentDelegationStatusPtrOutput() AssessmentDelegationStatusPtrOutput {
	return o.ToAssessmentDelegationStatusPtrOutputWithContext(context.Background())
}

func (o AssessmentDelegationStatusOutput) ToAssessmentDelegationStatusPtrOutputWithContext(ctx context.Context) AssessmentDelegationStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentDelegationStatus) *AssessmentDelegationStatus {
		return &v
	}).(AssessmentDelegationStatusPtrOutput)
}

func (o AssessmentDelegationStatusOutput) ToOutput(ctx context.Context) pulumix.Output[AssessmentDelegationStatus] {
	return pulumix.Output[AssessmentDelegationStatus]{
		OutputState: o.OutputState,
	}
}

func (o AssessmentDelegationStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssessmentDelegationStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentDelegationStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssessmentDelegationStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentDelegationStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentDelegationStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssessmentDelegationStatusPtrOutput struct{ *pulumi.OutputState }

func (AssessmentDelegationStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentDelegationStatus)(nil)).Elem()
}

func (o AssessmentDelegationStatusPtrOutput) ToAssessmentDelegationStatusPtrOutput() AssessmentDelegationStatusPtrOutput {
	return o
}

func (o AssessmentDelegationStatusPtrOutput) ToAssessmentDelegationStatusPtrOutputWithContext(ctx context.Context) AssessmentDelegationStatusPtrOutput {
	return o
}

func (o AssessmentDelegationStatusPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AssessmentDelegationStatus] {
	return pulumix.Output[*AssessmentDelegationStatus]{
		OutputState: o.OutputState,
	}
}

func (o AssessmentDelegationStatusPtrOutput) Elem() AssessmentDelegationStatusOutput {
	return o.ApplyT(func(v *AssessmentDelegationStatus) AssessmentDelegationStatus {
		if v != nil {
			return *v
		}
		var ret AssessmentDelegationStatus
		return ret
	}).(AssessmentDelegationStatusOutput)
}

func (o AssessmentDelegationStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentDelegationStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssessmentDelegationStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssessmentDelegationStatusInput is an input type that accepts AssessmentDelegationStatusArgs and AssessmentDelegationStatusOutput values.
// You can construct a concrete instance of `AssessmentDelegationStatusInput` via:
//
//	AssessmentDelegationStatusArgs{...}
type AssessmentDelegationStatusInput interface {
	pulumi.Input

	ToAssessmentDelegationStatusOutput() AssessmentDelegationStatusOutput
	ToAssessmentDelegationStatusOutputWithContext(context.Context) AssessmentDelegationStatusOutput
}

var assessmentDelegationStatusPtrType = reflect.TypeOf((**AssessmentDelegationStatus)(nil)).Elem()

type AssessmentDelegationStatusPtrInput interface {
	pulumi.Input

	ToAssessmentDelegationStatusPtrOutput() AssessmentDelegationStatusPtrOutput
	ToAssessmentDelegationStatusPtrOutputWithContext(context.Context) AssessmentDelegationStatusPtrOutput
}

type assessmentDelegationStatusPtr string

func AssessmentDelegationStatusPtr(v string) AssessmentDelegationStatusPtrInput {
	return (*assessmentDelegationStatusPtr)(&v)
}

func (*assessmentDelegationStatusPtr) ElementType() reflect.Type {
	return assessmentDelegationStatusPtrType
}

func (in *assessmentDelegationStatusPtr) ToAssessmentDelegationStatusPtrOutput() AssessmentDelegationStatusPtrOutput {
	return pulumi.ToOutput(in).(AssessmentDelegationStatusPtrOutput)
}

func (in *assessmentDelegationStatusPtr) ToAssessmentDelegationStatusPtrOutputWithContext(ctx context.Context) AssessmentDelegationStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssessmentDelegationStatusPtrOutput)
}

func (in *assessmentDelegationStatusPtr) ToOutput(ctx context.Context) pulumix.Output[*AssessmentDelegationStatus] {
	return pulumix.Output[*AssessmentDelegationStatus]{
		OutputState: in.ToAssessmentDelegationStatusPtrOutputWithContext(ctx).OutputState,
	}
}

// The destination type, such as Amazon S3.
type AssessmentReportDestinationType string

const (
	AssessmentReportDestinationTypeS3 = AssessmentReportDestinationType("S3")
)

func (AssessmentReportDestinationType) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentReportDestinationType)(nil)).Elem()
}

func (e AssessmentReportDestinationType) ToAssessmentReportDestinationTypeOutput() AssessmentReportDestinationTypeOutput {
	return pulumi.ToOutput(e).(AssessmentReportDestinationTypeOutput)
}

func (e AssessmentReportDestinationType) ToAssessmentReportDestinationTypeOutputWithContext(ctx context.Context) AssessmentReportDestinationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssessmentReportDestinationTypeOutput)
}

func (e AssessmentReportDestinationType) ToAssessmentReportDestinationTypePtrOutput() AssessmentReportDestinationTypePtrOutput {
	return e.ToAssessmentReportDestinationTypePtrOutputWithContext(context.Background())
}

func (e AssessmentReportDestinationType) ToAssessmentReportDestinationTypePtrOutputWithContext(ctx context.Context) AssessmentReportDestinationTypePtrOutput {
	return AssessmentReportDestinationType(e).ToAssessmentReportDestinationTypeOutputWithContext(ctx).ToAssessmentReportDestinationTypePtrOutputWithContext(ctx)
}

func (e AssessmentReportDestinationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentReportDestinationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentReportDestinationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentReportDestinationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssessmentReportDestinationTypeOutput struct{ *pulumi.OutputState }

func (AssessmentReportDestinationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentReportDestinationType)(nil)).Elem()
}

func (o AssessmentReportDestinationTypeOutput) ToAssessmentReportDestinationTypeOutput() AssessmentReportDestinationTypeOutput {
	return o
}

func (o AssessmentReportDestinationTypeOutput) ToAssessmentReportDestinationTypeOutputWithContext(ctx context.Context) AssessmentReportDestinationTypeOutput {
	return o
}

func (o AssessmentReportDestinationTypeOutput) ToAssessmentReportDestinationTypePtrOutput() AssessmentReportDestinationTypePtrOutput {
	return o.ToAssessmentReportDestinationTypePtrOutputWithContext(context.Background())
}

func (o AssessmentReportDestinationTypeOutput) ToAssessmentReportDestinationTypePtrOutputWithContext(ctx context.Context) AssessmentReportDestinationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentReportDestinationType) *AssessmentReportDestinationType {
		return &v
	}).(AssessmentReportDestinationTypePtrOutput)
}

func (o AssessmentReportDestinationTypeOutput) ToOutput(ctx context.Context) pulumix.Output[AssessmentReportDestinationType] {
	return pulumix.Output[AssessmentReportDestinationType]{
		OutputState: o.OutputState,
	}
}

func (o AssessmentReportDestinationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssessmentReportDestinationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentReportDestinationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssessmentReportDestinationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentReportDestinationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentReportDestinationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssessmentReportDestinationTypePtrOutput struct{ *pulumi.OutputState }

func (AssessmentReportDestinationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentReportDestinationType)(nil)).Elem()
}

func (o AssessmentReportDestinationTypePtrOutput) ToAssessmentReportDestinationTypePtrOutput() AssessmentReportDestinationTypePtrOutput {
	return o
}

func (o AssessmentReportDestinationTypePtrOutput) ToAssessmentReportDestinationTypePtrOutputWithContext(ctx context.Context) AssessmentReportDestinationTypePtrOutput {
	return o
}

func (o AssessmentReportDestinationTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AssessmentReportDestinationType] {
	return pulumix.Output[*AssessmentReportDestinationType]{
		OutputState: o.OutputState,
	}
}

func (o AssessmentReportDestinationTypePtrOutput) Elem() AssessmentReportDestinationTypeOutput {
	return o.ApplyT(func(v *AssessmentReportDestinationType) AssessmentReportDestinationType {
		if v != nil {
			return *v
		}
		var ret AssessmentReportDestinationType
		return ret
	}).(AssessmentReportDestinationTypeOutput)
}

func (o AssessmentReportDestinationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentReportDestinationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssessmentReportDestinationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssessmentReportDestinationTypeInput is an input type that accepts AssessmentReportDestinationTypeArgs and AssessmentReportDestinationTypeOutput values.
// You can construct a concrete instance of `AssessmentReportDestinationTypeInput` via:
//
//	AssessmentReportDestinationTypeArgs{...}
type AssessmentReportDestinationTypeInput interface {
	pulumi.Input

	ToAssessmentReportDestinationTypeOutput() AssessmentReportDestinationTypeOutput
	ToAssessmentReportDestinationTypeOutputWithContext(context.Context) AssessmentReportDestinationTypeOutput
}

var assessmentReportDestinationTypePtrType = reflect.TypeOf((**AssessmentReportDestinationType)(nil)).Elem()

type AssessmentReportDestinationTypePtrInput interface {
	pulumi.Input

	ToAssessmentReportDestinationTypePtrOutput() AssessmentReportDestinationTypePtrOutput
	ToAssessmentReportDestinationTypePtrOutputWithContext(context.Context) AssessmentReportDestinationTypePtrOutput
}

type assessmentReportDestinationTypePtr string

func AssessmentReportDestinationTypePtr(v string) AssessmentReportDestinationTypePtrInput {
	return (*assessmentReportDestinationTypePtr)(&v)
}

func (*assessmentReportDestinationTypePtr) ElementType() reflect.Type {
	return assessmentReportDestinationTypePtrType
}

func (in *assessmentReportDestinationTypePtr) ToAssessmentReportDestinationTypePtrOutput() AssessmentReportDestinationTypePtrOutput {
	return pulumi.ToOutput(in).(AssessmentReportDestinationTypePtrOutput)
}

func (in *assessmentReportDestinationTypePtr) ToAssessmentReportDestinationTypePtrOutputWithContext(ctx context.Context) AssessmentReportDestinationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssessmentReportDestinationTypePtrOutput)
}

func (in *assessmentReportDestinationTypePtr) ToOutput(ctx context.Context) pulumix.Output[*AssessmentReportDestinationType] {
	return pulumix.Output[*AssessmentReportDestinationType]{
		OutputState: in.ToAssessmentReportDestinationTypePtrOutputWithContext(ctx).OutputState,
	}
}

// The IAM role type.
type AssessmentRoleType string

const (
	AssessmentRoleTypeProcessOwner  = AssessmentRoleType("PROCESS_OWNER")
	AssessmentRoleTypeResourceOwner = AssessmentRoleType("RESOURCE_OWNER")
)

func (AssessmentRoleType) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentRoleType)(nil)).Elem()
}

func (e AssessmentRoleType) ToAssessmentRoleTypeOutput() AssessmentRoleTypeOutput {
	return pulumi.ToOutput(e).(AssessmentRoleTypeOutput)
}

func (e AssessmentRoleType) ToAssessmentRoleTypeOutputWithContext(ctx context.Context) AssessmentRoleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssessmentRoleTypeOutput)
}

func (e AssessmentRoleType) ToAssessmentRoleTypePtrOutput() AssessmentRoleTypePtrOutput {
	return e.ToAssessmentRoleTypePtrOutputWithContext(context.Background())
}

func (e AssessmentRoleType) ToAssessmentRoleTypePtrOutputWithContext(ctx context.Context) AssessmentRoleTypePtrOutput {
	return AssessmentRoleType(e).ToAssessmentRoleTypeOutputWithContext(ctx).ToAssessmentRoleTypePtrOutputWithContext(ctx)
}

func (e AssessmentRoleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentRoleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentRoleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentRoleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssessmentRoleTypeOutput struct{ *pulumi.OutputState }

func (AssessmentRoleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentRoleType)(nil)).Elem()
}

func (o AssessmentRoleTypeOutput) ToAssessmentRoleTypeOutput() AssessmentRoleTypeOutput {
	return o
}

func (o AssessmentRoleTypeOutput) ToAssessmentRoleTypeOutputWithContext(ctx context.Context) AssessmentRoleTypeOutput {
	return o
}

func (o AssessmentRoleTypeOutput) ToAssessmentRoleTypePtrOutput() AssessmentRoleTypePtrOutput {
	return o.ToAssessmentRoleTypePtrOutputWithContext(context.Background())
}

func (o AssessmentRoleTypeOutput) ToAssessmentRoleTypePtrOutputWithContext(ctx context.Context) AssessmentRoleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentRoleType) *AssessmentRoleType {
		return &v
	}).(AssessmentRoleTypePtrOutput)
}

func (o AssessmentRoleTypeOutput) ToOutput(ctx context.Context) pulumix.Output[AssessmentRoleType] {
	return pulumix.Output[AssessmentRoleType]{
		OutputState: o.OutputState,
	}
}

func (o AssessmentRoleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssessmentRoleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentRoleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssessmentRoleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentRoleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentRoleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssessmentRoleTypePtrOutput struct{ *pulumi.OutputState }

func (AssessmentRoleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentRoleType)(nil)).Elem()
}

func (o AssessmentRoleTypePtrOutput) ToAssessmentRoleTypePtrOutput() AssessmentRoleTypePtrOutput {
	return o
}

func (o AssessmentRoleTypePtrOutput) ToAssessmentRoleTypePtrOutputWithContext(ctx context.Context) AssessmentRoleTypePtrOutput {
	return o
}

func (o AssessmentRoleTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AssessmentRoleType] {
	return pulumix.Output[*AssessmentRoleType]{
		OutputState: o.OutputState,
	}
}

func (o AssessmentRoleTypePtrOutput) Elem() AssessmentRoleTypeOutput {
	return o.ApplyT(func(v *AssessmentRoleType) AssessmentRoleType {
		if v != nil {
			return *v
		}
		var ret AssessmentRoleType
		return ret
	}).(AssessmentRoleTypeOutput)
}

func (o AssessmentRoleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentRoleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssessmentRoleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssessmentRoleTypeInput is an input type that accepts AssessmentRoleTypeArgs and AssessmentRoleTypeOutput values.
// You can construct a concrete instance of `AssessmentRoleTypeInput` via:
//
//	AssessmentRoleTypeArgs{...}
type AssessmentRoleTypeInput interface {
	pulumi.Input

	ToAssessmentRoleTypeOutput() AssessmentRoleTypeOutput
	ToAssessmentRoleTypeOutputWithContext(context.Context) AssessmentRoleTypeOutput
}

var assessmentRoleTypePtrType = reflect.TypeOf((**AssessmentRoleType)(nil)).Elem()

type AssessmentRoleTypePtrInput interface {
	pulumi.Input

	ToAssessmentRoleTypePtrOutput() AssessmentRoleTypePtrOutput
	ToAssessmentRoleTypePtrOutputWithContext(context.Context) AssessmentRoleTypePtrOutput
}

type assessmentRoleTypePtr string

func AssessmentRoleTypePtr(v string) AssessmentRoleTypePtrInput {
	return (*assessmentRoleTypePtr)(&v)
}

func (*assessmentRoleTypePtr) ElementType() reflect.Type {
	return assessmentRoleTypePtrType
}

func (in *assessmentRoleTypePtr) ToAssessmentRoleTypePtrOutput() AssessmentRoleTypePtrOutput {
	return pulumi.ToOutput(in).(AssessmentRoleTypePtrOutput)
}

func (in *assessmentRoleTypePtr) ToAssessmentRoleTypePtrOutputWithContext(ctx context.Context) AssessmentRoleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssessmentRoleTypePtrOutput)
}

func (in *assessmentRoleTypePtr) ToOutput(ctx context.Context) pulumix.Output[*AssessmentRoleType] {
	return pulumix.Output[*AssessmentRoleType]{
		OutputState: in.ToAssessmentRoleTypePtrOutputWithContext(ctx).OutputState,
	}
}

// The status of the specified assessment.
type AssessmentStatus string

const (
	AssessmentStatusActive   = AssessmentStatus("ACTIVE")
	AssessmentStatusInactive = AssessmentStatus("INACTIVE")
)

func (AssessmentStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatus)(nil)).Elem()
}

func (e AssessmentStatus) ToAssessmentStatusOutput() AssessmentStatusOutput {
	return pulumi.ToOutput(e).(AssessmentStatusOutput)
}

func (e AssessmentStatus) ToAssessmentStatusOutputWithContext(ctx context.Context) AssessmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssessmentStatusOutput)
}

func (e AssessmentStatus) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return e.ToAssessmentStatusPtrOutputWithContext(context.Background())
}

func (e AssessmentStatus) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return AssessmentStatus(e).ToAssessmentStatusOutputWithContext(ctx).ToAssessmentStatusPtrOutputWithContext(ctx)
}

func (e AssessmentStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssessmentStatusOutput struct{ *pulumi.OutputState }

func (AssessmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatus)(nil)).Elem()
}

func (o AssessmentStatusOutput) ToAssessmentStatusOutput() AssessmentStatusOutput {
	return o
}

func (o AssessmentStatusOutput) ToAssessmentStatusOutputWithContext(ctx context.Context) AssessmentStatusOutput {
	return o
}

func (o AssessmentStatusOutput) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return o.ToAssessmentStatusPtrOutputWithContext(context.Background())
}

func (o AssessmentStatusOutput) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentStatus) *AssessmentStatus {
		return &v
	}).(AssessmentStatusPtrOutput)
}

func (o AssessmentStatusOutput) ToOutput(ctx context.Context) pulumix.Output[AssessmentStatus] {
	return pulumix.Output[AssessmentStatus]{
		OutputState: o.OutputState,
	}
}

func (o AssessmentStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssessmentStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssessmentStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssessmentStatusPtrOutput struct{ *pulumi.OutputState }

func (AssessmentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatus)(nil)).Elem()
}

func (o AssessmentStatusPtrOutput) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return o
}

func (o AssessmentStatusPtrOutput) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return o
}

func (o AssessmentStatusPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AssessmentStatus] {
	return pulumix.Output[*AssessmentStatus]{
		OutputState: o.OutputState,
	}
}

func (o AssessmentStatusPtrOutput) Elem() AssessmentStatusOutput {
	return o.ApplyT(func(v *AssessmentStatus) AssessmentStatus {
		if v != nil {
			return *v
		}
		var ret AssessmentStatus
		return ret
	}).(AssessmentStatusOutput)
}

func (o AssessmentStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssessmentStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssessmentStatusInput is an input type that accepts AssessmentStatusArgs and AssessmentStatusOutput values.
// You can construct a concrete instance of `AssessmentStatusInput` via:
//
//	AssessmentStatusArgs{...}
type AssessmentStatusInput interface {
	pulumi.Input

	ToAssessmentStatusOutput() AssessmentStatusOutput
	ToAssessmentStatusOutputWithContext(context.Context) AssessmentStatusOutput
}

var assessmentStatusPtrType = reflect.TypeOf((**AssessmentStatus)(nil)).Elem()

type AssessmentStatusPtrInput interface {
	pulumi.Input

	ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput
	ToAssessmentStatusPtrOutputWithContext(context.Context) AssessmentStatusPtrOutput
}

type assessmentStatusPtr string

func AssessmentStatusPtr(v string) AssessmentStatusPtrInput {
	return (*assessmentStatusPtr)(&v)
}

func (*assessmentStatusPtr) ElementType() reflect.Type {
	return assessmentStatusPtrType
}

func (in *assessmentStatusPtr) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return pulumi.ToOutput(in).(AssessmentStatusPtrOutput)
}

func (in *assessmentStatusPtr) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssessmentStatusPtrOutput)
}

func (in *assessmentStatusPtr) ToOutput(ctx context.Context) pulumix.Output[*AssessmentStatus] {
	return pulumix.Output[*AssessmentStatus]{
		OutputState: in.ToAssessmentStatusPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentDelegationStatusInput)(nil)).Elem(), AssessmentDelegationStatus("IN_PROGRESS"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentDelegationStatusPtrInput)(nil)).Elem(), AssessmentDelegationStatus("IN_PROGRESS"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentReportDestinationTypeInput)(nil)).Elem(), AssessmentReportDestinationType("S3"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentReportDestinationTypePtrInput)(nil)).Elem(), AssessmentReportDestinationType("S3"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentRoleTypeInput)(nil)).Elem(), AssessmentRoleType("PROCESS_OWNER"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentRoleTypePtrInput)(nil)).Elem(), AssessmentRoleType("PROCESS_OWNER"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentStatusInput)(nil)).Elem(), AssessmentStatus("ACTIVE"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentStatusPtrInput)(nil)).Elem(), AssessmentStatus("ACTIVE"))
	pulumi.RegisterOutputType(AssessmentDelegationStatusOutput{})
	pulumi.RegisterOutputType(AssessmentDelegationStatusPtrOutput{})
	pulumi.RegisterOutputType(AssessmentReportDestinationTypeOutput{})
	pulumi.RegisterOutputType(AssessmentReportDestinationTypePtrOutput{})
	pulumi.RegisterOutputType(AssessmentRoleTypeOutput{})
	pulumi.RegisterOutputType(AssessmentRoleTypePtrOutput{})
	pulumi.RegisterOutputType(AssessmentStatusOutput{})
	pulumi.RegisterOutputType(AssessmentStatusPtrOutput{})
}
