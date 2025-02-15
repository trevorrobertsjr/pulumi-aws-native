// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GatewayRouteArgs } from "./gatewayRoute";
export type GatewayRoute = import("./gatewayRoute").GatewayRoute;
export const GatewayRoute: typeof import("./gatewayRoute").GatewayRoute = null as any;
utilities.lazyLoad(exports, ["GatewayRoute"], () => require("./gatewayRoute"));

export { GetGatewayRouteArgs, GetGatewayRouteResult, GetGatewayRouteOutputArgs } from "./getGatewayRoute";
export const getGatewayRoute: typeof import("./getGatewayRoute").getGatewayRoute = null as any;
export const getGatewayRouteOutput: typeof import("./getGatewayRoute").getGatewayRouteOutput = null as any;
utilities.lazyLoad(exports, ["getGatewayRoute","getGatewayRouteOutput"], () => require("./getGatewayRoute"));

export { GetMeshArgs, GetMeshResult, GetMeshOutputArgs } from "./getMesh";
export const getMesh: typeof import("./getMesh").getMesh = null as any;
export const getMeshOutput: typeof import("./getMesh").getMeshOutput = null as any;
utilities.lazyLoad(exports, ["getMesh","getMeshOutput"], () => require("./getMesh"));

export { GetRouteArgs, GetRouteResult, GetRouteOutputArgs } from "./getRoute";
export const getRoute: typeof import("./getRoute").getRoute = null as any;
export const getRouteOutput: typeof import("./getRoute").getRouteOutput = null as any;
utilities.lazyLoad(exports, ["getRoute","getRouteOutput"], () => require("./getRoute"));

export { GetVirtualGatewayArgs, GetVirtualGatewayResult, GetVirtualGatewayOutputArgs } from "./getVirtualGateway";
export const getVirtualGateway: typeof import("./getVirtualGateway").getVirtualGateway = null as any;
export const getVirtualGatewayOutput: typeof import("./getVirtualGateway").getVirtualGatewayOutput = null as any;
utilities.lazyLoad(exports, ["getVirtualGateway","getVirtualGatewayOutput"], () => require("./getVirtualGateway"));

export { GetVirtualNodeArgs, GetVirtualNodeResult, GetVirtualNodeOutputArgs } from "./getVirtualNode";
export const getVirtualNode: typeof import("./getVirtualNode").getVirtualNode = null as any;
export const getVirtualNodeOutput: typeof import("./getVirtualNode").getVirtualNodeOutput = null as any;
utilities.lazyLoad(exports, ["getVirtualNode","getVirtualNodeOutput"], () => require("./getVirtualNode"));

export { GetVirtualRouterArgs, GetVirtualRouterResult, GetVirtualRouterOutputArgs } from "./getVirtualRouter";
export const getVirtualRouter: typeof import("./getVirtualRouter").getVirtualRouter = null as any;
export const getVirtualRouterOutput: typeof import("./getVirtualRouter").getVirtualRouterOutput = null as any;
utilities.lazyLoad(exports, ["getVirtualRouter","getVirtualRouterOutput"], () => require("./getVirtualRouter"));

export { GetVirtualServiceArgs, GetVirtualServiceResult, GetVirtualServiceOutputArgs } from "./getVirtualService";
export const getVirtualService: typeof import("./getVirtualService").getVirtualService = null as any;
export const getVirtualServiceOutput: typeof import("./getVirtualService").getVirtualServiceOutput = null as any;
utilities.lazyLoad(exports, ["getVirtualService","getVirtualServiceOutput"], () => require("./getVirtualService"));

export { MeshArgs } from "./mesh";
export type Mesh = import("./mesh").Mesh;
export const Mesh: typeof import("./mesh").Mesh = null as any;
utilities.lazyLoad(exports, ["Mesh"], () => require("./mesh"));

export { RouteArgs } from "./route";
export type Route = import("./route").Route;
export const Route: typeof import("./route").Route = null as any;
utilities.lazyLoad(exports, ["Route"], () => require("./route"));

export { VirtualGatewayArgs } from "./virtualGateway";
export type VirtualGateway = import("./virtualGateway").VirtualGateway;
export const VirtualGateway: typeof import("./virtualGateway").VirtualGateway = null as any;
utilities.lazyLoad(exports, ["VirtualGateway"], () => require("./virtualGateway"));

export { VirtualNodeArgs } from "./virtualNode";
export type VirtualNode = import("./virtualNode").VirtualNode;
export const VirtualNode: typeof import("./virtualNode").VirtualNode = null as any;
utilities.lazyLoad(exports, ["VirtualNode"], () => require("./virtualNode"));

export { VirtualRouterArgs } from "./virtualRouter";
export type VirtualRouter = import("./virtualRouter").VirtualRouter;
export const VirtualRouter: typeof import("./virtualRouter").VirtualRouter = null as any;
utilities.lazyLoad(exports, ["VirtualRouter"], () => require("./virtualRouter"));

export { VirtualServiceArgs } from "./virtualService";
export type VirtualService = import("./virtualService").VirtualService;
export const VirtualService: typeof import("./virtualService").VirtualService = null as any;
utilities.lazyLoad(exports, ["VirtualService"], () => require("./virtualService"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:appmesh:GatewayRoute":
                return new GatewayRoute(name, <any>undefined, { urn })
            case "aws-native:appmesh:Mesh":
                return new Mesh(name, <any>undefined, { urn })
            case "aws-native:appmesh:Route":
                return new Route(name, <any>undefined, { urn })
            case "aws-native:appmesh:VirtualGateway":
                return new VirtualGateway(name, <any>undefined, { urn })
            case "aws-native:appmesh:VirtualNode":
                return new VirtualNode(name, <any>undefined, { urn })
            case "aws-native:appmesh:VirtualRouter":
                return new VirtualRouter(name, <any>undefined, { urn })
            case "aws-native:appmesh:VirtualService":
                return new VirtualService(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "appmesh", _module)
