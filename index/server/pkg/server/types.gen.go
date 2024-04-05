// Copyright Red Hat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4. **DO NOT EDIT**
package server

import (
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/registry-support/index/generator/schema"
)

// Architectures Optional list of processor architectures that the devfile supports, empty list suggests that the devfile can be used on any architecture
type Architectures = []string

// AttributeNames List of the YAML free-form attribute names
type AttributeNames = []string

// CommandGroups List of command groups defined in devfile
type CommandGroups = []string

// Default Flag for default devfile registry entry version
type Default = bool

// Deprecated Flag for deprecated devfile registry entry
type Deprecated = bool

// Description Description of devfile registry entry
type Description = string

// Devfile Describes the structure of a cloud-native devworkspace and development environment.
type Devfile = v1alpha2.Devfile

// DisplayName User readable name of devfile registry entry
type DisplayName = string

// GitRemoteName Git repository remote name
type GitRemoteName = string

// GitRemoteNames List of git repository remote names
type GitRemoteNames = []GitRemoteName

// GitRemotes List of git repository remote urls
type GitRemotes = []Url

// GitRevision Branch, tag, or commit reference
type GitRevision = string

// GitSubDir Subdirectory of git repository to use as reference
type GitSubDir = string

// Icon Optional devfile icon encoding type
type Icon = string

// IconUri Optional devfile icon uri, can be a URL or a relative path in the project
type IconUri = string

// IndexParams IndexParams defines parameters for index endpoints.
type IndexParams struct {
	// Arch Optional list of processor architectures that the devfile supports, empty list suggests that the devfile can be used on any architecture
	Arch *Architectures `json:"arch,omitempty"`

	// AttributeNames List of the YAML free-form attribute names
	AttributeNames *AttributeNames `json:"attributeNames,omitempty"`

	// CommandGroups List of command groups defined in devfile
	CommandGroups *CommandGroups `json:"commandGroups,omitempty"`

	// Default Flag for default devfile registry entry version
	Default *Default `json:"default,omitempty"`

	// Deprecated Flag for deprecated devfile registry entry
	Deprecated *Deprecated `json:"deprecated,omitempty"`

	// Description Description of devfile registry entry
	Description *Description `json:"description,omitempty"`

	// DisplayName User readable name of devfile registry entry
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// GitRemoteName Git repository remote name
	GitRemoteName *GitRemoteName `json:"gitRemoteName,omitempty"`

	// GitRemoteNames List of git repository remote names
	GitRemoteNames *GitRemoteNames `json:"gitRemoteNames,omitempty"`

	// GitRemotes List of git repository remote urls
	GitRemotes *GitRemotes `json:"gitRemotes,omitempty"`

	// GitRevision Branch, tag, or commit reference
	GitRevision *GitRevision `json:"gitRevision,omitempty"`

	// GitSubDir Subdirectory of git repository to use as reference
	GitSubDir *GitSubDir `json:"gitSubDir,omitempty"`

	// GitUrl Url field type
	GitUrl *Url `json:"gitUrl,omitempty"`

	// Icon Optional devfile icon encoding type
	Icon *Icon `json:"icon,omitempty"`

	// IconUri Optional devfile icon uri, can be a URL or a relative path in the project
	IconUri *IconUri `json:"iconUri,omitempty"`

	// Language Programming language of the devfile workspace
	Language *Language `json:"language,omitempty"`

	// LinkNames Names of devfile links
	LinkNames *LinkNames `json:"linkNames,omitempty"`

	// Links List of devfile links
	Links *Links `json:"links,omitempty"`

	// MaxSchemaVersion Devfile schema version number
	MaxSchemaVersion *SchemaVersion `json:"maxSchemaVersion,omitempty"`

	// MaxVersion Devfile registry entry version number
	MaxVersion *Version `json:"maxVersion,omitempty"`

	// MinSchemaVersion Devfile schema version number
	MinSchemaVersion *SchemaVersion `json:"minSchemaVersion,omitempty"`

	// MinVersion Devfile registry entry version number
	MinVersion *Version `json:"minVersion,omitempty"`

	// Name Name of devfile registry entry
	Name *Name `json:"name,omitempty"`

	// ProjectType Type of project the devfile supports
	ProjectType *ProjectType `json:"projectType,omitempty"`

	// Provider Name of provider of the devfile registry entry
	Provider *Provider `json:"provider,omitempty"`

	// Resources List of file resources for the devfile
	Resources *Resources `json:"resources,omitempty"`

	// StarterProjects List of starter project names
	StarterProjects *StarterProjects `json:"starterProjects,omitempty"`

	// SupportUrl Url field type
	SupportUrl *Url `json:"supportUrl,omitempty"`

	// Tags List of devfile subject tags
	Tags *Tags `json:"tags,omitempty"`
}

