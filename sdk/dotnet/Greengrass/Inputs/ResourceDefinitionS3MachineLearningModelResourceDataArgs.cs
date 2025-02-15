// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Inputs
{

    public sealed class ResourceDefinitionS3MachineLearningModelResourceDataArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationPath", required: true)]
        public Input<string> DestinationPath { get; set; } = null!;

        [Input("ownerSetting")]
        public Input<Inputs.ResourceDefinitionResourceDownloadOwnerSettingArgs>? OwnerSetting { get; set; }

        [Input("s3Uri", required: true)]
        public Input<string> S3Uri { get; set; } = null!;

        public ResourceDefinitionS3MachineLearningModelResourceDataArgs()
        {
        }
        public static new ResourceDefinitionS3MachineLearningModelResourceDataArgs Empty => new ResourceDefinitionS3MachineLearningModelResourceDataArgs();
    }
}
