// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlabrunnerawsa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type S3Cache struct {
	BucketName    *string `pulumi:"bucketName"`
	Path          *int    `pulumi:"path"`
	ServerAddress *string `pulumi:"serverAddress"`
}

// S3CacheInput is an input type that accepts S3CacheArgs and S3CacheOutput values.
// You can construct a concrete instance of `S3CacheInput` via:
//
//          S3CacheArgs{...}
type S3CacheInput interface {
	pulumi.Input

	ToS3CacheOutput() S3CacheOutput
	ToS3CacheOutputWithContext(context.Context) S3CacheOutput
}

type S3CacheArgs struct {
	BucketName    pulumi.StringPtrInput `pulumi:"bucketName"`
	Path          pulumi.IntPtrInput    `pulumi:"path"`
	ServerAddress pulumi.StringPtrInput `pulumi:"serverAddress"`
}

func (S3CacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*S3Cache)(nil)).Elem()
}

func (i S3CacheArgs) ToS3CacheOutput() S3CacheOutput {
	return i.ToS3CacheOutputWithContext(context.Background())
}

func (i S3CacheArgs) ToS3CacheOutputWithContext(ctx context.Context) S3CacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3CacheOutput)
}

func (i S3CacheArgs) ToS3CachePtrOutput() S3CachePtrOutput {
	return i.ToS3CachePtrOutputWithContext(context.Background())
}

func (i S3CacheArgs) ToS3CachePtrOutputWithContext(ctx context.Context) S3CachePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3CacheOutput).ToS3CachePtrOutputWithContext(ctx)
}

// S3CachePtrInput is an input type that accepts S3CacheArgs, S3CachePtr and S3CachePtrOutput values.
// You can construct a concrete instance of `S3CachePtrInput` via:
//
//          S3CacheArgs{...}
//
//  or:
//
//          nil
type S3CachePtrInput interface {
	pulumi.Input

	ToS3CachePtrOutput() S3CachePtrOutput
	ToS3CachePtrOutputWithContext(context.Context) S3CachePtrOutput
}

type s3cachePtrType S3CacheArgs

func S3CachePtr(v *S3CacheArgs) S3CachePtrInput {
	return (*s3cachePtrType)(v)
}

func (*s3cachePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**S3Cache)(nil)).Elem()
}

func (i *s3cachePtrType) ToS3CachePtrOutput() S3CachePtrOutput {
	return i.ToS3CachePtrOutputWithContext(context.Background())
}

func (i *s3cachePtrType) ToS3CachePtrOutputWithContext(ctx context.Context) S3CachePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3CachePtrOutput)
}

type S3CacheOutput struct{ *pulumi.OutputState }

func (S3CacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*S3Cache)(nil)).Elem()
}

func (o S3CacheOutput) ToS3CacheOutput() S3CacheOutput {
	return o
}

func (o S3CacheOutput) ToS3CacheOutputWithContext(ctx context.Context) S3CacheOutput {
	return o
}

func (o S3CacheOutput) ToS3CachePtrOutput() S3CachePtrOutput {
	return o.ToS3CachePtrOutputWithContext(context.Background())
}

func (o S3CacheOutput) ToS3CachePtrOutputWithContext(ctx context.Context) S3CachePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v S3Cache) *S3Cache {
		return &v
	}).(S3CachePtrOutput)
}

func (o S3CacheOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v S3Cache) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o S3CacheOutput) Path() pulumi.IntPtrOutput {
	return o.ApplyT(func(v S3Cache) *int { return v.Path }).(pulumi.IntPtrOutput)
}

func (o S3CacheOutput) ServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v S3Cache) *string { return v.ServerAddress }).(pulumi.StringPtrOutput)
}

type S3CachePtrOutput struct{ *pulumi.OutputState }

func (S3CachePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3Cache)(nil)).Elem()
}

func (o S3CachePtrOutput) ToS3CachePtrOutput() S3CachePtrOutput {
	return o
}

func (o S3CachePtrOutput) ToS3CachePtrOutputWithContext(ctx context.Context) S3CachePtrOutput {
	return o
}

func (o S3CachePtrOutput) Elem() S3CacheOutput {
	return o.ApplyT(func(v *S3Cache) S3Cache {
		if v != nil {
			return *v
		}
		var ret S3Cache
		return ret
	}).(S3CacheOutput)
}

func (o S3CachePtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *S3Cache) *string {
		if v == nil {
			return nil
		}
		return v.BucketName
	}).(pulumi.StringPtrOutput)
}