// IndexSchema The index file schema
type IndexSchema = schema.Schema

// Language Programming language of the devfile workspace
type Language = string

// LinkNames Names of devfile links
type LinkNames = []string

// Links List of devfile links
type Links = []Url

// Name Name of devfile registry entry
type Name = string

// ProjectType Type of project the devfile supports
type ProjectType = string

// Provider Name of provider of the devfile registry entry
type Provider = string

// Resources List of file resources for the devfile
type Resources = []string

// SchemaVersion Devfile schema version number
type SchemaVersion = string

// StarterProjects List of starter project names
type StarterProjects = []string

// Tags List of devfile subject tags
type Tags = []string

// Url Url field type
type Url = string

// Version Devfile registry entry version number
type Version = string

// ArchParam Optional list of processor architectures that the devfile supports, empty list suggests that the devfile can be used on any architecture
type ArchParam = Architectures

// AttributeNamesParam defines model for attributeNamesParam.
type AttributeNamesParam = []string

// CommandGroupsParam List of command groups defined in devfile
type CommandGroupsParam = CommandGroups

// DefaultParam Flag for default devfile registry entry version
type DefaultParam = Default

// DeprecatedParam Flag for deprecated devfile registry entry
type DeprecatedParam = Deprecated

// DescriptionParam Description of devfile registry entry
type DescriptionParam = Description

// DisplayNameParam User readable name of devfile registry entry
type DisplayNameParam = DisplayName

// GitRemoteNameParam Git repository remote name
type GitRemoteNameParam = GitRemoteName

// GitRemoteNamesParam List of git repository remote names
type GitRemoteNamesParam = GitRemoteNames

// GitRemotesParam List of git repository remote urls
type GitRemotesParam = GitRemotes

// GitRevisionParam Branch, tag, or commit reference
type GitRevisionParam = GitRevision

// GitSubDirParam Subdirectory of git repository to use as reference
type GitSubDirParam = GitSubDir

// GitUrlParam Url field type
type GitUrlParam = Url

// IconParam Optional devfile icon encoding type
type IconParam = Icon

// IconUriParam Optional devfile icon uri, can be a URL or a relative path in the project
type IconUriParam = IconUri

// LanguageParam Programming language of the devfile workspace
type LanguageParam = Language

// LinkNamesParam Names of devfile links
type LinkNamesParam = LinkNames

// LinksParam List of devfile links
type LinksParam = Links

// MaxSchemaVersionParam Devfile schema version number
type MaxSchemaVersionParam = SchemaVersion

// MaxVersionParam Devfile registry entry version number
type MaxVersionParam = Version

// MinSchemaVersionParam Devfile schema version number
type MinSchemaVersionParam = SchemaVersion

// MinVersionParam Devfile registry entry version number
type MinVersionParam = Version

// NameParam Name of devfile registry entry
type NameParam = Name

// ProjectTypeParam Type of project the devfile supports
type ProjectTypeParam = ProjectType

// ProviderParam Name of provider of the devfile registry entry
type ProviderParam = Provider

// ResourcesParam List of file resources for the devfile
type ResourcesParam = Resources

// StarterProjectsParam List of starter project names
type StarterProjectsParam = StarterProjects

// SupportUrlParam Url field type
type SupportUrlParam = Url

// TagsParam defines model for tagsParam.
type TagsParam = []string

// DevfileErrorResponse defines model for devfileErrorResponse.
type DevfileErrorResponse struct {
	Error  *string `json:"error,omitempty"`
	Status *string `json:"status,omitempty"`
}

// DevfileNotFoundResponse defines model for devfileNotFoundResponse.
type DevfileNotFoundResponse struct {
	Status *string `json:"status,omitempty"`
}

