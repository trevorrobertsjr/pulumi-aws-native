// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker
{
    public static class GetSimulationApplicationVersion
    {
        /// <summary>
        /// AWS::RoboMaker::SimulationApplicationVersion resource creates an AWS RoboMaker SimulationApplicationVersion. This helps you control which code your simulation uses.
        /// </summary>
        public static Task<GetSimulationApplicationVersionResult> InvokeAsync(GetSimulationApplicationVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSimulationApplicationVersionResult>("aws-native:robomaker:getSimulationApplicationVersion", args ?? new GetSimulationApplicationVersionArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::RoboMaker::SimulationApplicationVersion resource creates an AWS RoboMaker SimulationApplicationVersion. This helps you control which code your simulation uses.
        /// </summary>
        public static Output<GetSimulationApplicationVersionResult> Invoke(GetSimulationApplicationVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSimulationApplicationVersionResult>("aws-native:robomaker:getSimulationApplicationVersion", args ?? new GetSimulationApplicationVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSimulationApplicationVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetSimulationApplicationVersionArgs()
        {
        }
        public static new GetSimulationApplicationVersionArgs Empty => new GetSimulationApplicationVersionArgs();
    }

    public sealed class GetSimulationApplicationVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetSimulationApplicationVersionInvokeArgs()
        {
        }
        public static new GetSimulationApplicationVersionInvokeArgs Empty => new GetSimulationApplicationVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSimulationApplicationVersionResult
    {
        public readonly string? ApplicationVersion;
        public readonly string? Arn;

        [OutputConstructor]
        private GetSimulationApplicationVersionResult(
            string? applicationVersion,

            string? arn)
        {
            ApplicationVersion = applicationVersion;
            Arn = arn;
        }
    }
}
