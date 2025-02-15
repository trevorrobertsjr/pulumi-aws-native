// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WorkSpaces
{
    public static class GetWorkspace
    {
        /// <summary>
        /// Resource Type definition for AWS::WorkSpaces::Workspace
        /// </summary>
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("aws-native:workspaces:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::WorkSpaces::Workspace
        /// </summary>
        public static Output<GetWorkspaceResult> Invoke(GetWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceResult>("aws-native:workspaces:getWorkspace", args ?? new GetWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
        public static new GetWorkspaceArgs Empty => new GetWorkspaceArgs();
    }

    public sealed class GetWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWorkspaceInvokeArgs()
        {
        }
        public static new GetWorkspaceInvokeArgs Empty => new GetWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        public readonly string? BundleId;
        public readonly string? DirectoryId;
        public readonly string? Id;
        public readonly bool? RootVolumeEncryptionEnabled;
        public readonly ImmutableArray<Outputs.WorkspaceTag> Tags;
        public readonly bool? UserVolumeEncryptionEnabled;
        public readonly string? VolumeEncryptionKey;
        public readonly Outputs.WorkspaceProperties? WorkspaceProperties;

        [OutputConstructor]
        private GetWorkspaceResult(
            string? bundleId,

            string? directoryId,

            string? id,

            bool? rootVolumeEncryptionEnabled,

            ImmutableArray<Outputs.WorkspaceTag> tags,

            bool? userVolumeEncryptionEnabled,

            string? volumeEncryptionKey,

            Outputs.WorkspaceProperties? workspaceProperties)
        {
            BundleId = bundleId;
            DirectoryId = directoryId;
            Id = id;
            RootVolumeEncryptionEnabled = rootVolumeEncryptionEnabled;
            Tags = tags;
            UserVolumeEncryptionEnabled = userVolumeEncryptionEnabled;
            VolumeEncryptionKey = volumeEncryptionKey;
            WorkspaceProperties = workspaceProperties;
        }
    }
}