func (o S3CachePtrOutput) Path() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *S3Cache) *int {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.IntPtrOutput)
}

func (o S3CachePtrOutput) ServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *S3Cache) *string {
		if v == nil {
			return nil
		}
		return v.ServerAddress
	}).(pulumi.StringPtrOutput)
}

type TagMap struct {
}

// TagMapInput is an input type that accepts TagMap and TagMapOutput values.
// You can construct a concrete instance of `TagMapInput` via:
//
//          TagMap{ "key": TagArgs{...} }
type TagMapInput interface {
	pulumi.Input

	ToTagMapOutput() TagMapOutput
	ToTagMapOutputWithContext(context.Context) TagMapOutput
}

type TagMapArgs struct {
}

func (TagMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagMap)(nil)).Elem()
}

func (i TagMapArgs) ToTagMapOutput() TagMapOutput {
	return i.ToTagMapOutputWithContext(context.Background())
}

func (i TagMapArgs) ToTagMapOutputWithContext(ctx context.Context) TagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagMapOutput)
}

func (i TagMapArgs) ToTagMapPtrOutput() TagMapPtrOutput {
	return i.ToTagMapPtrOutputWithContext(context.Background())
}

func (i TagMapArgs) ToTagMapPtrOutputWithContext(ctx context.Context) TagMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagMapOutput).ToTagMapPtrOutputWithContext(ctx)
}

// TagMapPtrInput is an input type that accepts TagMapArgs, TagMapPtr and TagMapPtrOutput values.
// You can construct a concrete instance of `TagMapPtrInput` via:
//
//          TagMapArgs{...}
//
//  or:
//
//          nil
type TagMapPtrInput interface {
	pulumi.Input

	ToTagMapPtrOutput() TagMapPtrOutput
	ToTagMapPtrOutputWithContext(context.Context) TagMapPtrOutput
}

type tagMapPtrType TagMapArgs

func TagMapPtr(v *TagMapArgs) TagMapPtrInput {
	return (*tagMapPtrType)(v)
}

func (*tagMapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagMap)(nil)).Elem()
}

func (i *tagMapPtrType) ToTagMapPtrOutput() TagMapPtrOutput {
	return i.ToTagMapPtrOutputWithContext(context.Background())
}

func (i *tagMapPtrType) ToTagMapPtrOutputWithContext(ctx context.Context) TagMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagMapPtrOutput)
}

type TagMapOutput struct{ *pulumi.OutputState }

func (TagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagMap)(nil)).Elem()
}

func (o TagMapOutput) ToTagMapOutput() TagMapOutput {
	return o
}

func (o TagMapOutput) ToTagMapOutputWithContext(ctx context.Context) TagMapOutput {
	return o
}

func (o TagMapOutput) ToTagMapPtrOutput() TagMapPtrOutput {
	return o.ToTagMapPtrOutputWithContext(context.Background())
}

func (o TagMapOutput) ToTagMapPtrOutputWithContext(ctx context.Context) TagMapPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagMap) *TagMap {
		return &v
	}).(TagMapPtrOutput)
}

type TagMapPtrOutput struct{ *pulumi.OutputState }

func (TagMapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagMap)(nil)).Elem()
}

func (o TagMapPtrOutput) ToTagMapPtrOutput() TagMapPtrOutput {
	return o
}

func (o TagMapPtrOutput) ToTagMapPtrOutputWithContext(ctx context.Context) TagMapPtrOutput {
	return o
}

func (o TagMapPtrOutput) Elem() TagMapOutput {
	return o.ApplyT(func(v *TagMap) TagMap {
		if v != nil {
			return *v
		}
		var ret TagMap
		return ret
	}).(TagMapOutput)
}

type VolumeArgs struct {
	Size *int    `pulumi:"size"`
	Type *string `pulumi:"type"`
}

// VolumeArgsInput is an input type that accepts VolumeArgsArgs and VolumeArgsOutput values.
// You can construct a concrete instance of `VolumeArgsInput` via:
//
//          VolumeArgsArgs{...}
type VolumeArgsInput interface {
	pulumi.Input

	ToVolumeArgsOutput() VolumeArgsOutput
	ToVolumeArgsOutputWithContext(context.Context) VolumeArgsOutput
}

