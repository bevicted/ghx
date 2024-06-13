package ghx

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-github/v62/github"
)

type Client struct {
	*github.Client

	Actions            *ActionsService
	Activity           *ActivityService
	Admin              *AdminService
	Apps               *AppsService
	Authorizations     *AuthorizationsService
	Billing            *BillingService
	Checks             *ChecksService
	CodeScanning       *CodeScanningService
	CodesOfConduct     *CodesOfConductService
	Codespaces         *CodespacesService
	Copilot            *CopilotService
	Dependabot         *DependabotService
	DependencyGraph    *DependencyGraphService
	Emojis             *EmojisService
	Enterprise         *EnterpriseService
	Gists              *GistsService
	Git                *GitService
	Gitignores         *GitignoresService
	Interactions       *InteractionsService
	IssueImport        *IssueImportService
	Issues             *IssuesService
	Licenses           *LicensesService
	Markdown           *MarkdownService
	Marketplace        *MarketplaceService
	Meta               *MetaService
	Migrations         *MigrationService
	Organizations      *OrganizationsService
	Projects           *ProjectsService
	PullRequests       *PullRequestsService
	RateLimit          *RateLimitService
	Reactions          *ReactionsService
	Repositories       *RepositoriesService
	SCIM               *SCIMService
	Search             *SearchService
	SecretScanning     *SecretScanningService
	SecurityAdvisories *SecurityAdvisoriesService
	Teams              *TeamsService
	Users              *UsersService
}

type ActionsService struct {
	AddEnabledOrgInEnterprise                    func(ctx context.Context, owner string, organizationID int64) (*github.Response, error)
	AddEnabledReposInOrg                         func(ctx context.Context, owner string, repositoryID int64) (*github.Response, error)
	AddRepoToRequiredWorkflow                    func(ctx context.Context, org string, requiredWorkflowID int64, repoID int64) (*github.Response, error)
	AddRepositoryAccessRunnerGroup               func(ctx context.Context, org string, groupID int64, repoID int64) (*github.Response, error)
	AddRunnerGroupRunners                        func(ctx context.Context, org string, groupID int64, runnerID int64) (*github.Response, error)
	AddSelectedRepoToOrgSecret                   func(ctx context.Context, org string, name string, repo *github.Repository) (*github.Response, error)
	AddSelectedRepoToOrgVariable                 func(ctx context.Context, org string, name string, repo *github.Repository) (*github.Response, error)
	CancelWorkflowRunByID                        func(ctx context.Context, owner string, repo string, runID int64) (*github.Response, error)
	CreateEnvVariable                            func(ctx context.Context, owner string, repo string, env string, variable *github.ActionsVariable) (*github.Response, error)
	CreateOrUpdateEnvSecret                      func(ctx context.Context, repoID int, env string, eSecret *github.EncryptedSecret) (*github.Response, error)
	CreateOrUpdateOrgSecret                      func(ctx context.Context, org string, eSecret *github.EncryptedSecret) (*github.Response, error)
	CreateOrUpdateRepoSecret                     func(ctx context.Context, owner string, repo string, eSecret *github.EncryptedSecret) (*github.Response, error)
	CreateOrgVariable                            func(ctx context.Context, org string, variable *github.ActionsVariable) (*github.Response, error)
	CreateOrganizationRegistrationToken          func(ctx context.Context, org string) (*github.RegistrationToken, *github.Response, error)
	CreateOrganizationRemoveToken                func(ctx context.Context, org string) (*github.RemoveToken, *github.Response, error)
	CreateOrganizationRunnerGroup                func(ctx context.Context, org string, createReq github.CreateRunnerGroupRequest) (*github.RunnerGroup, *github.Response, error)
	CreateRegistrationToken                      func(ctx context.Context, owner string, repo string) (*github.RegistrationToken, *github.Response, error)
	CreateRemoveToken                            func(ctx context.Context, owner string, repo string) (*github.RemoveToken, *github.Response, error)
	CreateRepoVariable                           func(ctx context.Context, owner string, repo string, variable *github.ActionsVariable) (*github.Response, error)
	CreateRequiredWorkflow                       func(ctx context.Context, org string, createRequiredWorkflowOptions *github.CreateUpdateRequiredWorkflowOptions) (*github.OrgRequiredWorkflow, *github.Response, error)
	CreateWorkflowDispatchEventByFileName        func(ctx context.Context, owner string, repo string, workflowFileName string, event github.CreateWorkflowDispatchEventRequest) (*github.Response, error)
	CreateWorkflowDispatchEventByID              func(ctx context.Context, owner string, repo string, workflowID int64, event github.CreateWorkflowDispatchEventRequest) (*github.Response, error)
	DeleteArtifact                               func(ctx context.Context, owner string, repo string, artifactID int64) (*github.Response, error)
	DeleteCachesByID                             func(ctx context.Context, owner string, repo string, cacheID int64) (*github.Response, error)
	DeleteCachesByKey                            func(ctx context.Context, owner string, repo string, key string, ref *string) (*github.Response, error)
	DeleteEnvSecret                              func(ctx context.Context, repoID int, env string, secretName string) (*github.Response, error)
	DeleteEnvVariable                            func(ctx context.Context, owner string, repo string, env string, variableName string) (*github.Response, error)
	DeleteOrgSecret                              func(ctx context.Context, org string, name string) (*github.Response, error)
	DeleteOrgVariable                            func(ctx context.Context, org string, name string) (*github.Response, error)
	DeleteOrganizationRunnerGroup                func(ctx context.Context, org string, groupID int64) (*github.Response, error)
	DeleteRepoSecret                             func(ctx context.Context, owner string, repo string, name string) (*github.Response, error)
	DeleteRepoVariable                           func(ctx context.Context, owner string, repo string, name string) (*github.Response, error)
	DeleteRequiredWorkflow                       func(ctx context.Context, org string, requiredWorkflowID int64) (*github.Response, error)
	DeleteWorkflowRun                            func(ctx context.Context, owner string, repo string, runID int64) (*github.Response, error)
	DeleteWorkflowRunLogs                        func(ctx context.Context, owner string, repo string, runID int64) (*github.Response, error)
	DisableWorkflowByFileName                    func(ctx context.Context, owner string, repo string, workflowFileName string) (*github.Response, error)
	DisableWorkflowByID                          func(ctx context.Context, owner string, repo string, workflowID int64) (*github.Response, error)
	DownloadArtifact                             func(ctx context.Context, owner string, repo string, artifactID int64, maxRedirects int) (*url.URL, *github.Response, error)
	EditActionsAllowed                           func(ctx context.Context, org string, actionsAllowed github.ActionsAllowed) (*github.ActionsAllowed, *github.Response, error)
	EditActionsAllowedInEnterprise               func(ctx context.Context, enterprise string, actionsAllowed github.ActionsAllowed) (*github.ActionsAllowed, *github.Response, error)
	EditActionsPermissions                       func(ctx context.Context, org string, actionsPermissions github.ActionsPermissions) (*github.ActionsPermissions, *github.Response, error)
	EditActionsPermissionsInEnterprise           func(ctx context.Context, enterprise string, actionsPermissionsEnterprise github.ActionsPermissionsEnterprise) (*github.ActionsPermissionsEnterprise, *github.Response, error)
	EditDefaultWorkflowPermissionsInEnterprise   func(ctx context.Context, enterprise string, permissions github.DefaultWorkflowPermissionEnterprise) (*github.DefaultWorkflowPermissionEnterprise, *github.Response, error)
	EditDefaultWorkflowPermissionsInOrganization func(ctx context.Context, org string, permissions github.DefaultWorkflowPermissionOrganization) (*github.DefaultWorkflowPermissionOrganization, *github.Response, error)
	EnableWorkflowByFileName                     func(ctx context.Context, owner string, repo string, workflowFileName string) (*github.Response, error)
	EnableWorkflowByID                           func(ctx context.Context, owner string, repo string, workflowID int64) (*github.Response, error)
	GenerateOrgJITConfig                         func(ctx context.Context, org string, request *github.GenerateJITConfigRequest) (*github.JITRunnerConfig, *github.Response, error)
	GenerateRepoJITConfig                        func(ctx context.Context, owner string, repo string, request *github.GenerateJITConfigRequest) (*github.JITRunnerConfig, *github.Response, error)
	GetActionsAllowed                            func(ctx context.Context, org string) (*github.ActionsAllowed, *github.Response, error)
	GetActionsAllowedInEnterprise                func(ctx context.Context, enterprise string) (*github.ActionsAllowed, *github.Response, error)
	GetActionsPermissions                        func(ctx context.Context, org string) (*github.ActionsPermissions, *github.Response, error)
	GetActionsPermissionsInEnterprise            func(ctx context.Context, enterprise string) (*github.ActionsPermissionsEnterprise, *github.Response, error)
	GetArtifact                                  func(ctx context.Context, owner string, repo string, artifactID int64) (*github.Artifact, *github.Response, error)
	GetCacheUsageForRepo                         func(ctx context.Context, owner string, repo string) (*github.ActionsCacheUsage, *github.Response, error)
	GetDefaultWorkflowPermissionsInEnterprise    func(ctx context.Context, enterprise string) (*github.DefaultWorkflowPermissionEnterprise, *github.Response, error)
	GetDefaultWorkflowPermissionsInOrganization  func(ctx context.Context, org string) (*github.DefaultWorkflowPermissionOrganization, *github.Response, error)
	GetEnvPublicKey                              func(ctx context.Context, repoID int, env string) (*github.PublicKey, *github.Response, error)
	GetEnvSecret                                 func(ctx context.Context, repoID int, env string, secretName string) (*github.Secret, *github.Response, error)
	GetEnvVariable                               func(ctx context.Context, owner string, repo string, env string, variableName string) (*github.ActionsVariable, *github.Response, error)
	GetOrgOIDCSubjectClaimCustomTemplate         func(ctx context.Context, org string) (*github.OIDCSubjectClaimCustomTemplate, *github.Response, error)
	GetOrgPublicKey                              func(ctx context.Context, org string) (*github.PublicKey, *github.Response, error)
	GetOrgSecret                                 func(ctx context.Context, org string, name string) (*github.Secret, *github.Response, error)
	GetOrgVariable                               func(ctx context.Context, org string, name string) (*github.ActionsVariable, *github.Response, error)
	GetOrganizationRunner                        func(ctx context.Context, org string, runnerID int64) (*github.Runner, *github.Response, error)
	GetOrganizationRunnerGroup                   func(ctx context.Context, org string, groupID int64) (*github.RunnerGroup, *github.Response, error)
	GetRepoOIDCSubjectClaimCustomTemplate        func(ctx context.Context, owner string, repo string) (*github.OIDCSubjectClaimCustomTemplate, *github.Response, error)
	GetRepoPublicKey                             func(ctx context.Context, owner string, repo string) (*github.PublicKey, *github.Response, error)
	GetRepoSecret                                func(ctx context.Context, owner string, repo string, name string) (*github.Secret, *github.Response, error)
	GetRepoVariable                              func(ctx context.Context, owner string, repo string, name string) (*github.ActionsVariable, *github.Response, error)
	GetRequiredWorkflowByID                      func(ctx context.Context, owner string, requiredWorkflowID int64) (*github.OrgRequiredWorkflow, *github.Response, error)
	GetRunner                                    func(ctx context.Context, owner string, repo string, runnerID int64) (*github.Runner, *github.Response, error)
	GetTotalCacheUsageForEnterprise              func(ctx context.Context, enterprise string) (*github.TotalCacheUsage, *github.Response, error)
	GetTotalCacheUsageForOrg                     func(ctx context.Context, org string) (*github.TotalCacheUsage, *github.Response, error)
	GetWorkflowByFileName                        func(ctx context.Context, owner string, repo string, workflowFileName string) (*github.Workflow, *github.Response, error)
	GetWorkflowByID                              func(ctx context.Context, owner string, repo string, workflowID int64) (*github.Workflow, *github.Response, error)
	GetWorkflowJobByID                           func(ctx context.Context, owner string, repo string, jobID int64) (*github.WorkflowJob, *github.Response, error)
	GetWorkflowJobLogs                           func(ctx context.Context, owner string, repo string, jobID int64, maxRedirects int) (*url.URL, *github.Response, error)
	GetWorkflowRunAttempt                        func(ctx context.Context, owner string, repo string, runID int64, attemptNumber int, opts *github.WorkflowRunAttemptOptions) (*github.WorkflowRun, *github.Response, error)
	GetWorkflowRunAttemptLogs                    func(ctx context.Context, owner string, repo string, runID int64, attemptNumber int, maxRedirects int) (*url.URL, *github.Response, error)
	GetWorkflowRunByID                           func(ctx context.Context, owner string, repo string, runID int64) (*github.WorkflowRun, *github.Response, error)
	GetWorkflowRunLogs                           func(ctx context.Context, owner string, repo string, runID int64, maxRedirects int) (*url.URL, *github.Response, error)
	GetWorkflowRunUsageByID                      func(ctx context.Context, owner string, repo string, runID int64) (*github.WorkflowRunUsage, *github.Response, error)
	GetWorkflowUsageByFileName                   func(ctx context.Context, owner string, repo string, workflowFileName string) (*github.WorkflowUsage, *github.Response, error)
	GetWorkflowUsageByID                         func(ctx context.Context, owner string, repo string, workflowID int64) (*github.WorkflowUsage, *github.Response, error)
	ListArtifacts                                func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.ArtifactList, *github.Response, error)
	ListCacheUsageByRepoForOrg                   func(ctx context.Context, org string, opts *github.ListOptions) (*github.ActionsCacheUsageList, *github.Response, error)
	ListCaches                                   func(ctx context.Context, owner string, repo string, opts *github.ActionsCacheListOptions) (*github.ActionsCacheList, *github.Response, error)
	ListEnabledOrgsInEnterprise                  func(ctx context.Context, owner string, opts *github.ListOptions) (*github.ActionsEnabledOnEnterpriseRepos, *github.Response, error)
	ListEnabledReposInOrg                        func(ctx context.Context, owner string, opts *github.ListOptions) (*github.ActionsEnabledOnOrgRepos, *github.Response, error)
	ListEnvSecrets                               func(ctx context.Context, repoID int, env string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	ListEnvVariables                             func(ctx context.Context, owner string, repo string, env string, opts *github.ListOptions) (*github.ActionsVariables, *github.Response, error)
	ListOrgRequiredWorkflows                     func(ctx context.Context, org string, opts *github.ListOptions) (*github.OrgRequiredWorkflows, *github.Response, error)
	ListOrgSecrets                               func(ctx context.Context, org string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	ListOrgVariables                             func(ctx context.Context, org string, opts *github.ListOptions) (*github.ActionsVariables, *github.Response, error)
	ListOrganizationRunnerApplicationDownloads   func(ctx context.Context, org string) ([]*github.RunnerApplicationDownload, *github.Response, error)
	ListOrganizationRunnerGroups                 func(ctx context.Context, org string, opts *github.ListOrgRunnerGroupOptions) (*github.RunnerGroups, *github.Response, error)
	ListOrganizationRunners                      func(ctx context.Context, org string, opts *github.ListRunnersOptions) (*github.Runners, *github.Response, error)
	ListRepoOrgSecrets                           func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	ListRepoOrgVariables                         func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.ActionsVariables, *github.Response, error)
	ListRepoRequiredWorkflows                    func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.RepoRequiredWorkflows, *github.Response, error)
	ListRepoSecrets                              func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	ListRepoVariables                            func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.ActionsVariables, *github.Response, error)
	ListRepositoryAccessRunnerGroup              func(ctx context.Context, org string, groupID int64, opts *github.ListOptions) (*github.ListRepositories, *github.Response, error)
	ListRepositoryWorkflowRuns                   func(ctx context.Context, owner string, repo string, opts *github.ListWorkflowRunsOptions) (*github.WorkflowRuns, *github.Response, error)
	ListRequiredWorkflowSelectedRepos            func(ctx context.Context, org string, requiredWorkflowID int64, opts *github.ListOptions) (*github.RequiredWorkflowSelectedRepos, *github.Response, error)
	ListRunnerApplicationDownloads               func(ctx context.Context, owner string, repo string) ([]*github.RunnerApplicationDownload, *github.Response, error)
	ListRunnerGroupRunners                       func(ctx context.Context, org string, groupID int64, opts *github.ListOptions) (*github.Runners, *github.Response, error)
	ListRunners                                  func(ctx context.Context, owner string, repo string, opts *github.ListRunnersOptions) (*github.Runners, *github.Response, error)
	ListSelectedReposForOrgSecret                func(ctx context.Context, org string, name string, opts *github.ListOptions) (*github.SelectedReposList, *github.Response, error)
	ListSelectedReposForOrgVariable              func(ctx context.Context, org string, name string, opts *github.ListOptions) (*github.SelectedReposList, *github.Response, error)
	ListWorkflowJobs                             func(ctx context.Context, owner string, repo string, runID int64, opts *github.ListWorkflowJobsOptions) (*github.Jobs, *github.Response, error)
	ListWorkflowJobsAttempt                      func(ctx context.Context, owner string, repo string, runID int64, attemptNumber int64, opts *github.ListOptions) (*github.Jobs, *github.Response, error)
	ListWorkflowRunArtifacts                     func(ctx context.Context, owner string, repo string, runID int64, opts *github.ListOptions) (*github.ArtifactList, *github.Response, error)
	ListWorkflowRunsByFileName                   func(ctx context.Context, owner string, repo string, workflowFileName string, opts *github.ListWorkflowRunsOptions) (*github.WorkflowRuns, *github.Response, error)
	ListWorkflowRunsByID                         func(ctx context.Context, owner string, repo string, workflowID int64, opts *github.ListWorkflowRunsOptions) (*github.WorkflowRuns, *github.Response, error)
	ListWorkflows                                func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.Workflows, *github.Response, error)
	PendingDeployments                           func(ctx context.Context, owner string, repo string, runID int64, request *github.PendingDeploymentsRequest) ([]*github.Deployment, *github.Response, error)
	RemoveEnabledOrgInEnterprise                 func(ctx context.Context, owner string, organizationID int64) (*github.Response, error)
	RemoveEnabledReposInOrg                      func(ctx context.Context, owner string, repositoryID int64) (*github.Response, error)
	RemoveOrganizationRunner                     func(ctx context.Context, org string, runnerID int64) (*github.Response, error)
	RemoveRepoFromRequiredWorkflow               func(ctx context.Context, org string, requiredWorkflowID int64, repoID int64) (*github.Response, error)
	RemoveRepositoryAccessRunnerGroup            func(ctx context.Context, org string, groupID int64, repoID int64) (*github.Response, error)
	RemoveRunner                                 func(ctx context.Context, owner string, repo string, runnerID int64) (*github.Response, error)
	RemoveRunnerGroupRunners                     func(ctx context.Context, org string, groupID int64, runnerID int64) (*github.Response, error)
	RemoveSelectedRepoFromOrgSecret              func(ctx context.Context, org string, name string, repo *github.Repository) (*github.Response, error)
	RemoveSelectedRepoFromOrgVariable            func(ctx context.Context, org string, name string, repo *github.Repository) (*github.Response, error)
	RerunFailedJobsByID                          func(ctx context.Context, owner string, repo string, runID int64) (*github.Response, error)
	RerunJobByID                                 func(ctx context.Context, owner string, repo string, jobID int64) (*github.Response, error)
	RerunWorkflowByID                            func(ctx context.Context, owner string, repo string, runID int64) (*github.Response, error)
	SetEnabledOrgsInEnterprise                   func(ctx context.Context, owner string, organizationIDs []int64) (*github.Response, error)
	SetEnabledReposInOrg                         func(ctx context.Context, owner string, repositoryIDs []int64) (*github.Response, error)
	SetOrgOIDCSubjectClaimCustomTemplate         func(ctx context.Context, org string, template *github.OIDCSubjectClaimCustomTemplate) (*github.Response, error)
	SetRepoOIDCSubjectClaimCustomTemplate        func(ctx context.Context, owner string, repo string, template *github.OIDCSubjectClaimCustomTemplate) (*github.Response, error)
	SetRepositoryAccessRunnerGroup               func(ctx context.Context, org string, groupID int64, ids github.SetRepoAccessRunnerGroupRequest) (*github.Response, error)
	SetRequiredWorkflowSelectedRepos             func(ctx context.Context, org string, requiredWorkflowID int64, ids github.SelectedRepoIDs) (*github.Response, error)
	SetRunnerGroupRunners                        func(ctx context.Context, org string, groupID int64, ids github.SetRunnerGroupRunnersRequest) (*github.Response, error)
	SetSelectedReposForOrgSecret                 func(ctx context.Context, org string, name string, ids github.SelectedRepoIDs) (*github.Response, error)
	SetSelectedReposForOrgVariable               func(ctx context.Context, org string, name string, ids github.SelectedRepoIDs) (*github.Response, error)
	UpdateEnvVariable                            func(ctx context.Context, owner string, repo string, env string, variable *github.ActionsVariable) (*github.Response, error)
	UpdateOrgVariable                            func(ctx context.Context, org string, variable *github.ActionsVariable) (*github.Response, error)
	UpdateOrganizationRunnerGroup                func(ctx context.Context, org string, groupID int64, updateReq github.UpdateRunnerGroupRequest) (*github.RunnerGroup, *github.Response, error)
	UpdateRepoVariable                           func(ctx context.Context, owner string, repo string, variable *github.ActionsVariable) (*github.Response, error)
	UpdateRequiredWorkflow                       func(ctx context.Context, org string, requiredWorkflowID int64, updateRequiredWorkflowOptions *github.CreateUpdateRequiredWorkflowOptions) (*github.OrgRequiredWorkflow, *github.Response, error)
}

