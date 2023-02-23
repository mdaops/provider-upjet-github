/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Branch.
func (mg *Branch) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BranchProtection.
func (mg *BranchProtection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RepositoryID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryIDRef,
		Selector:     mg.Spec.ForProvider.RepositoryIDSelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RepositoryID")
	}
	mg.Spec.ForProvider.RepositoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DefaultBranch.
func (mg *DefaultBranch) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Branch),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BranchRef,
		Selector:     mg.Spec.ForProvider.BranchSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Branch")
	}
	mg.Spec.ForProvider.Branch = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BranchRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DeployKey.
func (mg *DeployKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PullRequest.
func (mg *PullRequest) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BaseRepository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BaseRepositoryRef,
		Selector:     mg.Spec.ForProvider.BaseRepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BaseRepository")
	}
	mg.Spec.ForProvider.BaseRepository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BaseRepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HeadRef),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HeadRefRef,
		Selector:     mg.Spec.ForProvider.HeadRefSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HeadRef")
	}
	mg.Spec.ForProvider.HeadRef = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HeadRefRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RepositoryFile.
func (mg *RepositoryFile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Branch),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BranchRef,
		Selector:     mg.Spec.ForProvider.BranchSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Branch")
	}
	mg.Spec.ForProvider.Branch = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BranchRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}
