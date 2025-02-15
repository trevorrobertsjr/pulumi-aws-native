// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    public static class GetWaitCondition
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFormation::WaitCondition
        /// </summary>
        public static Task<GetWaitConditionResult> InvokeAsync(GetWaitConditionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWaitConditionResult>("aws-native:cloudformation:getWaitCondition", args ?? new GetWaitConditionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFormation::WaitCondition
        /// </summary>
        public static Output<GetWaitConditionResult> Invoke(GetWaitConditionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWaitConditionResult>("aws-native:cloudformation:getWaitCondition", args ?? new GetWaitConditionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWaitConditionArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWaitConditionArgs()
        {
        }
        public static new GetWaitConditionArgs Empty => new GetWaitConditionArgs();
    }

    public sealed class GetWaitConditionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWaitConditionInvokeArgs()
        {
        }
        public static new GetWaitConditionInvokeArgs Empty => new GetWaitConditionInvokeArgs();
    }


    [OutputType]
    public sealed class GetWaitConditionResult
    {
        public readonly int? Count;
        public readonly object? Data;
        public readonly string? Handle;
        public readonly string? Id;
        public readonly string? Timeout;

        [OutputConstructor]
        private GetWaitConditionResult(
            int? count,

            object? data,

            string? handle,

            string? id,

            string? timeout)
        {
            Count = count;
            Data = data;
            Handle = handle;
            Id = id;
            Timeout = timeout;
        }
    }
}
