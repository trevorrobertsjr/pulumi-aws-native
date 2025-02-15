// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight
{
    public static class GetTemplate
    {
        /// <summary>
        /// Definition of the AWS::QuickSight::Template Resource Type.
        /// </summary>
        public static Task<GetTemplateResult> InvokeAsync(GetTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTemplateResult>("aws-native:quicksight:getTemplate", args ?? new GetTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of the AWS::QuickSight::Template Resource Type.
        /// </summary>
        public static Output<GetTemplateResult> Invoke(GetTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateResult>("aws-native:quicksight:getTemplate", args ?? new GetTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemplateArgs : global::Pulumi.InvokeArgs
    {
        [Input("awsAccountId", required: true)]
        public string AwsAccountId { get; set; } = null!;

        [Input("templateId", required: true)]
        public string TemplateId { get; set; } = null!;

        public GetTemplateArgs()
        {
        }
        public static new GetTemplateArgs Empty => new GetTemplateArgs();
    }

    public sealed class GetTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        public GetTemplateInvokeArgs()
        {
        }
        public static new GetTemplateInvokeArgs Empty => new GetTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetTemplateResult
    {
        public readonly string? Arn;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.TemplateResourcePermission> Permissions;
        public readonly ImmutableArray<Outputs.TemplateTag> Tags;

        [OutputConstructor]
        private GetTemplateResult(
            string? arn,

            string? name,

            ImmutableArray<Outputs.TemplateResourcePermission> permissions,

            ImmutableArray<Outputs.TemplateTag> tags)
        {
            Arn = arn;
            Name = name;
            Permissions = permissions;
            Tags = tags;
        }
    }
}
