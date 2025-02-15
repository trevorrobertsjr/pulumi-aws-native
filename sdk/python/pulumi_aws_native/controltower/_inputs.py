# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'EnabledControlParameterArgs',
    'LandingZoneTagArgs',
]

@pulumi.input_type
class EnabledControlParameterArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[Union[Sequence[pulumi.Input[Union[str, float, Any, bool]]], str, float, Any, bool]]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[Union[Sequence[pulumi.Input[Union[str, float, Any, bool]]], str, float, Any, bool]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[Union[Sequence[pulumi.Input[Union[str, float, Any, bool]]], str, float, Any, bool]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class LandingZoneTagArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


