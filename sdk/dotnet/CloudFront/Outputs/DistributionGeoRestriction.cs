// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    [OutputType]
    public sealed class DistributionGeoRestriction
    {
        public readonly ImmutableArray<string> Locations;
        public readonly string RestrictionType;

        [OutputConstructor]
        private DistributionGeoRestriction(
            ImmutableArray<string> locations,

            string restrictionType)
        {
            Locations = locations;
            RestrictionType = restrictionType;
        }
    }
}