type ActivityService struct {
	DeleteRepositorySubscription    func(ctx context.Context, owner string, repo string) (*github.Response, error)
	DeleteThreadSubscription        func(ctx context.Context, id string) (*github.Response, error)
	GetRepositorySubscription       func(ctx context.Context, owner string, repo string) (*github.Subscription, *github.Response, error)
	GetThread                       func(ctx context.Context, id string) (*github.Notification, *github.Response, error)
	GetThreadSubscription           func(ctx context.Context, id string) (*github.Subscription, *github.Response, error)
	IsStarred                       func(ctx context.Context, owner string, repo string) (bool, *github.Response, error)
	ListEvents                      func(ctx context.Context, opts *github.ListOptions) ([]*github.Event, *github.Response, error)
	ListEventsForOrganization       func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Event, *github.Response, error)
	ListEventsForRepoNetwork        func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Event, *github.Response, error)
	ListEventsPerformedByUser       func(ctx context.Context, user string, publicOnly bool, opts *github.ListOptions) ([]*github.Event, *github.Response, error)
	ListEventsReceivedByUser        func(ctx context.Context, user string, publicOnly bool, opts *github.ListOptions) ([]*github.Event, *github.Response, error)
	ListFeeds                       func(ctx context.Context) (*github.Feeds, *github.Response, error)
	ListIssueEventsForRepository    func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.IssueEvent, *github.Response, error)
	ListNotifications               func(ctx context.Context, opts *github.NotificationListOptions) ([]*github.Notification, *github.Response, error)
	ListRepositoryEvents            func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Event, *github.Response, error)
	ListRepositoryNotifications     func(ctx context.Context, owner string, repo string, opts *github.NotificationListOptions) ([]*github.Notification, *github.Response, error)
	ListStargazers                  func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Stargazer, *github.Response, error)
	ListStarred                     func(ctx context.Context, user string, opts *github.ActivityListStarredOptions) ([]*github.StarredRepository, *github.Response, error)
	ListUserEventsForOrganization   func(ctx context.Context, org string, user string, opts *github.ListOptions) ([]*github.Event, *github.Response, error)
	ListWatched                     func(ctx context.Context, user string, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
	ListWatchers                    func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.User, *github.Response, error)
	MarkNotificationsRead           func(ctx context.Context, lastRead github.Timestamp) (*github.Response, error)
	MarkRepositoryNotificationsRead func(ctx context.Context, owner string, repo string, lastRead github.Timestamp) (*github.Response, error)
	MarkThreadRead                  func(ctx context.Context, id string) (*github.Response, error)
	SetRepositorySubscription       func(ctx context.Context, owner string, repo string, subscription *github.Subscription) (*github.Subscription, *github.Response, error)
	SetThreadSubscription           func(ctx context.Context, id string, subscription *github.Subscription) (*github.Subscription, *github.Response, error)
	Star                            func(ctx context.Context, owner string, repo string) (*github.Response, error)
	Unstar                          func(ctx context.Context, owner string, repo string) (*github.Response, error)
}

type AdminService struct {
	CreateOrg               func(ctx context.Context, org *github.Organization, admin string) (*github.Organization, *github.Response, error)
	CreateUser              func(ctx context.Context, userReq github.CreateUserRequest) (*github.User, *github.Response, error)
	CreateUserImpersonation func(ctx context.Context, username string, opts *github.ImpersonateUserOptions) (*github.UserAuthorization, *github.Response, error)
	DeleteUser              func(ctx context.Context, username string) (*github.Response, error)
	DeleteUserImpersonation func(ctx context.Context, username string) (*github.Response, error)
	GetAdminStats           func(ctx context.Context) (*github.AdminStats, *github.Response, error)
	RenameOrg               func(ctx context.Context, org *github.Organization, newName string) (*github.RenameOrgResponse, *github.Response, error)
	RenameOrgByName         func(ctx context.Context, org string, newName string) (*github.RenameOrgResponse, *github.Response, error)
	UpdateTeamLDAPMapping   func(ctx context.Context, team int64, mapping *github.TeamLDAPMapping) (*github.TeamLDAPMapping, *github.Response, error)
	UpdateUserLDAPMapping   func(ctx context.Context, user string, mapping *github.UserLDAPMapping) (*github.UserLDAPMapping, *github.Response, error)
}

type AppsService struct {
	AddRepository                    func(ctx context.Context, instID int64, repoID int64) (*github.Repository, *github.Response, error)
	CompleteAppManifest              func(ctx context.Context, code string) (*github.AppConfig, *github.Response, error)
	CreateAttachment                 func(ctx context.Context, contentReferenceID int64, title string, body string) (*github.Attachment, *github.Response, error)
	CreateInstallationToken          func(ctx context.Context, id int64, opts *github.InstallationTokenOptions) (*github.InstallationToken, *github.Response, error)
	CreateInstallationTokenListRepos func(ctx context.Context, id int64, opts *github.InstallationTokenListRepoOptions) (*github.InstallationToken, *github.Response, error)
	DeleteInstallation               func(ctx context.Context, id int64) (*github.Response, error)
	FindOrganizationInstallation     func(ctx context.Context, org string) (*github.Installation, *github.Response, error)
	FindRepositoryInstallation       func(ctx context.Context, owner string, repo string) (*github.Installation, *github.Response, error)
	FindRepositoryInstallationByID   func(ctx context.Context, id int64) (*github.Installation, *github.Response, error)
	FindUserInstallation             func(ctx context.Context, user string) (*github.Installation, *github.Response, error)
	Get                              func(ctx context.Context, appSlug string) (*github.App, *github.Response, error)
	GetHookConfig                    func(ctx context.Context) (*github.HookConfig, *github.Response, error)
	GetHookDelivery                  func(ctx context.Context, deliveryID int64) (*github.HookDelivery, *github.Response, error)
	GetInstallation                  func(ctx context.Context, id int64) (*github.Installation, *github.Response, error)
	ListHookDeliveries               func(ctx context.Context, opts *github.ListCursorOptions) ([]*github.HookDelivery, *github.Response, error)
	ListInstallationRequests         func(ctx context.Context, opts *github.ListOptions) ([]*github.InstallationRequest, *github.Response, error)
	ListInstallations                func(ctx context.Context, opts *github.ListOptions) ([]*github.Installation, *github.Response, error)
	ListRepos                        func(ctx context.Context, opts *github.ListOptions) (*github.ListRepositories, *github.Response, error)
	ListUserInstallations            func(ctx context.Context, opts *github.ListOptions) ([]*github.Installation, *github.Response, error)
	ListUserRepos                    func(ctx context.Context, id int64, opts *github.ListOptions) (*github.ListRepositories, *github.Response, error)
	RedeliverHookDelivery            func(ctx context.Context, deliveryID int64) (*github.HookDelivery, *github.Response, error)
	RemoveRepository                 func(ctx context.Context, instID int64, repoID int64) (*github.Response, error)
	RevokeInstallationToken          func(ctx context.Context) (*github.Response, error)
	SuspendInstallation              func(ctx context.Context, id int64) (*github.Response, error)
	UnsuspendInstallation            func(ctx context.Context, id int64) (*github.Response, error)
	UpdateHookConfig                 func(ctx context.Context, config *github.HookConfig) (*github.HookConfig, *github.Response, error)
}

type AuthorizationsService struct {
	Check               func(ctx context.Context, clientID string, accessToken string) (*github.Authorization, *github.Response, error)
	CreateImpersonation func(ctx context.Context, username string, authReq *github.AuthorizationRequest) (*github.Authorization, *github.Response, error)
	DeleteGrant         func(ctx context.Context, clientID string, accessToken string) (*github.Response, error)
	DeleteImpersonation func(ctx context.Context, username string) (*github.Response, error)
	Reset               func(ctx context.Context, clientID string, accessToken string) (*github.Authorization, *github.Response, error)
	Revoke              func(ctx context.Context, clientID string, accessToken string) (*github.Response, error)
}

type BillingService struct {
	GetActionsBillingOrg                   func(ctx context.Context, org string) (*github.ActionBilling, *github.Response, error)
	GetActionsBillingUser                  func(ctx context.Context, user string) (*github.ActionBilling, *github.Response, error)
	GetAdvancedSecurityActiveCommittersOrg func(ctx context.Context, org string, opts *github.ListOptions) (*github.ActiveCommitters, *github.Response, error)
	GetPackagesBillingOrg                  func(ctx context.Context, org string) (*github.PackageBilling, *github.Response, error)
	GetPackagesBillingUser                 func(ctx context.Context, user string) (*github.PackageBilling, *github.Response, error)
	GetStorageBillingOrg                   func(ctx context.Context, org string) (*github.StorageBilling, *github.Response, error)
	GetStorageBillingUser                  func(ctx context.Context, user string) (*github.StorageBilling, *github.Response, error)
}

type ChecksService struct {
	CreateCheckRun           func(ctx context.Context, owner string, repo string, opts github.CreateCheckRunOptions) (*github.CheckRun, *github.Response, error)
	CreateCheckSuite         func(ctx context.Context, owner string, repo string, opts github.CreateCheckSuiteOptions) (*github.CheckSuite, *github.Response, error)
	GetCheckRun              func(ctx context.Context, owner string, repo string, checkRunID int64) (*github.CheckRun, *github.Response, error)
	GetCheckSuite            func(ctx context.Context, owner string, repo string, checkSuiteID int64) (*github.CheckSuite, *github.Response, error)
	ListCheckRunAnnotations  func(ctx context.Context, owner string, repo string, checkRunID int64, opts *github.ListOptions) ([]*github.CheckRunAnnotation, *github.Response, error)
	ListCheckRunsCheckSuite  func(ctx context.Context, owner string, repo string, checkSuiteID int64, opts *github.ListCheckRunsOptions) (*github.ListCheckRunsResults, *github.Response, error)
	ListCheckRunsForRef      func(ctx context.Context, owner string, repo string, ref string, opts *github.ListCheckRunsOptions) (*github.ListCheckRunsResults, *github.Response, error)
	ListCheckSuitesForRef    func(ctx context.Context, owner string, repo string, ref string, opts *github.ListCheckSuiteOptions) (*github.ListCheckSuiteResults, *github.Response, error)
	ReRequestCheckRun        func(ctx context.Context, owner string, repo string, checkRunID int64) (*github.Response, error)
	ReRequestCheckSuite      func(ctx context.Context, owner string, repo string, checkSuiteID int64) (*github.Response, error)
	SetCheckSuitePreferences func(ctx context.Context, owner string, repo string, opts github.CheckSuitePreferenceOptions) (*github.CheckSuitePreferenceResults, *github.Response, error)
	UpdateCheckRun           func(ctx context.Context, owner string, repo string, checkRunID int64, opts github.UpdateCheckRunOptions) (*github.CheckRun, *github.Response, error)
}

type CodeScanningService struct {
	DeleteAnalysis                  func(ctx context.Context, owner string, repo string, id int64) (*github.DeleteAnalysis, *github.Response, error)
	GetAlert                        func(ctx context.Context, owner string, repo string, id int64) (*github.Alert, *github.Response, error)
	GetAnalysis                     func(ctx context.Context, owner string, repo string, id int64) (*github.ScanningAnalysis, *github.Response, error)
	GetCodeQLDatabase               func(ctx context.Context, owner string, repo string, language string) (*github.CodeQLDatabase, *github.Response, error)
	GetDefaultSetupConfiguration    func(ctx context.Context, owner string, repo string) (*github.DefaultSetupConfiguration, *github.Response, error)
	GetSARIF                        func(ctx context.Context, owner string, repo string, sarifID string) (*github.SARIFUpload, *github.Response, error)
	ListAlertInstances              func(ctx context.Context, owner string, repo string, id int64, opts *github.AlertInstancesListOptions) ([]*github.MostRecentInstance, *github.Response, error)
	ListAlertsForOrg                func(ctx context.Context, org string, opts *github.AlertListOptions) ([]*github.Alert, *github.Response, error)
	ListAlertsForRepo               func(ctx context.Context, owner string, repo string, opts *github.AlertListOptions) ([]*github.Alert, *github.Response, error)
	ListAnalysesForRepo             func(ctx context.Context, owner string, repo string, opts *github.AnalysesListOptions) ([]*github.ScanningAnalysis, *github.Response, error)
	ListCodeQLDatabases             func(ctx context.Context, owner string, repo string) ([]*github.CodeQLDatabase, *github.Response, error)
	UpdateAlert                     func(ctx context.Context, owner string, repo string, id int64, stateInfo *github.CodeScanningAlertState) (*github.Alert, *github.Response, error)
	UpdateDefaultSetupConfiguration func(ctx context.Context, owner string, repo string, options *github.UpdateDefaultSetupConfigurationOptions) (*github.UpdateDefaultSetupConfigurationResponse, *github.Response, error)
	UploadSarif                     func(ctx context.Context, owner string, repo string, sarif *github.SarifAnalysis) (*github.SarifID, *github.Response, error)
}

type CodesOfConductService struct {
	Get  func(ctx context.Context, key string) (*github.CodeOfConduct, *github.Response, error)
	List func(ctx context.Context) ([]*github.CodeOfConduct, *github.Response, error)
}

