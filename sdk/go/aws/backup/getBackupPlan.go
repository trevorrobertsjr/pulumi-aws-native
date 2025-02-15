// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Backup::BackupPlan
func LookupBackupPlan(ctx *pulumi.Context, args *LookupBackupPlanArgs, opts ...pulumi.InvokeOption) (*LookupBackupPlanResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupPlanResult
	err := ctx.Invoke("aws-native:backup:getBackupPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPlanArgs struct {
	BackupPlanId string `pulumi:"backupPlanId"`
}

type LookupBackupPlanResult struct {
	BackupPlan     *BackupPlanResourceType `pulumi:"backupPlan"`
	BackupPlanArn  *string                 `pulumi:"backupPlanArn"`
	BackupPlanId   *string                 `pulumi:"backupPlanId"`
	BackupPlanTags interface{}             `pulumi:"backupPlanTags"`
	VersionId      *string                 `pulumi:"versionId"`
}

func LookupBackupPlanOutput(ctx *pulumi.Context, args LookupBackupPlanOutputArgs, opts ...pulumi.InvokeOption) LookupBackupPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupPlanResult, error) {
			args := v.(LookupBackupPlanArgs)
			r, err := LookupBackupPlan(ctx, &args, opts...)
			var s LookupBackupPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupPlanResultOutput)
}

type LookupBackupPlanOutputArgs struct {
	BackupPlanId pulumi.StringInput `pulumi:"backupPlanId"`
}

func (LookupBackupPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPlanArgs)(nil)).Elem()
}

type LookupBackupPlanResultOutput struct{ *pulumi.OutputState }

func (LookupBackupPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPlanResult)(nil)).Elem()
}

func (o LookupBackupPlanResultOutput) ToLookupBackupPlanResultOutput() LookupBackupPlanResultOutput {
	return o
}

func (o LookupBackupPlanResultOutput) ToLookupBackupPlanResultOutputWithContext(ctx context.Context) LookupBackupPlanResultOutput {
	return o
}

func (o LookupBackupPlanResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBackupPlanResult] {
	return pulumix.Output[LookupBackupPlanResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBackupPlanResultOutput) BackupPlan() BackupPlanResourceTypePtrOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) *BackupPlanResourceType { return v.BackupPlan }).(BackupPlanResourceTypePtrOutput)
}

func (o LookupBackupPlanResultOutput) BackupPlanArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) *string { return v.BackupPlanArn }).(pulumi.StringPtrOutput)
}

func (o LookupBackupPlanResultOutput) BackupPlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) *string { return v.BackupPlanId }).(pulumi.StringPtrOutput)
}

func (o LookupBackupPlanResultOutput) BackupPlanTags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) interface{} { return v.BackupPlanTags }).(pulumi.AnyOutput)
}

func (o LookupBackupPlanResultOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) *string { return v.VersionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupPlanResultOutput{})
}
