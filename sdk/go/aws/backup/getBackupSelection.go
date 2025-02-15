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

// Resource Type definition for AWS::Backup::BackupSelection
func LookupBackupSelection(ctx *pulumi.Context, args *LookupBackupSelectionArgs, opts ...pulumi.InvokeOption) (*LookupBackupSelectionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupSelectionResult
	err := ctx.Invoke("aws-native:backup:getBackupSelection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupSelectionArgs struct {
	Id string `pulumi:"id"`
}

type LookupBackupSelectionResult struct {
	Id          *string `pulumi:"id"`
	SelectionId *string `pulumi:"selectionId"`
}

func LookupBackupSelectionOutput(ctx *pulumi.Context, args LookupBackupSelectionOutputArgs, opts ...pulumi.InvokeOption) LookupBackupSelectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupSelectionResult, error) {
			args := v.(LookupBackupSelectionArgs)
			r, err := LookupBackupSelection(ctx, &args, opts...)
			var s LookupBackupSelectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupSelectionResultOutput)
}

type LookupBackupSelectionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupBackupSelectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupSelectionArgs)(nil)).Elem()
}

type LookupBackupSelectionResultOutput struct{ *pulumi.OutputState }

func (LookupBackupSelectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupSelectionResult)(nil)).Elem()
}

func (o LookupBackupSelectionResultOutput) ToLookupBackupSelectionResultOutput() LookupBackupSelectionResultOutput {
	return o
}

func (o LookupBackupSelectionResultOutput) ToLookupBackupSelectionResultOutputWithContext(ctx context.Context) LookupBackupSelectionResultOutput {
	return o
}

func (o LookupBackupSelectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBackupSelectionResult] {
	return pulumix.Output[LookupBackupSelectionResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBackupSelectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupSelectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupBackupSelectionResultOutput) SelectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupSelectionResult) *string { return v.SelectionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupSelectionResultOutput{})
}
