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
    /// Determines whether Amazon Lex obscures slot values in conversation logs.
    /// </summary>
    public sealed class BotObfuscationSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value that determines whether Amazon Lex obscures slot values in conversation logs. The default is to obscure the values.
        /// </summary>
        [Input("obfuscationSettingType", required: true)]
        public Input<Pulumi.AwsNative.Lex.BotObfuscationSettingObfuscationSettingType> ObfuscationSettingType { get; set; } = null!;

        public BotObfuscationSettingArgs()
        {
        }
        public static new BotObfuscationSettingArgs Empty => new BotObfuscationSettingArgs();
    }
}
