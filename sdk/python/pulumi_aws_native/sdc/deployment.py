# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DeploymentArgs', 'Deployment']

@pulumi.input_type
class DeploymentArgs:
    def __init__(__self__, *,
                 config_name: pulumi.Input[str],
                 dimension: pulumi.Input[str],
                 s3_bucket: pulumi.Input[str],
                 s3_key: pulumi.Input[str],
                 stage: pulumi.Input[str],
                 pipeline_id: Optional[pulumi.Input[str]] = None,
                 target_region_override: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Deployment resource.
        """
        pulumi.set(__self__, "config_name", config_name)
        pulumi.set(__self__, "dimension", dimension)
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        pulumi.set(__self__, "s3_key", s3_key)
        pulumi.set(__self__, "stage", stage)
        if pipeline_id is not None:
            pulumi.set(__self__, "pipeline_id", pipeline_id)
        if target_region_override is not None:
            pulumi.set(__self__, "target_region_override", target_region_override)

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "config_name")

    @config_name.setter
    def config_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_name", value)

    @property
    @pulumi.getter
    def dimension(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dimension")

    @dimension.setter
    def dimension(self, value: pulumi.Input[str]):
        pulumi.set(self, "dimension", value)

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "s3_bucket")

    @s3_bucket.setter
    def s3_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_bucket", value)

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "s3_key")

    @s3_key.setter
    def s3_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_key", value)

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Input[str]:
        return pulumi.get(self, "stage")

    @stage.setter
    def stage(self, value: pulumi.Input[str]):
        pulumi.set(self, "stage", value)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_id", value)

    @property
    @pulumi.getter(name="targetRegionOverride")
    def target_region_override(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target_region_override")

    @target_region_override.setter
    def target_region_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_region_override", value)


warnings.warn("""Deployment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Deployment(pulumi.CustomResource):
    warnings.warn("""Deployment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_name: Optional[pulumi.Input[str]] = None,
                 dimension: Optional[pulumi.Input[str]] = None,
                 pipeline_id: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_key: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 target_region_override: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AMZN::SDC::Deployment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeploymentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AMZN::SDC::Deployment

        :param str resource_name: The name of the resource.
        :param DeploymentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeploymentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_name: Optional[pulumi.Input[str]] = None,
                 dimension: Optional[pulumi.Input[str]] = None,
                 pipeline_id: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_key: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 target_region_override: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Deployment is deprecated: Deployment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeploymentArgs.__new__(DeploymentArgs)

            if config_name is None and not opts.urn:
                raise TypeError("Missing required property 'config_name'")
            __props__.__dict__["config_name"] = config_name
            if dimension is None and not opts.urn:
                raise TypeError("Missing required property 'dimension'")
            __props__.__dict__["dimension"] = dimension
            __props__.__dict__["pipeline_id"] = pipeline_id
            if s3_bucket is None and not opts.urn:
                raise TypeError("Missing required property 's3_bucket'")
            __props__.__dict__["s3_bucket"] = s3_bucket
            if s3_key is None and not opts.urn:
                raise TypeError("Missing required property 's3_key'")
            __props__.__dict__["s3_key"] = s3_key
            if stage is None and not opts.urn:
                raise TypeError("Missing required property 'stage'")
            __props__.__dict__["stage"] = stage
            __props__.__dict__["target_region_override"] = target_region_override
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["s3_key"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Deployment, __self__).__init__(
            'aws-native:sdc:Deployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Deployment':
        """
        Get an existing Deployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeploymentArgs.__new__(DeploymentArgs)

        __props__.__dict__["config_name"] = None
        __props__.__dict__["dimension"] = None
        __props__.__dict__["pipeline_id"] = None
        __props__.__dict__["s3_bucket"] = None
        __props__.__dict__["s3_key"] = None
        __props__.__dict__["stage"] = None
        __props__.__dict__["target_region_override"] = None
        return Deployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "config_name")

    @property
    @pulumi.getter
    def dimension(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dimension")

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "pipeline_id")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Output[str]:
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "s3_key")

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stage")

    @property
    @pulumi.getter(name="targetRegionOverride")
    def target_region_override(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "target_region_override")

