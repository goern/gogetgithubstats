/* Copyright © 2015 Christoph Görn

This file is part of gogetgithubstats.

gogetgithubstats is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

gogetgithubstats is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with gogetgithubstats. If not, see <http://www.gnu.org/licenses/>.
*/

package githubstats

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

//GetStatsByRepo will get a few usefull statistics on a given user's repository
func GetStatsByRepo(user, repo, accessToken string) {
	var issues []github.Issue

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.Get(user, repo)
	client.Repositories.GetCombinedStatus(user, repo, "master", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	if *repos.HasIssues {
		opt := &github.IssueListByRepoOptions{
			State: "closed", ListOptions: github.ListOptions{PerPage: 10},
		}

		for {
			is, resp, err := client.Issues.ListByRepo(user, repo, opt)
			if err != nil {
				fmt.Println(err)
				return
			}

			issues = append(issues, is...)
			if resp.NextPage == 0 {
				break
			}
			opt.ListOptions.Page = resp.NextPage
		}
	}

	fmt.Printf("repository: %s\n\nstars: %d, open issues: %d\n",
		*repos.FullName, *repos.StargazersCount, len(issues))

	for _, issue := range issues {
		duration := (*issue.ClosedAt).Sub(*issue.CreatedAt)

		fmt.Printf("%03d, %v: created: %v, closed: %v, TTC: %v\n",
			*issue.Number, *issue.State,
			*issue.CreatedAt, *issue.ClosedAt, humanizeDuration(duration))
	}

	rate, _, err := client.RateLimits()
	if err != nil {
		fmt.Printf("Error fetching rate limit: %#v\n\n", err)
	} else {
		fmt.Printf("API Rate Limit: %05d/%05d, Search: %05d\n",
			rate.Core.Remaining, rate.Core.Limit, rate.Search.Limit)
	}

}

// humanizeDuration humanizes time.Duration output to a meaningful value,
// golang's default ``time.Duration`` output is badly formatted and unreadable.
func humanizeDuration(duration time.Duration) string {
	if duration.Seconds() < 60.0 {
		return fmt.Sprintf("%d seconds", int64(duration.Seconds()))
	}
	if duration.Minutes() < 60.0 {
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("%d minutes %d seconds", int64(duration.Minutes()), int64(remainingSeconds))
	}
	if duration.Hours() < 24.0 {
		remainingMinutes := math.Mod(duration.Minutes(), 60)
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("%d hours %d minutes %d seconds",
			int64(duration.Hours()), int64(remainingMinutes), int64(remainingSeconds))
	}
	remainingHours := math.Mod(duration.Hours(), 24)
	remainingMinutes := math.Mod(duration.Minutes(), 60)
	remainingSeconds := math.Mod(duration.Seconds(), 60)
	return fmt.Sprintf("%d days %d hours %d minutes %d seconds",
		int64(duration.Hours()/24), int64(remainingHours),
		int64(remainingMinutes), int64(remainingSeconds))
}

/*

{
   ID:32819266,
   Owner:github.User   {
      Login:"projectatomic",
      ID:6852258,
      AvatarURL:"https://avatars.githubusercontent.com/u/6852258?v=3",
      HTMLURL:"https://github.com/projectatomic",
      GravatarID:"",
      Type:"Organization",
      SiteAdmin:false,
      URL:"https://api.github.com/users/projectatomic",
      EventsURL:"https://api.github.com/users/projectatomic/events{/privacy}",
      FollowingURL:"https://api.github.com/users/projectatomic/following{/other_user}",
      FollowersURL:"https://api.github.com/users/projectatomic/followers",
      GistsURL:"https://api.github.com/users/projectatomic/gists{/gist_id}",
      OrganizationsURL:"https://api.github.com/users/projectatomic/orgs",
      ReceivedEventsURL:"https://api.github.com/users/projectatomic/received_events",
      ReposURL:"https://api.github.com/users/projectatomic/repos",
      StarredURL:"https://api.github.com/users/projectatomic/starred{/owner}{/repo}",
      SubscriptionsURL:"https://api.github.com/users/projectatomic/subscriptions"
   },
   Name:"atomicapp",
   FullName:"projectatomic/atomicapp",
   Description:"This is the reference implementation of the Nulecule container application Specification: Atomic App",
   Homepage:"",
   DefaultBranch:"master",
   MasterBranch:"master",
   CreatedAt:github.Timestamp   {
      2015-03-24 19:09:39+0000 UTC
   },
   PushedAt:github.Timestamp   {
      2015-11-20 15:53:39+0000 UTC
   },
   UpdatedAt:github.Timestamp   {
      2015-11-19 15:11:33+0000 UTC
   },
   HTMLURL:"https://github.com/projectatomic/atomicapp",
   CloneURL:"https://github.com/projectatomic/atomicapp.git",
   GitURL:"git://github.com/projectatomic/atomicapp.git",
   SSHURL:"git@github.com:projectatomic/atomicapp.git",
   SVNURL:"https://github.com/projectatomic/atomicapp",
   Language:"Python",
   Fork:false,
   ForksCount:41,
   NetworkCount:41,
   OpenIssuesCount:74,
   StargazersCount:37,
   SubscribersCount:20,
   WatchersCount:37,
   Size:1310,
   Organization:github.Organization   {
      Login:"projectatomic",
      ID:6852258,
      AvatarURL:"https://avatars.githubusercontent.com/u/6852258?v=3",
      HTMLURL:"https://github.com/projectatomic",
      Type:"Organization",
      URL:"https://api.github.com/users/projectatomic",
      EventsURL:"https://api.github.com/users/projectatomic/events{/privacy}",
      ReposURL:"https://api.github.com/users/projectatomic/repos"
   },
   Permissions:map   [
      pull:true      admin:true      push:true
   ],
   Private:false,
   HasIssues:true,
   HasWiki:false,
   HasDownloads:true,
   URL:"https://api.github.com/repos/projectatomic/atomicapp",
   ArchiveURL:"https://api.github.com/repos/projectatomic/atomicapp/{archive_format}{/ref}",
   AssigneesURL:"https://api.github.com/repos/projectatomic/atomicapp/assignees{/user}",
   BlobsURL:"https://api.github.com/repos/projectatomic/atomicapp/git/blobs{/sha}",
   BranchesURL:"https://api.github.com/repos/projectatomic/atomicapp/branches{/branch}",
   CollaboratorsURL:"https://api.github.com/repos/projectatomic/atomicapp/collaborators{/collaborator}",
   CommentsURL:"https://api.github.com/repos/projectatomic/atomicapp/comments{/number}",
   CommitsURL:"https://api.github.com/repos/projectatomic/atomicapp/commits{/sha}",
   CompareURL:"https://api.github.com/repos/projectatomic/atomicapp/compare/{base}...{head}",
   ContentsURL:"https://api.github.com/repos/projectatomic/atomicapp/contents/{+path}",
   ContributorsURL:"https://api.github.com/repos/projectatomic/atomicapp/contributors",
   DownloadsURL:"https://api.github.com/repos/projectatomic/atomicapp/downloads",
   EventsURL:"https://api.github.com/repos/projectatomic/atomicapp/events",
   ForksURL:"https://api.github.com/repos/projectatomic/atomicapp/forks",
   GitCommitsURL:"https://api.github.com/repos/projectatomic/atomicapp/git/commits{/sha}",
   GitRefsURL:"https://api.github.com/repos/projectatomic/atomicapp/git/refs{/sha}",
   GitTagsURL:"https://api.github.com/repos/projectatomic/atomicapp/git/tags{/sha}",
   HooksURL:"https://api.github.com/repos/projectatomic/atomicapp/hooks",
   IssueCommentURL:"https://api.github.com/repos/projectatomic/atomicapp/issues/comments{/number}",
   IssueEventsURL:"https://api.github.com/repos/projectatomic/atomicapp/issues/events{/number}",
   IssuesURL:"https://api.github.com/repos/projectatomic/atomicapp/issues{/number}",
   KeysURL:"https://api.github.com/repos/projectatomic/atomicapp/keys{/key_id}",
   LabelsURL:"https://api.github.com/repos/projectatomic/atomicapp/labels{/name}",
   LanguagesURL:"https://api.github.com/repos/projectatomic/atomicapp/languages",
   MergesURL:"https://api.github.com/repos/projectatomic/atomicapp/merges",
   MilestonesURL:"https://api.github.com/repos/projectatomic/atomicapp/milestones{/number}",
   NotificationsURL:"https://api.github.com/repos/projectatomic/atomicapp/notifications{?since,all,participating}",
   PullsURL:"https://api.github.com/repos/projectatomic/atomicapp/pulls{/number}",
   ReleasesURL:"https://api.github.com/repos/projectatomic/atomicapp/releases{/id}",
   StargazersURL:"https://api.github.com/repos/projectatomic/atomicapp/stargazers",
   StatusesURL:"https://api.github.com/repos/projectatomic/atomicapp/statuses/{sha}",
   SubscribersURL:"https://api.github.com/repos/projectatomic/atomicapp/subscribers",
   SubscriptionURL:"https://api.github.com/repos/projectatomic/atomicapp/subscription",
   TagsURL:"https://api.github.com/repos/projectatomic/atomicapp/tags",
   TreesURL:"https://api.github.com/repos/projectatomic/atomicapp/git/trees{/sha}",
   TeamsURL:"https://api.github.com/repos/projectatomic/atomicapp/teams"
}


{
   Number:(*int)(0xc82045ac00),
   State:(*string)(0xc82045ad70),
   Title:(*string)(0xc82045ac20),
   Body:(*string)(0xc82045b0f0),
   User:(*github.User)(0xc8204e8000),
   Labels:[

   ]   github.Label   {

   },
   Assignee:(*github.User)(nil),
   Comments:(*int)(0xc82045b0d0),
   ClosedAt:(*time.Time)(nil),
   CreatedAt:(*time.Time)(0xc8203c0da0),
   UpdatedAt:(*time.Time)(0xc8203c0de0),
   URL:(*string)(0xc82045aae0),
   HTMLURL:(*string)(0xc82045abb0),
   Milestone:(*github.Milestone)(0xc8201d0300),
   PullRequestLinks:(*github.PullRequestLinks)(nil),
   TextMatches:[

   ]   github.TextMatch(nil)
}
*/
