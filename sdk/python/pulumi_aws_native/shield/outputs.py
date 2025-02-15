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
from ._enums import *

__all__ = [
    'ProactiveEngagementEmergencyContact',
    'ProtectionApplicationLayerAutomaticResponseConfiguration',
    'ProtectionApplicationLayerAutomaticResponseConfigurationAction0Properties',
    'ProtectionApplicationLayerAutomaticResponseConfigurationAction1Properties',
    'ProtectionGroupTag',
    'ProtectionTag',
]

@pulumi.output_type
class ProactiveEngagementEmergencyContact(dict):
    """
    An emergency contact is used by Shield Response Team (SRT) to contact you for escalations to the SRT and to initiate proactive customer support. An emergency contact requires an email address.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "emailAddress":
            suggest = "email_address"
        elif key == "contactNotes":
            suggest = "contact_notes"
        elif key == "phoneNumber":
            suggest = "phone_number"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProactiveEngagementEmergencyContact. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProactiveEngagementEmergencyContact.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProactiveEngagementEmergencyContact.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 email_address: str,
                 contact_notes: Optional[str] = None,
                 phone_number: Optional[str] = None):
        """
        An emergency contact is used by Shield Response Team (SRT) to contact you for escalations to the SRT and to initiate proactive customer support. An emergency contact requires an email address.
        :param str email_address: The email address for the contact.
        :param str contact_notes: Additional notes regarding the contact.
        :param str phone_number: The phone number for the contact
        """
        pulumi.set(__self__, "email_address", email_address)
        if contact_notes is not None:
            pulumi.set(__self__, "contact_notes", contact_notes)
        if phone_number is not None:
            pulumi.set(__self__, "phone_number", phone_number)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> str:
        """
        The email address for the contact.
        """
        return pulumi.get(self, "email_address")

    @property
    @pulumi.getter(name="contactNotes")
    def contact_notes(self) -> Optional[str]:
        """
        Additional notes regarding the contact.
        """
        return pulumi.get(self, "contact_notes")

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> Optional[str]:
        """
        The phone number for the contact
        """
        return pulumi.get(self, "phone_number")


@pulumi.output_type
class ProtectionApplicationLayerAutomaticResponseConfiguration(dict):
    """
    The automatic application layer DDoS mitigation settings for a Protection. This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
    """
    def __init__(__self__, *,
                 action: Any,
                 status: 'ProtectionApplicationLayerAutomaticResponseConfigurationStatus'):
        """
        The automatic application layer DDoS mitigation settings for a Protection. This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
        :param Union['ProtectionApplicationLayerAutomaticResponseConfigurationAction0Properties', 'ProtectionApplicationLayerAutomaticResponseConfigurationAction1Properties'] action: Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
        :param 'ProtectionApplicationLayerAutomaticResponseConfigurationStatus' status: Indicates whether automatic application layer DDoS mitigation is enabled for the protection.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def action(self) -> Any:
        """
        Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def status(self) -> 'ProtectionApplicationLayerAutomaticResponseConfigurationStatus':
        """
        Indicates whether automatic application layer DDoS mitigation is enabled for the protection.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ProtectionApplicationLayerAutomaticResponseConfigurationAction0Properties(dict):
    """
    Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
    """
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


@pulumi.output_type
class ProtectionApplicationLayerAutomaticResponseConfigurationAction1Properties(dict):
    """
    Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
    """
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


@pulumi.output_type
class ProtectionGroupTag(dict):
    """
    A tag associated with an AWS resource. Tags are key:value pairs that you can use to categorize and manage your resources, for purposes like billing or other management. Typically, the tag key represents a category, such as "environment", and the tag value represents a specific value within that category, such as "test," "development," or "production". Or you might set the tag key to "customer" and the value to the customer name or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a resource.
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A tag associated with an AWS resource. Tags are key:value pairs that you can use to categorize and manage your resources, for purposes like billing or other management. Typically, the tag key represents a category, such as "environment", and the tag value represents a specific value within that category, such as "test," "development," or "production". Or you might set the tag key to "customer" and the value to the customer name or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a resource.
        :param str key: Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
        :param str value: Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ProtectionTag(dict):
    """
    A tag associated with an AWS resource. Tags are key:value pairs that you can use to categorize and manage your resources, for purposes like billing or other management. Typically, the tag key represents a category, such as "environment", and the tag value represents a specific value within that category, such as "test," "development," or "production". Or you might set the tag key to "customer" and the value to the customer name or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a resource.
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A tag associated with an AWS resource. Tags are key:value pairs that you can use to categorize and manage your resources, for purposes like billing or other management. Typically, the tag key represents a category, such as "environment", and the tag value represents a specific value within that category, such as "test," "development," or "production". Or you might set the tag key to "customer" and the value to the customer name or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a resource.
        :param str key: Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
        :param str value: Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
        """
        return pulumi.get(self, "value")


