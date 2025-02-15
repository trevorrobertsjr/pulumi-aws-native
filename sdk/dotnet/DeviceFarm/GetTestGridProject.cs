// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DeviceFarm
{
    public static class GetTestGridProject
    {
        /// <summary>
        /// AWS::DeviceFarm::TestGridProject creates a new TestGrid Project
        /// </summary>
        public static Task<GetTestGridProjectResult> InvokeAsync(GetTestGridProjectArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTestGridProjectResult>("aws-native:devicefarm:getTestGridProject", args ?? new GetTestGridProjectArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::DeviceFarm::TestGridProject creates a new TestGrid Project
        /// </summary>
        public static Output<GetTestGridProjectResult> Invoke(GetTestGridProjectInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTestGridProjectResult>("aws-native:devicefarm:getTestGridProject", args ?? new GetTestGridProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTestGridProjectArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetTestGridProjectArgs()
        {
        }
        public static new GetTestGridProjectArgs Empty => new GetTestGridProjectArgs();
    }

    public sealed class GetTestGridProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetTestGridProjectInvokeArgs()
        {
        }
        public static new GetTestGridProjectInvokeArgs Empty => new GetTestGridProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetTestGridProjectResult
    {
        public readonly string? Arn;
        public readonly string? Description;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.TestGridProjectTag> Tags;

        [OutputConstructor]
        private GetTestGridProjectResult(
            string? arn,

            string? description,

            string? name,

            ImmutableArray<Outputs.TestGridProjectTag> tags)
        {
            Arn = arn;
            Description = description;
            Name = name;
            Tags = tags;
        }
    }
}
