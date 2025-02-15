// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeTargetHttpParameters
    {
        public readonly Outputs.PipeHeaderParametersMap? HeaderParameters;
        public readonly ImmutableArray<string> PathParameterValues;
        public readonly Outputs.PipeQueryStringParametersMap? QueryStringParameters;

        [OutputConstructor]
        private PipeTargetHttpParameters(
            Outputs.PipeHeaderParametersMap? headerParameters,

            ImmutableArray<string> pathParameterValues,

            Outputs.PipeQueryStringParametersMap? queryStringParameters)
        {
            HeaderParameters = headerParameters;
            PathParameterValues = pathParameterValues;
            QueryStringParameters = queryStringParameters;
        }
    }
}
