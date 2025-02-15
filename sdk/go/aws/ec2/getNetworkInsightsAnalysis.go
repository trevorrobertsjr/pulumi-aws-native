// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::EC2::NetworkInsightsAnalysis
func LookupNetworkInsightsAnalysis(ctx *pulumi.Context, args *LookupNetworkInsightsAnalysisArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInsightsAnalysisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkInsightsAnalysisResult
	err := ctx.Invoke("aws-native:ec2:getNetworkInsightsAnalysis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInsightsAnalysisArgs struct {
	NetworkInsightsAnalysisId string `pulumi:"networkInsightsAnalysisId"`
}

type LookupNetworkInsightsAnalysisResult struct {
	AdditionalAccounts         []string                                   `pulumi:"additionalAccounts"`
	AlternatePathHints         []NetworkInsightsAnalysisAlternatePathHint `pulumi:"alternatePathHints"`
	Explanations               []NetworkInsightsAnalysisExplanation       `pulumi:"explanations"`
	ForwardPathComponents      []NetworkInsightsAnalysisPathComponent     `pulumi:"forwardPathComponents"`
	NetworkInsightsAnalysisArn *string                                    `pulumi:"networkInsightsAnalysisArn"`
	NetworkInsightsAnalysisId  *string                                    `pulumi:"networkInsightsAnalysisId"`
	NetworkPathFound           *bool                                      `pulumi:"networkPathFound"`
	ReturnPathComponents       []NetworkInsightsAnalysisPathComponent     `pulumi:"returnPathComponents"`
	StartDate                  *string                                    `pulumi:"startDate"`
	Status                     *NetworkInsightsAnalysisStatus             `pulumi:"status"`
	StatusMessage              *string                                    `pulumi:"statusMessage"`
	SuggestedAccounts          []string                                   `pulumi:"suggestedAccounts"`
	Tags                       []NetworkInsightsAnalysisTag               `pulumi:"tags"`
}

func LookupNetworkInsightsAnalysisOutput(ctx *pulumi.Context, args LookupNetworkInsightsAnalysisOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInsightsAnalysisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkInsightsAnalysisResult, error) {
			args := v.(LookupNetworkInsightsAnalysisArgs)
			r, err := LookupNetworkInsightsAnalysis(ctx, &args, opts...)
			var s LookupNetworkInsightsAnalysisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkInsightsAnalysisResultOutput)
}

type LookupNetworkInsightsAnalysisOutputArgs struct {
	NetworkInsightsAnalysisId pulumi.StringInput `pulumi:"networkInsightsAnalysisId"`
}

func (LookupNetworkInsightsAnalysisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInsightsAnalysisArgs)(nil)).Elem()
}

type LookupNetworkInsightsAnalysisResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInsightsAnalysisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInsightsAnalysisResult)(nil)).Elem()
}

func (o LookupNetworkInsightsAnalysisResultOutput) ToLookupNetworkInsightsAnalysisResultOutput() LookupNetworkInsightsAnalysisResultOutput {
	return o
}

func (o LookupNetworkInsightsAnalysisResultOutput) ToLookupNetworkInsightsAnalysisResultOutputWithContext(ctx context.Context) LookupNetworkInsightsAnalysisResultOutput {
	return o
}

func (o LookupNetworkInsightsAnalysisResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNetworkInsightsAnalysisResult] {
	return pulumix.Output[LookupNetworkInsightsAnalysisResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupNetworkInsightsAnalysisResultOutput) AdditionalAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) []string { return v.AdditionalAccounts }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) AlternatePathHints() NetworkInsightsAnalysisAlternatePathHintArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) []NetworkInsightsAnalysisAlternatePathHint {
		return v.AlternatePathHints
	}).(NetworkInsightsAnalysisAlternatePathHintArrayOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) Explanations() NetworkInsightsAnalysisExplanationArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) []NetworkInsightsAnalysisExplanation {
		return v.Explanations
	}).(NetworkInsightsAnalysisExplanationArrayOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) ForwardPathComponents() NetworkInsightsAnalysisPathComponentArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) []NetworkInsightsAnalysisPathComponent {
		return v.ForwardPathComponents
	}).(NetworkInsightsAnalysisPathComponentArrayOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) NetworkInsightsAnalysisArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) *string { return v.NetworkInsightsAnalysisArn }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) NetworkInsightsAnalysisId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) *string { return v.NetworkInsightsAnalysisId }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) NetworkPathFound() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) *bool { return v.NetworkPathFound }).(pulumi.BoolPtrOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) ReturnPathComponents() NetworkInsightsAnalysisPathComponentArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) []NetworkInsightsAnalysisPathComponent {
		return v.ReturnPathComponents
	}).(NetworkInsightsAnalysisPathComponentArrayOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) Status() NetworkInsightsAnalysisStatusPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) *NetworkInsightsAnalysisStatus { return v.Status }).(NetworkInsightsAnalysisStatusPtrOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) StatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) *string { return v.StatusMessage }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) SuggestedAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) []string { return v.SuggestedAccounts }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkInsightsAnalysisResultOutput) Tags() NetworkInsightsAnalysisTagArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAnalysisResult) []NetworkInsightsAnalysisTag { return v.Tags }).(NetworkInsightsAnalysisTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInsightsAnalysisResultOutput{})
}