// DevfileResponse Describes the structure of a cloud-native devworkspace and development environment.
type DevfileResponse = Devfile

// HealthResponse defines model for healthResponse.
type HealthResponse struct {
	Message string `json:"message"`
}

// IndexResponse The index file schema
type IndexResponse = IndexSchema

// MethodNotAllowedResponse defines model for methodNotAllowedResponse.
type MethodNotAllowedResponse struct {
	Message string `json:"message"`
}

// V2IndexResponse defines model for v2IndexResponse.
type V2IndexResponse = schema.Schema

// ServeDevfileParams defines parameters for ServeDevfile.
type ServeDevfileParams struct {
	// MinSchemaVersion The minimum devfile schema version
	MinSchemaVersion *MinSchemaVersionParam `form:"minSchemaVersion,omitempty" json:"minSchemaVersion,omitempty"`

	// MaxSchemaVersion The maximum devfile schema version
	MaxSchemaVersion *MaxSchemaVersionParam `form:"maxSchemaVersion,omitempty" json:"maxSchemaVersion,omitempty"`
}

// ServeDevfileStarterProjectParams defines parameters for ServeDevfileStarterProject.
type ServeDevfileStarterProjectParams struct {
	// MinSchemaVersion The minimum devfile schema version
	MinSchemaVersion *MinSchemaVersionParam `form:"minSchemaVersion,omitempty" json:"minSchemaVersion,omitempty"`

	// MaxSchemaVersion The maximum devfile schema version
	MaxSchemaVersion *MaxSchemaVersionParam `form:"maxSchemaVersion,omitempty" json:"maxSchemaVersion,omitempty"`
}

// ServeDevfileWithVersionParams defines parameters for ServeDevfileWithVersion.
type ServeDevfileWithVersionParams struct {
	// MinSchemaVersion The minimum devfile schema version
	MinSchemaVersion *MinSchemaVersionParam `form:"minSchemaVersion,omitempty" json:"minSchemaVersion,omitempty"`

	// MaxSchemaVersion The maximum devfile schema version
	MaxSchemaVersion *MaxSchemaVersionParam `form:"maxSchemaVersion,omitempty" json:"maxSchemaVersion,omitempty"`
}

// ServeDevfileStarterProjectWithVersionParams defines parameters for ServeDevfileStarterProjectWithVersion.
type ServeDevfileStarterProjectWithVersionParams struct {
	// MinSchemaVersion The minimum devfile schema version
	MinSchemaVersion *MinSchemaVersionParam `form:"minSchemaVersion,omitempty" json:"minSchemaVersion,omitempty"`

	// MaxSchemaVersion The maximum devfile schema version
	MaxSchemaVersion *MaxSchemaVersionParam `form:"maxSchemaVersion,omitempty" json:"maxSchemaVersion,omitempty"`
}

