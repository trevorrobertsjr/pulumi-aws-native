// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Response status codes that indicate success or failure of a login request
    /// </summary>
    [OutputType]
    public sealed class WebAclResponseInspectionStatusCode
    {
        public readonly ImmutableArray<int> FailureCodes;
        public readonly ImmutableArray<int> SuccessCodes;

        [OutputConstructor]
        private WebAclResponseInspectionStatusCode(
            ImmutableArray<int> failureCodes,

            ImmutableArray<int> successCodes)
        {
            FailureCodes = failureCodes;
            SuccessCodes = successCodes;
        }
    }
}
