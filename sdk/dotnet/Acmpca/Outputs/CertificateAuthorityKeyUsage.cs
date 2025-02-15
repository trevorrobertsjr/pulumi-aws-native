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
    /// Structure that contains X.509 KeyUsage information.
    /// </summary>
    [OutputType]
    public sealed class CertificateAuthorityKeyUsage
    {
        public readonly bool? CrlSign;
        public readonly bool? DataEncipherment;
        public readonly bool? DecipherOnly;
        public readonly bool? DigitalSignature;
        public readonly bool? EncipherOnly;
        public readonly bool? KeyAgreement;
        public readonly bool? KeyCertSign;
        public readonly bool? KeyEncipherment;
        public readonly bool? NonRepudiation;

        [OutputConstructor]
        private CertificateAuthorityKeyUsage(
            bool? crlSign,

            bool? dataEncipherment,

            bool? decipherOnly,

            bool? digitalSignature,

            bool? encipherOnly,

            bool? keyAgreement,

            bool? keyCertSign,

            bool? keyEncipherment,

            bool? nonRepudiation)
        {
            CrlSign = crlSign;
            DataEncipherment = dataEncipherment;
            DecipherOnly = decipherOnly;
            DigitalSignature = digitalSignature;
            EncipherOnly = encipherOnly;
            KeyAgreement = keyAgreement;
            KeyCertSign = keyCertSign;
            KeyEncipherment = keyEncipherment;
            NonRepudiation = nonRepudiation;
        }
    }
}