// ServeDevfileIndexV1Params defines parameters for ServeDevfileIndexV1.
type ServeDevfileIndexV1Params struct {
	// Name Search string to filter stacks by their name
	Name *NameParam `form:"name,omitempty" json:"name,omitempty"`

	// DisplayName Search string to filter stacks by their display names
	DisplayName *DisplayNameParam `form:"displayName,omitempty" json:"displayName,omitempty"`

	// Description Search string to filter stacks by the description text
	Description *DescriptionParam `form:"description,omitempty" json:"description,omitempty"`

	// AttributeNames Collection of search strings to filter stacks by the names of
	// defined free-form attributes
	AttributeNames *AttributeNamesParam `form:"attributeNames,omitempty" json:"attributeNames,omitempty"`

	// Tags Collection of search strings to filter stacks by their tags
	Tags *TagsParam `form:"tags,omitempty" json:"tags,omitempty"`

	// Arch Collection of search strings to filter stacks by their architectures
	Arch *ArchParam `form:"arch,omitempty" json:"arch,omitempty"`

	// Icon Toggle on encoding content passed
	Icon *IconParam `form:"icon,omitempty" json:"icon,omitempty"`

	// IconUri Search string to filter stacks by their icon uri
	IconUri *IconUriParam `form:"iconUri,omitempty" json:"iconUri,omitempty"`

	// ProjectType Search string to filter stacks by their project type
	ProjectType *ProjectTypeParam `form:"projectType,omitempty" json:"projectType,omitempty"`

	// Language Search string to filter stacks by their programming language
	Language *LanguageParam `form:"language,omitempty" json:"language,omitempty"`

	// Deprecated Boolean to filter stacks if they are deprecated or not
	Deprecated *DeprecatedParam `form:"deprecated,omitempty" json:"deprecated,omitempty"`

	// Resources Collection of search strings to filter stacks by their
	// resource files
	Resources *ResourcesParam `form:"resources,omitempty" json:"resources,omitempty"`

	// StarterProjects Collection of search strings to filter stacks by the names
	// of the starter projects
	StarterProjects *StarterProjectsParam `form:"starterProjects,omitempty" json:"starterProjects,omitempty"`

	// LinkNames Collection of search strings to filter stacks by the names
	// of the link sources
	LinkNames *LinkNamesParam `form:"linkNames,omitempty" json:"linkNames,omitempty"`

	// Links Collection of search strings to filter stacks by their link
	// sources
	Links *LinksParam `form:"links,omitempty" json:"links,omitempty"`

	// GitRemoteNames Collection of search strings to filter stacks by the names of
	// the git remotes
	GitRemoteNames *GitRemoteNamesParam `form:"gitRemoteNames,omitempty" json:"gitRemoteNames,omitempty"`

	// GitRemotes Collection of search strings to filter stacks by the URIs of
	// the git remotes
	GitRemotes *GitRemotesParam `form:"gitRemotes,omitempty" json:"gitRemotes,omitempty"`

	// GitUrl Search string to filter stacks by their git urls
	GitUrl *GitUrlParam `form:"gitUrl,omitempty" json:"gitUrl,omitempty"`

	// GitRemoteName Search string to filter stacks by their git remote name
	GitRemoteName *GitRemoteNameParam `form:"gitRemoteName,omitempty" json:"gitRemoteName,omitempty"`

	// GitSubDir Search string to filter stacks by their target subdirectory
	// of the git repository
	GitSubDir *GitSubDirParam `form:"gitSubDir,omitempty" json:"gitSubDir,omitempty"`

	// GitRevision Search string to filter stacks by their git revision
	GitRevision *GitRevisionParam `form:"gitRevision,omitempty" json:"gitRevision,omitempty"`

	// Provider Search string to filter stacks by the stack provider
	Provider *ProviderParam `form:"provider,omitempty" json:"provider,omitempty"`

	// SupportUrl Search string to filter stacks by their given support url
	SupportUrl *SupportUrlParam `form:"supportUrl,omitempty" json:"supportUrl,omitempty"`
}

