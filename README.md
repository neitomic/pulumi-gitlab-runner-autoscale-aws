# Pulumi Component For Gitlab Runner AWS Autoscale (TypeScript)

This repo contains Gitlab Runner AWS Autoscale component for Pulumi witten in TypeScript.


## Contribute
### Prerequisites

- Pulumi CLI
- Node.js
- Yarn
- Go 1.17 (to regenerate the SDKs)
- Python 3.6+ (to build the Python SDK)
- .NET Core SDK (to build the .NET SDK)

### Build and Test

```bash
# Build and install the provider
make install_provider

# Regenerate SDKs
make generate

# Ensure the pulumi-provider-gitlab-runner-awsa script is on PATH
$ export PATH=$PATH:$PWD/bin

# Test Node.js SDK
$ make install_nodejs_sdk
$ cd examples/simple
$ yarn install
$ yarn link @pulumi/gitlab-runner-awsa
$ pulumi stack init test
$ pulumi config set aws:region us-east-1
$ pulumi up
```

## Usage

See [example](examples/simple)

```typescript
import * as gitlabRunnerAwsa  from "@pulumi/gitlab-runner-awsa";

const runner = new gitlabRunnerAwsa.GitlabRunnerAwsa("test-runner-awsa", {
    gitlabUrl: "https://gitlab.com",
    gitlabRunnerToken: "",
    machineIdleNodes: 0,
    machineIdleTimeSecond: 1800,
    machineInstanceType: "t3.2xlarge",
    machineMaxBuilds: 20,
    runnerMaxConcurrentBuild: 10,
    subnetId: "subnet-xxxxxxx"
});
```

## TODOs

- [ ] Support more outputs
- [ ] Support S3 cache
- [ ] Support more machine options (instance profile, userdata,...)
- [ ] Support machine private/public IP addresses (enable/disable)
- [ ] Support machine spot instance 
- [ ] Support auto-scaling config
- [ ] Support Docker auth
- [ ] Support Docker Registry Proxy with rpardini/docker-registry-proxy
- [ ] Workaround private network environment (no Internet access)