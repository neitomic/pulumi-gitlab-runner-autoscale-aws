import * as gitlab-runner-awsa from "@pulumi/gitlab-runner-awsa";

const page = new gitlab-runner-awsa.StaticPage("page", {
    indexContent: "<html><body><p>Hello world!</p></body></html>",
});

export const bucket = page.bucket;
export const url = page.websiteUrl;
