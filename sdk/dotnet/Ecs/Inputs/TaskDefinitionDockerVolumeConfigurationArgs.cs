// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class TaskDefinitionDockerVolumeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoprovision")]
        public Input<bool>? Autoprovision { get; set; }

        [Input("driver")]
        public Input<string>? Driver { get; set; }

        [Input("driverOpts")]
        public Input<object>? DriverOpts { get; set; }

        [Input("labels")]
        public Input<object>? Labels { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public TaskDefinitionDockerVolumeConfigurationArgs()
        {
        }
        public static new TaskDefinitionDockerVolumeConfigurationArgs Empty => new TaskDefinitionDockerVolumeConfigurationArgs();
    }
}
