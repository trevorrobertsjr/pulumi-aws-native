// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    public sealed class WebAclAndStatementArgs : global::Pulumi.ResourceArgs
    {
        [Input("statements", required: true)]
        private InputList<Inputs.WebAclStatementArgs>? _statements;
        public InputList<Inputs.WebAclStatementArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.WebAclStatementArgs>());
            set => _statements = value;
        }

        public WebAclAndStatementArgs()
        {
        }
        public static new WebAclAndStatementArgs Empty => new WebAclAndStatementArgs();
    }
}