// ServeDevfileIndexV1WithTypeParams defines parameters for ServeDevfileIndexV1WithType.
type ServeDevfileIndexV1WithTypeParams struct {
	// Name Search string to filter stacks by their name
	Name *NameParam `form:"name,omitempty" json:"name,omitempty"`

	// DisplayName Search string to filter stacks by their display names
	DisplayName *DisplayNameParam `form:"displayName,omitempty" json:"displayName,omitempty"`

	// Description Search string to filter stacks by the description text
	Description *DescriptionParam `form:"description,omitempty" json:"description,omitempty"`

	// AttributeNames Collection of search strings to filter stacks by the names of
	// defined free-form attributes
	AttributeNames *AttributeNamesParam `form:"attributeNames,omitempty" json:"attributeNames,omitempty"`

	// Tags Collection of search strings to filter stacks by their tags
	Tags *TagsParam `form:"tags,omitempty" json:"tags,omitempty"`

	// Arch Collection of search strings to filter stacks by their architectures
	Arch *ArchParam `form:"arch,omitempty" json:"arch,omitempty"`

	// Icon Toggle on encoding content passed
	Icon *IconParam `form:"icon,omitempty" json:"icon,omitempty"`

	// IconUri Search string to filter stacks by their icon uri
	IconUri *IconUriParam `form:"iconUri,omitempty" json:"iconUri,omitempty"`

	// ProjectType Search string to filter stacks by their project type
	ProjectType *ProjectTypeParam `form:"projectType,omitempty" json:"projectType,omitempty"`

	// Language Search string to filter stacks by their programming language
	Language *LanguageParam `form:"language,omitempty" json:"language,omitempty"`

	// Deprecated Boolean to filter stacks if they are deprecated or not
	Deprecated *DeprecatedParam `form:"deprecated,omitempty" json:"deprecated,omitempty"`

	// Resources Collection of search strings to filter stacks by their
	// resource files
	Resources *ResourcesParam `form:"resources,omitempty" json:"resources,omitempty"`

	// StarterProjects Collection of search strings to filter stacks by the names
	// of the starter projects
	StarterProjects *StarterProjectsParam `form:"starterProjects,omitempty" json:"starterProjects,omitempty"`

	// LinkNames Collection of search strings to filter stacks by the names
	// of the link sources
	LinkNames *LinkNamesParam `form:"linkNames,omitempty" json:"linkNames,omitempty"`

	// Links Collection of search strings to filter stacks by their link
	// sources
	Links *LinksParam `form:"links,omitempty" json:"links,omitempty"`

	// GitRemoteNames Collection of search strings to filter stacks by the names of
	// the git remotes
	GitRemoteNames *GitRemoteNamesParam `form:"gitRemoteNames,omitempty" json:"gitRemoteNames,omitempty"`

	// GitRemotes Collection of search strings to filter stacks by the URIs of
	// the git remotes
	GitRemotes *GitRemotesParam `form:"gitRemotes,omitempty" json:"gitRemotes,omitempty"`

	// GitUrl Search string to filter stacks by their git urls
	GitUrl *GitUrlParam `form:"gitUrl,omitempty" json:"gitUrl,omitempty"`

	// GitRemoteName Search string to filter stacks by their git remote name
	GitRemoteName *GitRemoteNameParam `form:"gitRemoteName,omitempty" json:"gitRemoteName,omitempty"`

	// GitSubDir Search string to filter stacks by their target subdirectory
	// of the git repository
	GitSubDir *GitSubDirParam `form:"gitSubDir,omitempty" json:"gitSubDir,omitempty"`

	// GitRevision Search string to filter stacks by their git revision
	GitRevision *GitRevisionParam `form:"gitRevision,omitempty" json:"gitRevision,omitempty"`

	// Provider Search string to filter stacks by the stack provider
	Provider *ProviderParam `form:"provider,omitempty" json:"provider,omitempty"`

	// SupportUrl Search string to filter stacks by their given support url
	SupportUrl *SupportUrlParam `form:"supportUrl,omitempty" json:"supportUrl,omitempty"`
}

