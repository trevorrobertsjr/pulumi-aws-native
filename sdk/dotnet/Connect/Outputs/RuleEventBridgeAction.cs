// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// The definition for event bridge action.
    /// </summary>
    [OutputType]
    public sealed class RuleEventBridgeAction
    {
        /// <summary>
        /// The name of the event bridge action.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private RuleEventBridgeAction(string name)
        {
            Name = name;
        }
    }
}
