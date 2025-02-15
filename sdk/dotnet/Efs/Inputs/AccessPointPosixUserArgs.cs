// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Efs.Inputs
{

    public sealed class AccessPointPosixUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The POSIX group ID used for all file system operations using this access point.
        /// </summary>
        [Input("gid", required: true)]
        public Input<string> Gid { get; set; } = null!;

        [Input("secondaryGids")]
        private InputList<string>? _secondaryGids;

        /// <summary>
        /// Secondary POSIX group IDs used for all file system operations using this access point.
        /// </summary>
        public InputList<string> SecondaryGids
        {
            get => _secondaryGids ?? (_secondaryGids = new InputList<string>());
            set => _secondaryGids = value;
        }

        /// <summary>
        /// The POSIX user ID used for all file system operations using this access point.
        /// </summary>
        [Input("uid", required: true)]
        public Input<string> Uid { get; set; } = null!;

        public AccessPointPosixUserArgs()
        {
        }
        public static new AccessPointPosixUserArgs Empty => new AccessPointPosixUserArgs();
    }
}
