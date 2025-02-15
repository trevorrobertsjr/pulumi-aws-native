// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::FeatureGroup
 */
export function getFeatureGroup(args: GetFeatureGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getFeatureGroup", {
        "featureGroupName": args.featureGroupName,
    }, opts);
}

export interface GetFeatureGroupArgs {
    /**
     * The Name of the FeatureGroup.
     */
    featureGroupName: string;
}

export interface GetFeatureGroupResult {
    /**
     * A timestamp of FeatureGroup creation time.
     */
    readonly creationTime?: string;
    /**
     * An Array of Feature Definition
     */
    readonly featureDefinitions?: outputs.sagemaker.FeatureGroupFeatureDefinition[];
    /**
     * The status of the feature group.
     */
    readonly featureGroupStatus?: string;
}
/**
 * Resource Type definition for AWS::SageMaker::FeatureGroup
 */
export function getFeatureGroupOutput(args: GetFeatureGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureGroupResult> {
    return pulumi.output(args).apply((a: any) => getFeatureGroup(a, opts))
}

export interface GetFeatureGroupOutputArgs {
    /**
     * The Name of the FeatureGroup.
     */
    featureGroupName: pulumi.Input<string>;
}
