# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDbClusterParameterGroupResult',
    'AwaitableGetDbClusterParameterGroupResult',
    'get_db_cluster_parameter_group',
    'get_db_cluster_parameter_group_output',
]

@pulumi.output_type
class GetDbClusterParameterGroupResult:
    def __init__(__self__, id=None, parameters=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if parameters and not isinstance(parameters, dict):
            raise TypeError("Expected argument 'parameters' to be a dict")
        pulumi.set(__self__, "parameters", parameters)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Any]:
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DbClusterParameterGroupTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetDbClusterParameterGroupResult(GetDbClusterParameterGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDbClusterParameterGroupResult(
            id=self.id,
            parameters=self.parameters,
            tags=self.tags)


def get_db_cluster_parameter_group(id: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDbClusterParameterGroupResult:
    """
    Resource Type definition for AWS::Neptune::DBClusterParameterGroup
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:neptune:getDbClusterParameterGroup', __args__, opts=opts, typ=GetDbClusterParameterGroupResult).value

    return AwaitableGetDbClusterParameterGroupResult(
        id=pulumi.get(__ret__, 'id'),
        parameters=pulumi.get(__ret__, 'parameters'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_db_cluster_parameter_group)
def get_db_cluster_parameter_group_output(id: Optional[pulumi.Input[str]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDbClusterParameterGroupResult]:
    """
    Resource Type definition for AWS::Neptune::DBClusterParameterGroup
    """
    ...
