// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    public static class GetOriginAccessControl
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFront::OriginAccessControl
        /// </summary>
        public static Task<GetOriginAccessControlResult> InvokeAsync(GetOriginAccessControlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOriginAccessControlResult>("aws-native:cloudfront:getOriginAccessControl", args ?? new GetOriginAccessControlArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFront::OriginAccessControl
        /// </summary>
        public static Output<GetOriginAccessControlResult> Invoke(GetOriginAccessControlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOriginAccessControlResult>("aws-native:cloudfront:getOriginAccessControl", args ?? new GetOriginAccessControlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOriginAccessControlArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetOriginAccessControlArgs()
        {
        }
        public static new GetOriginAccessControlArgs Empty => new GetOriginAccessControlArgs();
    }

    public sealed class GetOriginAccessControlInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetOriginAccessControlInvokeArgs()
        {
        }
        public static new GetOriginAccessControlInvokeArgs Empty => new GetOriginAccessControlInvokeArgs();
    }


    [OutputType]
    public sealed class GetOriginAccessControlResult
    {
        public readonly string? Id;
        public readonly Outputs.OriginAccessControlConfig? OriginAccessControlConfig;

        [OutputConstructor]
        private GetOriginAccessControlResult(
            string? id,

            Outputs.OriginAccessControlConfig? originAccessControlConfig)
        {
            Id = id;
            OriginAccessControlConfig = originAccessControlConfig;
        }
    }
}
