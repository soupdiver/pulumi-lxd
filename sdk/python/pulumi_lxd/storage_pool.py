# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['StoragePoolArgs', 'StoragePool']

@pulumi.input_type
class StoragePoolArgs:
    def __init__(__self__, *,
                 driver: pulumi.Input[str],
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StoragePool resource.
        """
        pulumi.set(__self__, "driver", driver)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remote is not None:
            pulumi.set(__self__, "remote", remote)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def driver(self) -> pulumi.Input[str]:
        return pulumi.get(self, "driver")

    @driver.setter
    def driver(self, value: pulumi.Input[str]):
        pulumi.set(self, "driver", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remote(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "remote")

    @remote.setter
    def remote(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)


@pulumi.input_type
class _StoragePoolState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 driver: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StoragePool resources.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if driver is not None:
            pulumi.set(__self__, "driver", driver)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remote is not None:
            pulumi.set(__self__, "remote", remote)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def driver(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "driver")

    @driver.setter
    def driver(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "driver", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remote(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "remote")

    @remote.setter
    def remote(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)


class StoragePool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 driver: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a StoragePool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StoragePoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a StoragePool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param StoragePoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StoragePoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 driver: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StoragePoolArgs.__new__(StoragePoolArgs)

            __props__.__dict__["config"] = config
            if driver is None and not opts.urn:
                raise TypeError("Missing required property 'driver'")
            __props__.__dict__["driver"] = driver
            __props__.__dict__["name"] = name
            __props__.__dict__["remote"] = remote
            __props__.__dict__["target"] = target
        super(StoragePool, __self__).__init__(
            'lxd:index/storagePool:StoragePool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            driver: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            remote: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None) -> 'StoragePool':
        """
        Get an existing StoragePool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StoragePoolState.__new__(_StoragePoolState)

        __props__.__dict__["config"] = config
        __props__.__dict__["driver"] = driver
        __props__.__dict__["name"] = name
        __props__.__dict__["remote"] = remote
        __props__.__dict__["target"] = target
        return StoragePool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def driver(self) -> pulumi.Output[str]:
        return pulumi.get(self, "driver")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def remote(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "remote")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "target")

