// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Lxd
{
    [LxdResourceType("lxd:index/container:Container")]
    public partial class Container : Pulumi.CustomResource
    {
        [Output("config")]
        public Output<ImmutableDictionary<string, object>?> Config { get; private set; } = null!;

        [Output("devices")]
        public Output<ImmutableArray<Outputs.ContainerDevice>> Devices { get; private set; } = null!;

        [Output("ephemeral")]
        public Output<bool?> Ephemeral { get; private set; } = null!;

        [Output("files")]
        public Output<ImmutableArray<Outputs.ContainerFile>> Files { get; private set; } = null!;

        [Output("image")]
        public Output<string> Image { get; private set; } = null!;

        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        [Output("ipv4Address")]
        public Output<string> Ipv4Address { get; private set; } = null!;

        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        [Output("limits")]
        public Output<ImmutableDictionary<string, object>?> Limits { get; private set; } = null!;

        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("privileged")]
        public Output<bool?> Privileged { get; private set; } = null!;

        [Output("profiles")]
        public Output<ImmutableArray<string>> Profiles { get; private set; } = null!;

        [Output("remote")]
        public Output<string?> Remote { get; private set; } = null!;

        [Output("startContainer")]
        public Output<bool?> StartContainer { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("waitForNetwork")]
        public Output<bool?> WaitForNetwork { get; private set; } = null!;


        /// <summary>
        /// Create a Container resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Container(string name, ContainerArgs args, CustomResourceOptions? options = null)
            : base("lxd:index/container:Container", name, args ?? new ContainerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Container(string name, Input<string> id, ContainerState? state = null, CustomResourceOptions? options = null)
            : base("lxd:index/container:Container", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Container resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Container Get(string name, Input<string> id, ContainerState? state = null, CustomResourceOptions? options = null)
        {
            return new Container(name, id, state, options);
        }
    }

    public sealed class ContainerArgs : Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<object>? _config;
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        [Input("devices")]
        private InputList<Inputs.ContainerDeviceArgs>? _devices;
        public InputList<Inputs.ContainerDeviceArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.ContainerDeviceArgs>());
            set => _devices = value;
        }

        [Input("ephemeral")]
        public Input<bool>? Ephemeral { get; set; }

        [Input("files")]
        private InputList<Inputs.ContainerFileArgs>? _files;
        public InputList<Inputs.ContainerFileArgs> Files
        {
            get => _files ?? (_files = new InputList<Inputs.ContainerFileArgs>());
            set => _files = value;
        }

        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        [Input("limits")]
        private InputMap<object>? _limits;
        public InputMap<object> Limits
        {
            get => _limits ?? (_limits = new InputMap<object>());
            set => _limits = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privileged")]
        public Input<bool>? Privileged { get; set; }

        [Input("profiles")]
        private InputList<string>? _profiles;
        public InputList<string> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<string>());
            set => _profiles = value;
        }

        [Input("remote")]
        public Input<string>? Remote { get; set; }

        [Input("startContainer")]
        public Input<bool>? StartContainer { get; set; }

        [Input("target")]
        public Input<string>? Target { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("waitForNetwork")]
        public Input<bool>? WaitForNetwork { get; set; }

        public ContainerArgs()
        {
        }
    }

    public sealed class ContainerState : Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<object>? _config;
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        [Input("devices")]
        private InputList<Inputs.ContainerDeviceGetArgs>? _devices;
        public InputList<Inputs.ContainerDeviceGetArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.ContainerDeviceGetArgs>());
            set => _devices = value;
        }

        [Input("ephemeral")]
        public Input<bool>? Ephemeral { get; set; }

        [Input("files")]
        private InputList<Inputs.ContainerFileGetArgs>? _files;
        public InputList<Inputs.ContainerFileGetArgs> Files
        {
            get => _files ?? (_files = new InputList<Inputs.ContainerFileGetArgs>());
            set => _files = value;
        }

        [Input("image")]
        public Input<string>? Image { get; set; }

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("limits")]
        private InputMap<object>? _limits;
        public InputMap<object> Limits
        {
            get => _limits ?? (_limits = new InputMap<object>());
            set => _limits = value;
        }

        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privileged")]
        public Input<bool>? Privileged { get; set; }

        [Input("profiles")]
        private InputList<string>? _profiles;
        public InputList<string> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<string>());
            set => _profiles = value;
        }

        [Input("remote")]
        public Input<string>? Remote { get; set; }

        [Input("startContainer")]
        public Input<bool>? StartContainer { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("target")]
        public Input<string>? Target { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("waitForNetwork")]
        public Input<bool>? WaitForNetwork { get; set; }

        public ContainerState()
        {
        }
    }
}