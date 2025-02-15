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
    'GetStepResult',
    'AwaitableGetStepResult',
    'get_step',
    'get_step_output',
]

@pulumi.output_type
class GetStepResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        ID generated by service
        """
        return pulumi.get(self, "id")


class AwaitableGetStepResult(GetStepResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStepResult(
            id=self.id)


def get_step(id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStepResult:
    """
    Schema for AWS::EMR::Step


    :param str id: ID generated by service
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:emr:getStep', __args__, opts=opts, typ=GetStepResult).value

    return AwaitableGetStepResult(
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_step)
def get_step_output(id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStepResult]:
    """
    Schema for AWS::EMR::Step


    :param str id: ID generated by service
    """
    ...
