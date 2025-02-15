// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kms
{
    public static class GetKey
    {
        /// <summary>
        /// The AWS::KMS::Key resource specifies an AWS KMS key in AWS Key Management Service (AWS KMS). Authorized users can use the AWS KMS key to encrypt and decrypt small amounts of data (up to 4096 bytes), but they are more commonly used to generate data keys. You can also use AWS KMS keys to encrypt data stored in AWS services that are integrated with AWS KMS or within their applications.
        /// </summary>
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("aws-native:kms:getKey", args ?? new GetKeyArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::KMS::Key resource specifies an AWS KMS key in AWS Key Management Service (AWS KMS). Authorized users can use the AWS KMS key to encrypt and decrypt small amounts of data (up to 4096 bytes), but they are more commonly used to generate data keys. You can also use AWS KMS keys to encrypt data stored in AWS services that are integrated with AWS KMS or within their applications.
        /// </summary>
        public static Output<GetKeyResult> Invoke(GetKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKeyResult>("aws-native:kms:getKey", args ?? new GetKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public string KeyId { get; set; } = null!;

        public GetKeyArgs()
        {
        }
        public static new GetKeyArgs Empty => new GetKeyArgs();
    }

    public sealed class GetKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        public GetKeyInvokeArgs()
        {
        }
        public static new GetKeyInvokeArgs Empty => new GetKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        public readonly string? Arn;
        /// <summary>
        /// A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Enables automatic rotation of the key material for the specified AWS KMS key. By default, automation key rotation is not enabled.
        /// </summary>
        public readonly bool? EnableKeyRotation;
        /// <summary>
        /// Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
        /// </summary>
        public readonly bool? Enabled;
        public readonly string? KeyId;
        /// <summary>
        /// The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
        /// </summary>
        public readonly object? KeyPolicy;
        /// <summary>
        /// Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.
        /// </summary>
        public readonly Pulumi.AwsNative.Kms.KeySpec? KeySpec;
        /// <summary>
        /// Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.
        /// </summary>
        public readonly Pulumi.AwsNative.Kms.KeyUsage? KeyUsage;
        /// <summary>
        /// Specifies whether the AWS KMS key should be Multi-Region. You can't change the MultiRegion value after the AWS KMS key is created.
        /// </summary>
        public readonly bool? MultiRegion;
        /// <summary>
        /// The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that AWS KMS creates the key material.
        /// </summary>
        public readonly Pulumi.AwsNative.Kms.KeyOrigin? Origin;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.KeyTag> Tags;

        [OutputConstructor]
        private GetKeyResult(
            string? arn,

            string? description,

            bool? enableKeyRotation,

            bool? enabled,

            string? keyId,

            object? keyPolicy,

            Pulumi.AwsNative.Kms.KeySpec? keySpec,

            Pulumi.AwsNative.Kms.KeyUsage? keyUsage,

            bool? multiRegion,

            Pulumi.AwsNative.Kms.KeyOrigin? origin,

            ImmutableArray<Outputs.KeyTag> tags)
        {
            Arn = arn;
            Description = description;
            EnableKeyRotation = enableKeyRotation;
            Enabled = enabled;
            KeyId = keyId;
            KeyPolicy = keyPolicy;
            KeySpec = keySpec;
            KeyUsage = keyUsage;
            MultiRegion = multiRegion;
            Origin = origin;
            Tags = tags;
        }
    }
}
