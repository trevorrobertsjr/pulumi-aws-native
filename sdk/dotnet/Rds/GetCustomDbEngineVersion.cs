// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rds
{
    public static class GetCustomDbEngineVersion
    {
        /// <summary>
        /// The AWS::RDS::CustomDBEngineVersion resource creates an Amazon RDS custom DB engine version.
        /// </summary>
        public static Task<GetCustomDbEngineVersionResult> InvokeAsync(GetCustomDbEngineVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomDbEngineVersionResult>("aws-native:rds:getCustomDbEngineVersion", args ?? new GetCustomDbEngineVersionArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::RDS::CustomDBEngineVersion resource creates an Amazon RDS custom DB engine version.
        /// </summary>
        public static Output<GetCustomDbEngineVersionResult> Invoke(GetCustomDbEngineVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomDbEngineVersionResult>("aws-native:rds:getCustomDbEngineVersion", args ?? new GetCustomDbEngineVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomDbEngineVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database engine to use for your custom engine version (CEV). The only supported value is `custom-oracle-ee`.
        /// </summary>
        [Input("engine", required: true)]
        public string Engine { get; set; } = null!;

        /// <summary>
        /// The name of your CEV. The name format is 19.customized_string . For example, a valid name is 19.my_cev1. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of Engine and EngineVersion is unique per customer per Region.
        /// </summary>
        [Input("engineVersion", required: true)]
        public string EngineVersion { get; set; } = null!;

        public GetCustomDbEngineVersionArgs()
        {
        }
        public static new GetCustomDbEngineVersionArgs Empty => new GetCustomDbEngineVersionArgs();
    }

    public sealed class GetCustomDbEngineVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database engine to use for your custom engine version (CEV). The only supported value is `custom-oracle-ee`.
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// The name of your CEV. The name format is 19.customized_string . For example, a valid name is 19.my_cev1. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of Engine and EngineVersion is unique per customer per Region.
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        public GetCustomDbEngineVersionInvokeArgs()
        {
        }
        public static new GetCustomDbEngineVersionInvokeArgs Empty => new GetCustomDbEngineVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomDbEngineVersionResult
    {
        /// <summary>
        /// The ARN of the custom engine version.
        /// </summary>
        public readonly string? DbEngineVersionArn;
        /// <summary>
        /// An optional description of your CEV.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The availability status to be assigned to the CEV.
        /// </summary>
        public readonly Pulumi.AwsNative.Rds.CustomDbEngineVersionStatus? Status;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomDbEngineVersionTag> Tags;

        [OutputConstructor]
        private GetCustomDbEngineVersionResult(
            string? dbEngineVersionArn,

            string? description,

            Pulumi.AwsNative.Rds.CustomDbEngineVersionStatus? status,

            ImmutableArray<Outputs.CustomDbEngineVersionTag> tags)
        {
            DbEngineVersionArn = dbEngineVersionArn;
            Description = description;
            Status = status;
            Tags = tags;
        }
    }
}
