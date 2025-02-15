// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Inputs
{

    public sealed class DeviceDefinitionVersionDeviceArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateArn", required: true)]
        public Input<string> CertificateArn { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("syncShadow")]
        public Input<bool>? SyncShadow { get; set; }

        [Input("thingArn", required: true)]
        public Input<string> ThingArn { get; set; } = null!;

        public DeviceDefinitionVersionDeviceArgs()
        {
        }
        public static new DeviceDefinitionVersionDeviceArgs Empty => new DeviceDefinitionVersionDeviceArgs();
    }
}
