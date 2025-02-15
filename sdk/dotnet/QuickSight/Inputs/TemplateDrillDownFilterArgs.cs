// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateDrillDownFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryFilter")]
        public Input<Inputs.TemplateCategoryDrillDownFilterArgs>? CategoryFilter { get; set; }

        [Input("numericEqualityFilter")]
        public Input<Inputs.TemplateNumericEqualityDrillDownFilterArgs>? NumericEqualityFilter { get; set; }

        [Input("timeRangeFilter")]
        public Input<Inputs.TemplateTimeRangeDrillDownFilterArgs>? TimeRangeFilter { get; set; }

        public TemplateDrillDownFilterArgs()
        {
        }
        public static new TemplateDrillDownFilterArgs Empty => new TemplateDrillDownFilterArgs();
    }
}
