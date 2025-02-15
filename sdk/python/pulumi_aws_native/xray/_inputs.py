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
    'GroupInsightsConfigurationArgs',
    'GroupTagArgs',
    'SamplingRuleRecordArgs',
    'SamplingRuleTagArgs',
    'SamplingRuleUpdateArgs',
    'SamplingRuleArgs',
]

@pulumi.input_type
class GroupInsightsConfigurationArgs:
    def __init__(__self__, *,
                 insights_enabled: Optional[pulumi.Input[bool]] = None,
                 notifications_enabled: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] insights_enabled: Set the InsightsEnabled value to true to enable insights or false to disable insights.
        :param pulumi.Input[bool] notifications_enabled: Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.
        """
        if insights_enabled is not None:
            pulumi.set(__self__, "insights_enabled", insights_enabled)
        if notifications_enabled is not None:
            pulumi.set(__self__, "notifications_enabled", notifications_enabled)

    @property
    @pulumi.getter(name="insightsEnabled")
    def insights_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set the InsightsEnabled value to true to enable insights or false to disable insights.
        """
        return pulumi.get(self, "insights_enabled")

    @insights_enabled.setter
    def insights_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insights_enabled", value)

    @property
    @pulumi.getter(name="notificationsEnabled")
    def notifications_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.
        """
        return pulumi.get(self, "notifications_enabled")

    @notifications_enabled.setter
    def notifications_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notifications_enabled", value)


@pulumi.input_type
class GroupTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: The key name of the tag.
        :param pulumi.Input[str] value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class SamplingRuleRecordArgs:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 modified_at: Optional[pulumi.Input[str]] = None,
                 sampling_rule: Optional[pulumi.Input['SamplingRuleArgs']] = None):
        """
        :param pulumi.Input[str] created_at: When the rule was created, in Unix time seconds.
        :param pulumi.Input[str] modified_at: When the rule was modified, in Unix time seconds.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if modified_at is not None:
            pulumi.set(__self__, "modified_at", modified_at)
        if sampling_rule is not None:
            pulumi.set(__self__, "sampling_rule", sampling_rule)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        When the rule was created, in Unix time seconds.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="modifiedAt")
    def modified_at(self) -> Optional[pulumi.Input[str]]:
        """
        When the rule was modified, in Unix time seconds.
        """
        return pulumi.get(self, "modified_at")

    @modified_at.setter
    def modified_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modified_at", value)

    @property
    @pulumi.getter(name="samplingRule")
    def sampling_rule(self) -> Optional[pulumi.Input['SamplingRuleArgs']]:
        return pulumi.get(self, "sampling_rule")

    @sampling_rule.setter
    def sampling_rule(self, value: Optional[pulumi.Input['SamplingRuleArgs']]):
        pulumi.set(self, "sampling_rule", value)


@pulumi.input_type
class SamplingRuleTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: The key name of the tag.
        :param pulumi.Input[str] value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class SamplingRuleUpdateArgs:
    def __init__(__self__, *,
                 attributes: Optional[Any] = None,
                 fixed_rate: Optional[pulumi.Input[float]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 reservoir_size: Optional[pulumi.Input[int]] = None,
                 resource_arn: Optional[pulumi.Input[str]] = None,
                 rule_arn: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 service_type: Optional[pulumi.Input[str]] = None,
                 url_path: Optional[pulumi.Input[str]] = None):
        """
        :param Any attributes: Matches attributes derived from the request.
        :param pulumi.Input[float] fixed_rate: The percentage of matching requests to instrument, after the reservoir is exhausted.
        :param pulumi.Input[str] host: Matches the hostname from a request URL.
        :param pulumi.Input[str] http_method: Matches the HTTP method from a request URL.
        :param pulumi.Input[int] priority: The priority of the sampling rule.
        :param pulumi.Input[int] reservoir_size: A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
        :param pulumi.Input[str] resource_arn: Matches the ARN of the AWS resource on which the service runs.
        :param pulumi.Input[str] service_name: Matches the name that the service uses to identify itself in segments.
        :param pulumi.Input[str] service_type: Matches the origin that the service uses to identify its type in segments.
        :param pulumi.Input[str] url_path: Matches the path from a request URL.
        """
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if fixed_rate is not None:
            pulumi.set(__self__, "fixed_rate", fixed_rate)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if reservoir_size is not None:
            pulumi.set(__self__, "reservoir_size", reservoir_size)
        if resource_arn is not None:
            pulumi.set(__self__, "resource_arn", resource_arn)
        if rule_arn is not None:
            pulumi.set(__self__, "rule_arn", rule_arn)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if service_type is not None:
            pulumi.set(__self__, "service_type", service_type)
        if url_path is not None:
            pulumi.set(__self__, "url_path", url_path)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Any]:
        """
        Matches attributes derived from the request.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[Any]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter(name="fixedRate")
    def fixed_rate(self) -> Optional[pulumi.Input[float]]:
        """
        The percentage of matching requests to instrument, after the reservoir is exhausted.
        """
        return pulumi.get(self, "fixed_rate")

    @fixed_rate.setter
    def fixed_rate(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "fixed_rate", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Matches the hostname from a request URL.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Matches the HTTP method from a request URL.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the sampling rule.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="reservoirSize")
    def reservoir_size(self) -> Optional[pulumi.Input[int]]:
        """
        A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
        """
        return pulumi.get(self, "reservoir_size")

    @reservoir_size.setter
    def reservoir_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reservoir_size", value)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Matches the ARN of the AWS resource on which the service runs.
        """
        return pulumi.get(self, "resource_arn")

    @resource_arn.setter
    def resource_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_arn", value)

    @property
    @pulumi.getter(name="ruleArn")
    def rule_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rule_arn")

    @rule_arn.setter
    def rule_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_arn", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Matches the name that the service uses to identify itself in segments.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[pulumi.Input[str]]:
        """
        Matches the origin that the service uses to identify its type in segments.
        """
        return pulumi.get(self, "service_type")

    @service_type.setter
    def service_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_type", value)

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> Optional[pulumi.Input[str]]:
        """
        Matches the path from a request URL.
        """
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_path", value)


