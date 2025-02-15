// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca
{
    /// <summary>
    /// Private certificate authority.
    /// </summary>
    [AwsNativeResourceType("aws-native:acmpca:CertificateAuthority")]
    public partial class CertificateAuthority : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the certificate authority.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
        /// </summary>
        [Output("certificateSigningRequest")]
        public Output<string> CertificateSigningRequest { get; private set; } = null!;

        /// <summary>
        /// Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.
        /// </summary>
        [Output("csrExtensions")]
        public Output<Outputs.CertificateAuthorityCsrExtensions?> CsrExtensions { get; private set; } = null!;

        /// <summary>
        /// Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
        /// </summary>
        [Output("keyAlgorithm")]
        public Output<string> KeyAlgorithm { get; private set; } = null!;

        /// <summary>
        /// KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
        /// </summary>
        [Output("keyStorageSecurityStandard")]
        public Output<string?> KeyStorageSecurityStandard { get; private set; } = null!;

        /// <summary>
        /// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
        /// </summary>
        [Output("revocationConfiguration")]
        public Output<Outputs.CertificateAuthorityRevocationConfiguration?> RevocationConfiguration { get; private set; } = null!;

        /// <summary>
        /// Algorithm your CA uses to sign certificate requests.
        /// </summary>
        [Output("signingAlgorithm")]
        public Output<string> SigningAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Structure that contains X.500 distinguished name information for your CA.
        /// </summary>
        [Output("subject")]
        public Output<Outputs.CertificateAuthoritySubject> Subject { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.CertificateAuthorityTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the certificate authority.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Usage mode of the ceritificate authority.
        /// </summary>
        [Output("usageMode")]
        public Output<string?> UsageMode { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateAuthority resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateAuthority(string name, CertificateAuthorityArgs args, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:CertificateAuthority", name, args ?? new CertificateAuthorityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateAuthority(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:CertificateAuthority", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "csrExtensions",
                    "keyAlgorithm",
                    "keyStorageSecurityStandard",
                    "signingAlgorithm",
                    "subject",
                    "type",
                    "usageMode",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertificateAuthority resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateAuthority Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CertificateAuthority(name, id, options);
        }
    }

    public sealed class CertificateAuthorityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.
        /// </summary>
        [Input("csrExtensions")]
        public Input<Inputs.CertificateAuthorityCsrExtensionsArgs>? CsrExtensions { get; set; }

        /// <summary>
        /// Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
        /// </summary>
        [Input("keyAlgorithm", required: true)]
        public Input<string> KeyAlgorithm { get; set; } = null!;

        /// <summary>
        /// KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
        /// </summary>
        [Input("keyStorageSecurityStandard")]
        public Input<string>? KeyStorageSecurityStandard { get; set; }

        /// <summary>
        /// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
        /// </summary>
        [Input("revocationConfiguration")]
        public Input<Inputs.CertificateAuthorityRevocationConfigurationArgs>? RevocationConfiguration { get; set; }

        /// <summary>
        /// Algorithm your CA uses to sign certificate requests.
        /// </summary>
        [Input("signingAlgorithm", required: true)]
        public Input<string> SigningAlgorithm { get; set; } = null!;

        /// <summary>
        /// Structure that contains X.500 distinguished name information for your CA.
        /// </summary>
        [Input("subject", required: true)]
        public Input<Inputs.CertificateAuthoritySubjectArgs> Subject { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.CertificateAuthorityTagArgs>? _tags;
        public InputList<Inputs.CertificateAuthorityTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CertificateAuthorityTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the certificate authority.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Usage mode of the ceritificate authority.
        /// </summary>
        [Input("usageMode")]
        public Input<string>? UsageMode { get; set; }

        public CertificateAuthorityArgs()
        {
        }
        public static new CertificateAuthorityArgs Empty => new CertificateAuthorityArgs();
    }
}
