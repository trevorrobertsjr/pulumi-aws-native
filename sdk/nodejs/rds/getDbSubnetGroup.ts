// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::RDS::DBSubnetGroup resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region.
 */
export function getDbSubnetGroup(args: GetDbSubnetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDbSubnetGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:rds:getDbSubnetGroup", {
        "dbSubnetGroupName": args.dbSubnetGroupName,
    }, opts);
}

export interface GetDbSubnetGroupArgs {
    dbSubnetGroupName: string;
}

export interface GetDbSubnetGroupResult {
    readonly dbSubnetGroupDescription?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.rds.DbSubnetGroupTag[];
}
/**
 * The AWS::RDS::DBSubnetGroup resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region.
 */
export function getDbSubnetGroupOutput(args: GetDbSubnetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDbSubnetGroupResult> {
    return pulumi.output(args).apply((a: any) => getDbSubnetGroup(a, opts))
}

export interface GetDbSubnetGroupOutputArgs {
    dbSubnetGroupName: pulumi.Input<string>;
}
