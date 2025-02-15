// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A replication configuration that you later provide to configure and start a AWS DMS Serverless replication
 */
export function getReplicationConfig(args: GetReplicationConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationConfigResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:dms:getReplicationConfig", {
        "replicationConfigArn": args.replicationConfigArn,
    }, opts);
}

export interface GetReplicationConfigArgs {
    /**
     * The Amazon Resource Name (ARN) of the Replication Config
     */
    replicationConfigArn: string;
}

export interface GetReplicationConfigResult {
    readonly computeConfig?: outputs.dms.ReplicationConfigComputeConfig;
    /**
     * The Amazon Resource Name (ARN) of the Replication Config
     */
    readonly replicationConfigArn?: string;
    /**
     * A unique identifier of replication configuration
     */
    readonly replicationConfigIdentifier?: string;
    /**
     * JSON settings for Servereless replications that are provisioned using this replication configuration
     */
    readonly replicationSettings?: any;
    /**
     * The type of AWS DMS Serverless replication to provision using this replication configuration
     */
    readonly replicationType?: enums.dms.ReplicationConfigReplicationType;
    /**
     * The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration
     */
    readonly sourceEndpointArn?: string;
    /**
     * JSON settings for specifying supplemental data
     */
    readonly supplementalSettings?: any;
    /**
     * JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration
     */
    readonly tableMappings?: any;
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>
     */
    readonly tags?: outputs.dms.ReplicationConfigTag[];
    /**
     * The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration
     */
    readonly targetEndpointArn?: string;
}
/**
 * A replication configuration that you later provide to configure and start a AWS DMS Serverless replication
 */
export function getReplicationConfigOutput(args: GetReplicationConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplicationConfigResult> {
    return pulumi.output(args).apply((a: any) => getReplicationConfig(a, opts))
}

export interface GetReplicationConfigOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the Replication Config
     */
    replicationConfigArn: pulumi.Input<string>;
}
