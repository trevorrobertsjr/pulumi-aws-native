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
    public sealed class DashboardParameters
    {
        public readonly ImmutableArray<Outputs.DashboardDateTimeParameter> DateTimeParameters;
        public readonly ImmutableArray<Outputs.DashboardDecimalParameter> DecimalParameters;
        public readonly ImmutableArray<Outputs.DashboardIntegerParameter> IntegerParameters;
        public readonly ImmutableArray<Outputs.DashboardStringParameter> StringParameters;

        [OutputConstructor]
        private DashboardParameters(
            ImmutableArray<Outputs.DashboardDateTimeParameter> dateTimeParameters,

            ImmutableArray<Outputs.DashboardDecimalParameter> decimalParameters,

            ImmutableArray<Outputs.DashboardIntegerParameter> integerParameters,

            ImmutableArray<Outputs.DashboardStringParameter> stringParameters)
        {
            DateTimeParameters = dateTimeParameters;
            DecimalParameters = decimalParameters;
            IntegerParameters = integerParameters;
            StringParameters = stringParameters;
        }
    }
}
