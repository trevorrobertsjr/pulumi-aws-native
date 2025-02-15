// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class TrustStoreRevocationRevocationContent
    {
        public readonly string? RevocationType;
        public readonly string? S3Bucket;
        public readonly string? S3Key;
        public readonly string? S3ObjectVersion;

        [OutputConstructor]
        private TrustStoreRevocationRevocationContent(
            string? revocationType,

            string? s3Bucket,

            string? s3Key,

            string? s3ObjectVersion)
        {
            RevocationType = revocationType;
            S3Bucket = s3Bucket;
            S3Key = s3Key;
            S3ObjectVersion = s3ObjectVersion;
        }
    }
}
