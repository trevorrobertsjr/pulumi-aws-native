// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Inspect the request body as JSON. The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. 
    /// </summary>
    public sealed class LoggingConfigurationFieldToMatchJsonBodyPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// What AWS WAF should do if it fails to completely parse the JSON body.
        /// </summary>
        [Input("invalidFallbackBehavior")]
        public Input<Pulumi.AwsNative.WaFv2.LoggingConfigurationFieldToMatchJsonBodyPropertiesInvalidFallbackBehavior>? InvalidFallbackBehavior { get; set; }

        /// <summary>
        /// The patterns to look for in the JSON body. AWS WAF inspects the results of these pattern matches against the rule inspection criteria. 
        /// </summary>
        [Input("matchPattern", required: true)]
        public Input<Inputs.LoggingConfigurationFieldToMatchJsonBodyPropertiesMatchPatternPropertiesArgs> MatchPattern { get; set; } = null!;

        /// <summary>
        /// The parts of the JSON to match against using the MatchPattern. If you specify All, AWS WAF matches against keys and values. 
        /// </summary>
        [Input("matchScope", required: true)]
        public Input<Pulumi.AwsNative.WaFv2.LoggingConfigurationFieldToMatchJsonBodyPropertiesMatchScope> MatchScope { get; set; } = null!;

        public LoggingConfigurationFieldToMatchJsonBodyPropertiesArgs()
        {
        }
        public static new LoggingConfigurationFieldToMatchJsonBodyPropertiesArgs Empty => new LoggingConfigurationFieldToMatchJsonBodyPropertiesArgs();
    }
}