type VolumeArgsArgs struct {
	Size pulumi.IntPtrInput    `pulumi:"size"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (VolumeArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeArgs)(nil)).Elem()
}

func (i VolumeArgsArgs) ToVolumeArgsOutput() VolumeArgsOutput {
	return i.ToVolumeArgsOutputWithContext(context.Background())
}

func (i VolumeArgsArgs) ToVolumeArgsOutputWithContext(ctx context.Context) VolumeArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArgsOutput)
}

func (i VolumeArgsArgs) ToVolumeArgsPtrOutput() VolumeArgsPtrOutput {
	return i.ToVolumeArgsPtrOutputWithContext(context.Background())
}

func (i VolumeArgsArgs) ToVolumeArgsPtrOutputWithContext(ctx context.Context) VolumeArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArgsOutput).ToVolumeArgsPtrOutputWithContext(ctx)
}

// VolumeArgsPtrInput is an input type that accepts VolumeArgsArgs, VolumeArgsPtr and VolumeArgsPtrOutput values.
// You can construct a concrete instance of `VolumeArgsPtrInput` via:
//
//          VolumeArgsArgs{...}
//
//  or:
//
//          nil
type VolumeArgsPtrInput interface {
	pulumi.Input

	ToVolumeArgsPtrOutput() VolumeArgsPtrOutput
	ToVolumeArgsPtrOutputWithContext(context.Context) VolumeArgsPtrOutput
}

type volumeArgsPtrType VolumeArgsArgs

func VolumeArgsPtr(v *VolumeArgsArgs) VolumeArgsPtrInput {
	return (*volumeArgsPtrType)(v)
}

func (*volumeArgsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeArgs)(nil)).Elem()
}

func (i *volumeArgsPtrType) ToVolumeArgsPtrOutput() VolumeArgsPtrOutput {
	return i.ToVolumeArgsPtrOutputWithContext(context.Background())
}

func (i *volumeArgsPtrType) ToVolumeArgsPtrOutputWithContext(ctx context.Context) VolumeArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArgsPtrOutput)
}

type VolumeArgsOutput struct{ *pulumi.OutputState }

func (VolumeArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeArgs)(nil)).Elem()
}

func (o VolumeArgsOutput) ToVolumeArgsOutput() VolumeArgsOutput {
	return o
}

func (o VolumeArgsOutput) ToVolumeArgsOutputWithContext(ctx context.Context) VolumeArgsOutput {
	return o
}

func (o VolumeArgsOutput) ToVolumeArgsPtrOutput() VolumeArgsPtrOutput {
	return o.ToVolumeArgsPtrOutputWithContext(context.Background())
}

func (o VolumeArgsOutput) ToVolumeArgsPtrOutputWithContext(ctx context.Context) VolumeArgsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeArgs) *VolumeArgs {
		return &v
	}).(VolumeArgsPtrOutput)
}

func (o VolumeArgsOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VolumeArgs) *int { return v.Size }).(pulumi.IntPtrOutput)
}

func (o VolumeArgsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeArgs) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VolumeArgsPtrOutput struct{ *pulumi.OutputState }

func (VolumeArgsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeArgs)(nil)).Elem()
}

func (o VolumeArgsPtrOutput) ToVolumeArgsPtrOutput() VolumeArgsPtrOutput {
	return o
}

func (o VolumeArgsPtrOutput) ToVolumeArgsPtrOutputWithContext(ctx context.Context) VolumeArgsPtrOutput {
	return o
}

func (o VolumeArgsPtrOutput) Elem() VolumeArgsOutput {
	return o.ApplyT(func(v *VolumeArgs) VolumeArgs {
		if v != nil {
			return *v
		}
		var ret VolumeArgs
		return ret
	}).(VolumeArgsOutput)
}

func (o VolumeArgsPtrOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VolumeArgs) *int {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.IntPtrOutput)
}

func (o VolumeArgsPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeArgs) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*S3CacheInput)(nil)).Elem(), S3CacheArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3CachePtrInput)(nil)).Elem(), S3CacheArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagMapInput)(nil)).Elem(), TagMapArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagMapPtrInput)(nil)).Elem(), TagMapArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeArgsInput)(nil)).Elem(), VolumeArgsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeArgsPtrInput)(nil)).Elem(), VolumeArgsArgs{})
	pulumi.RegisterOutputType(S3CacheOutput{})
	pulumi.RegisterOutputType(S3CachePtrOutput{})
	pulumi.RegisterOutputType(TagMapOutput{})
	pulumi.RegisterOutputType(TagMapPtrOutput{})
	pulumi.RegisterOutputType(VolumeArgsOutput{})
	pulumi.RegisterOutputType(VolumeArgsPtrOutput{})
}
