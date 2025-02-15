// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::LakeFormation::DataLakeSettings
 */
export function getDataLakeSettings(args: GetDataLakeSettingsArgs, opts?: pulumi.InvokeOptions): Promise<GetDataLakeSettingsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:lakeformation:getDataLakeSettings", {
        "id": args.id,
    }, opts);
}

export interface GetDataLakeSettingsArgs {
    id: string;
}

export interface GetDataLakeSettingsResult {
    readonly admins?: outputs.lakeformation.DataLakeSettingsAdmins;
    readonly allowExternalDataFiltering?: boolean;
    readonly allowFullTableExternalDataAccess?: boolean;
    readonly authorizedSessionTagValueList?: string[];
    readonly createDatabaseDefaultPermissions?: outputs.lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissions;
    readonly createTableDefaultPermissions?: outputs.lakeformation.DataLakeSettingsCreateTableDefaultPermissions;
    readonly externalDataFilteringAllowList?: outputs.lakeformation.DataLakeSettingsExternalDataFilteringAllowList;
    readonly id?: string;
    readonly mutationType?: string;
    readonly parameters?: any;
    readonly trustedResourceOwners?: string[];
}
/**
 * Resource Type definition for AWS::LakeFormation::DataLakeSettings
 */
export function getDataLakeSettingsOutput(args: GetDataLakeSettingsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataLakeSettingsResult> {
    return pulumi.output(args).apply((a: any) => getDataLakeSettings(a, opts))
}

export interface GetDataLakeSettingsOutputArgs {
    id: pulumi.Input<string>;
}
