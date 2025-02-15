// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Config::StoredQuery
 */
export function getStoredQuery(args: GetStoredQueryArgs, opts?: pulumi.InvokeOptions): Promise<GetStoredQueryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:configuration:getStoredQuery", {
        "queryName": args.queryName,
    }, opts);
}

export interface GetStoredQueryArgs {
    queryName: string;
}

export interface GetStoredQueryResult {
    readonly queryArn?: string;
    readonly queryDescription?: string;
    readonly queryExpression?: string;
    readonly queryId?: string;
    /**
     * The tags for the stored query.
     */
    readonly tags?: outputs.configuration.StoredQueryTag[];
}
/**
 * Resource Type definition for AWS::Config::StoredQuery
 */
export function getStoredQueryOutput(args: GetStoredQueryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStoredQueryResult> {
    return pulumi.output(args).apply((a: any) => getStoredQuery(a, opts))
}

export interface GetStoredQueryOutputArgs {
    queryName: pulumi.Input<string>;
}
