// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateForecastScenarioArgs : global::Pulumi.ResourceArgs
    {
        [Input("whatIfPointScenario")]
        public Input<Inputs.TemplateWhatIfPointScenarioArgs>? WhatIfPointScenario { get; set; }

        [Input("whatIfRangeScenario")]
        public Input<Inputs.TemplateWhatIfRangeScenarioArgs>? WhatIfRangeScenario { get; set; }

        public TemplateForecastScenarioArgs()
        {
        }
        public static new TemplateForecastScenarioArgs Empty => new TemplateForecastScenarioArgs();
    }
}
