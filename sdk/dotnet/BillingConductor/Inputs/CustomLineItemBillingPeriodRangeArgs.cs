// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.BillingConductor.Inputs
{

    public sealed class CustomLineItemBillingPeriodRangeArgs : global::Pulumi.ResourceArgs
    {
        [Input("exclusiveEndBillingPeriod")]
        public Input<string>? ExclusiveEndBillingPeriod { get; set; }

        [Input("inclusiveStartBillingPeriod")]
        public Input<string>? InclusiveStartBillingPeriod { get; set; }

        public CustomLineItemBillingPeriodRangeArgs()
        {
        }
        public static new CustomLineItemBillingPeriodRangeArgs Empty => new CustomLineItemBillingPeriodRangeArgs();
    }
}