@pulumi.input_type
class SamplingRuleArgs:
    def __init__(__self__, *,
                 fixed_rate: pulumi.Input[float],
                 host: pulumi.Input[str],
                 http_method: pulumi.Input[str],
                 priority: pulumi.Input[int],
                 reservoir_size: pulumi.Input[int],
                 resource_arn: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 service_type: pulumi.Input[str],
                 url_path: pulumi.Input[str],
                 attributes: Optional[Any] = None,
                 rule_arn: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[float] fixed_rate: The percentage of matching requests to instrument, after the reservoir is exhausted.
        :param pulumi.Input[str] host: Matches the hostname from a request URL.
        :param pulumi.Input[str] http_method: Matches the HTTP method from a request URL.
        :param pulumi.Input[int] priority: The priority of the sampling rule.
        :param pulumi.Input[int] reservoir_size: A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
        :param pulumi.Input[str] resource_arn: Matches the ARN of the AWS resource on which the service runs.
        :param pulumi.Input[str] service_name: Matches the name that the service uses to identify itself in segments.
        :param pulumi.Input[str] service_type: Matches the origin that the service uses to identify its type in segments.
        :param pulumi.Input[str] url_path: Matches the path from a request URL.
        :param Any attributes: Matches attributes derived from the request.
        :param pulumi.Input[int] version: The version of the sampling rule format (1)
        """
        pulumi.set(__self__, "fixed_rate", fixed_rate)
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "http_method", http_method)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "reservoir_size", reservoir_size)
        pulumi.set(__self__, "resource_arn", resource_arn)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "service_type", service_type)
        pulumi.set(__self__, "url_path", url_path)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if rule_arn is not None:
            pulumi.set(__self__, "rule_arn", rule_arn)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="fixedRate")
    def fixed_rate(self) -> pulumi.Input[float]:
        """
        The percentage of matching requests to instrument, after the reservoir is exhausted.
        """
        return pulumi.get(self, "fixed_rate")

    @fixed_rate.setter
    def fixed_rate(self, value: pulumi.Input[float]):
        pulumi.set(self, "fixed_rate", value)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        Matches the hostname from a request URL.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> pulumi.Input[str]:
        """
        Matches the HTTP method from a request URL.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: pulumi.Input[str]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        """
        The priority of the sampling rule.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="reservoirSize")
    def reservoir_size(self) -> pulumi.Input[int]:
        """
        A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
        """
        return pulumi.get(self, "reservoir_size")

    @reservoir_size.setter
    def reservoir_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "reservoir_size", value)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Input[str]:
        """
        Matches the ARN of the AWS resource on which the service runs.
        """
        return pulumi.get(self, "resource_arn")

    @resource_arn.setter
    def resource_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_arn", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        Matches the name that the service uses to identify itself in segments.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> pulumi.Input[str]:
        """
        Matches the origin that the service uses to identify its type in segments.
        """
        return pulumi.get(self, "service_type")

    @service_type.setter
    def service_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_type", value)

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> pulumi.Input[str]:
        """
        Matches the path from a request URL.
        """
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "url_path", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Any]:
        """
        Matches attributes derived from the request.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[Any]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter(name="ruleArn")
    def rule_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rule_arn")

    @rule_arn.setter
    def rule_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_arn", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        The version of the sampling rule format (1)
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


