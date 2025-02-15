# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DestinationExpressionType',
    'NetworkAnalyzerConfigurationLogLevel',
    'NetworkAnalyzerConfigurationWirelessDeviceFrameInfo',
    'PartnerAccountPartnerType',
    'TaskDefinitionType',
    'WirelessDeviceImportTaskStatus',
    'WirelessDeviceType',
]


class DestinationExpressionType(str, Enum):
    """
    Must be RuleName
    """
    RULE_NAME = "RuleName"
    MQTT_TOPIC = "MqttTopic"
    SNS_TOPIC = "SnsTopic"


class NetworkAnalyzerConfigurationLogLevel(str, Enum):
    INFO = "INFO"
    ERROR = "ERROR"
    DISABLED = "DISABLED"


class NetworkAnalyzerConfigurationWirelessDeviceFrameInfo(str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class PartnerAccountPartnerType(str, Enum):
    """
    The partner type
    """
    SIDEWALK = "Sidewalk"


class TaskDefinitionType(str, Enum):
    """
    A filter to list only the wireless gateway task definitions that use this task definition type
    """
    UPDATE = "UPDATE"


class WirelessDeviceImportTaskStatus(str, Enum):
    """
    Status for import task
    """
    INITIALIZING = "INITIALIZING"
    INITIALIZED = "INITIALIZED"
    PENDING = "PENDING"
    COMPLETE = "COMPLETE"
    FAILED = "FAILED"
    DELETING = "DELETING"


class WirelessDeviceType(str, Enum):
    """
    Wireless device type, currently only Sidewalk and LoRa
    """
    SIDEWALK = "Sidewalk"
    LO_RA_WAN = "LoRaWAN"
