import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import { render } from "template-file";

export interface VolumeArgs {
    size: number;
    type: string;
}
//todo: render
export interface S3Cache {
    path: string;
    serverAddress?: string;
    bucketName: string;
}
export interface GitlabRunnerAwsaArgs {
    subnetId: pulumi.Input<string>;
    gitlabUrl: pulumi.Input<string>;
    gitlabRunnerToken: pulumi.Input<string>;
    runnerMaxConcurrentBuild: pulumi.Input<number>;
    gitlabHelperImage?: pulumi.Input<string>;
    machineVolume?: VolumeArgs;
    machineInstanceType: pulumi.Input<string>;
    machineTags?: {
        [key: string]: string;
    };
    cache?: S3Cache;
    defaultDockerImage?: pulumi.Input<string>,
    tags?: {
        [key: string]: string;
    };
    machineIdleNodes: pulumi.Input<number>;
    machineIdleTimeSecond: pulumi.Input<number>;
    machineMaxBuilds: pulumi.Input<number>;
    outputLimit: pulumi.Input<number>;
}

export class GitlabRunnerAwsa extends pulumi.ComponentResource {
    public readonly gitlabRunnerUserData: pulumi.Input<string>;

    constructor(name: string, args: GitlabRunnerAwsaArgs, opts?: pulumi.ComponentResourceOptions) {
        super("gitlab-runner-awsa:index:GitlabRunnerAwsa", name, args, opts);
        const subnetDetail = pulumi.all([args.subnetId]).apply(([subnetId]) => {
            return aws.ec2.getSubnet({ id: subnetId })
        })

        const gitlabRunnerScalerRole = new aws.iam.Role(`${name}-role`, {
            name: `${name}-role`,
            assumeRolePolicy: aws.iam.assumeRolePolicyForPrincipal({ Service: `ec2.amazonaws.com` }),
            managedPolicyArns: [
                "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
                "arn:aws:iam::aws:policy/CloudWatchAgentAdminPolicy"
            ]
        })
        const gitlabRunnerScalerRolePolicy = new aws.iam.RolePolicy(`${name}-role-policy`, {
            role: gitlabRunnerScalerRole.id,
            name: `${name}-role-policy`,
            policy: {
                Version: "2012-10-17",
                Statement: [
                    {
                        Effect: "Allow",
                        Action: "ec2:*",
                        Resource: "*",
                    },
                ],
            },
            //todo: fine-tune role
        });

        const gitlabRunnerScalerInstanceProfile = new aws.iam.InstanceProfile(`${name}-instance-profile`, {
            name: `${name}-instance-profile`,
            path: "/",
            role: gitlabRunnerScalerRole.name
        })

        const amz2Ami = aws.ec2.getAmi({
            mostRecent: true,
            filters: [
                {
                    name: "name",
                    values: ["amzn2-ami-kernel-5.10-hvm-*"]
                },
                {
                    name: "architecture",
                    values: ["x86_64"]
                },
                {
                    name: "virtualization-type",
                    values: ["hvm"],
                },
            ],
            owners: ["amazon"]
        })

        const gitlabRunnerSecGroup = new aws.ec2.SecurityGroup(`${name}-security-group`, {
            name: `${name}-security-group`,
            vpcId: subnetDetail.apply(subnet => subnet.vpcId),
            egress: [
                {
                    description: "allow all egress",
                    cidrBlocks: ["0.0.0.0/0"],
                    fromPort: 0,
                    toPort: 0,
                    protocol: "-1",
                }
            ],
            ingress: [
                {
                    description: "allow all from itself",
                    fromPort: 0,
                    toPort: 0,
                    protocol: "-1",
                    self: true
                }
            ]
        })

        const machineRootDeviceDisk = args.machineVolume ?? { size: 10, type: "gp3" }
        const machineOptions = pulumi.all([subnetDetail, amz2Ami]).apply(([subnet, ami]) => {
            return pulumi.all([
                pulumi.interpolate`amazonec2-region=ap-southeast-1`,
                pulumi.interpolate`amazonec2-vpc-id=${subnet.vpcId}`,
                pulumi.interpolate`amazonec2-subnet-id=${subnet.id}`,
                pulumi.interpolate`amazonec2-zone=${subnet.availabilityZone.slice(-1)}`,
                pulumi.interpolate`amazonec2-use-private-address=true`,
                pulumi.interpolate`amazonec2-tags=${Object.entries(args.machineTags ?? {}).join(",")}`,
                pulumi.interpolate`amazonec2-security-group=${gitlabRunnerSecGroup.name}`,
                pulumi.interpolate`amazonec2-instance-type=${args.machineInstanceType}`,
                pulumi.interpolate`amazonec2-ami=${ami.id}`,
                pulumi.interpolate`amazonec2-device-name=${ami.rootDeviceName}`,
                pulumi.interpolate`amazonec2-root-size=${machineRootDeviceDisk.size.toString()}`,
                pulumi.interpolate`amazonec2-volume-type=${machineRootDeviceDisk.type}`,
                pulumi.interpolate`amazonec2-security-group-readonly=true`,
                pulumi.interpolate`amazonec2-ssh-port=22`,
                pulumi.interpolate`amazonec2-ssh-user=ec2-user`,
            ])
        })
        const cacheParams: string[] = []
        const dockerVolumes: string[] = []

        const generalTemplateParams = pulumi.all([
            pulumi.interpolate`${args.gitlabUrl}`,
            pulumi.interpolate`${args.gitlabRunnerToken}`,
            pulumi.interpolate`${args.defaultDockerImage ?? "busybox"}`,
            pulumi.interpolate`${args.gitlabHelperImage ?? "gitlab/gitlab-runner-helper:x86_64-latest"}`,
            pulumi.interpolate`${args.outputLimit ?? 2048}`,
            pulumi.interpolate`${Object.entries(args.tags ?? {}).map(e => `${e[0]}=${e[1]}`).join(",")}`,
            pulumi.interpolate`${args.runnerMaxConcurrentBuild}`,
            pulumi.interpolate`${args.machineIdleNodes}`,
            pulumi.interpolate`${args.machineIdleTimeSecond}`,
            pulumi.interpolate`${args.machineMaxBuilds}`
        ]).apply(([gitlabUrl, gitlabRunnerToken, defaultDockerImage, gitlabHelperImage, outputLimit, tagList,
            runnerMaxConcurrentBuild, machineIdleNodes, machineIdleTimeSecond, machineMaxBuilds]) => {
            return {
                name: name,
                runnerMaxConcurrentBuild: runnerMaxConcurrentBuild,
                gitlabUrl: gitlabUrl,
                gitlabRunnerToken: gitlabRunnerToken,
                defaultDockerImage: defaultDockerImage,
                gitlabHelperImage: gitlabHelperImage,
                outputLimit: outputLimit,
                tagList: tagList,
                machineIdleNodes: machineIdleNodes,
                machineIdleTimeSecond: machineIdleTimeSecond,
                machineMaxBuilds: machineMaxBuilds
            }
        })

        const gitlabRunnerUserData = pulumi.all([generalTemplateParams, machineOptions]).apply(([general, machineOpts]) => {
            return `#!/bin/bash -xe
# Install Docker
sudo yum -y install docker
sudo service docker start

# Install Gitlab Runner
sudo curl -L --output /usr/bin/gitlab-runner "https://gitlab-runner-downloads.s3.amazonaws.com/latest/binaries/gitlab-runner-linux-amd64"
sudo chmod a+x /usr/bin/gitlab-runner
sudo useradd --comment 'GitLab Runner' --create-home gitlab-runner --shell /bin/bash || true
sudo gitlab-runner install --user=gitlab-runner --working-directory=/home/gitlab-runner
sudo gitlab-runner start

# Install Docker Machine
sudo curl -L --output /usr/bin/docker-machine "https://gitlab-docker-machine-downloads.s3.amazonaws.com/v0.16.2-gitlab.18/docker-machine-Linux-x86_64"
chmod +x /usr/bin/docker-machine

# Prepare config
cat <<EOF >/etc/gitlab-runner/config.toml
listen_address = ":9252"
concurrent = ${general.runnerMaxConcurrentBuild}
check_interval = 0
log_format = "json"

[session_server]
session_timeout = 1800
EOF

# Register runner
gitlab-runner register \\
--name "${name}" \\
--non-interactive \\
--url "${general.gitlabUrl}" \\
--registration-token "${general.gitlabRunnerToken}" \\
--executor "docker+machine" \\
--docker-image "${general.defaultDockerImage}" \\
--docker-helper-image "${general.gitlabHelperImage}" \\
${dockerVolumes.map(vol => `--docker-volumes "${vol}"`).join(" \\\n")} \\
--output-limit ${general.outputLimit} \\
--limit ${general.runnerMaxConcurrentBuild} \\
--tag-list "${general.tagList}" \\
--docker-privileged \\
${cacheParams.map(param => `--${param}`).join(" \\\n")} \\
--machine-idle-nodes ${general.machineIdleNodes} \\
--machine-idle-time ${general.machineIdleTimeSecond} \\
--machine-max-builds ${general.machineMaxBuilds} \\
--machine-machine-driver "amazonec2" \\
--machine-machine-name "${general.name}-%s" \\
${machineOpts.map(opt => `--machine-machine-options "${opt}"`).join(" \\\n")}`
        })


        const gitlabRunnerMachine = new aws.ec2.Instance(`${name}-orchestrator`, {
            ami: amz2Ami.then(amz2 => amz2.id),
            instanceType: "t3.micro",
            iamInstanceProfile: gitlabRunnerScalerInstanceProfile,
            subnetId: subnetDetail.apply(subnet => subnet.id),
            vpcSecurityGroupIds: [gitlabRunnerSecGroup.id],
            monitoring: true,
            userData: gitlabRunnerUserData,
            tags: {
                Name: `${name}-gitlab-runner-orchestrator`
            }
        }, { replaceOnChanges: ["userData"] })

        this.gitlabRunnerUserData = gitlabRunnerUserData

        this.registerOutputs({
            gitlabRunnerUserData
        });
    }
}