// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class VirtualNodeBackendDefaultsArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientPolicy")]
        public Input<Inputs.VirtualNodeClientPolicyArgs>? ClientPolicy { get; set; }

        public VirtualNodeBackendDefaultsArgs()
        {
        }
        public static new VirtualNodeBackendDefaultsArgs Empty => new VirtualNodeBackendDefaultsArgs();
    }
}
