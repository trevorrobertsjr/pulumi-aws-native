// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAF::XssMatchSet
 */
export function getXssMatchSet(args: GetXssMatchSetArgs, opts?: pulumi.InvokeOptions): Promise<GetXssMatchSetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:waf:getXssMatchSet", {
        "id": args.id,
    }, opts);
}

export interface GetXssMatchSetArgs {
    id: string;
}

export interface GetXssMatchSetResult {
    readonly id?: string;
    readonly xssMatchTuples?: outputs.waf.XssMatchSetXssMatchTuple[];
}
/**
 * Resource Type definition for AWS::WAF::XssMatchSet
 */
export function getXssMatchSetOutput(args: GetXssMatchSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetXssMatchSetResult> {
    return pulumi.output(args).apply((a: any) => getXssMatchSet(a, opts))
}

export interface GetXssMatchSetOutputArgs {
    id: pulumi.Input<string>;
}
