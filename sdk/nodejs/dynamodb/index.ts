// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetGlobalTableArgs, GetGlobalTableResult, GetGlobalTableOutputArgs } from "./getGlobalTable";
export const getGlobalTable: typeof import("./getGlobalTable").getGlobalTable = null as any;
export const getGlobalTableOutput: typeof import("./getGlobalTable").getGlobalTableOutput = null as any;
utilities.lazyLoad(exports, ["getGlobalTable","getGlobalTableOutput"], () => require("./getGlobalTable"));

export { GetTableArgs, GetTableResult, GetTableOutputArgs } from "./getTable";
export const getTable: typeof import("./getTable").getTable = null as any;
export const getTableOutput: typeof import("./getTable").getTableOutput = null as any;
utilities.lazyLoad(exports, ["getTable","getTableOutput"], () => require("./getTable"));

export { GlobalTableArgs } from "./globalTable";
export type GlobalTable = import("./globalTable").GlobalTable;
export const GlobalTable: typeof import("./globalTable").GlobalTable = null as any;
utilities.lazyLoad(exports, ["GlobalTable"], () => require("./globalTable"));

export { TableArgs } from "./table";
export type Table = import("./table").Table;
export const Table: typeof import("./table").Table = null as any;
utilities.lazyLoad(exports, ["Table"], () => require("./table"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:dynamodb:GlobalTable":
                return new GlobalTable(name, <any>undefined, { urn })
            case "aws-native:dynamodb:Table":
                return new Table(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "dynamodb", _module)
