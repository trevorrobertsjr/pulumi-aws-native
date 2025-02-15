// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateGrowthRateComputation
    {
        public readonly string ComputationId;
        public readonly string? Name;
        public readonly double? PeriodSize;
        public readonly Outputs.TemplateDimensionField? Time;
        public readonly Outputs.TemplateMeasureField? Value;

        [OutputConstructor]
        private TemplateGrowthRateComputation(
            string computationId,

            string? name,

            double? periodSize,

            Outputs.TemplateDimensionField? time,

            Outputs.TemplateMeasureField? value)
        {
            ComputationId = computationId;
            Name = name;
            PeriodSize = periodSize;
            Time = time;
            Value = value;
        }
    }
}
