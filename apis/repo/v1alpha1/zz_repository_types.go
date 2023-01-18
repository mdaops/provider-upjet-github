/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdvancedSecurityObservation struct {
}

type AdvancedSecurityParameters struct {

	// Set to enabled to enable secret scanning on the repository. Can be enabled or disabled.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type PagesObservation struct {

	// Whether the rendered GitHub Pages site has a custom 404 page.
	Custom404 *bool `json:"custom404,omitempty" tf:"custom_404,omitempty"`

	// URL to the repository on the web.
	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	// Set to enabled to enable secret scanning on the repository. Can be enabled or disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PagesParameters struct {

	// The custom domain for the repository. This can only be set after the repository has been created.
	// +kubebuilder:validation:Optional
	Cname *string `json:"cname,omitempty" tf:"cname,omitempty"`

	// The source branch and directory for the rendered Pages site. See GitHub Pages Source below for details.
	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`
}

type RepositoryObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A string of the form "orgname/reponame".
	FullName *string `json:"fullName,omitempty" tf:"full_name,omitempty"`

	// URL that can be provided to git clone to clone the repository anonymously via the git protocol.
	GitCloneURL *string `json:"gitCloneUrl,omitempty" tf:"git_clone_url,omitempty"`

	// URL to the repository on the web.
	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	// URL that can be provided to git clone to clone the repository via HTTPS.
	HTTPCloneURL *string `json:"httpCloneUrl,omitempty" tf:"http_clone_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// GraphQL global node id for use with v4 API
	NodeID *string `json:"nodeId,omitempty" tf:"node_id,omitempty"`

	// The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
	// +kubebuilder:validation:Optional
	Pages []PagesObservation `json:"pages,omitempty" tf:"pages,omitempty"`

	// GitHub ID for the repository
	RepoID *float64 `json:"repoId,omitempty" tf:"repo_id,omitempty"`

	// URL that can be provided to git clone to clone the repository via SSH.
	SSHCloneURL *string `json:"sshCloneUrl,omitempty" tf:"ssh_clone_url,omitempty"`

	// URL that can be provided to svn checkout to check out the repository via GitHub's Subversion protocol emulation.
	SvnURL *string `json:"svnUrl,omitempty" tf:"svn_url,omitempty"`
}

