// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VolumeContainerAttach extends pulumi.CustomResource {
    /**
     * Get an existing VolumeContainerAttach resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeContainerAttachState, opts?: pulumi.CustomResourceOptions): VolumeContainerAttach {
        return new VolumeContainerAttach(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'lxd:index/volumeContainerAttach:VolumeContainerAttach';

    /**
     * Returns true if the given object is an instance of VolumeContainerAttach.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeContainerAttach {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeContainerAttach.__pulumiType;
    }

    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    public readonly containerName!: pulumi.Output<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    public readonly deviceName!: pulumi.Output<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    public readonly pool!: pulumi.Output<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    public readonly remote!: pulumi.Output<string | undefined>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    public readonly volumeName!: pulumi.Output<string>;

    /**
     * Create a VolumeContainerAttach resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeContainerAttachArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeContainerAttachArgs | VolumeContainerAttachState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeContainerAttachState | undefined;
            inputs["containerName"] = state ? state.containerName : undefined;
            inputs["deviceName"] = state ? state.deviceName : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["pool"] = state ? state.pool : undefined;
            inputs["remote"] = state ? state.remote : undefined;
            inputs["volumeName"] = state ? state.volumeName : undefined;
        } else {
            const args = argsOrState as VolumeContainerAttachArgs | undefined;
            if ((!args || args.containerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containerName'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.pool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pool'");
            }
            if ((!args || args.volumeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeName'");
            }
            inputs["containerName"] = args ? args.containerName : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["pool"] = args ? args.pool : undefined;
            inputs["remote"] = args ? args.remote : undefined;
            inputs["volumeName"] = args ? args.volumeName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VolumeContainerAttach.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VolumeContainerAttach resources.
 */
export interface VolumeContainerAttachState {
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly containerName?: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly deviceName?: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly path?: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly pool?: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly remote?: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly volumeName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VolumeContainerAttach resource.
 */
export interface VolumeContainerAttachArgs {
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly containerName: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly deviceName?: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly path: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly pool: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly remote?: pulumi.Input<string>;
    /**
     * @deprecated lxd_volume_container_attach has been deprecated and will be removed
     */
    readonly volumeName: pulumi.Input<string>;
}
