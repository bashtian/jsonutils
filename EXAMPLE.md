Examples
=========
#https://api.github.com

	{
		current_user_url: "https://api.github.com/user",
		authorizations_url: "https://api.github.com/authorizations",
		emails_url: "https://api.github.com/user/emails",
		emojis_url: "https://api.github.com/emojis",
		events_url: "https://api.github.com/events",
		following_url: "https://api.github.com/user/following{/target}",
		gists_url: "https://api.github.com/gists{/gist_id}",
		hub_url: "https://api.github.com/hub",
		issue_search_url: "https://api.github.com/legacy/issues/search/{owner}/{repo}/{state}/{keyword}",
		issues_url: "https://api.github.com/issues",
		keys_url: "https://api.github.com/user/keys",
		notifications_url: "https://api.github.com/notifications",
		organization_repositories_url: "https://api.github.com/orgs/{org}/repos/{?type,page,per_page,sort}",
		organization_url: "https://api.github.com/orgs/{org}",
		public_gists_url: "https://api.github.com/gists/public",
		rate_limit_url: "https://api.github.com/rate_limit",
		repository_url: "https://api.github.com/repos/{owner}/{repo}",
		repository_search_url: "https://api.github.com/legacy/repos/search/{keyword}{?language,start_page}",
		current_user_repositories_url: "https://api.github.com/user/repos{?type,page,per_page,sort}",
		starred_url: "https://api.github.com/user/starred{/owner}{/repo}",
		starred_gists_url: "https://api.github.com/gists/starred",
		team_url: "https://api.github.com/teams",
		user_url: "https://api.github.com/users/{user}",
		user_organizations_url: "https://api.github.com/user/orgs",
		user_repositories_url: "https://api.github.com/users/{user}/repos{?type,page,per_page,sort}",
		user_search_url: "https://api.github.com/legacy/user/search/{keyword}"
	}

	jsonutils https://api.github.com

	type Data struct {
		AuthorizationsUrl           string `json:"authorizations_url"`
		CurrentUserRepositoriesUrl  string `json:"current_user_repositories_url"`
		CurrentUserUrl              string `json:"current_user_url"`
		EmailsUrl                   string `json:"emails_url"`
		EmojisUrl                   string `json:"emojis_url"`
		EventsUrl                   string `json:"events_url"`
		FollowingUrl                string `json:"following_url"`
		GistsUrl                    string `json:"gists_url"`
		HubUrl                      string `json:"hub_url"`
		IssueSearchUrl              string `json:"issue_search_url"`
		IssuesUrl                   string `json:"issues_url"`
		KeysUrl                     string `json:"keys_url"`
		NotificationsUrl            string `json:"notifications_url"`
		OrganizationRepositoriesUrl string `json:"organization_repositories_url"`
		OrganizationUrl             string `json:"organization_url"`
		PublicGistsUrl              string `json:"public_gists_url"`
		RateLimitUrl                string `json:"rate_limit_url"`
		RepositorySearchUrl         string `json:"repository_search_url"`
		RepositoryUrl               string `json:"repository_url"`
		StarredGistsUrl             string `json:"starred_gists_url"`
		StarredUrl                  string `json:"starred_url"`
		TeamUrl                     string `json:"team_url"`
		UserOrganizationsUrl        string `json:"user_organizations_url"`
		UserRepositoriesUrl         string `json:"user_repositories_url"`
		UserSearchUrl               string `json:"user_search_url"`
		UserUrl                     string `json:"user_url"`
	}

	jsonutils -x https://api.github.com

	type Data struct {
		AuthorizationsUrl           string `json:"authorizations_url"`            // https://api.github.com/authorizations
		CurrentUserRepositoriesUrl  string `json:"current_user_repositories_url"` // https://api.github.com/user/repos{?type,page,per_page,sort}
		CurrentUserUrl              string `json:"current_user_url"`              // https://api.github.com/user
		EmailsUrl                   string `json:"emails_url"`                    // https://api.github.com/user/emails
		EmojisUrl                   string `json:"emojis_url"`                    // https://api.github.com/emojis
		EventsUrl                   string `json:"events_url"`                    // https://api.github.com/events
		FollowingUrl                string `json:"following_url"`                 // https://api.github.com/user/following{/target}
		GistsUrl                    string `json:"gists_url"`                     // https://api.github.com/gists{/gist_id}
		HubUrl                      string `json:"hub_url"`                       // https://api.github.com/hub
		IssueSearchUrl              string `json:"issue_search_url"`              // https://api.github.com/legacy/issues/search/{owner}/{repo}/{state}/{keyword}
		IssuesUrl                   string `json:"issues_url"`                    // https://api.github.com/issues
		KeysUrl                     string `json:"keys_url"`                      // https://api.github.com/user/keys
		NotificationsUrl            string `json:"notifications_url"`             // https://api.github.com/notifications
		OrganizationRepositoriesUrl string `json:"organization_repositories_url"` // https://api.github.com/orgs/{org}/repos/{?type,page,per_page,sort}
		OrganizationUrl             string `json:"organization_url"`              // https://api.github.com/orgs/{org}
		PublicGistsUrl              string `json:"public_gists_url"`              // https://api.github.com/gists/public
		RateLimitUrl                string `json:"rate_limit_url"`                // https://api.github.com/rate_limit
		RepositorySearchUrl         string `json:"repository_search_url"`         // https://api.github.com/legacy/repos/search/{keyword}{?language,start_page}
		RepositoryUrl               string `json:"repository_url"`                // https://api.github.com/repos/{owner}/{repo}
		StarredGistsUrl             string `json:"starred_gists_url"`             // https://api.github.com/gists/starred
		StarredUrl                  string `json:"starred_url"`                   // https://api.github.com/user/starred{/owner}{/repo}
		TeamUrl                     string `json:"team_url"`                      // https://api.github.com/teams
		UserOrganizationsUrl        string `json:"user_organizations_url"`        // https://api.github.com/user/orgs
		UserRepositoriesUrl         string `json:"user_repositories_url"`         // https://api.github.com/users/{user}/repos{?type,page,per_page,sort}
		UserSearchUrl               string `json:"user_search_url"`               // https://api.github.com/legacy/user/search/{keyword}
		UserUrl                     string `json:"user_url"`                      // https://api.github.com/users/{user}
	}