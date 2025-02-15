// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Outputs
{

    /// <summary>
    /// Structure that contains X.509 ExtendedKeyUsage information.
    /// </summary>
    [OutputType]
    public sealed class CertificateExtendedKeyUsage
    {
        public readonly string? ExtendedKeyUsageObjectIdentifier;
        public readonly string? ExtendedKeyUsageType;

        [OutputConstructor]
        private CertificateExtendedKeyUsage(
            string? extendedKeyUsageObjectIdentifier,

            string? extendedKeyUsageType)
        {
            ExtendedKeyUsageObjectIdentifier = extendedKeyUsageObjectIdentifier;
            ExtendedKeyUsageType = extendedKeyUsageType;
        }
    }
}
