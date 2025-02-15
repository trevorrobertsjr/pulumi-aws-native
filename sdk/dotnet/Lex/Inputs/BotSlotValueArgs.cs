// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    /// <summary>
    /// The value to set in a slot.
    /// </summary>
    public sealed class BotSlotValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value that Amazon Lex determines for the slot.
        /// </summary>
        [Input("interpretedValue")]
        public Input<string>? InterpretedValue { get; set; }

        public BotSlotValueArgs()
        {
        }
        public static new BotSlotValueArgs Empty => new BotSlotValueArgs();
    }
}
