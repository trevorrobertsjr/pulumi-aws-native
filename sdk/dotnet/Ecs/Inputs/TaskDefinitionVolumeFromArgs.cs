// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class TaskDefinitionVolumeFromArgs : global::Pulumi.ResourceArgs
    {
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("sourceContainer")]
        public Input<string>? SourceContainer { get; set; }

        public TaskDefinitionVolumeFromArgs()
        {
        }
        public static new TaskDefinitionVolumeFromArgs Empty => new TaskDefinitionVolumeFromArgs();
    }
}
