# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetVoiceChannelResult',
    'AwaitableGetVoiceChannelResult',
    'get_voice_channel',
    'get_voice_channel_output',
]

@pulumi.output_type
class GetVoiceChannelResult:
    def __init__(__self__, enabled=None, id=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetVoiceChannelResult(GetVoiceChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVoiceChannelResult(
            enabled=self.enabled,
            id=self.id)


def get_voice_channel(id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVoiceChannelResult:
    """
    Resource Type definition for AWS::Pinpoint::VoiceChannel
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:pinpoint:getVoiceChannel', __args__, opts=opts, typ=GetVoiceChannelResult).value

    return AwaitableGetVoiceChannelResult(
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_voice_channel)
def get_voice_channel_output(id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVoiceChannelResult]:
    """
    Resource Type definition for AWS::Pinpoint::VoiceChannel
    """
    ...
