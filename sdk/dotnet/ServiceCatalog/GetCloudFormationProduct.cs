// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    public static class GetCloudFormationProduct
    {
        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::CloudFormationProduct
        /// </summary>
        public static Task<GetCloudFormationProductResult> InvokeAsync(GetCloudFormationProductArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudFormationProductResult>("aws-native:servicecatalog:getCloudFormationProduct", args ?? new GetCloudFormationProductArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::CloudFormationProduct
        /// </summary>
        public static Output<GetCloudFormationProductResult> Invoke(GetCloudFormationProductInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudFormationProductResult>("aws-native:servicecatalog:getCloudFormationProduct", args ?? new GetCloudFormationProductInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudFormationProductArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetCloudFormationProductArgs()
        {
        }
        public static new GetCloudFormationProductArgs Empty => new GetCloudFormationProductArgs();
    }

    public sealed class GetCloudFormationProductInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetCloudFormationProductInvokeArgs()
        {
        }
        public static new GetCloudFormationProductInvokeArgs Empty => new GetCloudFormationProductInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudFormationProductResult
    {
        public readonly string? AcceptLanguage;
        public readonly string? Description;
        public readonly string? Distributor;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? Owner;
        public readonly string? ProductName;
        public readonly string? ProductType;
        public readonly string? ProvisioningArtifactIds;
        public readonly string? ProvisioningArtifactNames;
        public readonly ImmutableArray<Outputs.CloudFormationProductProvisioningArtifactProperties> ProvisioningArtifactParameters;
        public readonly bool? ReplaceProvisioningArtifacts;
        public readonly Outputs.CloudFormationProductSourceConnection? SourceConnection;
        public readonly string? SupportDescription;
        public readonly string? SupportEmail;
        public readonly string? SupportUrl;
        public readonly ImmutableArray<Outputs.CloudFormationProductTag> Tags;

        [OutputConstructor]
        private GetCloudFormationProductResult(
            string? acceptLanguage,

            string? description,

            string? distributor,

            string? id,

            string? name,

            string? owner,

            string? productName,

            string? productType,

            string? provisioningArtifactIds,

            string? provisioningArtifactNames,

            ImmutableArray<Outputs.CloudFormationProductProvisioningArtifactProperties> provisioningArtifactParameters,

            bool? replaceProvisioningArtifacts,

            Outputs.CloudFormationProductSourceConnection? sourceConnection,

            string? supportDescription,

            string? supportEmail,

            string? supportUrl,

            ImmutableArray<Outputs.CloudFormationProductTag> tags)
        {
            AcceptLanguage = acceptLanguage;
            Description = description;
            Distributor = distributor;
            Id = id;
            Name = name;
            Owner = owner;
            ProductName = productName;
            ProductType = productType;
            ProvisioningArtifactIds = provisioningArtifactIds;
            ProvisioningArtifactNames = provisioningArtifactNames;
            ProvisioningArtifactParameters = provisioningArtifactParameters;
            ReplaceProvisioningArtifacts = replaceProvisioningArtifacts;
            SourceConnection = sourceConnection;
            SupportDescription = supportDescription;
            SupportEmail = supportEmail;
            SupportUrl = supportUrl;
            Tags = tags;
        }
    }
}
