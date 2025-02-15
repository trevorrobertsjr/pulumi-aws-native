// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WorkSpacesWeb
{
    public static class GetTrustStore
    {
        /// <summary>
        /// Definition of AWS::WorkSpacesWeb::TrustStore Resource Type
        /// </summary>
        public static Task<GetTrustStoreResult> InvokeAsync(GetTrustStoreArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrustStoreResult>("aws-native:workspacesweb:getTrustStore", args ?? new GetTrustStoreArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::WorkSpacesWeb::TrustStore Resource Type
        /// </summary>
        public static Output<GetTrustStoreResult> Invoke(GetTrustStoreInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrustStoreResult>("aws-native:workspacesweb:getTrustStore", args ?? new GetTrustStoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrustStoreArgs : global::Pulumi.InvokeArgs
    {
        [Input("trustStoreArn", required: true)]
        public string TrustStoreArn { get; set; } = null!;

        public GetTrustStoreArgs()
        {
        }
        public static new GetTrustStoreArgs Empty => new GetTrustStoreArgs();
    }

    public sealed class GetTrustStoreInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("trustStoreArn", required: true)]
        public Input<string> TrustStoreArn { get; set; } = null!;

        public GetTrustStoreInvokeArgs()
        {
        }
        public static new GetTrustStoreInvokeArgs Empty => new GetTrustStoreInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrustStoreResult
    {
        public readonly ImmutableArray<string> AssociatedPortalArns;
        public readonly ImmutableArray<string> CertificateList;
        public readonly ImmutableArray<Outputs.TrustStoreTag> Tags;
        public readonly string? TrustStoreArn;

        [OutputConstructor]
        private GetTrustStoreResult(
            ImmutableArray<string> associatedPortalArns,

            ImmutableArray<string> certificateList,

            ImmutableArray<Outputs.TrustStoreTag> tags,

            string? trustStoreArn)
        {
            AssociatedPortalArns = associatedPortalArns;
            CertificateList = certificateList;
            Tags = tags;
            TrustStoreArn = trustStoreArn;
        }
    }
}
