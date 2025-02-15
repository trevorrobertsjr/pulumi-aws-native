// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowRedshiftDestinationProperties
    {
        public readonly string? BucketPrefix;
        public readonly Outputs.FlowErrorHandlingConfig? ErrorHandlingConfig;
        public readonly string IntermediateBucketName;
        public readonly string Object;

        [OutputConstructor]
        private FlowRedshiftDestinationProperties(
            string? bucketPrefix,

            Outputs.FlowErrorHandlingConfig? errorHandlingConfig,

            string intermediateBucketName,

            string @object)
        {
            BucketPrefix = bucketPrefix;
            ErrorHandlingConfig = errorHandlingConfig;
            IntermediateBucketName = intermediateBucketName;
            Object = @object;
        }
    }
}
