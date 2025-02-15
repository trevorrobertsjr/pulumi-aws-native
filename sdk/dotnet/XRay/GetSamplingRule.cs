// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.XRay
{
    public static class GetSamplingRule
    {
        /// <summary>
        /// This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.
        /// </summary>
        public static Task<GetSamplingRuleResult> InvokeAsync(GetSamplingRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSamplingRuleResult>("aws-native:xray:getSamplingRule", args ?? new GetSamplingRuleArgs(), options.WithDefaults());

        /// <summary>
        /// This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.
        /// </summary>
        public static Output<GetSamplingRuleResult> Invoke(GetSamplingRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSamplingRuleResult>("aws-native:xray:getSamplingRule", args ?? new GetSamplingRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSamplingRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("ruleArn", required: true)]
        public string RuleArn { get; set; } = null!;

        public GetSamplingRuleArgs()
        {
        }
        public static new GetSamplingRuleArgs Empty => new GetSamplingRuleArgs();
    }

    public sealed class GetSamplingRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ruleArn", required: true)]
        public Input<string> RuleArn { get; set; } = null!;

        public GetSamplingRuleInvokeArgs()
        {
        }
        public static new GetSamplingRuleInvokeArgs Empty => new GetSamplingRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetSamplingRuleResult
    {
        public readonly string? RuleArn;
        public readonly string? RuleName;
        public readonly Outputs.SamplingRule? SamplingRuleValue;
        public readonly Outputs.SamplingRuleRecord? SamplingRuleRecord;
        public readonly Outputs.SamplingRuleUpdate? SamplingRuleUpdate;
        public readonly ImmutableArray<Outputs.SamplingRuleTag> Tags;

        [OutputConstructor]
        private GetSamplingRuleResult(
            string? ruleArn,

            string? ruleName,

            Outputs.SamplingRule? samplingRule,

            Outputs.SamplingRuleRecord? samplingRuleRecord,

            Outputs.SamplingRuleUpdate? samplingRuleUpdate,

            ImmutableArray<Outputs.SamplingRuleTag> tags)
        {
            RuleArn = ruleArn;
            RuleName = ruleName;
            SamplingRuleValue = samplingRule;
            SamplingRuleRecord = samplingRuleRecord;
            SamplingRuleUpdate = samplingRuleUpdate;
            Tags = tags;
        }
    }
}
