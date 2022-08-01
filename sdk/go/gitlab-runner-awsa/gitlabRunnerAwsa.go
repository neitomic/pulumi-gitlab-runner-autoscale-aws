// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlabrunnerawsa

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GitlabRunnerAwsa struct {
	pulumi.ResourceState

	// Gitlab Runner EC2 User Data
	GitlabRunnerUserData pulumi.StringOutput `pulumi:"gitlabRunnerUserData"`
}

// NewGitlabRunnerAwsa registers a new resource with the given unique name, arguments, and options.
func NewGitlabRunnerAwsa(ctx *pulumi.Context,
	name string, args *GitlabRunnerAwsaArgs, opts ...pulumi.ResourceOption) (*GitlabRunnerAwsa, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GitlabRunnerToken == nil {
		return nil, errors.New("invalid value for required argument 'GitlabRunnerToken'")
	}
	if args.GitlabUrl == nil {
		return nil, errors.New("invalid value for required argument 'GitlabUrl'")
	}
	if args.MachineIdleNodes == nil {
		return nil, errors.New("invalid value for required argument 'MachineIdleNodes'")
	}
	if args.MachineIdleTimeSecond == nil {
		return nil, errors.New("invalid value for required argument 'MachineIdleTimeSecond'")
	}
	if args.MachineInstanceType == nil {
		return nil, errors.New("invalid value for required argument 'MachineInstanceType'")
	}
	if args.MachineMaxBuilds == nil {
		return nil, errors.New("invalid value for required argument 'MachineMaxBuilds'")
	}
	if args.RunnerMaxConcurrentBuild == nil {
		return nil, errors.New("invalid value for required argument 'RunnerMaxConcurrentBuild'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource GitlabRunnerAwsa
	err := ctx.RegisterRemoteComponentResource("gitlab-runner-awsa:index:GitlabRunnerAwsa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type gitlabRunnerAwsaArgs struct {
	Cache              *S3Cache `pulumi:"cache"`
	DefaultDockerImage *string  `pulumi:"defaultDockerImage"`
	// Gitlab Runner helper image (default: gitlab-runner-helper:x86_64-latest).
	GitlabHelperImage *string `pulumi:"gitlabHelperImage"`
	// The Gitlab Runner token.
	GitlabRunnerToken string `pulumi:"gitlabRunnerToken"`
	// The Gitlab URL.
	GitlabUrl string `pulumi:"gitlabUrl"`
	// Number of iddle machines.
	MachineIdleNodes int `pulumi:"machineIdleNodes"`
	// Number of seconds a machine in idle mode after it's remove.
	MachineIdleTimeSecond int `pulumi:"machineIdleTimeSecond"`
	// Machine instance type.
	MachineInstanceType string `pulumi:"machineInstanceType"`
	// Maximum number of build a machine will run (machine is re-created after reach this number).
	MachineMaxBuilds int         `pulumi:"machineMaxBuilds"`
	MachineTags      *TagMap     `pulumi:"machineTags"`
	MachineVolume    *VolumeArgs `pulumi:"machineVolume"`
	// Max concurrent build.
	RunnerMaxConcurrentBuild int `pulumi:"runnerMaxConcurrentBuild"`
	// The Subnet ID.
	SubnetId string  `pulumi:"subnetId"`
	Tags     *TagMap `pulumi:"tags"`
}

// The set of arguments for constructing a GitlabRunnerAwsa resource.
type GitlabRunnerAwsaArgs struct {
	Cache              S3CachePtrInput
	DefaultDockerImage pulumi.StringPtrInput
	// Gitlab Runner helper image (default: gitlab-runner-helper:x86_64-latest).
	GitlabHelperImage pulumi.StringPtrInput
	// The Gitlab Runner token.
	GitlabRunnerToken pulumi.StringInput
	// The Gitlab URL.
	GitlabUrl pulumi.StringInput
	// Number of iddle machines.
	MachineIdleNodes pulumi.IntInput
	// Number of seconds a machine in idle mode after it's remove.
	MachineIdleTimeSecond pulumi.IntInput
	// Machine instance type.
	MachineInstanceType pulumi.StringInput
	// Maximum number of build a machine will run (machine is re-created after reach this number).
	MachineMaxBuilds pulumi.IntInput
	MachineTags      TagMapPtrInput
	MachineVolume    VolumeArgsPtrInput
	// Max concurrent build.
	RunnerMaxConcurrentBuild pulumi.IntInput
	// The Subnet ID.
	SubnetId pulumi.StringInput
	Tags     TagMapPtrInput
}

func (GitlabRunnerAwsaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitlabRunnerAwsaArgs)(nil)).Elem()
}

type GitlabRunnerAwsaInput interface {
	pulumi.Input

	ToGitlabRunnerAwsaOutput() GitlabRunnerAwsaOutput
	ToGitlabRunnerAwsaOutputWithContext(ctx context.Context) GitlabRunnerAwsaOutput
}

func (*GitlabRunnerAwsa) ElementType() reflect.Type {
	return reflect.TypeOf((**GitlabRunnerAwsa)(nil)).Elem()
}

func (i *GitlabRunnerAwsa) ToGitlabRunnerAwsaOutput() GitlabRunnerAwsaOutput {
	return i.ToGitlabRunnerAwsaOutputWithContext(context.Background())
}

func (i *GitlabRunnerAwsa) ToGitlabRunnerAwsaOutputWithContext(ctx context.Context) GitlabRunnerAwsaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitlabRunnerAwsaOutput)
}

