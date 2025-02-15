// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca
{
    public static class GetCertificateAuthority
    {
        /// <summary>
        /// Private certificate authority.
        /// </summary>
        public static Task<GetCertificateAuthorityResult> InvokeAsync(GetCertificateAuthorityArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateAuthorityResult>("aws-native:acmpca:getCertificateAuthority", args ?? new GetCertificateAuthorityArgs(), options.WithDefaults());

        /// <summary>
        /// Private certificate authority.
        /// </summary>
        public static Output<GetCertificateAuthorityResult> Invoke(GetCertificateAuthorityInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateAuthorityResult>("aws-native:acmpca:getCertificateAuthority", args ?? new GetCertificateAuthorityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateAuthorityArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the certificate authority.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetCertificateAuthorityArgs()
        {
        }
        public static new GetCertificateAuthorityArgs Empty => new GetCertificateAuthorityArgs();
    }

    public sealed class GetCertificateAuthorityInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the certificate authority.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetCertificateAuthorityInvokeArgs()
        {
        }
        public static new GetCertificateAuthorityInvokeArgs Empty => new GetCertificateAuthorityInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateAuthorityResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the certificate authority.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
        /// </summary>
        public readonly string? CertificateSigningRequest;
        /// <summary>
        /// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
        /// </summary>
        public readonly Outputs.CertificateAuthorityRevocationConfiguration? RevocationConfiguration;
        public readonly ImmutableArray<Outputs.CertificateAuthorityTag> Tags;

        [OutputConstructor]
        private GetCertificateAuthorityResult(
            string? arn,

            string? certificateSigningRequest,

            Outputs.CertificateAuthorityRevocationConfiguration? revocationConfiguration,

            ImmutableArray<Outputs.CertificateAuthorityTag> tags)
        {
            Arn = arn;
            CertificateSigningRequest = certificateSigningRequest;
            RevocationConfiguration = revocationConfiguration;
            Tags = tags;
        }
    }
}
