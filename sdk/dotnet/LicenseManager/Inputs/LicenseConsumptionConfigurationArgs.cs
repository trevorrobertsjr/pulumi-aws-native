// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LicenseManager.Inputs
{

    public sealed class LicenseConsumptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("borrowConfiguration")]
        public Input<Inputs.LicenseBorrowConfigurationArgs>? BorrowConfiguration { get; set; }

        [Input("provisionalConfiguration")]
        public Input<Inputs.LicenseProvisionalConfigurationArgs>? ProvisionalConfiguration { get; set; }

        [Input("renewType")]
        public Input<string>? RenewType { get; set; }

        public LicenseConsumptionConfigurationArgs()
        {
        }
        public static new LicenseConsumptionConfigurationArgs Empty => new LicenseConsumptionConfigurationArgs();
    }
}
