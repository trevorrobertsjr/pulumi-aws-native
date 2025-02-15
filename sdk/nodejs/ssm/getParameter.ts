// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SSM::Parameter
 */
export function getParameter(args: GetParameterArgs, opts?: pulumi.InvokeOptions): Promise<GetParameterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ssm:getParameter", {
        "name": args.name,
    }, opts);
}

export interface GetParameterArgs {
    /**
     * The name of the parameter.
     */
    name: string;
}

export interface GetParameterResult {
    /**
     * The corresponding DataType of the parameter.
     */
    readonly dataType?: enums.ssm.ParameterDataType;
    /**
     * The type of the parameter.
     */
    readonly type?: enums.ssm.ParameterType;
    /**
     * The value associated with the parameter.
     */
    readonly value?: string;
}
/**
 * Resource Type definition for AWS::SSM::Parameter
 */
export function getParameterOutput(args: GetParameterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetParameterResult> {
    return pulumi.output(args).apply((a: any) => getParameter(a, opts))
}

export interface GetParameterOutputArgs {
    /**
     * The name of the parameter.
     */
    name: pulumi.Input<string>;
}
