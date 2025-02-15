// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Omics::ReferenceStore Resource Type
 */
export function getReferenceStore(args: GetReferenceStoreArgs, opts?: pulumi.InvokeOptions): Promise<GetReferenceStoreResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:omics:getReferenceStore", {
        "referenceStoreId": args.referenceStoreId,
    }, opts);
}

export interface GetReferenceStoreArgs {
    referenceStoreId: string;
}

export interface GetReferenceStoreResult {
    /**
     * The store's ARN.
     */
    readonly arn?: string;
    /**
     * When the store was created.
     */
    readonly creationTime?: string;
    readonly referenceStoreId?: string;
}
/**
 * Definition of AWS::Omics::ReferenceStore Resource Type
 */
export function getReferenceStoreOutput(args: GetReferenceStoreOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReferenceStoreResult> {
    return pulumi.output(args).apply((a: any) => getReferenceStore(a, opts))
}

export interface GetReferenceStoreOutputArgs {
    referenceStoreId: pulumi.Input<string>;
}
