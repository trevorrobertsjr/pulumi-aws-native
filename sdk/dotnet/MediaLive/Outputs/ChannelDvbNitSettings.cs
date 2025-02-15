// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelDvbNitSettings
    {
        public readonly int? NetworkId;
        public readonly string? NetworkName;
        public readonly int? RepInterval;

        [OutputConstructor]
        private ChannelDvbNitSettings(
            int? networkId,

            string? networkName,

            int? repInterval)
        {
            NetworkId = networkId;
            NetworkName = networkName;
            RepInterval = repInterval;
        }
    }
}
