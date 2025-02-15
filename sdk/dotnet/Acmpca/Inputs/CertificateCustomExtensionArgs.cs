// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Inputs
{

    /// <summary>
    /// Structure that contains X.509 extension information for a certificate.
    /// </summary>
    public sealed class CertificateCustomExtensionArgs : global::Pulumi.ResourceArgs
    {
        [Input("critical")]
        public Input<bool>? Critical { get; set; }

        [Input("objectIdentifier", required: true)]
        public Input<string> ObjectIdentifier { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public CertificateCustomExtensionArgs()
        {
        }
        public static new CertificateCustomExtensionArgs Empty => new CertificateCustomExtensionArgs();
    }
}
