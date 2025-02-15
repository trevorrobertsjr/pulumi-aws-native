// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A Kendra FAQ resource
 */
export function getFaq(args: GetFaqArgs, opts?: pulumi.InvokeOptions): Promise<GetFaqResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:kendra:getFaq", {
        "id": args.id,
        "indexId": args.indexId,
    }, opts);
}

export interface GetFaqArgs {
    id: string;
    /**
     * Index ID
     */
    indexId: string;
}

export interface GetFaqResult {
    readonly arn?: string;
    readonly id?: string;
    /**
     * Tags for labeling the FAQ
     */
    readonly tags?: outputs.kendra.FaqTag[];
}
/**
 * A Kendra FAQ resource
 */
export function getFaqOutput(args: GetFaqOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFaqResult> {
    return pulumi.output(args).apply((a: any) => getFaq(a, opts))
}

export interface GetFaqOutputArgs {
    id: pulumi.Input<string>;
    /**
     * Index ID
     */
    indexId: pulumi.Input<string>;
}
