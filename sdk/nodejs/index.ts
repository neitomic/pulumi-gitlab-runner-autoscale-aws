// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./gitlabRunnerAwsa";
export * from "./provider";

// Export sub-modules:
import * as types from "./types";

export {
    types,
};

// Import resources to register:
import { GitlabRunnerAwsa } from "./gitlabRunnerAwsa";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gitlab-runner-awsa:index:GitlabRunnerAwsa":
                return new GitlabRunnerAwsa(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gitlab-runner-awsa", "index", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("gitlab-runner-awsa", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:gitlab-runner-awsa") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
