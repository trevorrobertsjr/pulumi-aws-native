// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeEcsEnvironmentFileArgs : global::Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Pipes.PipeEcsEnvironmentFileType> Type { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public PipeEcsEnvironmentFileArgs()
        {
        }
        public static new PipeEcsEnvironmentFileArgs Empty => new PipeEcsEnvironmentFileArgs();
    }
}
