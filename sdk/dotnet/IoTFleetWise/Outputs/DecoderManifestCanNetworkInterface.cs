// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTFleetWise.Outputs
{

    [OutputType]
    public sealed class DecoderManifestCanNetworkInterface
    {
        public readonly Outputs.DecoderManifestCanInterface CanInterface;
        public readonly string InterfaceId;
        public readonly Pulumi.AwsNative.IoTFleetWise.DecoderManifestCanNetworkInterfaceType Type;

        [OutputConstructor]
        private DecoderManifestCanNetworkInterface(
            Outputs.DecoderManifestCanInterface canInterface,

            string interfaceId,

            Pulumi.AwsNative.IoTFleetWise.DecoderManifestCanNetworkInterfaceType type)
        {
            CanInterface = canInterface;
            InterfaceId = interfaceId;
            Type = type;
        }
    }
}
