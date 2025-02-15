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
    'ProactiveEngagementEmergencyContactArgs',
    'ProtectionApplicationLayerAutomaticResponseConfigurationAction0PropertiesArgs',
    'ProtectionApplicationLayerAutomaticResponseConfigurationAction1PropertiesArgs',
    'ProtectionApplicationLayerAutomaticResponseConfigurationArgs',
    'ProtectionGroupTagArgs',
    'ProtectionTagArgs',
]

@pulumi.input_type
class ProactiveEngagementEmergencyContactArgs:
    def __init__(__self__, *,
                 email_address: pulumi.Input[str],
                 contact_notes: Optional[pulumi.Input[str]] = None,
                 phone_number: Optional[pulumi.Input[str]] = None):
        """
        An emergency contact is used by Shield Response Team (SRT) to contact you for escalations to the SRT and to initiate proactive customer support. An emergency contact requires an email address.
        :param pulumi.Input[str] email_address: The email address for the contact.
        :param pulumi.Input[str] contact_notes: Additional notes regarding the contact.
        :param pulumi.Input[str] phone_number: The phone number for the contact
        """
        pulumi.set(__self__, "email_address", email_address)
        if contact_notes is not None:
            pulumi.set(__self__, "contact_notes", contact_notes)
        if phone_number is not None:
            pulumi.set(__self__, "phone_number", phone_number)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> pulumi.Input[str]:
        """
        The email address for the contact.
        """
        return pulumi.get(self, "email_address")

    @email_address.setter
    def email_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "email_address", value)

    @property
    @pulumi.getter(name="contactNotes")
    def contact_notes(self) -> Optional[pulumi.Input[str]]:
        """
        Additional notes regarding the contact.
        """
        return pulumi.get(self, "contact_notes")

    @contact_notes.setter
    def contact_notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contact_notes", value)

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> Optional[pulumi.Input[str]]:
        """
        The phone number for the contact
        """
        return pulumi.get(self, "phone_number")

    @phone_number.setter
    def phone_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_number", value)


@pulumi.input_type
class ProtectionApplicationLayerAutomaticResponseConfigurationAction0PropertiesArgs:
    def __init__(__self__, *,
                 count: Optional[Any] = None):
        """
        Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
        :param Any count: Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Count` action.
               You must specify exactly one action, either `Block` or `Count`.
        """
        if count is not None:
            pulumi.set(__self__, "count", count)

    @property
    @pulumi.getter
    def count(self) -> Optional[Any]:
        """
        Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Count` action.
        You must specify exactly one action, either `Block` or `Count`.
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: Optional[Any]):
        pulumi.set(self, "count", value)


@pulumi.input_type
class ProtectionApplicationLayerAutomaticResponseConfigurationAction1PropertiesArgs:
    def __init__(__self__, *,
                 block: Optional[Any] = None):
        """
        Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
        :param Any block: Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Block` action.
               You must specify exactly one action, either `Block` or `Count`.
        """
        if block is not None:
            pulumi.set(__self__, "block", block)

    @property
    @pulumi.getter
    def block(self) -> Optional[Any]:
        """
        Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Block` action.
        You must specify exactly one action, either `Block` or `Count`.
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: Optional[Any]):
        pulumi.set(self, "block", value)


@pulumi.input_type
class ProtectionApplicationLayerAutomaticResponseConfigurationArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[Union['ProtectionApplicationLayerAutomaticResponseConfigurationAction0PropertiesArgs', 'ProtectionApplicationLayerAutomaticResponseConfigurationAction1PropertiesArgs']],
                 status: pulumi.Input['ProtectionApplicationLayerAutomaticResponseConfigurationStatus']):
        """
        The automatic application layer DDoS mitigation settings for a Protection. This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
        :param pulumi.Input[Union['ProtectionApplicationLayerAutomaticResponseConfigurationAction0PropertiesArgs', 'ProtectionApplicationLayerAutomaticResponseConfigurationAction1PropertiesArgs']] action: Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
        :param pulumi.Input['ProtectionApplicationLayerAutomaticResponseConfigurationStatus'] status: Indicates whether automatic application layer DDoS mitigation is enabled for the protection.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[Union['ProtectionApplicationLayerAutomaticResponseConfigurationAction0PropertiesArgs', 'ProtectionApplicationLayerAutomaticResponseConfigurationAction1PropertiesArgs']]:
        """
        Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[Union['ProtectionApplicationLayerAutomaticResponseConfigurationAction0PropertiesArgs', 'ProtectionApplicationLayerAutomaticResponseConfigurationAction1PropertiesArgs']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input['ProtectionApplicationLayerAutomaticResponseConfigurationStatus']:
        """
        Indicates whether automatic application layer DDoS mitigation is enabled for the protection.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input['ProtectionApplicationLayerAutomaticResponseConfigurationStatus']):
        pulumi.set(self, "status", value)


@pulumi.input_type
class ProtectionGroupTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A tag associated with an AWS resource. Tags are key:value pairs that you can use to categorize and manage your resources, for purposes like billing or other management. Typically, the tag key represents a category, such as "environment", and the tag value represents a specific value within that category, such as "test," "development," or "production". Or you might set the tag key to "customer" and the value to the customer name or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a resource.
        :param pulumi.Input[str] key: Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
        :param pulumi.Input[str] value: Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ProtectionTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A tag associated with an AWS resource. Tags are key:value pairs that you can use to categorize and manage your resources, for purposes like billing or other management. Typically, the tag key represents a category, such as "environment", and the tag value represents a specific value within that category, such as "test," "development," or "production". Or you might set the tag key to "customer" and the value to the customer name or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a resource.
        :param pulumi.Input[str] key: Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
        :param pulumi.Input[str] value: Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


