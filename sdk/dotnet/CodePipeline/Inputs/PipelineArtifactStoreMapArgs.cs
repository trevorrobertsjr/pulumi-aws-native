// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Inputs
{

    public sealed class PipelineArtifactStoreMapArgs : global::Pulumi.ResourceArgs
    {
        [Input("artifactStore", required: true)]
        public Input<Inputs.PipelineArtifactStoreArgs> ArtifactStore { get; set; } = null!;

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public PipelineArtifactStoreMapArgs()
        {
        }
        public static new PipelineArtifactStoreMapArgs Empty => new PipelineArtifactStoreMapArgs();
    }
}
