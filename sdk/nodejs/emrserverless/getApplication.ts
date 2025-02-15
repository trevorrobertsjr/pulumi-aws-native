// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::EMRServerless::Application Type
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:emrserverless:getApplication", {
        "applicationId": args.applicationId,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The ID of the EMR Serverless Application.
     */
    applicationId: string;
}

export interface GetApplicationResult {
    /**
     * The ID of the EMR Serverless Application.
     */
    readonly applicationId?: string;
    readonly architecture?: enums.emrserverless.ApplicationArchitecture;
    /**
     * The Amazon Resource Name (ARN) of the EMR Serverless Application.
     */
    readonly arn?: string;
    /**
     * Configuration for Auto Start of Application.
     */
    readonly autoStartConfiguration?: outputs.emrserverless.ApplicationAutoStartConfiguration;
    /**
     * Configuration for Auto Stop of Application.
     */
    readonly autoStopConfiguration?: outputs.emrserverless.ApplicationAutoStopConfiguration;
    readonly imageConfiguration?: outputs.emrserverless.ApplicationImageConfigurationInput;
    /**
     * Initial capacity initialized when an Application is started.
     */
    readonly initialCapacity?: outputs.emrserverless.ApplicationInitialCapacityConfigKeyValuePair[];
    /**
     * Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.
     */
    readonly maximumCapacity?: outputs.emrserverless.ApplicationMaximumAllowedResources;
    readonly monitoringConfiguration?: outputs.emrserverless.ApplicationMonitoringConfiguration;
    /**
     * Network Configuration for customer VPC connectivity.
     */
    readonly networkConfiguration?: outputs.emrserverless.ApplicationNetworkConfiguration;
    /**
     * EMR release label.
     */
    readonly releaseLabel?: string;
    readonly runtimeConfiguration?: outputs.emrserverless.ApplicationConfigurationObject[];
    /**
     * Tag map with key and value
     */
    readonly tags?: outputs.emrserverless.ApplicationTag[];
    /**
     * The key-value pairs that specify worker type to WorkerTypeSpecificationInput. This parameter must contain all valid worker types for a Spark or Hive application. Valid worker types include Driver and Executor for Spark applications and HiveDriver and TezTask for Hive applications. You can either set image details in this parameter for each worker type, or in imageConfiguration for all worker types.
     */
    readonly workerTypeSpecifications?: outputs.emrserverless.ApplicationWorkerTypeSpecificationInputMap;
}
/**
 * Resource schema for AWS::EMRServerless::Application Type
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply((a: any) => getApplication(a, opts))
}

export interface GetApplicationOutputArgs {
    /**
     * The ID of the EMR Serverless Application.
     */
    applicationId: pulumi.Input<string>;
}
