// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless
{
    public static class GetPartnerAccount
    {
        /// <summary>
        /// Create and manage partner account
        /// </summary>
        public static Task<GetPartnerAccountResult> InvokeAsync(GetPartnerAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPartnerAccountResult>("aws-native:iotwireless:getPartnerAccount", args ?? new GetPartnerAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Create and manage partner account
        /// </summary>
        public static Output<GetPartnerAccountResult> Invoke(GetPartnerAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPartnerAccountResult>("aws-native:iotwireless:getPartnerAccount", args ?? new GetPartnerAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPartnerAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The partner account ID to disassociate from the AWS account
        /// </summary>
        [Input("partnerAccountId", required: true)]
        public string PartnerAccountId { get; set; } = null!;

        public GetPartnerAccountArgs()
        {
        }
        public static new GetPartnerAccountArgs Empty => new GetPartnerAccountArgs();
    }

    public sealed class GetPartnerAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The partner account ID to disassociate from the AWS account
        /// </summary>
        [Input("partnerAccountId", required: true)]
        public Input<string> PartnerAccountId { get; set; } = null!;

        public GetPartnerAccountInvokeArgs()
        {
        }
        public static new GetPartnerAccountInvokeArgs Empty => new GetPartnerAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetPartnerAccountResult
    {
        /// <summary>
        /// Whether the partner account is linked to the AWS account.
        /// </summary>
        public readonly bool? AccountLinked;
        /// <summary>
        /// PartnerAccount arn. Returned after successful create.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The fingerprint of the Sidewalk application server private key.
        /// </summary>
        public readonly string? Fingerprint;
        /// <summary>
        /// The partner type
        /// </summary>
        public readonly Pulumi.AwsNative.IoTWireless.PartnerAccountPartnerType? PartnerType;
        /// <summary>
        /// The Sidewalk account credentials.
        /// </summary>
        public readonly Outputs.PartnerAccountSidewalkAccountInfoWithFingerprint? SidewalkResponse;
        /// <summary>
        /// A list of key-value pairs that contain metadata for the destination.
        /// </summary>
        public readonly ImmutableArray<Outputs.PartnerAccountTag> Tags;

        [OutputConstructor]
        private GetPartnerAccountResult(
            bool? accountLinked,

            string? arn,

            string? fingerprint,

            Pulumi.AwsNative.IoTWireless.PartnerAccountPartnerType? partnerType,

            Outputs.PartnerAccountSidewalkAccountInfoWithFingerprint? sidewalkResponse,

            ImmutableArray<Outputs.PartnerAccountTag> tags)
        {
            AccountLinked = accountLinked;
            Arn = arn;
            Fingerprint = fingerprint;
            PartnerType = partnerType;
            SidewalkResponse = sidewalkResponse;
            Tags = tags;
        }
    }
}
