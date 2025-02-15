// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::SSM::PatchBaseline
func LookupPatchBaseline(ctx *pulumi.Context, args *LookupPatchBaselineArgs, opts ...pulumi.InvokeOption) (*LookupPatchBaselineResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPatchBaselineResult
	err := ctx.Invoke("aws-native:ssm:getPatchBaseline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPatchBaselineArgs struct {
	Id string `pulumi:"id"`
}

type LookupPatchBaselineResult struct {
	ApprovalRules                    *PatchBaselineRuleGroup        `pulumi:"approvalRules"`
	ApprovedPatches                  []string                       `pulumi:"approvedPatches"`
	ApprovedPatchesComplianceLevel   *string                        `pulumi:"approvedPatchesComplianceLevel"`
	ApprovedPatchesEnableNonSecurity *bool                          `pulumi:"approvedPatchesEnableNonSecurity"`
	Description                      *string                        `pulumi:"description"`
	GlobalFilters                    *PatchBaselinePatchFilterGroup `pulumi:"globalFilters"`
	Id                               *string                        `pulumi:"id"`
	Name                             *string                        `pulumi:"name"`
	PatchGroups                      []string                       `pulumi:"patchGroups"`
	RejectedPatches                  []string                       `pulumi:"rejectedPatches"`
	RejectedPatchesAction            *string                        `pulumi:"rejectedPatchesAction"`
	Sources                          []PatchBaselinePatchSource     `pulumi:"sources"`
	Tags                             []PatchBaselineTag             `pulumi:"tags"`
}

func LookupPatchBaselineOutput(ctx *pulumi.Context, args LookupPatchBaselineOutputArgs, opts ...pulumi.InvokeOption) LookupPatchBaselineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPatchBaselineResult, error) {
			args := v.(LookupPatchBaselineArgs)
			r, err := LookupPatchBaseline(ctx, &args, opts...)
			var s LookupPatchBaselineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPatchBaselineResultOutput)
}

type LookupPatchBaselineOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPatchBaselineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchBaselineArgs)(nil)).Elem()
}

type LookupPatchBaselineResultOutput struct{ *pulumi.OutputState }

func (LookupPatchBaselineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchBaselineResult)(nil)).Elem()
}

func (o LookupPatchBaselineResultOutput) ToLookupPatchBaselineResultOutput() LookupPatchBaselineResultOutput {
	return o
}

func (o LookupPatchBaselineResultOutput) ToLookupPatchBaselineResultOutputWithContext(ctx context.Context) LookupPatchBaselineResultOutput {
	return o
}

func (o LookupPatchBaselineResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPatchBaselineResult] {
	return pulumix.Output[LookupPatchBaselineResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupPatchBaselineResultOutput) ApprovalRules() PatchBaselineRuleGroupPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *PatchBaselineRuleGroup { return v.ApprovalRules }).(PatchBaselineRuleGroupPtrOutput)
}

func (o LookupPatchBaselineResultOutput) ApprovedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []string { return v.ApprovedPatches }).(pulumi.StringArrayOutput)
}

func (o LookupPatchBaselineResultOutput) ApprovedPatchesComplianceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.ApprovedPatchesComplianceLevel }).(pulumi.StringPtrOutput)
}

func (o LookupPatchBaselineResultOutput) ApprovedPatchesEnableNonSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *bool { return v.ApprovedPatchesEnableNonSecurity }).(pulumi.BoolPtrOutput)
}

func (o LookupPatchBaselineResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPatchBaselineResultOutput) GlobalFilters() PatchBaselinePatchFilterGroupPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *PatchBaselinePatchFilterGroup { return v.GlobalFilters }).(PatchBaselinePatchFilterGroupPtrOutput)
}

func (o LookupPatchBaselineResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPatchBaselineResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupPatchBaselineResultOutput) PatchGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []string { return v.PatchGroups }).(pulumi.StringArrayOutput)
}

func (o LookupPatchBaselineResultOutput) RejectedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []string { return v.RejectedPatches }).(pulumi.StringArrayOutput)
}

func (o LookupPatchBaselineResultOutput) RejectedPatchesAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.RejectedPatchesAction }).(pulumi.StringPtrOutput)
}

func (o LookupPatchBaselineResultOutput) Sources() PatchBaselinePatchSourceArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []PatchBaselinePatchSource { return v.Sources }).(PatchBaselinePatchSourceArrayOutput)
}

func (o LookupPatchBaselineResultOutput) Tags() PatchBaselineTagArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []PatchBaselineTag { return v.Tags }).(PatchBaselineTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPatchBaselineResultOutput{})
}
