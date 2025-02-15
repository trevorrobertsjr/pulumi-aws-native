// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Efs.Inputs
{

    public sealed class AccessPointRootDirectoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Specifies the POSIX IDs and permissions to apply to the access point's RootDirectory. If the RootDirectory&gt;Path specified does not exist, EFS creates the root directory using the CreationInfo settings when a client connects to an access point. When specifying the CreationInfo, you must provide values for all properties.   If you do not provide CreationInfo and the specified RootDirectory&gt;Path does not exist, attempts to mount the file system using the access point will fail. 
        /// </summary>
        [Input("creationInfo")]
        public Input<Inputs.AccessPointCreationInfoArgs>? CreationInfo { get; set; }

        /// <summary>
        /// Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the CreationInfo.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public AccessPointRootDirectoryArgs()
        {
        }
        public static new AccessPointRootDirectoryArgs Empty => new AccessPointRootDirectoryArgs();
    }
}
