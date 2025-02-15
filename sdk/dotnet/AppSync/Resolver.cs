// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    /// <summary>
    /// Resource Type definition for AWS::AppSync::Resolver
    /// </summary>
    [AwsNativeResourceType("aws-native:appsync:Resolver")]
    public partial class Resolver : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS AppSync GraphQL API to which you want to attach this resolver.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The caching configuration for the resolver.
        /// </summary>
        [Output("cachingConfig")]
        public Output<Outputs.ResolverCachingConfig?> CachingConfig { get; private set; } = null!;

        /// <summary>
        /// The resolver code that contains the request and response functions. When code is used, the runtime is required.
        /// </summary>
        [Output("code")]
        public Output<string?> Code { get; private set; } = null!;

        /// <summary>
        /// The Amazon S3 endpoint.
        /// </summary>
        [Output("codeS3Location")]
        public Output<string?> CodeS3Location { get; private set; } = null!;

        /// <summary>
        /// The resolver data source name.
        /// </summary>
        [Output("dataSourceName")]
        public Output<string?> DataSourceName { get; private set; } = null!;

        /// <summary>
        /// The GraphQL field on a type that invokes the resolver.
        /// </summary>
        [Output("fieldName")]
        public Output<string> FieldName { get; private set; } = null!;

        /// <summary>
        /// The resolver type.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
        /// </summary>
        [Output("maxBatchSize")]
        public Output<int?> MaxBatchSize { get; private set; } = null!;

        /// <summary>
        /// Functions linked with the pipeline resolver.
        /// </summary>
        [Output("pipelineConfig")]
        public Output<Outputs.ResolverPipelineConfig?> PipelineConfig { get; private set; } = null!;

        /// <summary>
        /// Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.
        /// </summary>
        [Output("requestMappingTemplate")]
        public Output<string?> RequestMappingTemplate { get; private set; } = null!;

        /// <summary>
        /// The location of a request mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
        /// </summary>
        [Output("requestMappingTemplateS3Location")]
        public Output<string?> RequestMappingTemplateS3Location { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the resolver.
        /// </summary>
        [Output("resolverArn")]
        public Output<string> ResolverArn { get; private set; } = null!;

        /// <summary>
        /// The response mapping template.
        /// </summary>
        [Output("responseMappingTemplate")]
        public Output<string?> ResponseMappingTemplate { get; private set; } = null!;

        /// <summary>
        /// The location of a response mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
        /// </summary>
        [Output("responseMappingTemplateS3Location")]
        public Output<string?> ResponseMappingTemplateS3Location { get; private set; } = null!;

        /// <summary>
        /// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
        /// </summary>
        [Output("runtime")]
        public Output<Outputs.ResolverAppSyncRuntime?> Runtime { get; private set; } = null!;

        /// <summary>
        /// The SyncConfig for a resolver attached to a versioned data source.
        /// </summary>
        [Output("syncConfig")]
        public Output<Outputs.ResolverSyncConfig?> SyncConfig { get; private set; } = null!;

        /// <summary>
        /// The GraphQL type that invokes this resolver.
        /// </summary>
        [Output("typeName")]
        public Output<string> TypeName { get; private set; } = null!;


        /// <summary>
        /// Create a Resolver resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resolver(string name, ResolverArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appsync:Resolver", name, args ?? new ResolverArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Resolver(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appsync:Resolver", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "apiId",
                    "fieldName",
                    "typeName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Resolver resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resolver Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Resolver(name, id, options);
        }
    }

    public sealed class ResolverArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS AppSync GraphQL API to which you want to attach this resolver.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The caching configuration for the resolver.
        /// </summary>
        [Input("cachingConfig")]
        public Input<Inputs.ResolverCachingConfigArgs>? CachingConfig { get; set; }

        /// <summary>
        /// The resolver code that contains the request and response functions. When code is used, the runtime is required.
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// The Amazon S3 endpoint.
        /// </summary>
        [Input("codeS3Location")]
        public Input<string>? CodeS3Location { get; set; }

        /// <summary>
        /// The resolver data source name.
        /// </summary>
        [Input("dataSourceName")]
        public Input<string>? DataSourceName { get; set; }

        /// <summary>
        /// The GraphQL field on a type that invokes the resolver.
        /// </summary>
        [Input("fieldName", required: true)]
        public Input<string> FieldName { get; set; } = null!;

        /// <summary>
        /// The resolver type.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
        /// </summary>
        [Input("maxBatchSize")]
        public Input<int>? MaxBatchSize { get; set; }

        /// <summary>
        /// Functions linked with the pipeline resolver.
        /// </summary>
        [Input("pipelineConfig")]
        public Input<Inputs.ResolverPipelineConfigArgs>? PipelineConfig { get; set; }

        /// <summary>
        /// Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.
        /// </summary>
        [Input("requestMappingTemplate")]
        public Input<string>? RequestMappingTemplate { get; set; }

        /// <summary>
        /// The location of a request mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
        /// </summary>
        [Input("requestMappingTemplateS3Location")]
        public Input<string>? RequestMappingTemplateS3Location { get; set; }

        /// <summary>
        /// The response mapping template.
        /// </summary>
        [Input("responseMappingTemplate")]
        public Input<string>? ResponseMappingTemplate { get; set; }

        /// <summary>
        /// The location of a response mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
        /// </summary>
        [Input("responseMappingTemplateS3Location")]
        public Input<string>? ResponseMappingTemplateS3Location { get; set; }

        /// <summary>
        /// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
        /// </summary>
        [Input("runtime")]
        public Input<Inputs.ResolverAppSyncRuntimeArgs>? Runtime { get; set; }

        /// <summary>
        /// The SyncConfig for a resolver attached to a versioned data source.
        /// </summary>
        [Input("syncConfig")]
        public Input<Inputs.ResolverSyncConfigArgs>? SyncConfig { get; set; }

        /// <summary>
        /// The GraphQL type that invokes this resolver.
        /// </summary>
        [Input("typeName", required: true)]
        public Input<string> TypeName { get; set; } = null!;

        public ResolverArgs()
        {
        }
        public static new ResolverArgs Empty => new ResolverArgs();
    }
}
