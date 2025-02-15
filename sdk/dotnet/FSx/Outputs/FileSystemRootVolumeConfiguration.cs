// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Outputs
{

    [OutputType]
    public sealed class FileSystemRootVolumeConfiguration
    {
        public readonly bool? CopyTagsToSnapshots;
        public readonly string? DataCompressionType;
        public readonly ImmutableArray<Outputs.FileSystemNfsExports> NfsExports;
        public readonly bool? ReadOnly;
        public readonly int? RecordSizeKiB;
        public readonly ImmutableArray<Outputs.FileSystemUserAndGroupQuotas> UserAndGroupQuotas;

        [OutputConstructor]
        private FileSystemRootVolumeConfiguration(
            bool? copyTagsToSnapshots,

            string? dataCompressionType,

            ImmutableArray<Outputs.FileSystemNfsExports> nfsExports,

            bool? readOnly,

            int? recordSizeKiB,

            ImmutableArray<Outputs.FileSystemUserAndGroupQuotas> userAndGroupQuotas)
        {
            CopyTagsToSnapshots = copyTagsToSnapshots;
            DataCompressionType = dataCompressionType;
            NfsExports = nfsExports;
            ReadOnly = readOnly;
            RecordSizeKiB = recordSizeKiB;
            UserAndGroupQuotas = userAndGroupQuotas;
        }
    }
}
