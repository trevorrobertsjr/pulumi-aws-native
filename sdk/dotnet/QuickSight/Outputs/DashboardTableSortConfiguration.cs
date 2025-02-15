// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardTableSortConfiguration
    {
        public readonly Outputs.DashboardPaginationConfiguration? PaginationConfiguration;
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> RowSort;

        [OutputConstructor]
        private DashboardTableSortConfiguration(
            Outputs.DashboardPaginationConfiguration? paginationConfiguration,

            ImmutableArray<Outputs.DashboardFieldSortOptions> rowSort)
        {
            PaginationConfiguration = paginationConfiguration;
            RowSort = rowSort;
        }
    }
}
