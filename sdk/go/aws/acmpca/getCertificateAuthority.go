// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package acmpca

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Private certificate authority.
func LookupCertificateAuthority(ctx *pulumi.Context, args *LookupCertificateAuthorityArgs, opts ...pulumi.InvokeOption) (*LookupCertificateAuthorityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCertificateAuthorityResult
	err := ctx.Invoke("aws-native:acmpca:getCertificateAuthority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateAuthorityArgs struct {
	// The Amazon Resource Name (ARN) of the certificate authority.
	Arn string `pulumi:"arn"`
}

type LookupCertificateAuthorityResult struct {
	// The Amazon Resource Name (ARN) of the certificate authority.
	Arn *string `pulumi:"arn"`
	// The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
	CertificateSigningRequest *string `pulumi:"certificateSigningRequest"`
	// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
	RevocationConfiguration *CertificateAuthorityRevocationConfiguration `pulumi:"revocationConfiguration"`
	Tags                    []CertificateAuthorityTag                    `pulumi:"tags"`
}

func LookupCertificateAuthorityOutput(ctx *pulumi.Context, args LookupCertificateAuthorityOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateAuthorityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateAuthorityResult, error) {
			args := v.(LookupCertificateAuthorityArgs)
			r, err := LookupCertificateAuthority(ctx, &args, opts...)
			var s LookupCertificateAuthorityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateAuthorityResultOutput)
}

type LookupCertificateAuthorityOutputArgs struct {
	// The Amazon Resource Name (ARN) of the certificate authority.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupCertificateAuthorityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateAuthorityArgs)(nil)).Elem()
}

type LookupCertificateAuthorityResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateAuthorityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateAuthorityResult)(nil)).Elem()
}

func (o LookupCertificateAuthorityResultOutput) ToLookupCertificateAuthorityResultOutput() LookupCertificateAuthorityResultOutput {
	return o
}

func (o LookupCertificateAuthorityResultOutput) ToLookupCertificateAuthorityResultOutputWithContext(ctx context.Context) LookupCertificateAuthorityResultOutput {
	return o
}

func (o LookupCertificateAuthorityResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCertificateAuthorityResult] {
	return pulumix.Output[LookupCertificateAuthorityResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the certificate authority.
func (o LookupCertificateAuthorityResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
func (o LookupCertificateAuthorityResultOutput) CertificateSigningRequest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) *string { return v.CertificateSigningRequest }).(pulumi.StringPtrOutput)
}

// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
func (o LookupCertificateAuthorityResultOutput) RevocationConfiguration() CertificateAuthorityRevocationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) *CertificateAuthorityRevocationConfiguration {
		return v.RevocationConfiguration
	}).(CertificateAuthorityRevocationConfigurationPtrOutput)
}

func (o LookupCertificateAuthorityResultOutput) Tags() CertificateAuthorityTagArrayOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) []CertificateAuthorityTag { return v.Tags }).(CertificateAuthorityTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateAuthorityResultOutput{})
}
