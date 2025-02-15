// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Outputs
{

    /// <summary>
    /// API stage name of the associated API stage in a usage plan.
    /// </summary>
    [OutputType]
    public sealed class UsagePlanApiStage
    {
        /// <summary>
        /// API Id of the associated API stage in a usage plan.
        /// </summary>
        public readonly string? ApiId;
        /// <summary>
        /// API stage name of the associated API stage in a usage plan.
        /// </summary>
        public readonly string? Stage;
        /// <summary>
        /// Map containing method level throttling information for API stage in a usage plan.
        /// </summary>
        public readonly object? Throttle;

        [OutputConstructor]
        private UsagePlanApiStage(
            string? apiId,

            string? stage,

            object? throttle)
        {
            ApiId = apiId;
            Stage = stage;
            Throttle = throttle;
        }
    }
}
