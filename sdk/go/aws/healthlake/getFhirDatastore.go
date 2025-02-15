// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthlake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// HealthLake FHIR Datastore
func LookupFhirDatastore(ctx *pulumi.Context, args *LookupFhirDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupFhirDatastoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFhirDatastoreResult
	err := ctx.Invoke("aws-native:healthlake:getFhirDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFhirDatastoreArgs struct {
	DatastoreId string `pulumi:"datastoreId"`
}

type LookupFhirDatastoreResult struct {
	CreatedAt         *FhirDatastoreCreatedAt       `pulumi:"createdAt"`
	DatastoreArn      *string                       `pulumi:"datastoreArn"`
	DatastoreEndpoint *string                       `pulumi:"datastoreEndpoint"`
	DatastoreId       *string                       `pulumi:"datastoreId"`
	DatastoreStatus   *FhirDatastoreDatastoreStatus `pulumi:"datastoreStatus"`
	Tags              []FhirDatastoreTag            `pulumi:"tags"`
}

func LookupFhirDatastoreOutput(ctx *pulumi.Context, args LookupFhirDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupFhirDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFhirDatastoreResult, error) {
			args := v.(LookupFhirDatastoreArgs)
			r, err := LookupFhirDatastore(ctx, &args, opts...)
			var s LookupFhirDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFhirDatastoreResultOutput)
}

type LookupFhirDatastoreOutputArgs struct {
	DatastoreId pulumi.StringInput `pulumi:"datastoreId"`
}

func (LookupFhirDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFhirDatastoreArgs)(nil)).Elem()
}

type LookupFhirDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupFhirDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFhirDatastoreResult)(nil)).Elem()
}

func (o LookupFhirDatastoreResultOutput) ToLookupFhirDatastoreResultOutput() LookupFhirDatastoreResultOutput {
	return o
}

func (o LookupFhirDatastoreResultOutput) ToLookupFhirDatastoreResultOutputWithContext(ctx context.Context) LookupFhirDatastoreResultOutput {
	return o
}

func (o LookupFhirDatastoreResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFhirDatastoreResult] {
	return pulumix.Output[LookupFhirDatastoreResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFhirDatastoreResultOutput) CreatedAt() FhirDatastoreCreatedAtPtrOutput {
	return o.ApplyT(func(v LookupFhirDatastoreResult) *FhirDatastoreCreatedAt { return v.CreatedAt }).(FhirDatastoreCreatedAtPtrOutput)
}

func (o LookupFhirDatastoreResultOutput) DatastoreArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirDatastoreResult) *string { return v.DatastoreArn }).(pulumi.StringPtrOutput)
}

func (o LookupFhirDatastoreResultOutput) DatastoreEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirDatastoreResult) *string { return v.DatastoreEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupFhirDatastoreResultOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirDatastoreResult) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o LookupFhirDatastoreResultOutput) DatastoreStatus() FhirDatastoreDatastoreStatusPtrOutput {
	return o.ApplyT(func(v LookupFhirDatastoreResult) *FhirDatastoreDatastoreStatus { return v.DatastoreStatus }).(FhirDatastoreDatastoreStatusPtrOutput)
}

func (o LookupFhirDatastoreResultOutput) Tags() FhirDatastoreTagArrayOutput {
	return o.ApplyT(func(v LookupFhirDatastoreResult) []FhirDatastoreTag { return v.Tags }).(FhirDatastoreTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFhirDatastoreResultOutput{})
}
