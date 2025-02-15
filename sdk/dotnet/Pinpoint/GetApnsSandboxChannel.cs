// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetApnsSandboxChannel
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::APNSSandboxChannel
        /// </summary>
        public static Task<GetApnsSandboxChannelResult> InvokeAsync(GetApnsSandboxChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApnsSandboxChannelResult>("aws-native:pinpoint:getApnsSandboxChannel", args ?? new GetApnsSandboxChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::APNSSandboxChannel
        /// </summary>
        public static Output<GetApnsSandboxChannelResult> Invoke(GetApnsSandboxChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApnsSandboxChannelResult>("aws-native:pinpoint:getApnsSandboxChannel", args ?? new GetApnsSandboxChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApnsSandboxChannelArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetApnsSandboxChannelArgs()
        {
        }
        public static new GetApnsSandboxChannelArgs Empty => new GetApnsSandboxChannelArgs();
    }

    public sealed class GetApnsSandboxChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetApnsSandboxChannelInvokeArgs()
        {
        }
        public static new GetApnsSandboxChannelInvokeArgs Empty => new GetApnsSandboxChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetApnsSandboxChannelResult
    {
        public readonly string? BundleId;
        public readonly string? Certificate;
        public readonly string? DefaultAuthenticationMethod;
        public readonly bool? Enabled;
        public readonly string? Id;
        public readonly string? PrivateKey;
        public readonly string? TeamId;
        public readonly string? TokenKey;
        public readonly string? TokenKeyId;

        [OutputConstructor]
        private GetApnsSandboxChannelResult(
            string? bundleId,

            string? certificate,

            string? defaultAuthenticationMethod,

            bool? enabled,

            string? id,

            string? privateKey,

            string? teamId,

            string? tokenKey,

            string? tokenKeyId)
        {
            BundleId = bundleId;
            Certificate = certificate;
            DefaultAuthenticationMethod = defaultAuthenticationMethod;
            Enabled = enabled;
            Id = id;
            PrivateKey = privateKey;
            TeamId = teamId;
            TokenKey = tokenKey;
            TokenKeyId = tokenKeyId;
        }
    }
}
