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
    'GetWaitConditionHandleResult',
    'AwaitableGetWaitConditionHandleResult',
    'get_wait_condition_handle',
    'get_wait_condition_handle_output',
]

@pulumi.output_type
class GetWaitConditionHandleResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetWaitConditionHandleResult(GetWaitConditionHandleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWaitConditionHandleResult(
            id=self.id)


def get_wait_condition_handle(id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWaitConditionHandleResult:
    """
    Resource Type definition for AWS::CloudFormation::WaitConditionHandle
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudformation:getWaitConditionHandle', __args__, opts=opts, typ=GetWaitConditionHandleResult).value

    return AwaitableGetWaitConditionHandleResult(
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_wait_condition_handle)
def get_wait_condition_handle_output(id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWaitConditionHandleResult]:
    """
    Resource Type definition for AWS::CloudFormation::WaitConditionHandle
    """
    ...
