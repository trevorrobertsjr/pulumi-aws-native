// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    [OutputType]
    public sealed class JobDefinitionHost
    {
        public readonly string? SourcePath;

        [OutputConstructor]
        private JobDefinitionHost(string? sourcePath)
        {
            SourcePath = sourcePath;
        }
    }
}
