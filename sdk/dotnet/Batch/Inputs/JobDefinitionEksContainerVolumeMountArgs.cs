// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionEksContainerVolumeMountArgs : global::Pulumi.ResourceArgs
    {
        [Input("mountPath")]
        public Input<string>? MountPath { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        public JobDefinitionEksContainerVolumeMountArgs()
        {
        }
        public static new JobDefinitionEksContainerVolumeMountArgs Empty => new JobDefinitionEksContainerVolumeMountArgs();
    }
}