// GitlabRunnerAwsaArrayInput is an input type that accepts GitlabRunnerAwsaArray and GitlabRunnerAwsaArrayOutput values.
// You can construct a concrete instance of `GitlabRunnerAwsaArrayInput` via:
//
//          GitlabRunnerAwsaArray{ GitlabRunnerAwsaArgs{...} }
type GitlabRunnerAwsaArrayInput interface {
	pulumi.Input

	ToGitlabRunnerAwsaArrayOutput() GitlabRunnerAwsaArrayOutput
	ToGitlabRunnerAwsaArrayOutputWithContext(context.Context) GitlabRunnerAwsaArrayOutput
}

type GitlabRunnerAwsaArray []GitlabRunnerAwsaInput

func (GitlabRunnerAwsaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitlabRunnerAwsa)(nil)).Elem()
}

func (i GitlabRunnerAwsaArray) ToGitlabRunnerAwsaArrayOutput() GitlabRunnerAwsaArrayOutput {
	return i.ToGitlabRunnerAwsaArrayOutputWithContext(context.Background())
}

func (i GitlabRunnerAwsaArray) ToGitlabRunnerAwsaArrayOutputWithContext(ctx context.Context) GitlabRunnerAwsaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitlabRunnerAwsaArrayOutput)
}

// GitlabRunnerAwsaMapInput is an input type that accepts GitlabRunnerAwsaMap and GitlabRunnerAwsaMapOutput values.
// You can construct a concrete instance of `GitlabRunnerAwsaMapInput` via:
//
//          GitlabRunnerAwsaMap{ "key": GitlabRunnerAwsaArgs{...} }
type GitlabRunnerAwsaMapInput interface {
	pulumi.Input

	ToGitlabRunnerAwsaMapOutput() GitlabRunnerAwsaMapOutput
	ToGitlabRunnerAwsaMapOutputWithContext(context.Context) GitlabRunnerAwsaMapOutput
}

type GitlabRunnerAwsaMap map[string]GitlabRunnerAwsaInput

func (GitlabRunnerAwsaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitlabRunnerAwsa)(nil)).Elem()
}

func (i GitlabRunnerAwsaMap) ToGitlabRunnerAwsaMapOutput() GitlabRunnerAwsaMapOutput {
	return i.ToGitlabRunnerAwsaMapOutputWithContext(context.Background())
}

func (i GitlabRunnerAwsaMap) ToGitlabRunnerAwsaMapOutputWithContext(ctx context.Context) GitlabRunnerAwsaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitlabRunnerAwsaMapOutput)
}

type GitlabRunnerAwsaOutput struct{ *pulumi.OutputState }

func (GitlabRunnerAwsaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitlabRunnerAwsa)(nil)).Elem()
}

func (o GitlabRunnerAwsaOutput) ToGitlabRunnerAwsaOutput() GitlabRunnerAwsaOutput {
	return o
}

func (o GitlabRunnerAwsaOutput) ToGitlabRunnerAwsaOutputWithContext(ctx context.Context) GitlabRunnerAwsaOutput {
	return o
}

type GitlabRunnerAwsaArrayOutput struct{ *pulumi.OutputState }

func (GitlabRunnerAwsaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitlabRunnerAwsa)(nil)).Elem()
}

func (o GitlabRunnerAwsaArrayOutput) ToGitlabRunnerAwsaArrayOutput() GitlabRunnerAwsaArrayOutput {
	return o
}

func (o GitlabRunnerAwsaArrayOutput) ToGitlabRunnerAwsaArrayOutputWithContext(ctx context.Context) GitlabRunnerAwsaArrayOutput {
	return o
}

func (o GitlabRunnerAwsaArrayOutput) Index(i pulumi.IntInput) GitlabRunnerAwsaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GitlabRunnerAwsa {
		return vs[0].([]*GitlabRunnerAwsa)[vs[1].(int)]
	}).(GitlabRunnerAwsaOutput)
}

type GitlabRunnerAwsaMapOutput struct{ *pulumi.OutputState }

func (GitlabRunnerAwsaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitlabRunnerAwsa)(nil)).Elem()
}

func (o GitlabRunnerAwsaMapOutput) ToGitlabRunnerAwsaMapOutput() GitlabRunnerAwsaMapOutput {
	return o
}

func (o GitlabRunnerAwsaMapOutput) ToGitlabRunnerAwsaMapOutputWithContext(ctx context.Context) GitlabRunnerAwsaMapOutput {
	return o
}

func (o GitlabRunnerAwsaMapOutput) MapIndex(k pulumi.StringInput) GitlabRunnerAwsaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GitlabRunnerAwsa {
		return vs[0].(map[string]*GitlabRunnerAwsa)[vs[1].(string)]
	}).(GitlabRunnerAwsaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GitlabRunnerAwsaInput)(nil)).Elem(), &GitlabRunnerAwsa{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitlabRunnerAwsaArrayInput)(nil)).Elem(), GitlabRunnerAwsaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitlabRunnerAwsaMapInput)(nil)).Elem(), GitlabRunnerAwsaMap{})
	pulumi.RegisterOutputType(GitlabRunnerAwsaOutput{})
	pulumi.RegisterOutputType(GitlabRunnerAwsaArrayOutput{})
	pulumi.RegisterOutputType(GitlabRunnerAwsaMapOutput{})
}
