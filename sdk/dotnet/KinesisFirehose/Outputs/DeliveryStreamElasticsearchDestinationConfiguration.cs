// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamElasticsearchDestinationConfiguration
    {
        public readonly Outputs.DeliveryStreamElasticsearchBufferingHints? BufferingHints;
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        public readonly string? ClusterEndpoint;
        public readonly Outputs.DeliveryStreamDocumentIdOptions? DocumentIdOptions;
        public readonly string? DomainArn;
        public readonly string IndexName;
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod? IndexRotationPeriod;
        public readonly Outputs.DeliveryStreamProcessingConfiguration? ProcessingConfiguration;
        public readonly Outputs.DeliveryStreamElasticsearchRetryOptions? RetryOptions;
        public readonly string RoleArn;
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode? S3BackupMode;
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration S3Configuration;
        public readonly string? TypeName;
        public readonly Outputs.DeliveryStreamVpcConfiguration? VpcConfiguration;

        [OutputConstructor]
        private DeliveryStreamElasticsearchDestinationConfiguration(
            Outputs.DeliveryStreamElasticsearchBufferingHints? bufferingHints,

            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            string? clusterEndpoint,

            Outputs.DeliveryStreamDocumentIdOptions? documentIdOptions,

            string? domainArn,

            string indexName,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod? indexRotationPeriod,

            Outputs.DeliveryStreamProcessingConfiguration? processingConfiguration,

            Outputs.DeliveryStreamElasticsearchRetryOptions? retryOptions,

            string roleArn,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode? s3BackupMode,

            Outputs.DeliveryStreamS3DestinationConfiguration s3Configuration,

            string? typeName,

            Outputs.DeliveryStreamVpcConfiguration? vpcConfiguration)
        {
            BufferingHints = bufferingHints;
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            ClusterEndpoint = clusterEndpoint;
            DocumentIdOptions = documentIdOptions;
            DomainArn = domainArn;
            IndexName = indexName;
            IndexRotationPeriod = indexRotationPeriod;
            ProcessingConfiguration = processingConfiguration;
            RetryOptions = retryOptions;
            RoleArn = roleArn;
            S3BackupMode = s3BackupMode;
            S3Configuration = s3Configuration;
            TypeName = typeName;
            VpcConfiguration = vpcConfiguration;
        }
    }
}
