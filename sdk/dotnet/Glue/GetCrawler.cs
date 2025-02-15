// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetCrawler
    {
        /// <summary>
        /// Resource Type definition for AWS::Glue::Crawler
        /// </summary>
        public static Task<GetCrawlerResult> InvokeAsync(GetCrawlerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCrawlerResult>("aws-native:glue:getCrawler", args ?? new GetCrawlerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Glue::Crawler
        /// </summary>
        public static Output<GetCrawlerResult> Invoke(GetCrawlerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCrawlerResult>("aws-native:glue:getCrawler", args ?? new GetCrawlerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCrawlerArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetCrawlerArgs()
        {
        }
        public static new GetCrawlerArgs Empty => new GetCrawlerArgs();
    }

    public sealed class GetCrawlerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetCrawlerInvokeArgs()
        {
        }
        public static new GetCrawlerInvokeArgs Empty => new GetCrawlerInvokeArgs();
    }


    [OutputType]
    public sealed class GetCrawlerResult
    {
        public readonly ImmutableArray<string> Classifiers;
        public readonly string? Configuration;
        public readonly string? CrawlerSecurityConfiguration;
        public readonly string? DatabaseName;
        public readonly string? Description;
        public readonly string? Id;
        public readonly Outputs.CrawlerRecrawlPolicy? RecrawlPolicy;
        public readonly string? Role;
        public readonly Outputs.CrawlerSchedule? Schedule;
        public readonly Outputs.CrawlerSchemaChangePolicy? SchemaChangePolicy;
        public readonly string? TablePrefix;
        public readonly object? Tags;
        public readonly Outputs.CrawlerTargets? Targets;

        [OutputConstructor]
        private GetCrawlerResult(
            ImmutableArray<string> classifiers,

            string? configuration,

            string? crawlerSecurityConfiguration,

            string? databaseName,

            string? description,

            string? id,

            Outputs.CrawlerRecrawlPolicy? recrawlPolicy,

            string? role,

            Outputs.CrawlerSchedule? schedule,

            Outputs.CrawlerSchemaChangePolicy? schemaChangePolicy,

            string? tablePrefix,

            object? tags,

            Outputs.CrawlerTargets? targets)
        {
            Classifiers = classifiers;
            Configuration = configuration;
            CrawlerSecurityConfiguration = crawlerSecurityConfiguration;
            DatabaseName = databaseName;
            Description = description;
            Id = id;
            RecrawlPolicy = recrawlPolicy;
            Role = role;
            Schedule = schedule;
            SchemaChangePolicy = schemaChangePolicy;
            TablePrefix = tablePrefix;
            Tags = tags;
            Targets = targets;
        }
    }
}
