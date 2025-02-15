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
    public sealed class FileSystemWindowsConfiguration
    {
        public readonly string? ActiveDirectoryId;
        public readonly ImmutableArray<string> Aliases;
        public readonly Outputs.FileSystemAuditLogConfiguration? AuditLogConfiguration;
        public readonly int? AutomaticBackupRetentionDays;
        public readonly bool? CopyTagsToBackups;
        public readonly string? DailyAutomaticBackupStartTime;
        public readonly string? DeploymentType;
        public readonly Outputs.FileSystemDiskIopsConfiguration? DiskIopsConfiguration;
        public readonly string? PreferredSubnetId;
        public readonly Outputs.FileSystemSelfManagedActiveDirectoryConfiguration? SelfManagedActiveDirectoryConfiguration;
        public readonly int ThroughputCapacity;
        public readonly string? WeeklyMaintenanceStartTime;

        [OutputConstructor]
        private FileSystemWindowsConfiguration(
            string? activeDirectoryId,

            ImmutableArray<string> aliases,

            Outputs.FileSystemAuditLogConfiguration? auditLogConfiguration,

            int? automaticBackupRetentionDays,

            bool? copyTagsToBackups,

            string? dailyAutomaticBackupStartTime,

            string? deploymentType,

            Outputs.FileSystemDiskIopsConfiguration? diskIopsConfiguration,

            string? preferredSubnetId,

            Outputs.FileSystemSelfManagedActiveDirectoryConfiguration? selfManagedActiveDirectoryConfiguration,

            int throughputCapacity,

            string? weeklyMaintenanceStartTime)
        {
            ActiveDirectoryId = activeDirectoryId;
            Aliases = aliases;
            AuditLogConfiguration = auditLogConfiguration;
            AutomaticBackupRetentionDays = automaticBackupRetentionDays;
            CopyTagsToBackups = copyTagsToBackups;
            DailyAutomaticBackupStartTime = dailyAutomaticBackupStartTime;
            DeploymentType = deploymentType;
            DiskIopsConfiguration = diskIopsConfiguration;
            PreferredSubnetId = preferredSubnetId;
            SelfManagedActiveDirectoryConfiguration = selfManagedActiveDirectoryConfiguration;
            ThroughputCapacity = throughputCapacity;
            WeeklyMaintenanceStartTime = weeklyMaintenanceStartTime;
        }
    }
}
