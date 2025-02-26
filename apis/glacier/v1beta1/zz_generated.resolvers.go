/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Vault.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Vault) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Notification); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Notification[i3].SnsTopic),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Notification[i3].SnsTopicRef,
				Selector:     mg.Spec.ForProvider.Notification[i3].SnsTopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Notification[i3].SnsTopic")
		}
		mg.Spec.ForProvider.Notification[i3].SnsTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Notification[i3].SnsTopicRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Notification); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Notification[i3].SnsTopic),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Notification[i3].SnsTopicRef,
				Selector:     mg.Spec.InitProvider.Notification[i3].SnsTopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Notification[i3].SnsTopic")
		}
		mg.Spec.InitProvider.Notification[i3].SnsTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Notification[i3].SnsTopicRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this VaultLock.
func (mg *VaultLock) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("glacier.aws.upbound.io", "v1beta1", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VaultNameRef,
			Selector:     mg.Spec.ForProvider.VaultNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultName")
	}
	mg.Spec.ForProvider.VaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("glacier.aws.upbound.io", "v1beta1", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.VaultNameRef,
			Selector:     mg.Spec.InitProvider.VaultNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VaultName")
	}
	mg.Spec.InitProvider.VaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VaultNameRef = rsp.ResolvedReference

	return nil
}