type RepositoryParameters struct {

	// Set to true to allow auto-merging pull requests on the repository.
	// +kubebuilder:validation:Optional
	AllowAutoMerge *bool `json:"allowAutoMerge,omitempty" tf:"allow_auto_merge,omitempty"`

	// Set to false to disable merge commits on the repository.
	// +kubebuilder:validation:Optional
	AllowMergeCommit *bool `json:"allowMergeCommit,omitempty" tf:"allow_merge_commit,omitempty"`

	// Set to false to disable rebase merges on the repository.
	// +kubebuilder:validation:Optional
	AllowRebaseMerge *bool `json:"allowRebaseMerge,omitempty" tf:"allow_rebase_merge,omitempty"`

	// Set to false to disable squash merges on the repository.
	// +kubebuilder:validation:Optional
	AllowSquashMerge *bool `json:"allowSquashMerge,omitempty" tf:"allow_squash_merge,omitempty"`

	// Set to true to always suggest updating pull request branches.
	// +kubebuilder:validation:Optional
	AllowUpdateBranch *bool `json:"allowUpdateBranch,omitempty" tf:"allow_update_branch,omitempty"`

	// Set to true to archive the repository instead of deleting on destroy.
	// +kubebuilder:validation:Optional
	ArchiveOnDestroy *bool `json:"archiveOnDestroy,omitempty" tf:"archive_on_destroy,omitempty"`

	// Specifies if the repository should be archived. Defaults to false. NOTE Currently, the API does not support unarchiving.
	// +kubebuilder:validation:Optional
	Archived *bool `json:"archived,omitempty" tf:"archived,omitempty"`

	// Set to true to produce an initial commit in the repository.
	// +kubebuilder:validation:Optional
	AutoInit *bool `json:"autoInit,omitempty" tf:"auto_init,omitempty"`

	// (Deprecated: Use github_branch_default resource instead) The name of the default branch of the repository. NOTE: This can only be set after a repository has already been created,
	// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
	// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
	// Can only be set after initial repository creation, and only if the target branch exists
	// +kubebuilder:validation:Optional
	DefaultBranch *string `json:"defaultBranch,omitempty" tf:"default_branch,omitempty"`

	// Automatically delete head branch after a pull request is merged. Defaults to false.
	// +kubebuilder:validation:Optional
	DeleteBranchOnMerge *bool `json:"deleteBranchOnMerge,omitempty" tf:"delete_branch_on_merge,omitempty"`

	// A description of the repository.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Use the name of the template without the extension. For example, "Haskell".
	// +kubebuilder:validation:Optional
	GitignoreTemplate *string `json:"gitignoreTemplate,omitempty" tf:"gitignore_template,omitempty"`

	// Set to true to enable the (deprecated) downloads features on the repository.
	// +kubebuilder:validation:Optional
	HasDownloads *bool `json:"hasDownloads,omitempty" tf:"has_downloads,omitempty"`

	// Set to true to enable the GitHub Issues features
	// on the repository.
	// +kubebuilder:validation:Optional
	HasIssues *bool `json:"hasIssues,omitempty" tf:"has_issues,omitempty"`

	// Set to true to enable the GitHub Projects features on the repository. Per the GitHub documentation when in an organization that has disabled repository projects it will default to false and will otherwise default to true. If you specify true when it has been disabled it will return an error.
	// +kubebuilder:validation:Optional
	HasProjects *bool `json:"hasProjects,omitempty" tf:"has_projects,omitempty"`

	// Set to true to enable the GitHub Wiki features on
	// the repository.
	// +kubebuilder:validation:Optional
	HasWiki *bool `json:"hasWiki,omitempty" tf:"has_wiki,omitempty"`

	// URL of a page describing the project.
	// +kubebuilder:validation:Optional
	HomepageURL *string `json:"homepageUrl,omitempty" tf:"homepage_url,omitempty"`

	// Set to true to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
	// +kubebuilder:validation:Optional
	IgnoreVulnerabilityAlertsDuringRead *bool `json:"ignoreVulnerabilityAlertsDuringRead,omitempty" tf:"ignore_vulnerability_alerts_during_read,omitempty"`

	// Set to true to tell GitHub that this is a template repository.
	// +kubebuilder:validation:Optional
	IsTemplate *bool `json:"isTemplate,omitempty" tf:"is_template,omitempty"`

	// Use the name of the template without the extension. For example, "mit" or "mpl-2.0".
	// +kubebuilder:validation:Optional
	LicenseTemplate *string `json:"licenseTemplate,omitempty" tf:"license_template,omitempty"`

	// Can be PR_BODY, PR_TITLE, or BLANK for a default merge commit message.
	// +kubebuilder:validation:Optional
	MergeCommitMessage *string `json:"mergeCommitMessage,omitempty" tf:"merge_commit_message,omitempty"`

	// Can be PR_TITLE or MERGE_MESSAGE for a default merge commit title.
	// +kubebuilder:validation:Optional
	MergeCommitTitle *string `json:"mergeCommitTitle,omitempty" tf:"merge_commit_title,omitempty"`

	// The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
	// +kubebuilder:validation:Optional
	Pages []PagesParameters `json:"pages,omitempty" tf:"pages,omitempty"`

	// Set to true to create a private repository.
	// Repositories are created as public (e.g. open source) by default.
	// +kubebuilder:validation:Optional
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// The repository's security and analysis configuration. See Security and Analysis Configuration below for details.
	// Security and analysis settings for the repository. To use this parameter you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository.
	// +kubebuilder:validation:Optional
	SecurityAndAnalysis []SecurityAndAnalysisParameters `json:"securityAndAnalysis,omitempty" tf:"security_and_analysis,omitempty"`

	// Can be PR_BODY, COMMIT_MESSAGES, or BLANK for a default squash merge commit message.
	// +kubebuilder:validation:Optional
	SquashMergeCommitMessage *string `json:"squashMergeCommitMessage,omitempty" tf:"squash_merge_commit_message,omitempty"`

	// Can be PR_TITLE or COMMIT_OR_PR_TITLE for a default squash merge commit title.
	// +kubebuilder:validation:Optional
	SquashMergeCommitTitle *string `json:"squashMergeCommitTitle,omitempty" tf:"squash_merge_commit_title,omitempty"`

	// Use a template repository to create this resource. See Template Repositories below for details.
	// +kubebuilder:validation:Optional
	Template []TemplateParameters `json:"template,omitempty" tf:"template,omitempty"`

	// The list of topics of the repository.
	// +kubebuilder:validation:Optional
	Topics []*string `json:"topics,omitempty" tf:"topics,omitempty"`

	// Can be public or private. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be internal. The visibility parameter overrides the private parameter.
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// Set to true to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See GitHub Documentation for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
	// +kubebuilder:validation:Optional
	VulnerabilityAlerts *bool `json:"vulnerabilityAlerts,omitempty" tf:"vulnerability_alerts,omitempty"`
}

type SecretScanningObservation struct {
}

type SecretScanningParameters struct {

	// Set to enabled to enable secret scanning on the repository. Can be enabled or disabled.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type SecretScanningPushProtectionObservation struct {
}

type SecretScanningPushProtectionParameters struct {

	// Set to enabled to enable secret scanning on the repository. Can be enabled or disabled.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type SecurityAndAnalysisObservation struct {
}

type SecurityAndAnalysisParameters struct {

	// The advanced security configuration for the repository. See Advanced Security Configuration below for details.
	// +kubebuilder:validation:Required
	AdvancedSecurity []AdvancedSecurityParameters `json:"advancedSecurity" tf:"advanced_security,omitempty"`

	// The secret scanning configuration for the repository. See Secret Scanning Configuration below for details.
	// +kubebuilder:validation:Required
	SecretScanning []SecretScanningParameters `json:"secretScanning" tf:"secret_scanning,omitempty"`

	// The secret scanning push protection configuration for the repository. See Secret Scanning Push Protection Configuration below for details.
	// +kubebuilder:validation:Required
	SecretScanningPushProtection []SecretScanningPushProtectionParameters `json:"secretScanningPushProtection" tf:"secret_scanning_push_protection,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// The repository branch used to publish the site's source files. (i.e. main or gh-pages.
	// +kubebuilder:validation:Required
	Branch *string `json:"branch" tf:"branch,omitempty"`

	// The repository directory from which the site publishes (Default: /).
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type TemplateObservation struct {
}

type TemplateParameters struct {

	// : Whether the new repository should include all the branches from the template repository (defaults to false, which includes only the default branch from the template).
	// +kubebuilder:validation:Optional
	IncludeAllBranches *bool `json:"includeAllBranches,omitempty" tf:"include_all_branches,omitempty"`

	// : The GitHub organization or user the template repository is owned by.
	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// : The name of the template repository.
	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Repository is the Schema for the Repositorys API. Creates and manages repositories within GitHub organizations or personal accounts
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