// ServeDevfileIndexV2Params defines parameters for ServeDevfileIndexV2.
type ServeDevfileIndexV2Params struct {
	// Name Search string to filter stacks by their name
	Name *NameParam `form:"name,omitempty" json:"name,omitempty"`

	// DisplayName Search string to filter stacks by their display names
	DisplayName *DisplayNameParam `form:"displayName,omitempty" json:"displayName,omitempty"`

	// Description Search string to filter stacks by the description text
	Description *DescriptionParam `form:"description,omitempty" json:"description,omitempty"`

	// AttributeNames Collection of search strings to filter stacks by the names of
	// defined free-form attributes
	AttributeNames *AttributeNamesParam `form:"attributeNames,omitempty" json:"attributeNames,omitempty"`

	// Tags Collection of search strings to filter stacks by their tags
	Tags *TagsParam `form:"tags,omitempty" json:"tags,omitempty"`

	// Arch Collection of search strings to filter stacks by their architectures
	Arch *ArchParam `form:"arch,omitempty" json:"arch,omitempty"`

	// Icon Toggle on encoding content passed
	Icon *IconParam `form:"icon,omitempty" json:"icon,omitempty"`

	// IconUri Search string to filter stacks by their icon uri
	IconUri *IconUriParam `form:"iconUri,omitempty" json:"iconUri,omitempty"`

	// ProjectType Search string to filter stacks by their project type
	ProjectType *ProjectTypeParam `form:"projectType,omitempty" json:"projectType,omitempty"`

	// Language Search string to filter stacks by their programming language
	Language *LanguageParam `form:"language,omitempty" json:"language,omitempty"`

	// MinVersion The minimum stack version
	MinVersion *MinVersionParam `form:"minVersion,omitempty" json:"minVersion,omitempty"`

	// MaxVersion The maximum stack version
	MaxVersion *MaxVersionParam `form:"maxVersion,omitempty" json:"maxVersion,omitempty"`

	// MinSchemaVersion The minimum devfile schema version
	MinSchemaVersion *MinSchemaVersionParam `form:"minSchemaVersion,omitempty" json:"minSchemaVersion,omitempty"`

	// MaxSchemaVersion The maximum devfile schema version
	MaxSchemaVersion *MaxSchemaVersionParam `form:"maxSchemaVersion,omitempty" json:"maxSchemaVersion,omitempty"`

	// Deprecated Boolean to filter stacks if they are deprecated or not
	Deprecated *DeprecatedParam `form:"deprecated,omitempty" json:"deprecated,omitempty"`

	// Default Boolean to filter stacks if they are default or not
	Default *DefaultParam `form:"default,omitempty" json:"default,omitempty"`

	// Resources Collection of search strings to filter stacks by their
	// resource files
	Resources *ResourcesParam `form:"resources,omitempty" json:"resources,omitempty"`

	// StarterProjects Collection of search strings to filter stacks by the names
	// of the starter projects
	StarterProjects *StarterProjectsParam `form:"starterProjects,omitempty" json:"starterProjects,omitempty"`

	// LinkNames Collection of search strings to filter stacks by the names
	// of the link sources
	LinkNames *LinkNamesParam `form:"linkNames,omitempty" json:"linkNames,omitempty"`

	// Links Collection of search strings to filter stacks by their link
	// sources
	Links *LinksParam `form:"links,omitempty" json:"links,omitempty"`

	// CommandGroups Collection of search strings to filter stacks by their present command
	// groups
	CommandGroups *CommandGroupsParam `form:"commandGroups,omitempty" json:"commandGroups,omitempty"`

	// GitRemoteNames Collection of search strings to filter stacks by the names of
	// the git remotes
	GitRemoteNames *GitRemoteNamesParam `form:"gitRemoteNames,omitempty" json:"gitRemoteNames,omitempty"`

	// GitRemotes Collection of search strings to filter stacks by the URIs of
	// the git remotes
	GitRemotes *GitRemotesParam `form:"gitRemotes,omitempty" json:"gitRemotes,omitempty"`

	// GitUrl Search string to filter stacks by their git urls
	GitUrl *GitUrlParam `form:"gitUrl,omitempty" json:"gitUrl,omitempty"`

	// GitRemoteName Search string to filter stacks by their git remote name
	GitRemoteName *GitRemoteNameParam `form:"gitRemoteName,omitempty" json:"gitRemoteName,omitempty"`

	// GitSubDir Search string to filter stacks by their target subdirectory
	// of the git repository
	GitSubDir *GitSubDirParam `form:"gitSubDir,omitempty" json:"gitSubDir,omitempty"`

	// GitRevision Search string to filter stacks by their git revision
	GitRevision *GitRevisionParam `form:"gitRevision,omitempty" json:"gitRevision,omitempty"`

	// Provider Search string to filter stacks by the stack provider
	Provider *ProviderParam `form:"provider,omitempty" json:"provider,omitempty"`

	// SupportUrl Search string to filter stacks by their given support url
	SupportUrl *SupportUrlParam `form:"supportUrl,omitempty" json:"supportUrl,omitempty"`
}

