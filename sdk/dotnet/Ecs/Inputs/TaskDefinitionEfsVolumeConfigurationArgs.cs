// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class TaskDefinitionEfsVolumeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorizationConfig")]
        public Input<Inputs.TaskDefinitionAuthorizationConfigArgs>? AuthorizationConfig { get; set; }

        [Input("filesystemId", required: true)]
        public Input<string> FilesystemId { get; set; } = null!;

        [Input("rootDirectory")]
        public Input<string>? RootDirectory { get; set; }

        [Input("transitEncryption")]
        public Input<Pulumi.AwsNative.Ecs.TaskDefinitionEfsVolumeConfigurationTransitEncryption>? TransitEncryption { get; set; }

        [Input("transitEncryptionPort")]
        public Input<int>? TransitEncryptionPort { get; set; }

        public TaskDefinitionEfsVolumeConfigurationArgs()
        {
        }
        public static new TaskDefinitionEfsVolumeConfigurationArgs Empty => new TaskDefinitionEfsVolumeConfigurationArgs();
    }
}
