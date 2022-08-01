import * as gitlabRunnerAwsa  from "@pulumi/gitlab-runner-awsa";

const page = new gitlabRunnerAwsa.GitlabRunnerAwsa("test-runner-awsa", {
    gitlabUrl: "https://gitlab.com",
    gitlabRunnerToken: "",
    machineIdleNodes: 0,
    machineIdleTimeSecond: 1800,
    machineInstanceType: "t3.2xlarge",
    machineMaxBuilds: 20,
    runnerMaxConcurrentBuild: 10,
    subnetId: "subnet-xxxxxxx"
});

export const gitlabRunnerUserData = page.gitlabRunnerUserData;