type CodespacesService struct {
	AddSelectedRepoToOrgSecret       func(ctx context.Context, org string, name string, repo *github.Repository) (*github.Response, error)
	AddSelectedRepoToUserSecret      func(ctx context.Context, name string, repo *github.Repository) (*github.Response, error)
	CreateInRepo                     func(ctx context.Context, owner string, repo string, request *github.CreateCodespaceOptions) (*github.Codespace, *github.Response, error)
	CreateOrUpdateOrgSecret          func(ctx context.Context, org string, eSecret *github.EncryptedSecret) (*github.Response, error)
	CreateOrUpdateRepoSecret         func(ctx context.Context, owner string, repo string, eSecret *github.EncryptedSecret) (*github.Response, error)
	CreateOrUpdateUserSecret         func(ctx context.Context, eSecret *github.EncryptedSecret) (*github.Response, error)
	Delete                           func(ctx context.Context, codespaceName string) (*github.Response, error)
	DeleteOrgSecret                  func(ctx context.Context, org string, name string) (*github.Response, error)
	DeleteRepoSecret                 func(ctx context.Context, owner string, repo string, name string) (*github.Response, error)
	DeleteUserSecret                 func(ctx context.Context, name string) (*github.Response, error)
	GetOrgPublicKey                  func(ctx context.Context, org string) (*github.PublicKey, *github.Response, error)
	GetOrgSecret                     func(ctx context.Context, org string, name string) (*github.Secret, *github.Response, error)
	GetRepoPublicKey                 func(ctx context.Context, owner string, repo string) (*github.PublicKey, *github.Response, error)
	GetRepoSecret                    func(ctx context.Context, owner string, repo string, name string) (*github.Secret, *github.Response, error)
	GetUserPublicKey                 func(ctx context.Context) (*github.PublicKey, *github.Response, error)
	GetUserSecret                    func(ctx context.Context, name string) (*github.Secret, *github.Response, error)
	List                             func(ctx context.Context, opts *github.ListCodespacesOptions) (*github.ListCodespaces, *github.Response, error)
	ListInRepo                       func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.ListCodespaces, *github.Response, error)
	ListOrgSecrets                   func(ctx context.Context, org string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	ListRepoSecrets                  func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	ListSelectedReposForOrgSecret    func(ctx context.Context, org string, name string, opts *github.ListOptions) (*github.SelectedReposList, *github.Response, error)
	ListSelectedReposForUserSecret   func(ctx context.Context, name string, opts *github.ListOptions) (*github.SelectedReposList, *github.Response, error)
	ListUserSecrets                  func(ctx context.Context, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	RemoveSelectedRepoFromOrgSecret  func(ctx context.Context, org string, name string, repo *github.Repository) (*github.Response, error)
	RemoveSelectedRepoFromUserSecret func(ctx context.Context, name string, repo *github.Repository) (*github.Response, error)
	SetSelectedReposForOrgSecret     func(ctx context.Context, org string, name string, ids github.SelectedRepoIDs) (*github.Response, error)
	SetSelectedReposForUserSecret    func(ctx context.Context, name string, ids github.SelectedRepoIDs) (*github.Response, error)
	Start                            func(ctx context.Context, codespaceName string) (*github.Codespace, *github.Response, error)
	Stop                             func(ctx context.Context, codespaceName string) (*github.Codespace, *github.Response, error)
}

type CopilotService struct {
	AddCopilotTeams    func(ctx context.Context, org string, teamNames []string) (*github.SeatAssignments, *github.Response, error)
	AddCopilotUsers    func(ctx context.Context, org string, users []string) (*github.SeatAssignments, *github.Response, error)
	GetCopilotBilling  func(ctx context.Context, org string) (*github.CopilotOrganizationDetails, *github.Response, error)
	GetSeatDetails     func(ctx context.Context, org string, user string) (*github.CopilotSeatDetails, *github.Response, error)
	ListCopilotSeats   func(ctx context.Context, org string, opts *github.ListOptions) (*github.ListCopilotSeatsResponse, *github.Response, error)
	RemoveCopilotTeams func(ctx context.Context, org string, teamNames []string) (*github.SeatCancellations, *github.Response, error)
	RemoveCopilotUsers func(ctx context.Context, org string, users []string) (*github.SeatCancellations, *github.Response, error)
}

type DependabotService struct {
	AddSelectedRepoToOrgSecret      func(ctx context.Context, org string, name string, repo *github.Repository) (*github.Response, error)
	CreateOrUpdateOrgSecret         func(ctx context.Context, org string, eSecret *github.DependabotEncryptedSecret) (*github.Response, error)
	CreateOrUpdateRepoSecret        func(ctx context.Context, owner string, repo string, eSecret *github.DependabotEncryptedSecret) (*github.Response, error)
	DeleteOrgSecret                 func(ctx context.Context, org string, name string) (*github.Response, error)
	DeleteRepoSecret                func(ctx context.Context, owner string, repo string, name string) (*github.Response, error)
	GetOrgPublicKey                 func(ctx context.Context, org string) (*github.PublicKey, *github.Response, error)
	GetOrgSecret                    func(ctx context.Context, org string, name string) (*github.Secret, *github.Response, error)
	GetRepoAlert                    func(ctx context.Context, owner string, repo string, number int) (*github.DependabotAlert, *github.Response, error)
	GetRepoPublicKey                func(ctx context.Context, owner string, repo string) (*github.PublicKey, *github.Response, error)
	GetRepoSecret                   func(ctx context.Context, owner string, repo string, name string) (*github.Secret, *github.Response, error)
	ListOrgAlerts                   func(ctx context.Context, org string, opts *github.ListAlertsOptions) ([]*github.DependabotAlert, *github.Response, error)
	ListOrgSecrets                  func(ctx context.Context, org string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	ListRepoAlerts                  func(ctx context.Context, owner string, repo string, opts *github.ListAlertsOptions) ([]*github.DependabotAlert, *github.Response, error)
	ListRepoSecrets                 func(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	ListSelectedReposForOrgSecret   func(ctx context.Context, org string, name string, opts *github.ListOptions) (*github.SelectedReposList, *github.Response, error)
	RemoveSelectedRepoFromOrgSecret func(ctx context.Context, org string, name string, repo *github.Repository) (*github.Response, error)
	SetSelectedReposForOrgSecret    func(ctx context.Context, org string, name string, ids github.DependabotSecretsSelectedRepoIDs) (*github.Response, error)
	UpdateAlert                     func(ctx context.Context, owner string, repo string, number int, stateInfo *github.DependabotAlertState) (*github.DependabotAlert, *github.Response, error)
}

type DependencyGraphService struct {
	CreateSnapshot func(ctx context.Context, owner string, repo string, dependencyGraphSnapshot *github.DependencyGraphSnapshot) (*github.DependencyGraphSnapshotCreationData, *github.Response, error)
	GetSBOM        func(ctx context.Context, owner string, repo string) (*github.SBOM, *github.Response, error)
}

type EmojisService struct {
	List func(ctx context.Context) (map[string]string, *github.Response, error)
}

type EnterpriseService struct {
	AddOrganizationAccessRunnerGroup    func(ctx context.Context, enterprise string, groupID int64, orgID int64) (*github.Response, error)
	AddRunnerGroupRunners               func(ctx context.Context, enterprise string, groupID int64, runnerID int64) (*github.Response, error)
	CreateEnterpriseRunnerGroup         func(ctx context.Context, enterprise string, createReq github.CreateEnterpriseRunnerGroupRequest) (*github.EnterpriseRunnerGroup, *github.Response, error)
	CreateRegistrationToken             func(ctx context.Context, enterprise string) (*github.RegistrationToken, *github.Response, error)
	DeleteEnterpriseRunnerGroup         func(ctx context.Context, enterprise string, groupID int64) (*github.Response, error)
	EnableDisableSecurityFeature        func(ctx context.Context, enterprise string, securityProduct string, enablement string) (*github.Response, error)
	GenerateEnterpriseJITConfig         func(ctx context.Context, enterprise string, request *github.GenerateJITConfigRequest) (*github.JITRunnerConfig, *github.Response, error)
	GetAuditLog                         func(ctx context.Context, enterprise string, opts *github.GetAuditLogOptions) ([]*github.AuditEntry, *github.Response, error)
	GetCodeSecurityAndAnalysis          func(ctx context.Context, enterprise string) (*github.EnterpriseSecurityAnalysisSettings, *github.Response, error)
	GetEnterpriseRunnerGroup            func(ctx context.Context, enterprise string, groupID int64) (*github.EnterpriseRunnerGroup, *github.Response, error)
	ListOrganizationAccessRunnerGroup   func(ctx context.Context, enterprise string, groupID int64, opts *github.ListOptions) (*github.ListOrganizations, *github.Response, error)
	ListRunnerApplicationDownloads      func(ctx context.Context, enterprise string) ([]*github.RunnerApplicationDownload, *github.Response, error)
	ListRunnerGroupRunners              func(ctx context.Context, enterprise string, groupID int64, opts *github.ListOptions) (*github.Runners, *github.Response, error)
	ListRunnerGroups                    func(ctx context.Context, enterprise string, opts *github.ListEnterpriseRunnerGroupOptions) (*github.EnterpriseRunnerGroups, *github.Response, error)
	ListRunners                         func(ctx context.Context, enterprise string, opts *github.ListOptions) (*github.Runners, *github.Response, error)
	RemoveOrganizationAccessRunnerGroup func(ctx context.Context, enterprise string, groupID int64, orgID int64) (*github.Response, error)
	RemoveRunner                        func(ctx context.Context, enterprise string, runnerID int64) (*github.Response, error)
	RemoveRunnerGroupRunners            func(ctx context.Context, enterprise string, groupID int64, runnerID int64) (*github.Response, error)
	SetOrganizationAccessRunnerGroup    func(ctx context.Context, enterprise string, groupID int64, ids github.SetOrgAccessRunnerGroupRequest) (*github.Response, error)
	SetRunnerGroupRunners               func(ctx context.Context, enterprise string, groupID int64, ids github.SetRunnerGroupRunnersRequest) (*github.Response, error)
	UpdateCodeSecurityAndAnalysis       func(ctx context.Context, enterprise string, settings *github.EnterpriseSecurityAnalysisSettings) (*github.Response, error)
	UpdateEnterpriseRunnerGroup         func(ctx context.Context, enterprise string, groupID int64, updateReq github.UpdateEnterpriseRunnerGroupRequest) (*github.EnterpriseRunnerGroup, *github.Response, error)
}

type GistsService struct {
	Create        func(ctx context.Context, gist *github.Gist) (*github.Gist, *github.Response, error)
	CreateComment func(ctx context.Context, gistID string, comment *github.GistComment) (*github.GistComment, *github.Response, error)
	Delete        func(ctx context.Context, id string) (*github.Response, error)
	DeleteComment func(ctx context.Context, gistID string, commentID int64) (*github.Response, error)
	Edit          func(ctx context.Context, id string, gist *github.Gist) (*github.Gist, *github.Response, error)
	EditComment   func(ctx context.Context, gistID string, commentID int64, comment *github.GistComment) (*github.GistComment, *github.Response, error)
	Fork          func(ctx context.Context, id string) (*github.Gist, *github.Response, error)
	Get           func(ctx context.Context, id string) (*github.Gist, *github.Response, error)
	GetComment    func(ctx context.Context, gistID string, commentID int64) (*github.GistComment, *github.Response, error)
	GetRevision   func(ctx context.Context, id string, sha string) (*github.Gist, *github.Response, error)
	IsStarred     func(ctx context.Context, id string) (bool, *github.Response, error)
	List          func(ctx context.Context, user string, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error)
	ListAll       func(ctx context.Context, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error)
	ListComments  func(ctx context.Context, gistID string, opts *github.ListOptions) ([]*github.GistComment, *github.Response, error)
	ListCommits   func(ctx context.Context, id string, opts *github.ListOptions) ([]*github.GistCommit, *github.Response, error)
	ListForks     func(ctx context.Context, id string, opts *github.ListOptions) ([]*github.GistFork, *github.Response, error)
	ListStarred   func(ctx context.Context, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error)
	Star          func(ctx context.Context, id string) (*github.Response, error)
	Unstar        func(ctx context.Context, id string) (*github.Response, error)
}

type GitService struct {
	CreateBlob       func(ctx context.Context, owner string, repo string, blob *github.Blob) (*github.Blob, *github.Response, error)
	CreateCommit     func(ctx context.Context, owner string, repo string, commit *github.Commit, opts *github.CreateCommitOptions) (*github.Commit, *github.Response, error)
	CreateRef        func(ctx context.Context, owner string, repo string, ref *github.Reference) (*github.Reference, *github.Response, error)
	CreateTag        func(ctx context.Context, owner string, repo string, tag *github.Tag) (*github.Tag, *github.Response, error)
	CreateTree       func(ctx context.Context, owner string, repo string, baseTree string, entries []*github.TreeEntry) (*github.Tree, *github.Response, error)
	DeleteRef        func(ctx context.Context, owner string, repo string, ref string) (*github.Response, error)
	GetBlob          func(ctx context.Context, owner string, repo string, sha string) (*github.Blob, *github.Response, error)
	GetBlobRaw       func(ctx context.Context, owner string, repo string, sha string) ([]byte, *github.Response, error)
	GetCommit        func(ctx context.Context, owner string, repo string, sha string) (*github.Commit, *github.Response, error)
	GetRef           func(ctx context.Context, owner string, repo string, ref string) (*github.Reference, *github.Response, error)
	GetTag           func(ctx context.Context, owner string, repo string, sha string) (*github.Tag, *github.Response, error)
	GetTree          func(ctx context.Context, owner string, repo string, sha string, recursive bool) (*github.Tree, *github.Response, error)
	ListMatchingRefs func(ctx context.Context, owner string, repo string, opts *github.ReferenceListOptions) ([]*github.Reference, *github.Response, error)
	UpdateRef        func(ctx context.Context, owner string, repo string, ref *github.Reference, force bool) (*github.Reference, *github.Response, error)
}

type GitignoresService struct {
	Get  func(ctx context.Context, name string) (*github.Gitignore, *github.Response, error)
	List func(ctx context.Context) ([]string, *github.Response, error)
}

type InteractionsService struct {
	GetRestrictionsForOrg      func(ctx context.Context, organization string) (*github.InteractionRestriction, *github.Response, error)
	GetRestrictionsForRepo     func(ctx context.Context, owner string, repo string) (*github.InteractionRestriction, *github.Response, error)
	RemoveRestrictionsFromOrg  func(ctx context.Context, organization string) (*github.Response, error)
	RemoveRestrictionsFromRepo func(ctx context.Context, owner string, repo string) (*github.Response, error)
	UpdateRestrictionsForOrg   func(ctx context.Context, organization string, limit string) (*github.InteractionRestriction, *github.Response, error)
	UpdateRestrictionsForRepo  func(ctx context.Context, owner string, repo string, limit string) (*github.InteractionRestriction, *github.Response, error)
}

type IssueImportService struct {
	CheckStatus      func(ctx context.Context, owner string, repo string, issueID int64) (*github.IssueImportResponse, *github.Response, error)
	CheckStatusSince func(ctx context.Context, owner string, repo string, since github.Timestamp) ([]*github.IssueImportResponse, *github.Response, error)
	Create           func(ctx context.Context, owner string, repo string, issue *github.IssueImportRequest) (*github.IssueImportResponse, *github.Response, error)
}

type LicensesService struct {
	Get  func(ctx context.Context, licenseName string) (*github.License, *github.Response, error)
	List func(ctx context.Context) ([]*github.License, *github.Response, error)
}

type MarkdownService struct {
	Render func(ctx context.Context, text string, opts *github.MarkdownOptions) (string, *github.Response, error)
}

type MarketplaceService struct {
	GetPlanAccountForAccount        func(ctx context.Context, accountID int64) (*github.MarketplacePlanAccount, *github.Response, error)
	ListMarketplacePurchasesForUser func(ctx context.Context, opts *github.ListOptions) ([]*github.MarketplacePurchase, *github.Response, error)
	ListPlanAccountsForPlan         func(ctx context.Context, planID int64, opts *github.ListOptions) ([]*github.MarketplacePlanAccount, *github.Response, error)
	ListPlans                       func(ctx context.Context, opts *github.ListOptions) ([]*github.MarketplacePlan, *github.Response, error)
}

type MetaService struct {
	Get     func(ctx context.Context) (*github.APIMeta, *github.Response, error)
	Octocat func(ctx context.Context, message string) (string, *github.Response, error)
	Zen     func(ctx context.Context) (string, *github.Response, error)
}

type MigrationService struct {
	CancelImport            func(ctx context.Context, owner string, repo string) (*github.Response, error)
	CommitAuthors           func(ctx context.Context, owner string, repo string) ([]*github.SourceImportAuthor, *github.Response, error)
	DeleteMigration         func(ctx context.Context, org string, id int64) (*github.Response, error)
	DeleteUserMigration     func(ctx context.Context, id int64) (*github.Response, error)
	ImportProgress          func(ctx context.Context, owner string, repo string) (*github.Import, *github.Response, error)
	LargeFiles              func(ctx context.Context, owner string, repo string) ([]*github.LargeFile, *github.Response, error)
	ListMigrations          func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Migration, *github.Response, error)
	ListUserMigrations      func(ctx context.Context, opts *github.ListOptions) ([]*github.UserMigration, *github.Response, error)
	MapCommitAuthor         func(ctx context.Context, owner string, repo string, id int64, author *github.SourceImportAuthor) (*github.SourceImportAuthor, *github.Response, error)
	MigrationArchiveURL     func(ctx context.Context, org string, id int64) (url string, err error)
	MigrationStatus         func(ctx context.Context, org string, id int64) (*github.Migration, *github.Response, error)
	SetLFSPreference        func(ctx context.Context, owner string, repo string, in *github.Import) (*github.Import, *github.Response, error)
	StartImport             func(ctx context.Context, owner string, repo string, in *github.Import) (*github.Import, *github.Response, error)
	StartMigration          func(ctx context.Context, org string, repos []string, opts *github.MigrationOptions) (*github.Migration, *github.Response, error)
	StartUserMigration      func(ctx context.Context, repos []string, opts *github.UserMigrationOptions) (*github.UserMigration, *github.Response, error)
	UnlockRepo              func(ctx context.Context, org string, id int64, repo string) (*github.Response, error)
	UnlockUserRepo          func(ctx context.Context, id int64, repo string) (*github.Response, error)
	UpdateImport            func(ctx context.Context, owner string, repo string, in *github.Import) (*github.Import, *github.Response, error)
	UserMigrationArchiveURL func(ctx context.Context, id int64) (string, error)
	UserMigrationStatus     func(ctx context.Context, id int64) (*github.UserMigration, *github.Response, error)
}

type OrganizationsService struct {
	AddSecurityManagerTeam                 func(ctx context.Context, org string, team string) (*github.Response, error)
	BlockUser                              func(ctx context.Context, org string, user string) (*github.Response, error)
	ConcealMembership                      func(ctx context.Context, org string, user string) (*github.Response, error)
	ConvertMemberToOutsideCollaborator     func(ctx context.Context, org string, user string) (*github.Response, error)
	CreateCustomRepoRole                   func(ctx context.Context, org string, opts *github.CreateOrUpdateCustomRoleOptions) (*github.CustomRepoRoles, *github.Response, error)
	CreateHook                             func(ctx context.Context, org string, hook *github.Hook) (*github.Hook, *github.Response, error)
	CreateOrUpdateCustomProperties         func(ctx context.Context, org string, properties []*github.CustomProperty) ([]*github.CustomProperty, *github.Response, error)
	CreateOrUpdateCustomProperty           func(ctx context.Context, org string, customPropertyName string, property *github.CustomProperty) (*github.CustomProperty, *github.Response, error)
	CreateOrUpdateRepoCustomPropertyValues func(ctx context.Context, org string, repoNames []string, properties []*github.CustomPropertyValue) (*github.Response, error)
	CreateOrgInvitation                    func(ctx context.Context, org string, opts *github.CreateOrgInvitationOptions) (*github.Invitation, *github.Response, error)
	CreateOrganizationRuleset              func(ctx context.Context, org string, rs *github.Ruleset) (*github.Ruleset, *github.Response, error)
	CreateProject                          func(ctx context.Context, org string, opts *github.ProjectOptions) (*github.Project, *github.Response, error)
	Delete                                 func(ctx context.Context, org string) (*github.Response, error)
	DeleteCustomRepoRole                   func(ctx context.Context, org string, roleID string) (*github.Response, error)
	DeleteHook                             func(ctx context.Context, org string, id int64) (*github.Response, error)
	DeleteOrganizationRuleset              func(ctx context.Context, org string, rulesetID int64) (*github.Response, error)
	DeletePackage                          func(ctx context.Context, org string, packageType string, packageName string) (*github.Response, error)
	Edit                                   func(ctx context.Context, name string, org *github.Organization) (*github.Organization, *github.Response, error)
	EditHook                               func(ctx context.Context, org string, id int64, hook *github.Hook) (*github.Hook, *github.Response, error)
	EditHookConfiguration                  func(ctx context.Context, org string, id int64, config *github.HookConfig) (*github.HookConfig, *github.Response, error)
	EditOrgMembership                      func(ctx context.Context, user string, org string, membership *github.Membership) (*github.Membership, *github.Response, error)
	Get                                    func(ctx context.Context, org string) (*github.Organization, *github.Response, error)
	GetAllCustomProperties                 func(ctx context.Context, org string) ([]*github.CustomProperty, *github.Response, error)
	GetAllOrganizationRulesets             func(ctx context.Context, org string) ([]*github.Ruleset, *github.Response, error)
	GetAuditLog                            func(ctx context.Context, org string, opts *github.GetAuditLogOptions) ([]*github.AuditEntry, *github.Response, error)
	GetByID                                func(ctx context.Context, id int64) (*github.Organization, *github.Response, error)
	GetCustomProperty                      func(ctx context.Context, org string, name string) (*github.CustomProperty, *github.Response, error)
	GetHook                                func(ctx context.Context, org string, id int64) (*github.Hook, *github.Response, error)
	GetHookConfiguration                   func(ctx context.Context, org string, id int64) (*github.HookConfig, *github.Response, error)
	GetHookDelivery                        func(ctx context.Context, owner string, hookID int64, deliveryID int64) (*github.HookDelivery, *github.Response, error)
	GetOrgMembership                       func(ctx context.Context, user string, org string) (*github.Membership, *github.Response, error)
	GetOrganizationRuleset                 func(ctx context.Context, org string, rulesetID int64) (*github.Ruleset, *github.Response, error)
	GetPackage                             func(ctx context.Context, org string, packageType string, packageName string) (*github.Package, *github.Response, error)
	IsBlocked                              func(ctx context.Context, org string, user string) (bool, *github.Response, error)
	IsMember                               func(ctx context.Context, org string, user string) (bool, *github.Response, error)
	IsPublicMember                         func(ctx context.Context, org string, user string) (bool, *github.Response, error)
	List                                   func(ctx context.Context, user string, opts *github.ListOptions) ([]*github.Organization, *github.Response, error)
	ListAll                                func(ctx context.Context, opts *github.OrganizationsListOptions) ([]*github.Organization, *github.Response, error)
	ListBlockedUsers                       func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.User, *github.Response, error)
	ListCredentialAuthorizations           func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.CredentialAuthorization, *github.Response, error)
	ListCustomPropertyValues               func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.RepoCustomPropertyValue, *github.Response, error)
	ListCustomRepoRoles                    func(ctx context.Context, org string) (*github.OrganizationCustomRepoRoles, *github.Response, error)
	ListFailedOrgInvitations               func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Invitation, *github.Response, error)
	ListHookDeliveries                     func(ctx context.Context, org string, id int64, opts *github.ListCursorOptions) ([]*github.HookDelivery, *github.Response, error)
	ListHooks                              func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Hook, *github.Response, error)
	ListInstallations                      func(ctx context.Context, org string, opts *github.ListOptions) (*github.OrganizationInstallations, *github.Response, error)
	ListMembers                            func(ctx context.Context, org string, opts *github.ListMembersOptions) ([]*github.User, *github.Response, error)
	ListOrgInvitationTeams                 func(ctx context.Context, org string, invitationID string, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	ListOrgMemberships                     func(ctx context.Context, opts *github.ListOrgMembershipsOptions) ([]*github.Membership, *github.Response, error)
	ListOutsideCollaborators               func(ctx context.Context, org string, opts *github.ListOutsideCollaboratorsOptions) ([]*github.User, *github.Response, error)
	ListPackages                           func(ctx context.Context, org string, opts *github.PackageListOptions) ([]*github.Package, *github.Response, error)
	ListPendingOrgInvitations              func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Invitation, *github.Response, error)
	ListProjects                           func(ctx context.Context, org string, opts *github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListSecurityManagerTeams               func(ctx context.Context, org string) ([]*github.Team, *github.Response, error)
	ListTeamsAssignedToOrgRole             func(ctx context.Context, org string, roleID int64, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	ListUsersAssignedToOrgRole             func(ctx context.Context, org string, roleID int64, opts *github.ListOptions) ([]*github.User, *github.Response, error)
	PackageDeleteVersion                   func(ctx context.Context, org string, packageType string, packageName string, packageVersionID int64) (*github.Response, error)
	PackageGetAllVersions                  func(ctx context.Context, org string, packageType string, packageName string, opts *github.PackageListOptions) ([]*github.PackageVersion, *github.Response, error)
	PackageGetVersion                      func(ctx context.Context, org string, packageType string, packageName string, packageVersionID int64) (*github.PackageVersion, *github.Response, error)
	PackageRestoreVersion                  func(ctx context.Context, org string, packageType string, packageName string, packageVersionID int64) (*github.Response, error)
	PingHook                               func(ctx context.Context, org string, id int64) (*github.Response, error)
	PublicizeMembership                    func(ctx context.Context, org string, user string) (*github.Response, error)
	RedeliverHookDelivery                  func(ctx context.Context, owner string, hookID int64, deliveryID int64) (*github.HookDelivery, *github.Response, error)
	RemoveCredentialAuthorization          func(ctx context.Context, org string, credentialID int64) (*github.Response, error)
	RemoveCustomProperty                   func(ctx context.Context, org string, customPropertyName string) (*github.Response, error)
	RemoveMember                           func(ctx context.Context, org string, user string) (*github.Response, error)
	RemoveOrgMembership                    func(ctx context.Context, user string, org string) (*github.Response, error)
	RemoveOutsideCollaborator              func(ctx context.Context, org string, user string) (*github.Response, error)
	RemoveSecurityManagerTeam              func(ctx context.Context, org string, team string) (*github.Response, error)
	RestorePackage                         func(ctx context.Context, org string, packageType string, packageName string) (*github.Response, error)
	ReviewPersonalAccessTokenRequest       func(ctx context.Context, org string, requestID int64, opts github.ReviewPersonalAccessTokenRequestOptions) (*github.Response, error)
	UnblockUser                            func(ctx context.Context, org string, user string) (*github.Response, error)
	UpdateCustomRepoRole                   func(ctx context.Context, org string, roleID string, opts *github.CreateOrUpdateCustomRoleOptions) (*github.CustomRepoRoles, *github.Response, error)
	UpdateOrganizationRuleset              func(ctx context.Context, org string, rulesetID int64, rs *github.Ruleset) (*github.Ruleset, *github.Response, error)
}

type ProjectsService struct {
	AddProjectCollaborator              func(ctx context.Context, id int64, username string, opts *github.ProjectCollaboratorOptions) (*github.Response, error)
	CreateProjectCard                   func(ctx context.Context, columnID int64, opts *github.ProjectCardOptions) (*github.ProjectCard, *github.Response, error)
	CreateProjectColumn                 func(ctx context.Context, projectID int64, opts *github.ProjectColumnOptions) (*github.ProjectColumn, *github.Response, error)
	DeleteProject                       func(ctx context.Context, id int64) (*github.Response, error)
	DeleteProjectCard                   func(ctx context.Context, cardID int64) (*github.Response, error)
	DeleteProjectColumn                 func(ctx context.Context, columnID int64) (*github.Response, error)
	GetProject                          func(ctx context.Context, id int64) (*github.Project, *github.Response, error)
	GetProjectCard                      func(ctx context.Context, cardID int64) (*github.ProjectCard, *github.Response, error)
	GetProjectColumn                    func(ctx context.Context, id int64) (*github.ProjectColumn, *github.Response, error)
	ListProjectCards                    func(ctx context.Context, columnID int64, opts *github.ProjectCardListOptions) ([]*github.ProjectCard, *github.Response, error)
	ListProjectCollaborators            func(ctx context.Context, id int64, opts *github.ListCollaboratorOptions) ([]*github.User, *github.Response, error)
	ListProjectColumns                  func(ctx context.Context, projectID int64, opts *github.ListOptions) ([]*github.ProjectColumn, *github.Response, error)
	MoveProjectCard                     func(ctx context.Context, cardID int64, opts *github.ProjectCardMoveOptions) (*github.Response, error)
	MoveProjectColumn                   func(ctx context.Context, columnID int64, opts *github.ProjectColumnMoveOptions) (*github.Response, error)
	RemoveProjectCollaborator           func(ctx context.Context, id int64, username string) (*github.Response, error)
	ReviewProjectCollaboratorPermission func(ctx context.Context, id int64, username string) (*github.ProjectPermissionLevel, *github.Response, error)
	UpdateProject                       func(ctx context.Context, id int64, opts *github.ProjectOptions) (*github.Project, *github.Response, error)
	UpdateProjectCard                   func(ctx context.Context, cardID int64, opts *github.ProjectCardOptions) (*github.ProjectCard, *github.Response, error)
	UpdateProjectColumn                 func(ctx context.Context, columnID int64, opts *github.ProjectColumnOptions) (*github.ProjectColumn, *github.Response, error)
}

type PullRequestsService struct {
	Create                     func(ctx context.Context, owner string, repo string, pull *github.NewPullRequest) (*github.PullRequest, *github.Response, error)
	CreateComment              func(ctx context.Context, owner string, repo string, number int, comment *github.PullRequestComment) (*github.PullRequestComment, *github.Response, error)
	CreateCommentInReplyTo     func(ctx context.Context, owner string, repo string, number int, body string, commentID int64) (*github.PullRequestComment, *github.Response, error)
	CreateReview               func(ctx context.Context, owner string, repo string, number int, review *github.PullRequestReviewRequest) (*github.PullRequestReview, *github.Response, error)
	DeleteComment              func(ctx context.Context, owner string, repo string, commentID int64) (*github.Response, error)
	DeletePendingReview        func(ctx context.Context, owner string, repo string, number int, reviewID int64) (*github.PullRequestReview, *github.Response, error)
	DismissReview              func(ctx context.Context, owner string, repo string, number int, reviewID int64, review *github.PullRequestReviewDismissalRequest) (*github.PullRequestReview, *github.Response, error)
	Edit                       func(ctx context.Context, owner string, repo string, number int, pull *github.PullRequest) (*github.PullRequest, *github.Response, error)
	EditComment                func(ctx context.Context, owner string, repo string, commentID int64, comment *github.PullRequestComment) (*github.PullRequestComment, *github.Response, error)
	Get                        func(ctx context.Context, owner string, repo string, number int) (*github.PullRequest, *github.Response, error)
	GetComment                 func(ctx context.Context, owner string, repo string, commentID int64) (*github.PullRequestComment, *github.Response, error)
	GetRaw                     func(ctx context.Context, owner string, repo string, number int, opts github.RawOptions) (string, *github.Response, error)
	GetReview                  func(ctx context.Context, owner string, repo string, number int, reviewID int64) (*github.PullRequestReview, *github.Response, error)
	IsMerged                   func(ctx context.Context, owner string, repo string, number int) (bool, *github.Response, error)
	List                       func(ctx context.Context, owner string, repo string, opts *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error)
	ListComments               func(ctx context.Context, owner string, repo string, number int, opts *github.PullRequestListCommentsOptions) ([]*github.PullRequestComment, *github.Response, error)
	ListCommits                func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.RepositoryCommit, *github.Response, error)
	ListFiles                  func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.CommitFile, *github.Response, error)
	ListPullRequestsWithCommit func(ctx context.Context, owner string, repo string, sha string, opts *github.ListOptions) ([]*github.PullRequest, *github.Response, error)
	ListReviewComments         func(ctx context.Context, owner string, repo string, number int, reviewID int64, opts *github.ListOptions) ([]*github.PullRequestComment, *github.Response, error)
	ListReviewers              func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) (*github.Reviewers, *github.Response, error)
	ListReviews                func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.PullRequestReview, *github.Response, error)
	Merge                      func(ctx context.Context, owner string, repo string, number int, commitMessage string, options *github.PullRequestOptions) (*github.PullRequestMergeResult, *github.Response, error)
	RemoveReviewers            func(ctx context.Context, owner string, repo string, number int, reviewers github.ReviewersRequest) (*github.Response, error)
	RequestReviewers           func(ctx context.Context, owner string, repo string, number int, reviewers github.ReviewersRequest) (*github.PullRequest, *github.Response, error)
	SubmitReview               func(ctx context.Context, owner string, repo string, number int, reviewID int64, review *github.PullRequestReviewRequest) (*github.PullRequestReview, *github.Response, error)
	UpdateBranch               func(ctx context.Context, owner string, repo string, number int, opts *github.PullRequestBranchUpdateOptions) (*github.PullRequestBranchUpdateResponse, *github.Response, error)
	UpdateReview               func(ctx context.Context, owner string, repo string, number int, reviewID int64, body string) (*github.PullRequestReview, *github.Response, error)
}

