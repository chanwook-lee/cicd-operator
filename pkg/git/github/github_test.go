package github

import (
	"github.com/bmizerany/assert"
	"github.com/gorilla/mux"
	cicdv1 "github.com/tmax-cloud/cicd-operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"net/http"
	"net/http/httptest"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

const (
	samplePRList = "[{\"url\":\"https://api.github.com/repos/vingsu/cicd-test/pulls/25\",\"id\":611161419,\"node_id\":\"MDExOlB1bGxSZXF1ZXN0NjExMTYxNDE5\",\"html_url\":\"https://github.com/vingsu/cicd-test/pull/25\",\"number\":25,\"state\":\"open\",\"locked\":false,\"title\":\"newnew\",\"user\":{\"login\":\"cqbqdd11519\",\"id\":6166781,\"node_id\":\"MDQ6VXNlcjYxNjY3ODE=\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/6166781?v=4\",\"gravatar_id\":\"\",\"type\":\"User\",\"site_admin\":false},\"body\":\"\",\"created_at\":\"2021-04-08T02:35:17Z\",\"updated_at\":\"2021-04-13T04:54:16Z\",\"closed_at\":null,\"merged_at\":null,\"merge_commit_sha\":\"b6d9abd3254a6b3da35200f9cdbb307cea7db91a\",\"assignee\":null,\"assignees\":[],\"requested_reviewers\":[{\"login\":\"sunghyunkim3\",\"id\":66240202,\"node_id\":\"MDQ6VXNlcjY2MjQwMjAy\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/66240202?v=4\",\"gravatar_id\":\"\",\"type\":\"User\",\"site_admin\":false}],\"requested_teams\":[],\"labels\":[{\"id\":2905890093,\"node_id\":\"MDU6TGFiZWwyOTA1ODkwMDkz\",\"url\":\"https://api.github.com/repos/vingsu/cicd-test/labels/kind/test\",\"name\":\"kind/test\",\"color\":\"CF61D3\",\"default\":false,\"description\":\"\"}],\"milestone\":null,\"draft\":false,\"head\":{\"label\":\"vingsu:newnew\",\"ref\":\"newnew\",\"sha\":\"3196ccc37bcae94852079b04fcbfaf928341d6e9\",\"user\":{\"login\":\"vingsu\",\"id\":71878727,\"node_id\":\"MDEyOk9yZ2FuaXphdGlvbjcxODc4NzI3\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/71878727?v=4\",\"gravatar_id\":\"\",\"type\":\"Organization\",\"site_admin\":false},\"repo\":{\"id\":319253224,\"node_id\":\"MDEwOlJlcG9zaXRvcnkzMTkyNTMyMjQ=\",\"name\":\"cicd-test\",\"full_name\":\"vingsu/cicd-test\",\"private\":false,\"owner\":{\"login\":\"vingsu\",\"id\":71878727,\"node_id\":\"MDEyOk9yZ2FuaXphdGlvbjcxODc4NzI3\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/71878727?v=4\",\"gravatar_id\":\"\",\"type\":\"Organization\",\"site_admin\":false},\"html_url\":\"https://github.com/vingsu/cicd-test\",\"description\":null,\"fork\":false,\"created_at\":\"2020-12-07T08:31:55Z\",\"updated_at\":\"2021-01-27T04:29:32Z\",\"pushed_at\":\"2021-04-09T04:46:39Z\",\"git_url\":\"git://github.com/vingsu/cicd-test.git\",\"ssh_url\":\"git@github.com:vingsu/cicd-test.git\",\"clone_url\":\"https://github.com/vingsu/cicd-test.git\",\"svn_url\":\"https://github.com/vingsu/cicd-test\",\"homepage\":null,\"size\":10,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"HTML\",\"has_issues\":true,\"has_projects\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":0,\"mirror_url\":null,\"archived\":false,\"disabled\":false,\"open_issues_count\":1,\"license\":null,\"forks\":0,\"open_issues\":1,\"watchers\":0,\"default_branch\":\"master\"}},\"base\":{\"label\":\"vingsu:master\",\"ref\":\"master\",\"sha\":\"22ccae53032027186ba739dfaa473ee61a82b298\",\"user\":{\"login\":\"vingsu\",\"id\":71878727,\"node_id\":\"MDEyOk9yZ2FuaXphdGlvbjcxODc4NzI3\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/71878727?v=4\",\"gravatar_id\":\"\",\"type\":\"Organization\",\"site_admin\":false},\"repo\":{\"id\":319253224,\"node_id\":\"MDEwOlJlcG9zaXRvcnkzMTkyNTMyMjQ=\",\"name\":\"cicd-test\",\"full_name\":\"vingsu/cicd-test\",\"private\":false,\"owner\":{\"login\":\"vingsu\",\"id\":71878727,\"node_id\":\"MDEyOk9yZ2FuaXphdGlvbjcxODc4NzI3\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/71878727?v=4\",\"gravatar_id\":\"\",\"type\":\"Organization\",\"site_admin\":false},\"html_url\":\"https://github.com/vingsu/cicd-test\",\"description\":null,\"fork\":false,\"created_at\":\"2020-12-07T08:31:55Z\",\"updated_at\":\"2021-01-27T04:29:32Z\",\"pushed_at\":\"2021-04-09T04:46:39Z\",\"git_url\":\"git://github.com/vingsu/cicd-test.git\",\"ssh_url\":\"git@github.com:vingsu/cicd-test.git\",\"clone_url\":\"https://github.com/vingsu/cicd-test.git\",\"svn_url\":\"https://github.com/vingsu/cicd-test\",\"homepage\":null,\"size\":10,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"HTML\",\"has_issues\":true,\"has_projects\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":0,\"mirror_url\":null,\"archived\":false,\"disabled\":false,\"open_issues_count\":1,\"license\":null,\"forks\":0,\"open_issues\":1,\"watchers\":0,\"default_branch\":\"master\"}},\"author_association\":\"CONTRIBUTOR\",\"auto_merge\":null,\"active_lock_reason\":null}]"
)

func TestClient_ListPullRequests(t *testing.T) {
	c, err := testEnv()
	if err != nil {
		t.Fatal(err)
	}

	prs, err := c.ListPullRequests(false)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(prs), "Length")
	assert.Equal(t, 25, prs[0].ID, "ID")
	assert.Equal(t, "newnew", prs[0].Title, "Title")
}

func testEnv() (*Client, error) {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(req.URL.String()))
	})
	r.HandleFunc("/repos/{org}/{repo}/pulls", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(samplePRList))
	})
	testSrv := httptest.NewServer(r)

	s := runtime.NewScheme()
	utilruntime.Must(cicdv1.AddToScheme(s))

	ic := &cicdv1.IntegrationConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ic",
			Namespace: "default",
		},
		Spec: cicdv1.IntegrationConfigSpec{
			Git: cicdv1.GitConfig{
				Type:       "github",
				Repository: "tmax-cloud/cicd-test",
				APIUrl:     testSrv.URL,
				Token:      &cicdv1.GitToken{Value: "dummy"},
			},
		},
	}

	c := &Client{
		IntegrationConfig: ic,
		K8sClient:         fake.NewFakeClientWithScheme(s, ic),
	}
	if err := c.Init(); err != nil {
		return nil, err
	}

	return c, nil
}