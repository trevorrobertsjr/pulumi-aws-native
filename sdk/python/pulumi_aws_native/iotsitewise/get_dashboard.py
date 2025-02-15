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
    'GetDashboardResult',
    'AwaitableGetDashboardResult',
    'get_dashboard',
    'get_dashboard_output',
]

@pulumi.output_type
class GetDashboardResult:
    def __init__(__self__, dashboard_arn=None, dashboard_definition=None, dashboard_description=None, dashboard_id=None, dashboard_name=None, tags=None):
        if dashboard_arn and not isinstance(dashboard_arn, str):
            raise TypeError("Expected argument 'dashboard_arn' to be a str")
        pulumi.set(__self__, "dashboard_arn", dashboard_arn)
        if dashboard_definition and not isinstance(dashboard_definition, str):
            raise TypeError("Expected argument 'dashboard_definition' to be a str")
        pulumi.set(__self__, "dashboard_definition", dashboard_definition)
        if dashboard_description and not isinstance(dashboard_description, str):
            raise TypeError("Expected argument 'dashboard_description' to be a str")
        pulumi.set(__self__, "dashboard_description", dashboard_description)
        if dashboard_id and not isinstance(dashboard_id, str):
            raise TypeError("Expected argument 'dashboard_id' to be a str")
        pulumi.set(__self__, "dashboard_id", dashboard_id)
        if dashboard_name and not isinstance(dashboard_name, str):
            raise TypeError("Expected argument 'dashboard_name' to be a str")
        pulumi.set(__self__, "dashboard_name", dashboard_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dashboardArn")
    def dashboard_arn(self) -> Optional[str]:
        """
        The ARN of the dashboard.
        """
        return pulumi.get(self, "dashboard_arn")

    @property
    @pulumi.getter(name="dashboardDefinition")
    def dashboard_definition(self) -> Optional[str]:
        """
        The dashboard definition specified in a JSON literal.
        """
        return pulumi.get(self, "dashboard_definition")

    @property
    @pulumi.getter(name="dashboardDescription")
    def dashboard_description(self) -> Optional[str]:
        """
        A description for the dashboard.
        """
        return pulumi.get(self, "dashboard_description")

    @property
    @pulumi.getter(name="dashboardId")
    def dashboard_id(self) -> Optional[str]:
        """
        The ID of the dashboard.
        """
        return pulumi.get(self, "dashboard_id")

    @property
    @pulumi.getter(name="dashboardName")
    def dashboard_name(self) -> Optional[str]:
        """
        A friendly name for the dashboard.
        """
        return pulumi.get(self, "dashboard_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DashboardTag']]:
        """
        A list of key-value pairs that contain metadata for the dashboard.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDashboardResult(GetDashboardResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDashboardResult(
            dashboard_arn=self.dashboard_arn,
            dashboard_definition=self.dashboard_definition,
            dashboard_description=self.dashboard_description,
            dashboard_id=self.dashboard_id,
            dashboard_name=self.dashboard_name,
            tags=self.tags)


def get_dashboard(dashboard_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDashboardResult:
    """
    Resource schema for AWS::IoTSiteWise::Dashboard


    :param str dashboard_id: The ID of the dashboard.
    """
    __args__ = dict()
    __args__['dashboardId'] = dashboard_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iotsitewise:getDashboard', __args__, opts=opts, typ=GetDashboardResult).value

    return AwaitableGetDashboardResult(
        dashboard_arn=pulumi.get(__ret__, 'dashboard_arn'),
        dashboard_definition=pulumi.get(__ret__, 'dashboard_definition'),
        dashboard_description=pulumi.get(__ret__, 'dashboard_description'),
        dashboard_id=pulumi.get(__ret__, 'dashboard_id'),
        dashboard_name=pulumi.get(__ret__, 'dashboard_name'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_dashboard)
def get_dashboard_output(dashboard_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDashboardResult]:
    """
    Resource schema for AWS::IoTSiteWise::Dashboard


    :param str dashboard_id: The ID of the dashboard.
    """
    ...