type RateLimitService struct {
	Get func(ctx context.Context) (*github.RateLimits, *github.Response, error)
}

type ReactionsService struct {
	CreateCommentReaction                               func(ctx context.Context, owner string, repo string, id int64, content string) (*github.Reaction, *github.Response, error)
	CreateIssueCommentReaction                          func(ctx context.Context, owner string, repo string, id int64, content string) (*github.Reaction, *github.Response, error)
	CreateIssueReaction                                 func(ctx context.Context, owner string, repo string, number int, content string) (*github.Reaction, *github.Response, error)
	CreatePullRequestCommentReaction                    func(ctx context.Context, owner string, repo string, id int64, content string) (*github.Reaction, *github.Response, error)
	CreateReleaseReaction                               func(ctx context.Context, owner string, repo string, releaseID int64, content string) (*github.Reaction, *github.Response, error)
	CreateTeamDiscussionCommentReaction                 func(ctx context.Context, teamID int64, discussionNumber int, commentNumber int, content string) (*github.Reaction, *github.Response, error)
	CreateTeamDiscussionReaction                        func(ctx context.Context, teamID int64, discussionNumber int, content string) (*github.Reaction, *github.Response, error)
	DeleteCommentReaction                               func(ctx context.Context, owner string, repo string, commentID int64, reactionID int64) (*github.Response, error)
	DeleteCommentReactionByID                           func(ctx context.Context, repoID int64, commentID int64, reactionID int64) (*github.Response, error)
	DeleteIssueCommentReaction                          func(ctx context.Context, owner string, repo string, commentID int64, reactionID int64) (*github.Response, error)
	DeleteIssueCommentReactionByID                      func(ctx context.Context, repoID int64, commentID int64, reactionID int64) (*github.Response, error)
	DeleteIssueReaction                                 func(ctx context.Context, owner string, repo string, issueNumber int, reactionID int64) (*github.Response, error)
	DeleteIssueReactionByID                             func(ctx context.Context, repoID int, issueNumber int, reactionID int64) (*github.Response, error)
	DeletePullRequestCommentReaction                    func(ctx context.Context, owner string, repo string, commentID int64, reactionID int64) (*github.Response, error)
	DeletePullRequestCommentReactionByID                func(ctx context.Context, repoID int64, commentID int64, reactionID int64) (*github.Response, error)
	DeleteTeamDiscussionCommentReaction                 func(ctx context.Context, org string, teamSlug string, discussionNumber int, commentNumber int, reactionID int64) (*github.Response, error)
	DeleteTeamDiscussionCommentReactionByOrgIDAndTeamID func(ctx context.Context, orgID int, teamID int, discussionNumber int, commentNumber int, reactionID int64) (*github.Response, error)
	DeleteTeamDiscussionReaction                        func(ctx context.Context, org string, teamSlug string, discussionNumber int, reactionID int64) (*github.Response, error)
	DeleteTeamDiscussionReactionByOrgIDAndTeamID        func(ctx context.Context, orgID int, teamID int, discussionNumber int, reactionID int64) (*github.Response, error)
	ListCommentReactions                                func(ctx context.Context, owner string, repo string, id int64, opts *github.ListCommentReactionOptions) ([]*github.Reaction, *github.Response, error)
	ListIssueCommentReactions                           func(ctx context.Context, owner string, repo string, id int64, opts *github.ListOptions) ([]*github.Reaction, *github.Response, error)
	ListIssueReactions                                  func(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.Reaction, *github.Response, error)
	ListPullRequestCommentReactions                     func(ctx context.Context, owner string, repo string, id int64, opts *github.ListOptions) ([]*github.Reaction, *github.Response, error)
	ListTeamDiscussionCommentReactions                  func(ctx context.Context, teamID int64, discussionNumber int, commentNumber int, opts *github.ListOptions) ([]*github.Reaction, *github.Response, error)
	ListTeamDiscussionReactions                         func(ctx context.Context, teamID int64, discussionNumber int, opts *github.ListOptions) ([]*github.Reaction, *github.Response, error)
}

