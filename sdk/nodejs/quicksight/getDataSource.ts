// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::DataSource Resource Type.
 */
export function getDataSource(args: GetDataSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetDataSourceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:quicksight:getDataSource", {
        "awsAccountId": args.awsAccountId,
        "dataSourceId": args.dataSourceId,
    }, opts);
}

export interface GetDataSourceArgs {
    awsAccountId: string;
    dataSourceId: string;
}

export interface GetDataSourceResult {
    /**
     * <p>A set of alternate data source parameters that you want to share for the credentials
     *             stored with this data source. The credentials are applied in tandem with the data source
     *             parameters when you copy a data source by using a create or update request. The API
     *             operation compares the <code>DataSourceParameters</code> structure that's in the request
     *             with the structures in the <code>AlternateDataSourceParameters</code> allow list. If the
     *             structures are an exact match, the request is allowed to use the credentials from this
     *             existing data source. If the <code>AlternateDataSourceParameters</code> list is null,
     *             the <code>Credentials</code> originally used with this <code>DataSourceParameters</code>
     *             are automatically allowed.</p>
     */
    readonly alternateDataSourceParameters?: outputs.quicksight.DataSourceParameters[];
    /**
     * <p>The Amazon Resource Name (ARN) of the data source.</p>
     */
    readonly arn?: string;
    /**
     * <p>The time that this data source was created.</p>
     */
    readonly createdTime?: string;
    readonly dataSourceParameters?: outputs.quicksight.DataSourceParameters;
    readonly errorInfo?: outputs.quicksight.DataSourceErrorInfo;
    /**
     * <p>The last time that this data source was updated.</p>
     */
    readonly lastUpdatedTime?: string;
    /**
     * <p>A display name for the data source.</p>
     */
    readonly name?: string;
    /**
     * <p>A list of resource permissions on the data source.</p>
     */
    readonly permissions?: outputs.quicksight.DataSourceResourcePermission[];
    readonly sslProperties?: outputs.quicksight.DataSourceSslProperties;
    readonly status?: enums.quicksight.DataSourceResourceStatus;
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.</p>
     */
    readonly tags?: outputs.quicksight.DataSourceTag[];
    readonly vpcConnectionProperties?: outputs.quicksight.DataSourceVpcConnectionProperties;
}
/**
 * Definition of the AWS::QuickSight::DataSource Resource Type.
 */
export function getDataSourceOutput(args: GetDataSourceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataSourceResult> {
    return pulumi.output(args).apply((a: any) => getDataSource(a, opts))
}

export interface GetDataSourceOutputArgs {
    awsAccountId: pulumi.Input<string>;
    dataSourceId: pulumi.Input<string>;
}
