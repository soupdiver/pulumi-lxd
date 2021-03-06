// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class PublishImage extends pulumi.CustomResource {
    /**
     * Get an existing PublishImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublishImageState, opts?: pulumi.CustomResourceOptions): PublishImage {
        return new PublishImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'lxd:index/publishImage:PublishImage';

    /**
     * Returns true if the given object is an instance of PublishImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublishImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublishImage.__pulumiType;
    }

    public readonly aliases!: pulumi.Output<string[] | undefined>;
    public readonly compressionAlgorithm!: pulumi.Output<string | undefined>;
    public readonly container!: pulumi.Output<string>;
    public readonly filename!: pulumi.Output<string | undefined>;
    public readonly properties!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly public!: pulumi.Output<boolean | undefined>;
    /**
     * A map of arbitrary strings that, when changed, will force the resource to be replaced.
     */
    public readonly triggers!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a PublishImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublishImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublishImageArgs | PublishImageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublishImageState | undefined;
            inputs["aliases"] = state ? state.aliases : undefined;
            inputs["compressionAlgorithm"] = state ? state.compressionAlgorithm : undefined;
            inputs["container"] = state ? state.container : undefined;
            inputs["filename"] = state ? state.filename : undefined;
            inputs["properties"] = state ? state.properties : undefined;
            inputs["public"] = state ? state.public : undefined;
            inputs["triggers"] = state ? state.triggers : undefined;
        } else {
            const args = argsOrState as PublishImageArgs | undefined;
            if ((!args || args.container === undefined) && !opts.urn) {
                throw new Error("Missing required property 'container'");
            }
            inputs["aliases"] = args ? args.aliases : undefined;
            inputs["compressionAlgorithm"] = args ? args.compressionAlgorithm : undefined;
            inputs["container"] = args ? args.container : undefined;
            inputs["filename"] = args ? args.filename : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["public"] = args ? args.public : undefined;
            inputs["triggers"] = args ? args.triggers : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PublishImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublishImage resources.
 */
export interface PublishImageState {
    readonly aliases?: pulumi.Input<pulumi.Input<string>[]>;
    readonly compressionAlgorithm?: pulumi.Input<string>;
    readonly container?: pulumi.Input<string>;
    readonly filename?: pulumi.Input<string>;
    readonly properties?: pulumi.Input<{[key: string]: any}>;
    readonly public?: pulumi.Input<boolean>;
    /**
     * A map of arbitrary strings that, when changed, will force the resource to be replaced.
     */
    readonly triggers?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a PublishImage resource.
 */
export interface PublishImageArgs {
    readonly aliases?: pulumi.Input<pulumi.Input<string>[]>;
    readonly compressionAlgorithm?: pulumi.Input<string>;
    readonly container: pulumi.Input<string>;
    readonly filename?: pulumi.Input<string>;
    readonly properties?: pulumi.Input<{[key: string]: any}>;
    readonly public?: pulumi.Input<boolean>;
    /**
     * A map of arbitrary strings that, when changed, will force the resource to be replaced.
     */
    readonly triggers?: pulumi.Input<{[key: string]: any}>;
}