type RepositoriesService struct {
	AddAdminEnforcement                   func(ctx context.Context, owner string, repo string, branch string) (*github.AdminEnforcement, *github.Response, error)
	AddAppRestrictions                    func(ctx context.Context, owner string, repo string, branch string, apps []string) ([]*github.App, *github.Response, error)
	AddAutolink                           func(ctx context.Context, owner string, repo string, opts *github.AutolinkOptions) (*github.Autolink, *github.Response, error)
	AddCollaborator                       func(ctx context.Context, owner string, repo string, user string, opts *github.RepositoryAddCollaboratorOptions) (*github.CollaboratorInvitation, *github.Response, error)
	AddTeamRestrictions                   func(ctx context.Context, owner string, repo string, branch string, teams []string) ([]*github.Team, *github.Response, error)
	AddUserRestrictions                   func(ctx context.Context, owner string, repo string, branch string, users []string) ([]*github.User, *github.Response, error)
	CompareCommits                        func(ctx context.Context, owner string, repo string, base string, head string, opts *github.ListOptions) (*github.CommitsComparison, *github.Response, error)
	CompareCommitsRaw                     func(ctx context.Context, owner string, repo string, base string, head string, opts github.RawOptions) (string, *github.Response, error)
	Create                                func(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error)
	CreateComment                         func(ctx context.Context, owner string, repo string, sha string, comment *github.RepositoryComment) (*github.RepositoryComment, *github.Response, error)
	CreateCustomDeploymentProtectionRule  func(ctx context.Context, owner string, repo string, environment string, request *github.CustomDeploymentProtectionRuleRequest) (*github.CustomDeploymentProtectionRule, *github.Response, error)
	CreateDeployment                      func(ctx context.Context, owner string, repo string, request *github.DeploymentRequest) (*github.Deployment, *github.Response, error)
	CreateDeploymentBranchPolicy          func(ctx context.Context, owner string, repo string, environment string, request *github.DeploymentBranchPolicyRequest) (*github.DeploymentBranchPolicy, *github.Response, error)
	CreateDeploymentStatus                func(ctx context.Context, owner string, repo string, deployment int64, request *github.DeploymentStatusRequest) (*github.DeploymentStatus, *github.Response, error)
	CreateFile                            func(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)
	CreateFork                            func(ctx context.Context, owner string, repo string, opts *github.RepositoryCreateForkOptions) (*github.Repository, *github.Response, error)
	CreateFromTemplate                    func(ctx context.Context, templateOwner string, templateRepo string, templateRepoReq *github.TemplateRepoRequest) (*github.Repository, *github.Response, error)
	CreateHook                            func(ctx context.Context, owner string, repo string, hook *github.Hook) (*github.Hook, *github.Response, error)
	CreateKey                             func(ctx context.Context, owner string, repo string, key *github.Key) (*github.Key, *github.Response, error)
	CreateOrUpdateCustomProperties        func(ctx context.Context, org string, repo string, customPropertyValues []*github.CustomPropertyValue) (*github.Response, error)
	CreateProject                         func(ctx context.Context, owner string, repo string, opts *github.ProjectOptions) (*github.Project, *github.Response, error)
	CreateRelease                         func(ctx context.Context, owner string, repo string, release *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error)
	CreateRuleset                         func(ctx context.Context, owner string, repo string, rs *github.Ruleset) (*github.Ruleset, *github.Response, error)
	CreateStatus                          func(ctx context.Context, owner string, repo string, ref string, status *github.RepoStatus) (*github.RepoStatus, *github.Response, error)
	CreateTagProtection                   func(ctx context.Context, owner string, repo string, pattern string) (*github.TagProtection, *github.Response, error)
	CreateUpdateEnvironment               func(ctx context.Context, owner string, repo string, name string, environment *github.CreateUpdateEnvironment) (*github.Environment, *github.Response, error)
	Delete                                func(ctx context.Context, owner string, repo string) (*github.Response, error)
	DeleteAutolink                        func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	DeleteComment                         func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	DeleteDeployment                      func(ctx context.Context, owner string, repo string, deploymentID int64) (*github.Response, error)
	DeleteDeploymentBranchPolicy          func(ctx context.Context, owner string, repo string, environment string, branchPolicyID int64) (*github.Response, error)
	DeleteEnvironment                     func(ctx context.Context, owner string, repo string, name string) (*github.Response, error)
	DeleteFile                            func(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)
	DeleteHook                            func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	DeleteInvitation                      func(ctx context.Context, owner string, repo string, invitationID int64) (*github.Response, error)
	DeleteKey                             func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	DeletePreReceiveHook                  func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	DeleteRelease                         func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	DeleteReleaseAsset                    func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	DeleteRuleset                         func(ctx context.Context, owner string, repo string, rulesetID int64) (*github.Response, error)
	DeleteTagProtection                   func(ctx context.Context, owner string, repo string, tagProtectionID int64) (*github.Response, error)
	DisableAutomatedSecurityFixes         func(ctx context.Context, owner string, repository string) (*github.Response, error)
	DisableCustomDeploymentProtectionRule func(ctx context.Context, owner string, repo string, environment string, protectionRuleID int64) (*github.Response, error)
	DisableDismissalRestrictions          func(ctx context.Context, owner string, repo string, branch string) (*github.PullRequestReviewsEnforcement, *github.Response, error)
	DisableLFS                            func(ctx context.Context, owner string, repo string) (*github.Response, error)
	DisablePages                          func(ctx context.Context, owner string, repo string) (*github.Response, error)
	DisablePrivateReporting               func(ctx context.Context, owner string, repo string) (*github.Response, error)
	DisableVulnerabilityAlerts            func(ctx context.Context, owner string, repository string) (*github.Response, error)
	Dispatch                              func(ctx context.Context, owner string, repo string, opts github.DispatchRequestOptions) (*github.Repository, *github.Response, error)
	DownloadContents                      func(ctx context.Context, owner string, repo string, filepath string, opts *github.RepositoryContentGetOptions) (io.ReadCloser, *github.Response, error)
	DownloadContentsWithMeta              func(ctx context.Context, owner string, repo string, filepath string, opts *github.RepositoryContentGetOptions) (io.ReadCloser, *github.RepositoryContent, *github.Response, error)
	DownloadReleaseAsset                  func(ctx context.Context, owner string, repo string, id int64, followRedirectsClient *http.Client) (rc io.ReadCloser, redirectURL string, err error)
	Edit                                  func(ctx context.Context, owner string, repo string, repository *github.Repository) (*github.Repository, *github.Response, error)
	EditActionsAccessLevel                func(ctx context.Context, owner string, repo string, repositoryActionsAccessLevel github.RepositoryActionsAccessLevel) (*github.Response, error)
	EditActionsAllowed                    func(ctx context.Context, org string, repo string, actionsAllowed github.ActionsAllowed) (*github.ActionsAllowed, *github.Response, error)
	EditActionsPermissions                func(ctx context.Context, owner string, repo string, actionsPermissionsRepository github.ActionsPermissionsRepository) (*github.ActionsPermissionsRepository, *github.Response, error)
	EditDefaultWorkflowPermissions        func(ctx context.Context, owner string, repo string, permissions github.DefaultWorkflowPermissionRepository) (*github.DefaultWorkflowPermissionRepository, *github.Response, error)
	EditHook                              func(ctx context.Context, owner string, repo string, id int64, hook *github.Hook) (*github.Hook, *github.Response, error)
	EditHookConfiguration                 func(ctx context.Context, owner string, repo string, id int64, config *github.HookConfig) (*github.HookConfig, *github.Response, error)
	EditRelease                           func(ctx context.Context, owner string, repo string, id int64, release *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error)
	EditReleaseAsset                      func(ctx context.Context, owner string, repo string, id int64, release *github.ReleaseAsset) (*github.ReleaseAsset, *github.Response, error)
	EnableAutomatedSecurityFixes          func(ctx context.Context, owner string, repository string) (*github.Response, error)
	EnableLFS                             func(ctx context.Context, owner string, repo string) (*github.Response, error)
	EnablePages                           func(ctx context.Context, owner string, repo string, pages *github.Pages) (*github.Pages, *github.Response, error)
	EnablePrivateReporting                func(ctx context.Context, owner string, repo string) (*github.Response, error)
	EnableVulnerabilityAlerts             func(ctx context.Context, owner string, repository string) (*github.Response, error)
	GenerateReleaseNotes                  func(ctx context.Context, owner string, repo string, opts *github.GenerateNotesOptions) (*github.RepositoryReleaseNotes, *github.Response, error)
	Get                                   func(ctx context.Context, owner string, repo string) (*github.Repository, *github.Response, error)
	GetActionsAccessLevel                 func(ctx context.Context, owner string, repo string) (*github.RepositoryActionsAccessLevel, *github.Response, error)
	GetActionsAllowed                     func(ctx context.Context, org string, repo string) (*github.ActionsAllowed, *github.Response, error)
	GetActionsPermissions                 func(ctx context.Context, owner string, repo string) (*github.ActionsPermissionsRepository, *github.Response, error)
	GetAdminEnforcement                   func(ctx context.Context, owner string, repo string, branch string) (*github.AdminEnforcement, *github.Response, error)
	GetAllCustomPropertyValues            func(ctx context.Context, org string, repo string) ([]*github.CustomPropertyValue, *github.Response, error)
	GetAllDeploymentProtectionRules       func(ctx context.Context, owner string, repo string, environment string) (*github.ListDeploymentProtectionRuleResponse, *github.Response, error)
	GetAllRulesets                        func(ctx context.Context, owner string, repo string, includesParents bool) ([]*github.Ruleset, *github.Response, error)
	GetArchiveLink                        func(ctx context.Context, owner string, repo string, archiveformat github.ArchiveFormat, opts *github.RepositoryContentGetOptions, maxRedirects int) (*url.URL, *github.Response, error)
	GetAutolink                           func(ctx context.Context, owner string, repo string, id int64) (*github.Autolink, *github.Response, error)
	GetAutomatedSecurityFixes             func(ctx context.Context, owner string, repository string) (*github.AutomatedSecurityFixes, *github.Response, error)
	GetBranch                             func(ctx context.Context, owner string, repo string, branch string, maxRedirects int) (*github.Branch, *github.Response, error)
	GetBranchProtection                   func(ctx context.Context, owner string, repo string, branch string) (*github.Protection, *github.Response, error)
	GetByID                               func(ctx context.Context, id int64) (*github.Repository, *github.Response, error)
	GetCodeOfConduct                      func(ctx context.Context, owner string, repo string) (*github.CodeOfConduct, *github.Response, error)
	GetCodeownersErrors                   func(ctx context.Context, owner string, repo string, opts *github.GetCodeownersErrorsOptions) (*github.CodeownersErrors, *github.Response, error)
	GetCombinedStatus                     func(ctx context.Context, owner string, repo string, ref string, opts *github.ListOptions) (*github.CombinedStatus, *github.Response, error)
	GetComment                            func(ctx context.Context, owner string, repo string, id int64) (*github.RepositoryComment, *github.Response, error)
	GetCommit                             func(ctx context.Context, owner string, repo string, sha string, opts *github.ListOptions) (*github.RepositoryCommit, *github.Response, error)
	GetCommitRaw                          func(ctx context.Context, owner string, repo string, sha string, opts github.RawOptions) (string, *github.Response, error)
	GetCommitSHA1                         func(ctx context.Context, owner string, repo string, ref string, lastSHA string) (string, *github.Response, error)
	GetCommunityHealthMetrics             func(ctx context.Context, owner string, repo string) (*github.CommunityHealthMetrics, *github.Response, error)
	GetContents                           func(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error)
	GetCustomDeploymentProtectionRule     func(ctx context.Context, owner string, repo string, environment string, protectionRuleID int64) (*github.CustomDeploymentProtectionRule, *github.Response, error)
	GetDefaultWorkflowPermissions         func(ctx context.Context, owner string, repo string) (*github.DefaultWorkflowPermissionRepository, *github.Response, error)
	GetDeployment                         func(ctx context.Context, owner string, repo string, deploymentID int64) (*github.Deployment, *github.Response, error)
	GetDeploymentBranchPolicy             func(ctx context.Context, owner string, repo string, environment string, branchPolicyID int64) (*github.DeploymentBranchPolicy, *github.Response, error)
	GetDeploymentStatus                   func(ctx context.Context, owner string, repo string, deploymentID int64, deploymentStatusID int64) (*github.DeploymentStatus, *github.Response, error)
	GetEnvironment                        func(ctx context.Context, owner string, repo string, name string) (*github.Environment, *github.Response, error)
	GetHook                               func(ctx context.Context, owner string, repo string, id int64) (*github.Hook, *github.Response, error)
	GetHookConfiguration                  func(ctx context.Context, owner string, repo string, id int64) (*github.HookConfig, *github.Response, error)
	GetHookDelivery                       func(ctx context.Context, owner string, repo string, hookID int64, deliveryID int64) (*github.HookDelivery, *github.Response, error)
	GetKey                                func(ctx context.Context, owner string, repo string, id int64) (*github.Key, *github.Response, error)
	GetLatestPagesBuild                   func(ctx context.Context, owner string, repo string) (*github.PagesBuild, *github.Response, error)
	GetLatestRelease                      func(ctx context.Context, owner string, repo string) (*github.RepositoryRelease, *github.Response, error)
	GetPageBuild                          func(ctx context.Context, owner string, repo string, id int64) (*github.PagesBuild, *github.Response, error)
	GetPageHealthCheck                    func(ctx context.Context, owner string, repo string) (*github.PagesHealthCheckResponse, *github.Response, error)
	GetPagesInfo                          func(ctx context.Context, owner string, repo string) (*github.Pages, *github.Response, error)
	GetPermissionLevel                    func(ctx context.Context, owner string, repo string, user string) (*github.RepositoryPermissionLevel, *github.Response, error)
	GetPreReceiveHook                     func(ctx context.Context, owner string, repo string, id int64) (*github.PreReceiveHook, *github.Response, error)
	GetPullRequestReviewEnforcement       func(ctx context.Context, owner string, repo string, branch string) (*github.PullRequestReviewsEnforcement, *github.Response, error)
	GetReadme                             func(ctx context.Context, owner string, repo string, opts *github.RepositoryContentGetOptions) (*github.RepositoryContent, *github.Response, error)
	GetRelease                            func(ctx context.Context, owner string, repo string, id int64) (*github.RepositoryRelease, *github.Response, error)
	GetReleaseAsset                       func(ctx context.Context, owner string, repo string, id int64) (*github.ReleaseAsset, *github.Response, error)
	GetReleaseByTag                       func(ctx context.Context, owner string, repo string, tag string) (*github.RepositoryRelease, *github.Response, error)
	GetRequiredStatusChecks               func(ctx context.Context, owner string, repo string, branch string) (*github.RequiredStatusChecks, *github.Response, error)
	GetRulesForBranch                     func(ctx context.Context, owner string, repo string, branch string) ([]*github.RepositoryRule, *github.Response, error)
	GetRuleset                            func(ctx context.Context, owner string, repo string, rulesetID int64, includesParents bool) (*github.Ruleset, *github.Response, error)
	GetSignaturesProtectedBranch          func(ctx context.Context, owner string, repo string, branch string) (*github.SignaturesProtectedBranch, *github.Response, error)
	GetVulnerabilityAlerts                func(ctx context.Context, owner string, repository string) (bool, *github.Response, error)
	IsCollaborator                        func(ctx context.Context, owner string, repo string, user string) (bool, *github.Response, error)
	IsPrivateReportingEnabled             func(ctx context.Context, owner string, repo string) (bool, *github.Response, error)
	License                               func(ctx context.Context, owner string, repo string) (*github.RepositoryLicense, *github.Response, error)
	ListAll                               func(ctx context.Context, opts *github.RepositoryListAllOptions) ([]*github.Repository, *github.Response, error)
	ListAllTopics                         func(ctx context.Context, owner string, repo string) ([]string, *github.Response, error)
	ListAppRestrictions                   func(ctx context.Context, owner string, repo string, branch string) ([]*github.App, *github.Response, error)
	ListAutolinks                         func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Autolink, *github.Response, error)
	ListBranches                          func(ctx context.Context, owner string, repo string, opts *github.BranchListOptions) ([]*github.Branch, *github.Response, error)
	ListBranchesHeadCommit                func(ctx context.Context, owner string, repo string, sha string) ([]*github.BranchCommit, *github.Response, error)
	ListByAuthenticatedUser               func(ctx context.Context, opts *github.RepositoryListByAuthenticatedUserOptions) ([]*github.Repository, *github.Response, error)
	ListByOrg                             func(ctx context.Context, org string, opts *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error)
	ListByUser                            func(ctx context.Context, user string, opts *github.RepositoryListByUserOptions) ([]*github.Repository, *github.Response, error)
	ListCodeFrequency                     func(ctx context.Context, owner string, repo string) ([]*github.WeeklyStats, *github.Response, error)
	ListCollaborators                     func(ctx context.Context, owner string, repo string, opts *github.ListCollaboratorsOptions) ([]*github.User, *github.Response, error)
	ListComments                          func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryComment, *github.Response, error)
	ListCommitActivity                    func(ctx context.Context, owner string, repo string) ([]*github.WeeklyCommitActivity, *github.Response, error)
	ListCommitComments                    func(ctx context.Context, owner string, repo string, sha string, opts *github.ListOptions) ([]*github.RepositoryComment, *github.Response, error)
	ListCommits                           func(ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error)
	ListContributors                      func(ctx context.Context, owner string, repository string, opts *github.ListContributorsOptions) ([]*github.Contributor, *github.Response, error)
	ListContributorsStats                 func(ctx context.Context, owner string, repo string) ([]*github.ContributorStats, *github.Response, error)
	ListCustomDeploymentRuleIntegrations  func(ctx context.Context, owner string, repo string, environment string) (*github.ListCustomDeploymentRuleIntegrationsResponse, *github.Response, error)
	ListDeploymentBranchPolicies          func(ctx context.Context, owner string, repo string, environment string) (*github.DeploymentBranchPolicyResponse, *github.Response, error)
	ListDeploymentStatuses                func(ctx context.Context, owner string, repo string, deployment int64, opts *github.ListOptions) ([]*github.DeploymentStatus, *github.Response, error)
	ListDeployments                       func(ctx context.Context, owner string, repo string, opts *github.DeploymentsListOptions) ([]*github.Deployment, *github.Response, error)
	ListEnvironments                      func(ctx context.Context, owner string, repo string, opts *github.EnvironmentListOptions) (*github.EnvResponse, *github.Response, error)
	ListForks                             func(ctx context.Context, owner string, repo string, opts *github.RepositoryListForksOptions) ([]*github.Repository, *github.Response, error)
	ListHookDeliveries                    func(ctx context.Context, owner string, repo string, id int64, opts *github.ListCursorOptions) ([]*github.HookDelivery, *github.Response, error)
	ListHooks                             func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Hook, *github.Response, error)
	ListInvitations                       func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryInvitation, *github.Response, error)
	ListKeys                              func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Key, *github.Response, error)
	ListLanguages                         func(ctx context.Context, owner string, repo string) (map[string]int, *github.Response, error)
	ListPagesBuilds                       func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.PagesBuild, *github.Response, error)
	ListParticipation                     func(ctx context.Context, owner string, repo string) (*github.RepositoryParticipation, *github.Response, error)
	ListPreReceiveHooks                   func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.PreReceiveHook, *github.Response, error)
	ListProjects                          func(ctx context.Context, owner string, repo string, opts *github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListPunchCard                         func(ctx context.Context, owner string, repo string) ([]*github.PunchCard, *github.Response, error)
	ListReleaseAssets                     func(ctx context.Context, owner string, repo string, id int64, opts *github.ListOptions) ([]*github.ReleaseAsset, *github.Response, error)
	ListReleases                          func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error)
	ListRequiredStatusChecksContexts      func(ctx context.Context, owner string, repo string, branch string) (contexts []string, resp *github.Response, err error)
	ListStatuses                          func(ctx context.Context, owner string, repo string, ref string, opts *github.ListOptions) ([]*github.RepoStatus, *github.Response, error)
	ListTagProtection                     func(ctx context.Context, owner string, repo string) ([]*github.TagProtection, *github.Response, error)
	ListTags                              func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryTag, *github.Response, error)
	ListTeamRestrictions                  func(ctx context.Context, owner string, repo string, branch string) ([]*github.Team, *github.Response, error)
	ListTeams                             func(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	ListTrafficClones                     func(ctx context.Context, owner string, repo string, opts *github.TrafficBreakdownOptions) (*github.TrafficClones, *github.Response, error)
	ListTrafficPaths                      func(ctx context.Context, owner string, repo string) ([]*github.TrafficPath, *github.Response, error)
	ListTrafficReferrers                  func(ctx context.Context, owner string, repo string) ([]*github.TrafficReferrer, *github.Response, error)
	ListTrafficViews                      func(ctx context.Context, owner string, repo string, opts *github.TrafficBreakdownOptions) (*github.TrafficViews, *github.Response, error)
	ListUserRestrictions                  func(ctx context.Context, owner string, repo string, branch string) ([]*github.User, *github.Response, error)
	Merge                                 func(ctx context.Context, owner string, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error)
	MergeUpstream                         func(ctx context.Context, owner string, repo string, request *github.RepoMergeUpstreamRequest) (*github.RepoMergeUpstreamResult, *github.Response, error)
	OptionalSignaturesOnProtectedBranch   func(ctx context.Context, owner string, repo string, branch string) (*github.Response, error)
	PingHook                              func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	RedeliverHookDelivery                 func(ctx context.Context, owner string, repo string, hookID int64, deliveryID int64) (*github.HookDelivery, *github.Response, error)
	RemoveAdminEnforcement                func(ctx context.Context, owner string, repo string, branch string) (*github.Response, error)
	RemoveAppRestrictions                 func(ctx context.Context, owner string, repo string, branch string, apps []string) ([]*github.App, *github.Response, error)
	RemoveBranchProtection                func(ctx context.Context, owner string, repo string, branch string) (*github.Response, error)
	RemoveCollaborator                    func(ctx context.Context, owner string, repo string, user string) (*github.Response, error)
	RemovePullRequestReviewEnforcement    func(ctx context.Context, owner string, repo string, branch string) (*github.Response, error)
	RemoveRequiredStatusChecks            func(ctx context.Context, owner string, repo string, branch string) (*github.Response, error)
	RemoveTeamRestrictions                func(ctx context.Context, owner string, repo string, branch string, teams []string) ([]*github.Team, *github.Response, error)
	RemoveUserRestrictions                func(ctx context.Context, owner string, repo string, branch string, users []string) ([]*github.User, *github.Response, error)
	RenameBranch                          func(ctx context.Context, owner string, repo string, branch string, newName string) (*github.Branch, *github.Response, error)
	ReplaceAllTopics                      func(ctx context.Context, owner string, repo string, topics []string) ([]string, *github.Response, error)
	ReplaceAppRestrictions                func(ctx context.Context, owner string, repo string, branch string, apps []string) ([]*github.App, *github.Response, error)
	ReplaceTeamRestrictions               func(ctx context.Context, owner string, repo string, branch string, teams []string) ([]*github.Team, *github.Response, error)
	ReplaceUserRestrictions               func(ctx context.Context, owner string, repo string, branch string, users []string) ([]*github.User, *github.Response, error)
	RequestPageBuild                      func(ctx context.Context, owner string, repo string) (*github.PagesBuild, *github.Response, error)
	RequireSignaturesOnProtectedBranch    func(ctx context.Context, owner string, repo string, branch string) (*github.SignaturesProtectedBranch, *github.Response, error)
	Subscribe                             func(ctx context.Context, owner string, repo string, event string, callback string, secret []byte) (*github.Response, error)
	TestHook                              func(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	Transfer                              func(ctx context.Context, owner string, repo string, transfer github.TransferRequest) (*github.Repository, *github.Response, error)
	Unsubscribe                           func(ctx context.Context, owner string, repo string, event string, callback string, secret []byte) (*github.Response, error)
	UpdateBranchProtection                func(ctx context.Context, owner string, repo string, branch string, preq *github.ProtectionRequest) (*github.Protection, *github.Response, error)
	UpdateComment                         func(ctx context.Context, owner string, repo string, id int64, comment *github.RepositoryComment) (*github.RepositoryComment, *github.Response, error)
	UpdateDeploymentBranchPolicy          func(ctx context.Context, owner string, repo string, environment string, branchPolicyID int64, request *github.DeploymentBranchPolicyRequest) (*github.DeploymentBranchPolicy, *github.Response, error)
	UpdateFile                            func(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)
	UpdateInvitation                      func(ctx context.Context, owner string, repo string, invitationID int64, permissions string) (*github.RepositoryInvitation, *github.Response, error)
	UpdatePages                           func(ctx context.Context, owner string, repo string, opts *github.PagesUpdate) (*github.Response, error)
	UpdatePreReceiveHook                  func(ctx context.Context, owner string, repo string, id int64, hook *github.PreReceiveHook) (*github.PreReceiveHook, *github.Response, error)
	UpdatePullRequestReviewEnforcement    func(ctx context.Context, owner string, repo string, branch string, patch *github.PullRequestReviewsEnforcementUpdate) (*github.PullRequestReviewsEnforcement, *github.Response, error)
	UpdateRequiredStatusChecks            func(ctx context.Context, owner string, repo string, branch string, sreq *github.RequiredStatusChecksRequest) (*github.RequiredStatusChecks, *github.Response, error)
	UpdateRuleset                         func(ctx context.Context, owner string, repo string, rulesetID int64, rs *github.Ruleset) (*github.Ruleset, *github.Response, error)
	UploadReleaseAsset                    func(ctx context.Context, owner string, repo string, id int64, opts *github.UploadOptions, file *os.File) (*github.ReleaseAsset, *github.Response, error)
}

type SCIMService struct {
	DeleteSCIMUserFromOrg          func(ctx context.Context, org string, scimUserID string) (*github.Response, error)
	GetSCIMProvisioningInfoForUser func(ctx context.Context, org string, scimUserID string) (*github.SCIMUserAttributes, *github.Response, error)
	ListSCIMProvisionedIdentities  func(ctx context.Context, org string, opts *github.ListSCIMProvisionedIdentitiesOptions) (*github.SCIMProvisionedIdentities, *github.Response, error)
	ProvisionAndInviteSCIMUser     func(ctx context.Context, org string, opts *github.SCIMUserAttributes) (*github.Response, error)
	UpdateAttributeForSCIMUser     func(ctx context.Context, org string, scimUserID string, opts *github.UpdateAttributeForSCIMUserOptions) (*github.Response, error)
	UpdateProvisionedOrgMembership func(ctx context.Context, org string, scimUserID string, opts *github.SCIMUserAttributes) (*github.Response, error)
}

type SecretScanningService struct {
	GetAlert                func(ctx context.Context, owner string, repo string, number int64) (*github.SecretScanningAlert, *github.Response, error)
	ListAlertsForEnterprise func(ctx context.Context, enterprise string, opts *github.SecretScanningAlertListOptions) ([]*github.SecretScanningAlert, *github.Response, error)
	ListAlertsForOrg        func(ctx context.Context, org string, opts *github.SecretScanningAlertListOptions) ([]*github.SecretScanningAlert, *github.Response, error)
	ListAlertsForRepo       func(ctx context.Context, owner string, repo string, opts *github.SecretScanningAlertListOptions) ([]*github.SecretScanningAlert, *github.Response, error)
	ListLocationsForAlert   func(ctx context.Context, owner string, repo string, number int64, opts *github.ListOptions) ([]*github.SecretScanningAlertLocation, *github.Response, error)
	UpdateAlert             func(ctx context.Context, owner string, repo string, number int64, opts *github.SecretScanningAlertUpdateOptions) (*github.SecretScanningAlert, *github.Response, error)
}

type SecurityAdvisoriesService struct {
	CreateTemporaryPrivateFork             func(ctx context.Context, owner string, repo string, ghsaID string) (*github.Repository, *github.Response, error)
	GetGlobalSecurityAdvisories            func(ctx context.Context, ghsaID string) (*github.GlobalSecurityAdvisory, *github.Response, error)
	ListGlobalSecurityAdvisories           func(ctx context.Context, opts *github.ListGlobalSecurityAdvisoriesOptions) ([]*github.GlobalSecurityAdvisory, *github.Response, error)
	ListRepositorySecurityAdvisories       func(ctx context.Context, owner string, repo string, opt *github.ListRepositorySecurityAdvisoriesOptions) ([]*github.SecurityAdvisory, *github.Response, error)
	ListRepositorySecurityAdvisoriesForOrg func(ctx context.Context, org string, opt *github.ListRepositorySecurityAdvisoriesOptions) ([]*github.SecurityAdvisory, *github.Response, error)
	RequestCVE                             func(ctx context.Context, owner string, repo string, ghsaID string) (*github.Response, error)
}

type TeamsService struct {
	AddTeamMembershipByID                   func(ctx context.Context, orgID int64, teamID int64, user string, opts *github.TeamAddTeamMembershipOptions) (*github.Membership, *github.Response, error)
	AddTeamMembershipBySlug                 func(ctx context.Context, org string, slug string, user string, opts *github.TeamAddTeamMembershipOptions) (*github.Membership, *github.Response, error)
	AddTeamProjectByID                      func(ctx context.Context, orgID int64, teamID int64, projectID int64, opts *github.TeamProjectOptions) (*github.Response, error)
	AddTeamProjectBySlug                    func(ctx context.Context, org string, slug string, projectID int64, opts *github.TeamProjectOptions) (*github.Response, error)
	AddTeamRepoByID                         func(ctx context.Context, orgID int64, teamID int64, owner string, repo string, opts *github.TeamAddTeamRepoOptions) (*github.Response, error)
	AddTeamRepoBySlug                       func(ctx context.Context, org string, slug string, owner string, repo string, opts *github.TeamAddTeamRepoOptions) (*github.Response, error)
	CreateCommentByID                       func(ctx context.Context, orgID int64, teamID int64, discsusionNumber int, comment github.DiscussionComment) (*github.DiscussionComment, *github.Response, error)
	CreateCommentBySlug                     func(ctx context.Context, org string, slug string, discsusionNumber int, comment github.DiscussionComment) (*github.DiscussionComment, *github.Response, error)
	CreateDiscussionByID                    func(ctx context.Context, orgID int64, teamID int64, discussion github.TeamDiscussion) (*github.TeamDiscussion, *github.Response, error)
	CreateDiscussionBySlug                  func(ctx context.Context, org string, slug string, discussion github.TeamDiscussion) (*github.TeamDiscussion, *github.Response, error)
	CreateOrUpdateIDPGroupConnectionsByID   func(ctx context.Context, orgID int64, teamID int64, opts github.IDPGroupList) (*github.IDPGroupList, *github.Response, error)
	CreateOrUpdateIDPGroupConnectionsBySlug func(ctx context.Context, org string, slug string, opts github.IDPGroupList) (*github.IDPGroupList, *github.Response, error)
	CreateTeam                              func(ctx context.Context, org string, team github.NewTeam) (*github.Team, *github.Response, error)
	DeleteCommentByID                       func(ctx context.Context, orgID int64, teamID int64, discussionNumber int, commentNumber int) (*github.Response, error)
	DeleteCommentBySlug                     func(ctx context.Context, org string, slug string, discussionNumber int, commentNumber int) (*github.Response, error)
	DeleteDiscussionByID                    func(ctx context.Context, orgID int64, teamID int64, discussionNumber int) (*github.Response, error)
	DeleteDiscussionBySlug                  func(ctx context.Context, org string, slug string, discussionNumber int) (*github.Response, error)
	DeleteTeamByID                          func(ctx context.Context, orgID int64, teamID int64) (*github.Response, error)
	DeleteTeamBySlug                        func(ctx context.Context, org string, slug string) (*github.Response, error)
	EditCommentByID                         func(ctx context.Context, orgID int64, teamID int64, discussionNumber int, commentNumber int, comment github.DiscussionComment) (*github.DiscussionComment, *github.Response, error)
	EditCommentBySlug                       func(ctx context.Context, org string, slug string, discussionNumber int, commentNumber int, comment github.DiscussionComment) (*github.DiscussionComment, *github.Response, error)
	EditDiscussionByID                      func(ctx context.Context, orgID int64, teamID int64, discussionNumber int, discussion github.TeamDiscussion) (*github.TeamDiscussion, *github.Response, error)
	EditDiscussionBySlug                    func(ctx context.Context, org string, slug string, discussionNumber int, discussion github.TeamDiscussion) (*github.TeamDiscussion, *github.Response, error)
	EditTeamByID                            func(ctx context.Context, orgID int64, teamID int64, team github.NewTeam, removeParent bool) (*github.Team, *github.Response, error)
	EditTeamBySlug                          func(ctx context.Context, org string, slug string, team github.NewTeam, removeParent bool) (*github.Team, *github.Response, error)
	GetCommentByID                          func(ctx context.Context, orgID int64, teamID int64, discussionNumber int, commentNumber int) (*github.DiscussionComment, *github.Response, error)
	GetCommentBySlug                        func(ctx context.Context, org string, slug string, discussionNumber int, commentNumber int) (*github.DiscussionComment, *github.Response, error)
	GetDiscussionByID                       func(ctx context.Context, orgID int64, teamID int64, discussionNumber int) (*github.TeamDiscussion, *github.Response, error)
	GetDiscussionBySlug                     func(ctx context.Context, org string, slug string, discussionNumber int) (*github.TeamDiscussion, *github.Response, error)
	GetExternalGroup                        func(ctx context.Context, org string, groupID int64) (*github.ExternalGroup, *github.Response, error)
	GetTeamByID                             func(ctx context.Context, orgID int64, teamID int64) (*github.Team, *github.Response, error)
	GetTeamBySlug                           func(ctx context.Context, org string, slug string) (*github.Team, *github.Response, error)
	GetTeamMembershipByID                   func(ctx context.Context, orgID int64, teamID int64, user string) (*github.Membership, *github.Response, error)
	GetTeamMembershipBySlug                 func(ctx context.Context, org string, slug string, user string) (*github.Membership, *github.Response, error)
	IsTeamRepoByID                          func(ctx context.Context, orgID int64, teamID int64, owner string, repo string) (*github.Repository, *github.Response, error)
	IsTeamRepoBySlug                        func(ctx context.Context, org string, slug string, owner string, repo string) (*github.Repository, *github.Response, error)
	ListChildTeamsByParentID                func(ctx context.Context, orgID int64, teamID int64, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	ListChildTeamsByParentSlug              func(ctx context.Context, org string, slug string, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	ListCommentsByID                        func(ctx context.Context, orgID int64, teamID int64, discussionNumber int, options *github.DiscussionCommentListOptions) ([]*github.DiscussionComment, *github.Response, error)
	ListCommentsBySlug                      func(ctx context.Context, org string, slug string, discussionNumber int, options *github.DiscussionCommentListOptions) ([]*github.DiscussionComment, *github.Response, error)
	ListDiscussionsByID                     func(ctx context.Context, orgID int64, teamID int64, opts *github.DiscussionListOptions) ([]*github.TeamDiscussion, *github.Response, error)
	ListDiscussionsBySlug                   func(ctx context.Context, org string, slug string, opts *github.DiscussionListOptions) ([]*github.TeamDiscussion, *github.Response, error)
	ListExternalGroups                      func(ctx context.Context, org string, opts *github.ListExternalGroupsOptions) (*github.ExternalGroupList, *github.Response, error)
	ListExternalGroupsForTeamBySlug         func(ctx context.Context, org string, slug string) (*github.ExternalGroupList, *github.Response, error)
	ListIDPGroupsForTeamByID                func(ctx context.Context, orgID int64, teamID int64) (*github.IDPGroupList, *github.Response, error)
	ListIDPGroupsForTeamBySlug              func(ctx context.Context, org string, slug string) (*github.IDPGroupList, *github.Response, error)
	ListIDPGroupsInOrganization             func(ctx context.Context, org string, opts *github.ListCursorOptions) (*github.IDPGroupList, *github.Response, error)
	ListPendingTeamInvitationsByID          func(ctx context.Context, orgID int64, teamID int64, opts *github.ListOptions) ([]*github.Invitation, *github.Response, error)
	ListPendingTeamInvitationsBySlug        func(ctx context.Context, org string, slug string, opts *github.ListOptions) ([]*github.Invitation, *github.Response, error)
	ListTeamMembersByID                     func(ctx context.Context, orgID int64, teamID int64, opts *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error)
	ListTeamMembersBySlug                   func(ctx context.Context, org string, slug string, opts *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error)
	ListTeamProjectsByID                    func(ctx context.Context, orgID int64, teamID int64) ([]*github.Project, *github.Response, error)
	ListTeamProjectsBySlug                  func(ctx context.Context, org string, slug string) ([]*github.Project, *github.Response, error)
	ListTeamReposByID                       func(ctx context.Context, orgID int64, teamID int64, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
	ListTeamReposBySlug                     func(ctx context.Context, org string, slug string, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
	ListTeams                               func(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	ListUserTeams                           func(ctx context.Context, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	RemoveConnectedExternalGroup            func(ctx context.Context, org string, slug string) (*github.Response, error)
	RemoveTeamMembershipByID                func(ctx context.Context, orgID int64, teamID int64, user string) (*github.Response, error)
	RemoveTeamMembershipBySlug              func(ctx context.Context, org string, slug string, user string) (*github.Response, error)
	RemoveTeamProjectByID                   func(ctx context.Context, orgID int64, teamID int64, projectID int64) (*github.Response, error)
	RemoveTeamProjectBySlug                 func(ctx context.Context, org string, slug string, projectID int64) (*github.Response, error)
	RemoveTeamRepoByID                      func(ctx context.Context, orgID int64, teamID int64, owner string, repo string) (*github.Response, error)
	RemoveTeamRepoBySlug                    func(ctx context.Context, org string, slug string, owner string, repo string) (*github.Response, error)
	ReviewTeamProjectsByID                  func(ctx context.Context, orgID int64, teamID int64, projectID int64) (*github.Project, *github.Response, error)
	ReviewTeamProjectsBySlug                func(ctx context.Context, org string, slug string, projectID int64) (*github.Project, *github.Response, error)
	UpdateConnectedExternalGroup            func(ctx context.Context, org string, slug string, eg *github.ExternalGroup) (*github.ExternalGroup, *github.Response, error)
}

type UsersService struct {
	AcceptInvitation      func(ctx context.Context, invitationID int64) (*github.Response, error)
	AddEmails             func(ctx context.Context, emails []string) ([]*github.UserEmail, *github.Response, error)
	BlockUser             func(ctx context.Context, user string) (*github.Response, error)
	CreateGPGKey          func(ctx context.Context, armoredPublicKey string) (*github.GPGKey, *github.Response, error)
	CreateKey             func(ctx context.Context, key *github.Key) (*github.Key, *github.Response, error)
	CreateProject         func(ctx context.Context, opts *github.CreateUserProjectOptions) (*github.Project, *github.Response, error)
	CreateSSHSigningKey   func(ctx context.Context, key *github.Key) (*github.SSHSigningKey, *github.Response, error)
	DeclineInvitation     func(ctx context.Context, invitationID int64) (*github.Response, error)
	DeleteEmails          func(ctx context.Context, emails []string) (*github.Response, error)
	DeleteGPGKey          func(ctx context.Context, id int64) (*github.Response, error)
	DeleteKey             func(ctx context.Context, id int64) (*github.Response, error)
	DeletePackage         func(ctx context.Context, user string, packageType string, packageName string) (*github.Response, error)
	DeleteSSHSigningKey   func(ctx context.Context, id int64) (*github.Response, error)
	DemoteSiteAdmin       func(ctx context.Context, user string) (*github.Response, error)
	Edit                  func(ctx context.Context, user *github.User) (*github.User, *github.Response, error)
	Follow                func(ctx context.Context, user string) (*github.Response, error)
	Get                   func(ctx context.Context, user string) (*github.User, *github.Response, error)
	GetByID               func(ctx context.Context, id int64) (*github.User, *github.Response, error)
	GetGPGKey             func(ctx context.Context, id int64) (*github.GPGKey, *github.Response, error)
	GetHovercard          func(ctx context.Context, user string, opts *github.HovercardOptions) (*github.Hovercard, *github.Response, error)
	GetKey                func(ctx context.Context, id int64) (*github.Key, *github.Response, error)
	GetPackage            func(ctx context.Context, user string, packageType string, packageName string) (*github.Package, *github.Response, error)
	GetSSHSigningKey      func(ctx context.Context, id int64) (*github.SSHSigningKey, *github.Response, error)
	IsBlocked             func(ctx context.Context, user string) (bool, *github.Response, error)
	IsFollowing           func(ctx context.Context, user string, target string) (bool, *github.Response, error)
	ListAll               func(ctx context.Context, opts *github.UserListOptions) ([]*github.User, *github.Response, error)
	ListBlockedUsers      func(ctx context.Context, opts *github.ListOptions) ([]*github.User, *github.Response, error)
	ListEmails            func(ctx context.Context, opts *github.ListOptions) ([]*github.UserEmail, *github.Response, error)
	ListFollowers         func(ctx context.Context, user string, opts *github.ListOptions) ([]*github.User, *github.Response, error)
	ListFollowing         func(ctx context.Context, user string, opts *github.ListOptions) ([]*github.User, *github.Response, error)
	ListGPGKeys           func(ctx context.Context, user string, opts *github.ListOptions) ([]*github.GPGKey, *github.Response, error)
	ListInvitations       func(ctx context.Context, opts *github.ListOptions) ([]*github.RepositoryInvitation, *github.Response, error)
	ListKeys              func(ctx context.Context, user string, opts *github.ListOptions) ([]*github.Key, *github.Response, error)
	ListPackages          func(ctx context.Context, user string, opts *github.PackageListOptions) ([]*github.Package, *github.Response, error)
	ListProjects          func(ctx context.Context, user string, opts *github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListSSHSigningKeys    func(ctx context.Context, user string, opts *github.ListOptions) ([]*github.SSHSigningKey, *github.Response, error)
	PackageDeleteVersion  func(ctx context.Context, user string, packageType string, packageName string, packageVersionID int64) (*github.Response, error)
	PackageGetAllVersions func(ctx context.Context, user string, packageType string, packageName string, opts *github.PackageListOptions) ([]*github.PackageVersion, *github.Response, error)
	PackageGetVersion     func(ctx context.Context, user string, packageType string, packageName string, packageVersionID int64) (*github.PackageVersion, *github.Response, error)
	PackageRestoreVersion func(ctx context.Context, user string, packageType string, packageName string, packageVersionID int64) (*github.Response, error)
	PromoteSiteAdmin      func(ctx context.Context, user string) (*github.Response, error)
	RestorePackage        func(ctx context.Context, user string, packageType string, packageName string) (*github.Response, error)
	SetEmailVisibility    func(ctx context.Context, visibility string) ([]*github.UserEmail, *github.Response, error)
	Suspend               func(ctx context.Context, user string, opts *github.UserSuspendOptions) (*github.Response, error)
	UnblockUser           func(ctx context.Context, user string) (*github.Response, error)
	Unfollow              func(ctx context.Context, user string) (*github.Response, error)
	Unsuspend             func(ctx context.Context, user string) (*github.Response, error)
}

// NewClient returns a new *ghx.Client with all calls routed to *github.Client
func NewClient(client *github.Client) *Client {
	return &Client{
		Client: client,
		Issues: newIssuesService(client),
		Licenses: &LicensesService{
			Get:  client.Licenses.Get,
			List: client.Licenses.List,
		},
		Markdown: &MarkdownService{
			Render: client.Markdown.Render,
		},
		Marketplace: &MarketplaceService{
			GetPlanAccountForAccount:        client.Marketplace.GetPlanAccountForAccount,
			ListMarketplacePurchasesForUser: client.Marketplace.ListMarketplacePurchasesForUser,
			ListPlanAccountsForPlan:         client.Marketplace.ListPlanAccountsForPlan,
			ListPlans:                       client.Marketplace.ListPlans,
		},
		Meta: &MetaService{
			Get:     client.Meta.Get,
			Octocat: client.Meta.Octocat,
			Zen:     client.Meta.Zen,
		},
		Migrations: &MigrationService{
			CancelImport:            client.Migrations.CancelImport,
			CommitAuthors:           client.Migrations.CommitAuthors,
			DeleteMigration:         client.Migrations.DeleteMigration,
			DeleteUserMigration:     client.Migrations.DeleteUserMigration,
			ImportProgress:          client.Migrations.ImportProgress,
			LargeFiles:              client.Migrations.LargeFiles,
			ListMigrations:          client.Migrations.ListMigrations,
			ListUserMigrations:      client.Migrations.ListUserMigrations,
			MapCommitAuthor:         client.Migrations.MapCommitAuthor,
			MigrationArchiveURL:     client.Migrations.MigrationArchiveURL,
			MigrationStatus:         client.Migrations.MigrationStatus,
			SetLFSPreference:        client.Migrations.SetLFSPreference,
			StartImport:             client.Migrations.StartImport,
			StartMigration:          client.Migrations.StartMigration,
			StartUserMigration:      client.Migrations.StartUserMigration,
			UnlockRepo:              client.Migrations.UnlockRepo,
			UnlockUserRepo:          client.Migrations.UnlockUserRepo,
			UpdateImport:            client.Migrations.UpdateImport,
			UserMigrationArchiveURL: client.Migrations.UserMigrationArchiveURL,
			UserMigrationStatus:     client.Migrations.UserMigrationStatus,
		},
		Organizations: &OrganizationsService{
			AddSecurityManagerTeam:                 client.Organizations.AddSecurityManagerTeam,
			BlockUser:                              client.Organizations.BlockUser,
			ConcealMembership:                      client.Organizations.ConcealMembership,
			ConvertMemberToOutsideCollaborator:     client.Organizations.ConvertMemberToOutsideCollaborator,
			CreateCustomRepoRole:                   client.Organizations.CreateCustomRepoRole,
			CreateHook:                             client.Organizations.CreateHook,
			CreateOrUpdateCustomProperties:         client.Organizations.CreateOrUpdateCustomProperties,
			CreateOrUpdateCustomProperty:           client.Organizations.CreateOrUpdateCustomProperty,
			CreateOrUpdateRepoCustomPropertyValues: client.Organizations.CreateOrUpdateRepoCustomPropertyValues,
			CreateOrgInvitation:                    client.Organizations.CreateOrgInvitation,
			CreateOrganizationRuleset:              client.Organizations.CreateOrganizationRuleset,
			CreateProject:                          client.Organizations.CreateProject,
			Delete:                                 client.Organizations.Delete,
			DeleteCustomRepoRole:                   client.Organizations.DeleteCustomRepoRole,
			DeleteHook:                             client.Organizations.DeleteHook,
			DeleteOrganizationRuleset:              client.Organizations.DeleteOrganizationRuleset,
			DeletePackage:                          client.Organizations.DeletePackage,
			Edit:                                   client.Organizations.Edit,
			EditHook:                               client.Organizations.EditHook,
			EditHookConfiguration:                  client.Organizations.EditHookConfiguration,
			EditOrgMembership:                      client.Organizations.EditOrgMembership,
			Get:                                    client.Organizations.Get,
			GetAllCustomProperties:                 client.Organizations.GetAllCustomProperties,
			GetAllOrganizationRulesets:             client.Organizations.GetAllOrganizationRulesets,
			GetAuditLog:                            client.Organizations.GetAuditLog,
			GetByID:                                client.Organizations.GetByID,
			GetCustomProperty:                      client.Organizations.GetCustomProperty,
			GetHook:                                client.Organizations.GetHook,
			GetHookConfiguration:                   client.Organizations.GetHookConfiguration,
			GetHookDelivery:                        client.Organizations.GetHookDelivery,
			GetOrgMembership:                       client.Organizations.GetOrgMembership,
			GetOrganizationRuleset:                 client.Organizations.GetOrganizationRuleset,
			GetPackage:                             client.Organizations.GetPackage,
			IsBlocked:                              client.Organizations.IsBlocked,
			IsMember:                               client.Organizations.IsMember,
			IsPublicMember:                         client.Organizations.IsPublicMember,
			List:                                   client.Organizations.List,
			ListAll:                                client.Organizations.ListAll,
			ListBlockedUsers:                       client.Organizations.ListBlockedUsers,
			ListCredentialAuthorizations:           client.Organizations.ListCredentialAuthorizations,
			ListCustomPropertyValues:               client.Organizations.ListCustomPropertyValues,
			ListCustomRepoRoles:                    client.Organizations.ListCustomRepoRoles,
			ListFailedOrgInvitations:               client.Organizations.ListFailedOrgInvitations,
			ListHookDeliveries:                     client.Organizations.ListHookDeliveries,
			ListHooks:                              client.Organizations.ListHooks,
			ListInstallations:                      client.Organizations.ListInstallations,
			ListMembers:                            client.Organizations.ListMembers,
			ListOrgInvitationTeams:                 client.Organizations.ListOrgInvitationTeams,
			ListOrgMemberships:                     client.Organizations.ListOrgMemberships,
			ListOutsideCollaborators:               client.Organizations.ListOutsideCollaborators,
			ListPackages:                           client.Organizations.ListPackages,
			ListPendingOrgInvitations:              client.Organizations.ListPendingOrgInvitations,
			ListProjects:                           client.Organizations.ListProjects,
			ListSecurityManagerTeams:               client.Organizations.ListSecurityManagerTeams,
			ListTeamsAssignedToOrgRole:             client.Organizations.ListTeamsAssignedToOrgRole,
			ListUsersAssignedToOrgRole:             client.Organizations.ListUsersAssignedToOrgRole,
			PackageDeleteVersion:                   client.Organizations.PackageDeleteVersion,
			PackageGetAllVersions:                  client.Organizations.PackageGetAllVersions,
			PackageGetVersion:                      client.Organizations.PackageGetVersion,
			PackageRestoreVersion:                  client.Organizations.PackageRestoreVersion,
			PingHook:                               client.Organizations.PingHook,
			PublicizeMembership:                    client.Organizations.PublicizeMembership,
			RedeliverHookDelivery:                  client.Organizations.RedeliverHookDelivery,
			RemoveCredentialAuthorization:          client.Organizations.RemoveCredentialAuthorization,
			RemoveCustomProperty:                   client.Organizations.RemoveCustomProperty,
			RemoveMember:                           client.Organizations.RemoveMember,
			RemoveOrgMembership:                    client.Organizations.RemoveOrgMembership,
			RemoveOutsideCollaborator:              client.Organizations.RemoveOutsideCollaborator,
			RemoveSecurityManagerTeam:              client.Organizations.RemoveSecurityManagerTeam,
			RestorePackage:                         client.Organizations.RestorePackage,
			ReviewPersonalAccessTokenRequest:       client.Organizations.ReviewPersonalAccessTokenRequest,
			UnblockUser:                            client.Organizations.UnblockUser,
			UpdateCustomRepoRole:                   client.Organizations.UpdateCustomRepoRole,
			UpdateOrganizationRuleset:              client.Organizations.UpdateOrganizationRuleset,
		},
		Projects: &ProjectsService{
			AddProjectCollaborator:              client.Projects.AddProjectCollaborator,
			CreateProjectCard:                   client.Projects.CreateProjectCard,
			CreateProjectColumn:                 client.Projects.CreateProjectColumn,
			DeleteProject:                       client.Projects.DeleteProject,
			DeleteProjectCard:                   client.Projects.DeleteProjectCard,
			DeleteProjectColumn:                 client.Projects.DeleteProjectColumn,
			GetProject:                          client.Projects.GetProject,
			GetProjectCard:                      client.Projects.GetProjectCard,
			GetProjectColumn:                    client.Projects.GetProjectColumn,
			ListProjectCards:                    client.Projects.ListProjectCards,
			ListProjectCollaborators:            client.Projects.ListProjectCollaborators,
			ListProjectColumns:                  client.Projects.ListProjectColumns,
			MoveProjectCard:                     client.Projects.MoveProjectCard,
			MoveProjectColumn:                   client.Projects.MoveProjectColumn,
			RemoveProjectCollaborator:           client.Projects.RemoveProjectCollaborator,
			ReviewProjectCollaboratorPermission: client.Projects.ReviewProjectCollaboratorPermission,
			UpdateProject:                       client.Projects.UpdateProject,
			UpdateProjectCard:                   client.Projects.UpdateProjectCard,
			UpdateProjectColumn:                 client.Projects.UpdateProjectColumn,
		},
		PullRequests: &PullRequestsService{
			Create:                     client.PullRequests.Create,
			CreateComment:              client.PullRequests.CreateComment,
			CreateCommentInReplyTo:     client.PullRequests.CreateCommentInReplyTo,
			CreateReview:               client.PullRequests.CreateReview,
			DeleteComment:              client.PullRequests.DeleteComment,
			DeletePendingReview:        client.PullRequests.DeletePendingReview,
			DismissReview:              client.PullRequests.DismissReview,
			Edit:                       client.PullRequests.Edit,
			EditComment:                client.PullRequests.EditComment,
			Get:                        client.PullRequests.Get,
			GetComment:                 client.PullRequests.GetComment,
			GetRaw:                     client.PullRequests.GetRaw,
			GetReview:                  client.PullRequests.GetReview,
			IsMerged:                   client.PullRequests.IsMerged,
			List:                       client.PullRequests.List,
			ListComments:               client.PullRequests.ListComments,
			ListCommits:                client.PullRequests.ListCommits,
			ListFiles:                  client.PullRequests.ListFiles,
			ListPullRequestsWithCommit: client.PullRequests.ListPullRequestsWithCommit,
			ListReviewComments:         client.PullRequests.ListReviewComments,
			ListReviewers:              client.PullRequests.ListReviewers,
			ListReviews:                client.PullRequests.ListReviews,
			Merge:                      client.PullRequests.Merge,
			RemoveReviewers:            client.PullRequests.RemoveReviewers,
			RequestReviewers:           client.PullRequests.RequestReviewers,
			SubmitReview:               client.PullRequests.SubmitReview,
			UpdateBranch:               client.PullRequests.UpdateBranch,
			UpdateReview:               client.PullRequests.UpdateReview,
		},
		RateLimit: &RateLimitService{
			Get: client.RateLimit.Get,
		},
		Reactions: &ReactionsService{
			CreateCommentReaction:                               client.Reactions.CreateCommentReaction,
			CreateIssueCommentReaction:                          client.Reactions.CreateIssueCommentReaction,
			CreateIssueReaction:                                 client.Reactions.CreateIssueReaction,
			CreatePullRequestCommentReaction:                    client.Reactions.CreatePullRequestCommentReaction,
			CreateReleaseReaction:                               client.Reactions.CreateReleaseReaction,
			CreateTeamDiscussionCommentReaction:                 client.Reactions.CreateTeamDiscussionCommentReaction,
			CreateTeamDiscussionReaction:                        client.Reactions.CreateTeamDiscussionReaction,
			DeleteCommentReaction:                               client.Reactions.DeleteCommentReaction,
			DeleteCommentReactionByID:                           client.Reactions.DeleteCommentReactionByID,
			DeleteIssueCommentReaction:                          client.Reactions.DeleteIssueCommentReaction,
			DeleteIssueCommentReactionByID:                      client.Reactions.DeleteIssueCommentReactionByID,
			DeleteIssueReaction:                                 client.Reactions.DeleteIssueReaction,
			DeleteIssueReactionByID:                             client.Reactions.DeleteIssueReactionByID,
			DeletePullRequestCommentReaction:                    client.Reactions.DeletePullRequestCommentReaction,
			DeletePullRequestCommentReactionByID:                client.Reactions.DeletePullRequestCommentReactionByID,
			DeleteTeamDiscussionCommentReaction:                 client.Reactions.DeleteTeamDiscussionCommentReaction,
			DeleteTeamDiscussionCommentReactionByOrgIDAndTeamID: client.Reactions.DeleteTeamDiscussionCommentReactionByOrgIDAndTeamID,
			DeleteTeamDiscussionReaction:                        client.Reactions.DeleteTeamDiscussionReaction,
			DeleteTeamDiscussionReactionByOrgIDAndTeamID:        client.Reactions.DeleteTeamDiscussionReactionByOrgIDAndTeamID,
			ListCommentReactions:                                client.Reactions.ListCommentReactions,
			ListIssueCommentReactions:                           client.Reactions.ListIssueCommentReactions,
			ListIssueReactions:                                  client.Reactions.ListIssueReactions,
			ListPullRequestCommentReactions:                     client.Reactions.ListPullRequestCommentReactions,
			ListTeamDiscussionCommentReactions:                  client.Reactions.ListTeamDiscussionCommentReactions,
			ListTeamDiscussionReactions:                         client.Reactions.ListTeamDiscussionReactions,
		},
		Repositories: &RepositoriesService{
			AddAdminEnforcement:                   client.Repositories.AddAdminEnforcement,
			AddAppRestrictions:                    client.Repositories.AddAppRestrictions,
			AddAutolink:                           client.Repositories.AddAutolink,
			AddCollaborator:                       client.Repositories.AddCollaborator,
			AddTeamRestrictions:                   client.Repositories.AddTeamRestrictions,
			AddUserRestrictions:                   client.Repositories.AddUserRestrictions,
			CompareCommits:                        client.Repositories.CompareCommits,
			CompareCommitsRaw:                     client.Repositories.CompareCommitsRaw,
			Create:                                client.Repositories.Create,
			CreateComment:                         client.Repositories.CreateComment,
			CreateCustomDeploymentProtectionRule:  client.Repositories.CreateCustomDeploymentProtectionRule,
			CreateDeployment:                      client.Repositories.CreateDeployment,
			CreateDeploymentBranchPolicy:          client.Repositories.CreateDeploymentBranchPolicy,
			CreateDeploymentStatus:                client.Repositories.CreateDeploymentStatus,
			CreateFile:                            client.Repositories.CreateFile,
			CreateFork:                            client.Repositories.CreateFork,
			CreateFromTemplate:                    client.Repositories.CreateFromTemplate,
			CreateHook:                            client.Repositories.CreateHook,
			CreateKey:                             client.Repositories.CreateKey,
			CreateOrUpdateCustomProperties:        client.Repositories.CreateOrUpdateCustomProperties,
			CreateProject:                         client.Repositories.CreateProject,
			CreateRelease:                         client.Repositories.CreateRelease,
			CreateRuleset:                         client.Repositories.CreateRuleset,
			CreateStatus:                          client.Repositories.CreateStatus,
			CreateTagProtection:                   client.Repositories.CreateTagProtection,
			CreateUpdateEnvironment:               client.Repositories.CreateUpdateEnvironment,
			Delete:                                client.Repositories.Delete,
			DeleteAutolink:                        client.Repositories.DeleteAutolink,
			DeleteComment:                         client.Repositories.DeleteComment,
			DeleteDeployment:                      client.Repositories.DeleteDeployment,
			DeleteDeploymentBranchPolicy:          client.Repositories.DeleteDeploymentBranchPolicy,
			DeleteEnvironment:                     client.Repositories.DeleteEnvironment,
			DeleteFile:                            client.Repositories.DeleteFile,
			DeleteHook:                            client.Repositories.DeleteHook,
			DeleteInvitation:                      client.Repositories.DeleteInvitation,
			DeleteKey:                             client.Repositories.DeleteKey,
			DeletePreReceiveHook:                  client.Repositories.DeletePreReceiveHook,
			DeleteRelease:                         client.Repositories.DeleteRelease,
			DeleteReleaseAsset:                    client.Repositories.DeleteReleaseAsset,
			DeleteRuleset:                         client.Repositories.DeleteRuleset,
			DeleteTagProtection:                   client.Repositories.DeleteTagProtection,
			DisableAutomatedSecurityFixes:         client.Repositories.DisableAutomatedSecurityFixes,
			DisableCustomDeploymentProtectionRule: client.Repositories.DisableCustomDeploymentProtectionRule,
			DisableDismissalRestrictions:          client.Repositories.DisableDismissalRestrictions,
			DisableLFS:                            client.Repositories.DisableLFS,
			DisablePages:                          client.Repositories.DisablePages,
			DisablePrivateReporting:               client.Repositories.DisablePrivateReporting,
			DisableVulnerabilityAlerts:            client.Repositories.DisableVulnerabilityAlerts,
			Dispatch:                              client.Repositories.Dispatch,
			DownloadContents:                      client.Repositories.DownloadContents,
			DownloadContentsWithMeta:              client.Repositories.DownloadContentsWithMeta,
			DownloadReleaseAsset:                  client.Repositories.DownloadReleaseAsset,
			Edit:                                  client.Repositories.Edit,
			EditActionsAccessLevel:                client.Repositories.EditActionsAccessLevel,
			EditActionsAllowed:                    client.Repositories.EditActionsAllowed,
			EditActionsPermissions:                client.Repositories.EditActionsPermissions,
			EditDefaultWorkflowPermissions:        client.Repositories.EditDefaultWorkflowPermissions,
			EditHook:                              client.Repositories.EditHook,
			EditHookConfiguration:                 client.Repositories.EditHookConfiguration,
			EditRelease:                           client.Repositories.EditRelease,
			EditReleaseAsset:                      client.Repositories.EditReleaseAsset,
			EnableAutomatedSecurityFixes:          client.Repositories.EnableAutomatedSecurityFixes,
			EnableLFS:                             client.Repositories.EnableLFS,
			EnablePages:                           client.Repositories.EnablePages,
			EnablePrivateReporting:                client.Repositories.EnablePrivateReporting,
			EnableVulnerabilityAlerts:             client.Repositories.EnableVulnerabilityAlerts,
			GenerateReleaseNotes:                  client.Repositories.GenerateReleaseNotes,
			Get:                                   client.Repositories.Get,
			GetActionsAccessLevel:                 client.Repositories.GetActionsAccessLevel,
			GetActionsAllowed:                     client.Repositories.GetActionsAllowed,
			GetActionsPermissions:                 client.Repositories.GetActionsPermissions,
			GetAdminEnforcement:                   client.Repositories.GetAdminEnforcement,
			GetAllCustomPropertyValues:            client.Repositories.GetAllCustomPropertyValues,
			GetAllDeploymentProtectionRules:       client.Repositories.GetAllDeploymentProtectionRules,
			GetAllRulesets:                        client.Repositories.GetAllRulesets,
			GetArchiveLink:                        client.Repositories.GetArchiveLink,
			GetAutolink:                           client.Repositories.GetAutolink,
			GetAutomatedSecurityFixes:             client.Repositories.GetAutomatedSecurityFixes,
			GetBranch:                             client.Repositories.GetBranch,
			GetBranchProtection:                   client.Repositories.GetBranchProtection,
			GetByID:                               client.Repositories.GetByID,
			GetCodeOfConduct:                      client.Repositories.GetCodeOfConduct,
			GetCodeownersErrors:                   client.Repositories.GetCodeownersErrors,
			GetCombinedStatus:                     client.Repositories.GetCombinedStatus,
			GetComment:                            client.Repositories.GetComment,
			GetCommit:                             client.Repositories.GetCommit,
			GetCommitRaw:                          client.Repositories.GetCommitRaw,
			GetCommitSHA1:                         client.Repositories.GetCommitSHA1,
			GetCommunityHealthMetrics:             client.Repositories.GetCommunityHealthMetrics,
			GetContents:                           client.Repositories.GetContents,
			GetCustomDeploymentProtectionRule:     client.Repositories.GetCustomDeploymentProtectionRule,
			GetDefaultWorkflowPermissions:         client.Repositories.GetDefaultWorkflowPermissions,
			GetDeployment:                         client.Repositories.GetDeployment,
			GetDeploymentBranchPolicy:             client.Repositories.GetDeploymentBranchPolicy,
			GetDeploymentStatus:                   client.Repositories.GetDeploymentStatus,
			GetEnvironment:                        client.Repositories.GetEnvironment,
			GetHook:                               client.Repositories.GetHook,
			GetHookConfiguration:                  client.Repositories.GetHookConfiguration,
			GetHookDelivery:                       client.Repositories.GetHookDelivery,
			GetKey:                                client.Repositories.GetKey,
			GetLatestPagesBuild:                   client.Repositories.GetLatestPagesBuild,
			GetLatestRelease:                      client.Repositories.GetLatestRelease,
			GetPageBuild:                          client.Repositories.GetPageBuild,
			GetPageHealthCheck:                    client.Repositories.GetPageHealthCheck,
			GetPagesInfo:                          client.Repositories.GetPagesInfo,
			GetPermissionLevel:                    client.Repositories.GetPermissionLevel,
			GetPreReceiveHook:                     client.Repositories.GetPreReceiveHook,
			GetPullRequestReviewEnforcement:       client.Repositories.GetPullRequestReviewEnforcement,
			GetReadme:                             client.Repositories.GetReadme,
			GetRelease:                            client.Repositories.GetRelease,
			GetReleaseAsset:                       client.Repositories.GetReleaseAsset,
			GetReleaseByTag:                       client.Repositories.GetReleaseByTag,
			GetRequiredStatusChecks:               client.Repositories.GetRequiredStatusChecks,
			GetRulesForBranch:                     client.Repositories.GetRulesForBranch,
			GetRuleset:                            client.Repositories.GetRuleset,
			GetSignaturesProtectedBranch:          client.Repositories.GetSignaturesProtectedBranch,
			GetVulnerabilityAlerts:                client.Repositories.GetVulnerabilityAlerts,
			IsCollaborator:                        client.Repositories.IsCollaborator,
			IsPrivateReportingEnabled:             client.Repositories.IsPrivateReportingEnabled,
			License:                               client.Repositories.License,
			ListAll:                               client.Repositories.ListAll,
			ListAllTopics:                         client.Repositories.ListAllTopics,
			ListAppRestrictions:                   client.Repositories.ListAppRestrictions,
			ListAutolinks:                         client.Repositories.ListAutolinks,
			ListBranches:                          client.Repositories.ListBranches,
			ListBranchesHeadCommit:                client.Repositories.ListBranchesHeadCommit,
			ListByAuthenticatedUser:               client.Repositories.ListByAuthenticatedUser,
			ListByOrg:                             client.Repositories.ListByOrg,
			ListByUser:                            client.Repositories.ListByUser,
			ListCodeFrequency:                     client.Repositories.ListCodeFrequency,
			ListCollaborators:                     client.Repositories.ListCollaborators,
			ListComments:                          client.Repositories.ListComments,
			ListCommitActivity:                    client.Repositories.ListCommitActivity,
			ListCommitComments:                    client.Repositories.ListCommitComments,
			ListCommits:                           client.Repositories.ListCommits,
			ListContributors:                      client.Repositories.ListContributors,
			ListContributorsStats:                 client.Repositories.ListContributorsStats,
			ListCustomDeploymentRuleIntegrations:  client.Repositories.ListCustomDeploymentRuleIntegrations,
			ListDeploymentBranchPolicies:          client.Repositories.ListDeploymentBranchPolicies,
			ListDeploymentStatuses:                client.Repositories.ListDeploymentStatuses,
			ListDeployments:                       client.Repositories.ListDeployments,
			ListEnvironments:                      client.Repositories.ListEnvironments,
			ListForks:                             client.Repositories.ListForks,
			ListHookDeliveries:                    client.Repositories.ListHookDeliveries,
			ListHooks:                             client.Repositories.ListHooks,
			ListInvitations:                       client.Repositories.ListInvitations,
			ListKeys:                              client.Repositories.ListKeys,
			ListLanguages:                         client.Repositories.ListLanguages,
			ListPagesBuilds:                       client.Repositories.ListPagesBuilds,
			ListParticipation:                     client.Repositories.ListParticipation,
			ListPreReceiveHooks:                   client.Repositories.ListPreReceiveHooks,
			ListProjects:                          client.Repositories.ListProjects,
			ListPunchCard:                         client.Repositories.ListPunchCard,
			ListReleaseAssets:                     client.Repositories.ListReleaseAssets,
			ListReleases:                          client.Repositories.ListReleases,
			ListRequiredStatusChecksContexts:      client.Repositories.ListRequiredStatusChecksContexts,
			ListStatuses:                          client.Repositories.ListStatuses,
			ListTagProtection:                     client.Repositories.ListTagProtection,
			ListTags:                              client.Repositories.ListTags,
			ListTeamRestrictions:                  client.Repositories.ListTeamRestrictions,
			ListTeams:                             client.Repositories.ListTeams,
			ListTrafficClones:                     client.Repositories.ListTrafficClones,
			ListTrafficPaths:                      client.Repositories.ListTrafficPaths,
			ListTrafficReferrers:                  client.Repositories.ListTrafficReferrers,
			ListTrafficViews:                      client.Repositories.ListTrafficViews,
			ListUserRestrictions:                  client.Repositories.ListUserRestrictions,
			Merge:                                 client.Repositories.Merge,
			MergeUpstream:                         client.Repositories.MergeUpstream,
			OptionalSignaturesOnProtectedBranch:   client.Repositories.OptionalSignaturesOnProtectedBranch,
			PingHook:                              client.Repositories.PingHook,
			RedeliverHookDelivery:                 client.Repositories.RedeliverHookDelivery,
			RemoveAdminEnforcement:                client.Repositories.RemoveAdminEnforcement,
			RemoveAppRestrictions:                 client.Repositories.RemoveAppRestrictions,
			RemoveBranchProtection:                client.Repositories.RemoveBranchProtection,
			RemoveCollaborator:                    client.Repositories.RemoveCollaborator,
			RemovePullRequestReviewEnforcement:    client.Repositories.RemovePullRequestReviewEnforcement,
			RemoveRequiredStatusChecks:            client.Repositories.RemoveRequiredStatusChecks,
			RemoveTeamRestrictions:                client.Repositories.RemoveTeamRestrictions,
			RemoveUserRestrictions:                client.Repositories.RemoveUserRestrictions,
			RenameBranch:                          client.Repositories.RenameBranch,
			ReplaceAllTopics:                      client.Repositories.ReplaceAllTopics,
			ReplaceAppRestrictions:                client.Repositories.ReplaceAppRestrictions,
			ReplaceTeamRestrictions:               client.Repositories.ReplaceTeamRestrictions,
			ReplaceUserRestrictions:               client.Repositories.ReplaceUserRestrictions,
			RequestPageBuild:                      client.Repositories.RequestPageBuild,
			RequireSignaturesOnProtectedBranch:    client.Repositories.RequireSignaturesOnProtectedBranch,
			Subscribe:                             client.Repositories.Subscribe,
			TestHook:                              client.Repositories.TestHook,
			Transfer:                              client.Repositories.Transfer,
			Unsubscribe:                           client.Repositories.Unsubscribe,
			UpdateBranchProtection:                client.Repositories.UpdateBranchProtection,
			UpdateComment:                         client.Repositories.UpdateComment,
			UpdateDeploymentBranchPolicy:          client.Repositories.UpdateDeploymentBranchPolicy,
			UpdateFile:                            client.Repositories.UpdateFile,
			UpdateInvitation:                      client.Repositories.UpdateInvitation,
			UpdatePages:                           client.Repositories.UpdatePages,
			UpdatePreReceiveHook:                  client.Repositories.UpdatePreReceiveHook,
			UpdatePullRequestReviewEnforcement:    client.Repositories.UpdatePullRequestReviewEnforcement,
			UpdateRequiredStatusChecks:            client.Repositories.UpdateRequiredStatusChecks,
			UpdateRuleset:                         client.Repositories.UpdateRuleset,
			UploadReleaseAsset:                    client.Repositories.UploadReleaseAsset,
		},
		SCIM: &SCIMService{
			DeleteSCIMUserFromOrg:          client.SCIM.DeleteSCIMUserFromOrg,
			GetSCIMProvisioningInfoForUser: client.SCIM.GetSCIMProvisioningInfoForUser,
			ListSCIMProvisionedIdentities:  client.SCIM.ListSCIMProvisionedIdentities,
			ProvisionAndInviteSCIMUser:     client.SCIM.ProvisionAndInviteSCIMUser,
			UpdateAttributeForSCIMUser:     client.SCIM.UpdateAttributeForSCIMUser,
			UpdateProvisionedOrgMembership: client.SCIM.UpdateProvisionedOrgMembership,
		},
		Search: newSearchServicePassthrough(client),
		SecretScanning: &SecretScanningService{
			GetAlert:                client.SecretScanning.GetAlert,
			ListAlertsForEnterprise: client.SecretScanning.ListAlertsForEnterprise,
			ListAlertsForOrg:        client.SecretScanning.ListAlertsForOrg,
			ListAlertsForRepo:       client.SecretScanning.ListAlertsForRepo,
			ListLocationsForAlert:   client.SecretScanning.ListLocationsForAlert,
			UpdateAlert:             client.SecretScanning.UpdateAlert,
		},
		SecurityAdvisories: &SecurityAdvisoriesService{
			CreateTemporaryPrivateFork:             client.SecurityAdvisories.CreateTemporaryPrivateFork,
			GetGlobalSecurityAdvisories:            client.SecurityAdvisories.GetGlobalSecurityAdvisories,
			ListGlobalSecurityAdvisories:           client.SecurityAdvisories.ListGlobalSecurityAdvisories,
			ListRepositorySecurityAdvisories:       client.SecurityAdvisories.ListRepositorySecurityAdvisories,
			ListRepositorySecurityAdvisoriesForOrg: client.SecurityAdvisories.ListRepositorySecurityAdvisoriesForOrg,
			RequestCVE:                             client.SecurityAdvisories.RequestCVE,
		},
		Teams: &TeamsService{
			AddTeamMembershipByID:                   client.Teams.AddTeamMembershipByID,
			AddTeamMembershipBySlug:                 client.Teams.AddTeamMembershipBySlug,
			AddTeamProjectByID:                      client.Teams.AddTeamProjectByID,
			AddTeamProjectBySlug:                    client.Teams.AddTeamProjectBySlug,
			AddTeamRepoByID:                         client.Teams.AddTeamRepoByID,
			AddTeamRepoBySlug:                       client.Teams.AddTeamRepoBySlug,
			CreateCommentByID:                       client.Teams.CreateCommentByID,
			CreateCommentBySlug:                     client.Teams.CreateCommentBySlug,
			CreateDiscussionByID:                    client.Teams.CreateDiscussionByID,
			CreateDiscussionBySlug:                  client.Teams.CreateDiscussionBySlug,
			CreateOrUpdateIDPGroupConnectionsByID:   client.Teams.CreateOrUpdateIDPGroupConnectionsByID,
			CreateOrUpdateIDPGroupConnectionsBySlug: client.Teams.CreateOrUpdateIDPGroupConnectionsBySlug,
			CreateTeam:                              client.Teams.CreateTeam,
			DeleteCommentByID:                       client.Teams.DeleteCommentByID,
			DeleteCommentBySlug:                     client.Teams.DeleteCommentBySlug,
			DeleteDiscussionByID:                    client.Teams.DeleteDiscussionByID,
			DeleteDiscussionBySlug:                  client.Teams.DeleteDiscussionBySlug,
			DeleteTeamByID:                          client.Teams.DeleteTeamByID,
			DeleteTeamBySlug:                        client.Teams.DeleteTeamBySlug,
			EditCommentByID:                         client.Teams.EditCommentByID,
			EditCommentBySlug:                       client.Teams.EditCommentBySlug,
			EditDiscussionByID:                      client.Teams.EditDiscussionByID,
			EditDiscussionBySlug:                    client.Teams.EditDiscussionBySlug,
			EditTeamByID:                            client.Teams.EditTeamByID,
			EditTeamBySlug:                          client.Teams.EditTeamBySlug,
			GetCommentByID:                          client.Teams.GetCommentByID,
			GetCommentBySlug:                        client.Teams.GetCommentBySlug,
			GetDiscussionByID:                       client.Teams.GetDiscussionByID,
			GetDiscussionBySlug:                     client.Teams.GetDiscussionBySlug,
			GetExternalGroup:                        client.Teams.GetExternalGroup,
			GetTeamByID:                             client.Teams.GetTeamByID,
			GetTeamBySlug:                           client.Teams.GetTeamBySlug,
			GetTeamMembershipByID:                   client.Teams.GetTeamMembershipByID,
			GetTeamMembershipBySlug:                 client.Teams.GetTeamMembershipBySlug,
			IsTeamRepoByID:                          client.Teams.IsTeamRepoByID,
			IsTeamRepoBySlug:                        client.Teams.IsTeamRepoBySlug,
			ListChildTeamsByParentID:                client.Teams.ListChildTeamsByParentID,
			ListChildTeamsByParentSlug:              client.Teams.ListChildTeamsByParentSlug,
			ListCommentsByID:                        client.Teams.ListCommentsByID,
			ListCommentsBySlug:                      client.Teams.ListCommentsBySlug,
			ListDiscussionsByID:                     client.Teams.ListDiscussionsByID,
			ListDiscussionsBySlug:                   client.Teams.ListDiscussionsBySlug,
			ListExternalGroups:                      client.Teams.ListExternalGroups,
			ListExternalGroupsForTeamBySlug:         client.Teams.ListExternalGroupsForTeamBySlug,
			ListIDPGroupsForTeamByID:                client.Teams.ListIDPGroupsForTeamByID,
			ListIDPGroupsForTeamBySlug:              client.Teams.ListIDPGroupsForTeamBySlug,
			ListIDPGroupsInOrganization:             client.Teams.ListIDPGroupsInOrganization,
			ListPendingTeamInvitationsByID:          client.Teams.ListPendingTeamInvitationsByID,
			ListPendingTeamInvitationsBySlug:        client.Teams.ListPendingTeamInvitationsBySlug,
			ListTeamMembersByID:                     client.Teams.ListTeamMembersByID,
			ListTeamMembersBySlug:                   client.Teams.ListTeamMembersBySlug,
			ListTeamProjectsByID:                    client.Teams.ListTeamProjectsByID,
			ListTeamProjectsBySlug:                  client.Teams.ListTeamProjectsBySlug,
			ListTeamReposByID:                       client.Teams.ListTeamReposByID,
			ListTeamReposBySlug:                     client.Teams.ListTeamReposBySlug,
			ListTeams:                               client.Teams.ListTeams,
			ListUserTeams:                           client.Teams.ListUserTeams,
			RemoveConnectedExternalGroup:            client.Teams.RemoveConnectedExternalGroup,
			RemoveTeamMembershipByID:                client.Teams.RemoveTeamMembershipByID,
			RemoveTeamMembershipBySlug:              client.Teams.RemoveTeamMembershipBySlug,
			RemoveTeamProjectByID:                   client.Teams.RemoveTeamProjectByID,
			RemoveTeamProjectBySlug:                 client.Teams.RemoveTeamProjectBySlug,
			RemoveTeamRepoByID:                      client.Teams.RemoveTeamRepoByID,
			RemoveTeamRepoBySlug:                    client.Teams.RemoveTeamRepoBySlug,
			ReviewTeamProjectsByID:                  client.Teams.ReviewTeamProjectsByID,
			ReviewTeamProjectsBySlug:                client.Teams.ReviewTeamProjectsBySlug,
			UpdateConnectedExternalGroup:            client.Teams.UpdateConnectedExternalGroup,
		},
		Users: &UsersService{
			AcceptInvitation:      client.Users.AcceptInvitation,
			AddEmails:             client.Users.AddEmails,
			BlockUser:             client.Users.BlockUser,
			CreateGPGKey:          client.Users.CreateGPGKey,
			CreateKey:             client.Users.CreateKey,
			CreateProject:         client.Users.CreateProject,
			CreateSSHSigningKey:   client.Users.CreateSSHSigningKey,
			DeclineInvitation:     client.Users.DeclineInvitation,
			DeleteEmails:          client.Users.DeleteEmails,
			DeleteGPGKey:          client.Users.DeleteGPGKey,
			DeleteKey:             client.Users.DeleteKey,
			DeletePackage:         client.Users.DeletePackage,
			DeleteSSHSigningKey:   client.Users.DeleteSSHSigningKey,
			DemoteSiteAdmin:       client.Users.DemoteSiteAdmin,
			Edit:                  client.Users.Edit,
			Follow:                client.Users.Follow,
			Get:                   client.Users.Get,
			GetByID:               client.Users.GetByID,
			GetGPGKey:             client.Users.GetGPGKey,
			GetHovercard:          client.Users.GetHovercard,
			GetKey:                client.Users.GetKey,
			GetPackage:            client.Users.GetPackage,
			GetSSHSigningKey:      client.Users.GetSSHSigningKey,
			IsBlocked:             client.Users.IsBlocked,
			IsFollowing:           client.Users.IsFollowing,
			ListAll:               client.Users.ListAll,
			ListBlockedUsers:      client.Users.ListBlockedUsers,
			ListEmails:            client.Users.ListEmails,
			ListFollowers:         client.Users.ListFollowers,
			ListFollowing:         client.Users.ListFollowing,
			ListGPGKeys:           client.Users.ListGPGKeys,
			ListInvitations:       client.Users.ListInvitations,
			ListKeys:              client.Users.ListKeys,
			ListPackages:          client.Users.ListPackages,
			ListProjects:          client.Users.ListProjects,
			ListSSHSigningKeys:    client.Users.ListSSHSigningKeys,
			PackageDeleteVersion:  client.Users.PackageDeleteVersion,
			PackageGetAllVersions: client.Users.PackageGetAllVersions,
			PackageGetVersion:     client.Users.PackageGetVersion,
			PackageRestoreVersion: client.Users.PackageRestoreVersion,
			PromoteSiteAdmin:      client.Users.PromoteSiteAdmin,
			RestorePackage:        client.Users.RestorePackage,
			SetEmailVisibility:    client.Users.SetEmailVisibility,
			Suspend:               client.Users.Suspend,
			UnblockUser:           client.Users.UnblockUser,
			Unfollow:              client.Users.Unfollow,
			Unsuspend:             client.Users.Unsuspend,
		},
	}
}
