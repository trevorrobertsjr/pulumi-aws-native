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
    'GetMaintenanceWindowTaskResult',
    'AwaitableGetMaintenanceWindowTaskResult',
    'get_maintenance_window_task',
    'get_maintenance_window_task_output',
]

@pulumi.output_type
class GetMaintenanceWindowTaskResult:
    def __init__(__self__, cutoff_behavior=None, description=None, id=None, logging_info=None, max_concurrency=None, max_errors=None, name=None, priority=None, service_role_arn=None, targets=None, task_arn=None, task_invocation_parameters=None, task_parameters=None):
        if cutoff_behavior and not isinstance(cutoff_behavior, str):
            raise TypeError("Expected argument 'cutoff_behavior' to be a str")
        pulumi.set(__self__, "cutoff_behavior", cutoff_behavior)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if logging_info and not isinstance(logging_info, dict):
            raise TypeError("Expected argument 'logging_info' to be a dict")
        pulumi.set(__self__, "logging_info", logging_info)
        if max_concurrency and not isinstance(max_concurrency, str):
            raise TypeError("Expected argument 'max_concurrency' to be a str")
        pulumi.set(__self__, "max_concurrency", max_concurrency)
        if max_errors and not isinstance(max_errors, str):
            raise TypeError("Expected argument 'max_errors' to be a str")
        pulumi.set(__self__, "max_errors", max_errors)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if service_role_arn and not isinstance(service_role_arn, str):
            raise TypeError("Expected argument 'service_role_arn' to be a str")
        pulumi.set(__self__, "service_role_arn", service_role_arn)
        if targets and not isinstance(targets, list):
            raise TypeError("Expected argument 'targets' to be a list")
        pulumi.set(__self__, "targets", targets)
        if task_arn and not isinstance(task_arn, str):
            raise TypeError("Expected argument 'task_arn' to be a str")
        pulumi.set(__self__, "task_arn", task_arn)
        if task_invocation_parameters and not isinstance(task_invocation_parameters, dict):
            raise TypeError("Expected argument 'task_invocation_parameters' to be a dict")
        pulumi.set(__self__, "task_invocation_parameters", task_invocation_parameters)
        if task_parameters and not isinstance(task_parameters, dict):
            raise TypeError("Expected argument 'task_parameters' to be a dict")
        pulumi.set(__self__, "task_parameters", task_parameters)

    @property
    @pulumi.getter(name="cutoffBehavior")
    def cutoff_behavior(self) -> Optional[str]:
        return pulumi.get(self, "cutoff_behavior")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loggingInfo")
    def logging_info(self) -> Optional['outputs.MaintenanceWindowTaskLoggingInfo']:
        return pulumi.get(self, "logging_info")

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> Optional[str]:
        return pulumi.get(self, "max_concurrency")

    @property
    @pulumi.getter(name="maxErrors")
    def max_errors(self) -> Optional[str]:
        return pulumi.get(self, "max_errors")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> Optional[int]:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="serviceRoleArn")
    def service_role_arn(self) -> Optional[str]:
        return pulumi.get(self, "service_role_arn")

    @property
    @pulumi.getter
    def targets(self) -> Optional[Sequence['outputs.MaintenanceWindowTaskTarget']]:
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="taskArn")
    def task_arn(self) -> Optional[str]:
        return pulumi.get(self, "task_arn")

    @property
    @pulumi.getter(name="taskInvocationParameters")
    def task_invocation_parameters(self) -> Optional['outputs.MaintenanceWindowTaskTaskInvocationParameters']:
        return pulumi.get(self, "task_invocation_parameters")

    @property
    @pulumi.getter(name="taskParameters")
    def task_parameters(self) -> Optional[Any]:
        return pulumi.get(self, "task_parameters")


class AwaitableGetMaintenanceWindowTaskResult(GetMaintenanceWindowTaskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMaintenanceWindowTaskResult(
            cutoff_behavior=self.cutoff_behavior,
            description=self.description,
            id=self.id,
            logging_info=self.logging_info,
            max_concurrency=self.max_concurrency,
            max_errors=self.max_errors,
            name=self.name,
            priority=self.priority,
            service_role_arn=self.service_role_arn,
            targets=self.targets,
            task_arn=self.task_arn,
            task_invocation_parameters=self.task_invocation_parameters,
            task_parameters=self.task_parameters)


def get_maintenance_window_task(id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMaintenanceWindowTaskResult:
    """
    Resource Type definition for AWS::SSM::MaintenanceWindowTask
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ssm:getMaintenanceWindowTask', __args__, opts=opts, typ=GetMaintenanceWindowTaskResult).value

    return AwaitableGetMaintenanceWindowTaskResult(
        cutoff_behavior=pulumi.get(__ret__, 'cutoff_behavior'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        logging_info=pulumi.get(__ret__, 'logging_info'),
        max_concurrency=pulumi.get(__ret__, 'max_concurrency'),
        max_errors=pulumi.get(__ret__, 'max_errors'),
        name=pulumi.get(__ret__, 'name'),
        priority=pulumi.get(__ret__, 'priority'),
        service_role_arn=pulumi.get(__ret__, 'service_role_arn'),
        targets=pulumi.get(__ret__, 'targets'),
        task_arn=pulumi.get(__ret__, 'task_arn'),
        task_invocation_parameters=pulumi.get(__ret__, 'task_invocation_parameters'),
        task_parameters=pulumi.get(__ret__, 'task_parameters'))


@_utilities.lift_output_func(get_maintenance_window_task)
def get_maintenance_window_task_output(id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMaintenanceWindowTaskResult]:
    """
    Resource Type definition for AWS::SSM::MaintenanceWindowTask
    """
    ...
