// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ChannelArgs } from "./channel";
export type Channel = import("./channel").Channel;
export const Channel: typeof import("./channel").Channel = null as any;
utilities.lazyLoad(exports, ["Channel"], () => require("./channel"));

export { GetChannelArgs, GetChannelResult, GetChannelOutputArgs } from "./getChannel";
export const getChannel: typeof import("./getChannel").getChannel = null as any;
export const getChannelOutput: typeof import("./getChannel").getChannelOutput = null as any;
utilities.lazyLoad(exports, ["getChannel","getChannelOutput"], () => require("./getChannel"));

export { GetPlaybackKeyPairArgs, GetPlaybackKeyPairResult, GetPlaybackKeyPairOutputArgs } from "./getPlaybackKeyPair";
export const getPlaybackKeyPair: typeof import("./getPlaybackKeyPair").getPlaybackKeyPair = null as any;
export const getPlaybackKeyPairOutput: typeof import("./getPlaybackKeyPair").getPlaybackKeyPairOutput = null as any;
utilities.lazyLoad(exports, ["getPlaybackKeyPair","getPlaybackKeyPairOutput"], () => require("./getPlaybackKeyPair"));

export { GetRecordingConfigurationArgs, GetRecordingConfigurationResult, GetRecordingConfigurationOutputArgs } from "./getRecordingConfiguration";
export const getRecordingConfiguration: typeof import("./getRecordingConfiguration").getRecordingConfiguration = null as any;
export const getRecordingConfigurationOutput: typeof import("./getRecordingConfiguration").getRecordingConfigurationOutput = null as any;
utilities.lazyLoad(exports, ["getRecordingConfiguration","getRecordingConfigurationOutput"], () => require("./getRecordingConfiguration"));

export { GetStreamKeyArgs, GetStreamKeyResult, GetStreamKeyOutputArgs } from "./getStreamKey";
export const getStreamKey: typeof import("./getStreamKey").getStreamKey = null as any;
export const getStreamKeyOutput: typeof import("./getStreamKey").getStreamKeyOutput = null as any;
utilities.lazyLoad(exports, ["getStreamKey","getStreamKeyOutput"], () => require("./getStreamKey"));

export { PlaybackKeyPairArgs } from "./playbackKeyPair";
export type PlaybackKeyPair = import("./playbackKeyPair").PlaybackKeyPair;
export const PlaybackKeyPair: typeof import("./playbackKeyPair").PlaybackKeyPair = null as any;
utilities.lazyLoad(exports, ["PlaybackKeyPair"], () => require("./playbackKeyPair"));

export { RecordingConfigurationArgs } from "./recordingConfiguration";
export type RecordingConfiguration = import("./recordingConfiguration").RecordingConfiguration;
export const RecordingConfiguration: typeof import("./recordingConfiguration").RecordingConfiguration = null as any;
utilities.lazyLoad(exports, ["RecordingConfiguration"], () => require("./recordingConfiguration"));

export { StreamKeyArgs } from "./streamKey";
export type StreamKey = import("./streamKey").StreamKey;
export const StreamKey: typeof import("./streamKey").StreamKey = null as any;
utilities.lazyLoad(exports, ["StreamKey"], () => require("./streamKey"));


// Export enums:
export * from "../types/enums/ivs";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:ivs:Channel":
                return new Channel(name, <any>undefined, { urn })
            case "aws-native:ivs:PlaybackKeyPair":
                return new PlaybackKeyPair(name, <any>undefined, { urn })
            case "aws-native:ivs:RecordingConfiguration":
                return new RecordingConfiguration(name, <any>undefined, { urn })
            case "aws-native:ivs:StreamKey":
                return new StreamKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "ivs", _module)
