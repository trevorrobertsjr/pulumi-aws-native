// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Metrics that measure the quality of a model.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageModelQuality
    {
        public readonly Outputs.ModelPackageMetricsSource? Constraints;
        public readonly Outputs.ModelPackageMetricsSource? Statistics;

        [OutputConstructor]
        private ModelPackageModelQuality(
            Outputs.ModelPackageMetricsSource? constraints,

            Outputs.ModelPackageMetricsSource? statistics)
        {
            Constraints = constraints;
            Statistics = statistics;
        }
    }
}