// ServeDevfileIndexV2WithTypeParams defines parameters for ServeDevfileIndexV2WithType.
type ServeDevfileIndexV2WithTypeParams struct {
	// Name Search string to filter stacks by their name
	Name *NameParam `form:"name,omitempty" json:"name,omitempty"`

	// DisplayName Search string to filter stacks by their display names
	DisplayName *DisplayNameParam `form:"displayName,omitempty" json:"displayName,omitempty"`

	// Description Search string to filter stacks by the description text
	Description *DescriptionParam `form:"description,omitempty" json:"description,omitempty"`

	// AttributeNames Collection of search strings to filter stacks by the names of
	// defined free-form attributes
	AttributeNames *AttributeNamesParam `form:"attributeNames,omitempty" json:"attributeNames,omitempty"`

	// Tags Collection of search strings to filter stacks by their tags
	Tags *TagsParam `form:"tags,omitempty" json:"tags,omitempty"`

	// Arch Collection of search strings to filter stacks by their architectures
	Arch *ArchParam `form:"arch,omitempty" json:"arch,omitempty"`

	// Icon Toggle on encoding content passed
	Icon *IconParam `form:"icon,omitempty" json:"icon,omitempty"`

	// IconUri Search string to filter stacks by their icon uri
	IconUri *IconUriParam `form:"iconUri,omitempty" json:"iconUri,omitempty"`

	// ProjectType Search string to filter stacks by their project type
	ProjectType *ProjectTypeParam `form:"projectType,omitempty" json:"projectType,omitempty"`

	// Language Search string to filter stacks by their programming language
	Language *LanguageParam `form:"language,omitempty" json:"language,omitempty"`

	// MinVersion The minimum stack version
	MinVersion *MinVersionParam `form:"minVersion,omitempty" json:"minVersion,omitempty"`

	// MaxVersion The maximum stack version
	MaxVersion *MaxVersionParam `form:"maxVersion,omitempty" json:"maxVersion,omitempty"`

	// MinSchemaVersion The minimum devfile schema version
	MinSchemaVersion *MinSchemaVersionParam `form:"minSchemaVersion,omitempty" json:"minSchemaVersion,omitempty"`

	// MaxSchemaVersion The maximum devfile schema version
	MaxSchemaVersion *MaxSchemaVersionParam `form:"maxSchemaVersion,omitempty" json:"maxSchemaVersion,omitempty"`

	// Deprecated Boolean to filter stacks if they are deprecated or not
	Deprecated *DeprecatedParam `form:"deprecated,omitempty" json:"deprecated,omitempty"`

	// Default Boolean to filter stacks if they are default or not
	Default *DefaultParam `form:"default,omitempty" json:"default,omitempty"`

	// Resources Collection of search strings to filter stacks by their
	// resource files
	Resources *ResourcesParam `form:"resources,omitempty" json:"resources,omitempty"`

	// StarterProjects Collection of search strings to filter stacks by the names
	// of the starter projects
	StarterProjects *StarterProjectsParam `form:"starterProjects,omitempty" json:"starterProjects,omitempty"`

	// LinkNames Collection of search strings to filter stacks by the names
	// of the link sources
	LinkNames *LinkNamesParam `form:"linkNames,omitempty" json:"linkNames,omitempty"`

	// Links Collection of search strings to filter stacks by their link
	// sources
	Links *LinksParam `form:"links,omitempty" json:"links,omitempty"`

	// CommandGroups Collection of search strings to filter stacks by their present command
	// groups
	CommandGroups *CommandGroupsParam `form:"commandGroups,omitempty" json:"commandGroups,omitempty"`

	// GitRemoteNames Collection of search strings to filter stacks by the names of
	// the git remotes
	GitRemoteNames *GitRemoteNamesParam `form:"gitRemoteNames,omitempty" json:"gitRemoteNames,omitempty"`

	// GitRemotes Collection of search strings to filter stacks by the URIs of
	// the git remotes
	GitRemotes *GitRemotesParam `form:"gitRemotes,omitempty" json:"gitRemotes,omitempty"`

	// GitUrl Search string to filter stacks by their git urls
	GitUrl *GitUrlParam `form:"gitUrl,omitempty" json:"gitUrl,omitempty"`

	// GitRemoteName Search string to filter stacks by their git remote name
	GitRemoteName *GitRemoteNameParam `form:"gitRemoteName,omitempty" json:"gitRemoteName,omitempty"`

	// GitSubDir Search string to filter stacks by their target subdirectory
	// of the git repository
	GitSubDir *GitSubDirParam `form:"gitSubDir,omitempty" json:"gitSubDir,omitempty"`

	// GitRevision Search string to filter stacks by their git revision
	GitRevision *GitRevisionParam `form:"gitRevision,omitempty" json:"gitRevision,omitempty"`

	// Provider Search string to filter stacks by the stack provider
	Provider *ProviderParam `form:"provider,omitempty" json:"provider,omitempty"`

	// SupportUrl Search string to filter stacks by their given support url
	SupportUrl *SupportUrlParam `form:"supportUrl,omitempty" json:"supportUrl,omitempty"`
}
