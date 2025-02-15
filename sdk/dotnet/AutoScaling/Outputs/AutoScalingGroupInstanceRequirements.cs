// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class AutoScalingGroupInstanceRequirements
    {
        public readonly Outputs.AutoScalingGroupAcceleratorCountRequest? AcceleratorCount;
        public readonly ImmutableArray<string> AcceleratorManufacturers;
        public readonly ImmutableArray<string> AcceleratorNames;
        public readonly Outputs.AutoScalingGroupAcceleratorTotalMemoryMiBRequest? AcceleratorTotalMemoryMiB;
        public readonly ImmutableArray<string> AcceleratorTypes;
        public readonly ImmutableArray<string> AllowedInstanceTypes;
        public readonly string? BareMetal;
        public readonly Outputs.AutoScalingGroupBaselineEbsBandwidthMbpsRequest? BaselineEbsBandwidthMbps;
        public readonly string? BurstablePerformance;
        public readonly ImmutableArray<string> CpuManufacturers;
        public readonly ImmutableArray<string> ExcludedInstanceTypes;
        public readonly ImmutableArray<string> InstanceGenerations;
        public readonly string? LocalStorage;
        public readonly ImmutableArray<string> LocalStorageTypes;
        public readonly Outputs.AutoScalingGroupMemoryGiBPerVCpuRequest? MemoryGiBPerVCpu;
        public readonly Outputs.AutoScalingGroupMemoryMiBRequest MemoryMiB;
        public readonly Outputs.AutoScalingGroupNetworkBandwidthGbpsRequest? NetworkBandwidthGbps;
        public readonly Outputs.AutoScalingGroupNetworkInterfaceCountRequest? NetworkInterfaceCount;
        public readonly int? OnDemandMaxPricePercentageOverLowestPrice;
        public readonly bool? RequireHibernateSupport;
        public readonly int? SpotMaxPricePercentageOverLowestPrice;
        public readonly Outputs.AutoScalingGroupTotalLocalStorageGbRequest? TotalLocalStorageGb;
        public readonly Outputs.AutoScalingGroupVCpuCountRequest VCpuCount;

        [OutputConstructor]
        private AutoScalingGroupInstanceRequirements(
            Outputs.AutoScalingGroupAcceleratorCountRequest? acceleratorCount,

            ImmutableArray<string> acceleratorManufacturers,

            ImmutableArray<string> acceleratorNames,

            Outputs.AutoScalingGroupAcceleratorTotalMemoryMiBRequest? acceleratorTotalMemoryMiB,

            ImmutableArray<string> acceleratorTypes,

            ImmutableArray<string> allowedInstanceTypes,

            string? bareMetal,

            Outputs.AutoScalingGroupBaselineEbsBandwidthMbpsRequest? baselineEbsBandwidthMbps,

            string? burstablePerformance,

            ImmutableArray<string> cpuManufacturers,

            ImmutableArray<string> excludedInstanceTypes,

            ImmutableArray<string> instanceGenerations,

            string? localStorage,

            ImmutableArray<string> localStorageTypes,

            Outputs.AutoScalingGroupMemoryGiBPerVCpuRequest? memoryGiBPerVCpu,

            Outputs.AutoScalingGroupMemoryMiBRequest memoryMiB,

            Outputs.AutoScalingGroupNetworkBandwidthGbpsRequest? networkBandwidthGbps,

            Outputs.AutoScalingGroupNetworkInterfaceCountRequest? networkInterfaceCount,

            int? onDemandMaxPricePercentageOverLowestPrice,

            bool? requireHibernateSupport,

            int? spotMaxPricePercentageOverLowestPrice,

            Outputs.AutoScalingGroupTotalLocalStorageGbRequest? totalLocalStorageGb,

            Outputs.AutoScalingGroupVCpuCountRequest vCpuCount)
        {
            AcceleratorCount = acceleratorCount;
            AcceleratorManufacturers = acceleratorManufacturers;
            AcceleratorNames = acceleratorNames;
            AcceleratorTotalMemoryMiB = acceleratorTotalMemoryMiB;
            AcceleratorTypes = acceleratorTypes;
            AllowedInstanceTypes = allowedInstanceTypes;
            BareMetal = bareMetal;
            BaselineEbsBandwidthMbps = baselineEbsBandwidthMbps;
            BurstablePerformance = burstablePerformance;
            CpuManufacturers = cpuManufacturers;
            ExcludedInstanceTypes = excludedInstanceTypes;
            InstanceGenerations = instanceGenerations;
            LocalStorage = localStorage;
            LocalStorageTypes = localStorageTypes;
            MemoryGiBPerVCpu = memoryGiBPerVCpu;
            MemoryMiB = memoryMiB;
            NetworkBandwidthGbps = networkBandwidthGbps;
            NetworkInterfaceCount = networkInterfaceCount;
            OnDemandMaxPricePercentageOverLowestPrice = onDemandMaxPricePercentageOverLowestPrice;
            RequireHibernateSupport = requireHibernateSupport;
            SpotMaxPricePercentageOverLowestPrice = spotMaxPricePercentageOverLowestPrice;
            TotalLocalStorageGb = totalLocalStorageGb;
            VCpuCount = vCpuCount;
        }
    }
}
