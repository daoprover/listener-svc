package github

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/daoprover/listener-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func NewGithubListener(ctx context.Context, db data.MasterQ, id string, name string, address string) GithubListener {

	return &GithubListenerData{
		Ctx:     ctx,
		DB:      db,
		ID:      id,
		Name:    name,
		Address: address,
	}
}

func (g *GithubListenerData) Run() error {
	preparedURL := g.prepareOrgURL(g.Name)
	_ = preparedURL
	org, err := g.getOrganizations(g.Name)
	if err != nil {
		return errors.Wrap(err, "failed to get org by name ")
	}

	itemsCount := len(org.Items)
	if itemsCount != 0 {
		//todo  insert  to db
		return nil
	}

	repos, err := g.getRepositories(g.Name, "")
	if err != nil {
		return errors.Wrap(err, "failed to get repos by name")
	}

	reposCount := len(repos.Items)
	fmt.Println("repos: ", repos)
	if reposCount != 0 {
		//todo  insert  to db
		return nil
	}

	//todo  insert  to db empty data

	return nil
}

func (g *GithubListenerData) prepareOrgURL(name string) string {
	return fmt.Sprintf("%s/search/users?q=%s", githubSearchApi, name)
}

func (g *GithubListenerData) prepareReposURL(name string, language string) string {
	return fmt.Sprintf("%s/search/repositories/?q=%s+in:name+language:%s&sort=stars&order=desc", githubSearchApi, name, language)
}

func (g *GithubListenerData) getOrganizations(name string) (*SearchUserResponse, error) {
	response, err := http.Get(g.prepareOrgURL(name))
	if err != nil {
		return nil, errors.Wrap(err, "failed to call api")
	}

	respData := new(SearchUserResponse)
	err = json.NewDecoder(response.Body).Decode(respData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to  decode response")
	}

	return respData, nil
}

func (g *GithubListenerData) getRepositories(name string, language string) (*SearchRepoResponse, error) {
	link := g.prepareReposURL(name, language)
	response, err := http.Get(link)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call api")
	}

	respData := new(SearchRepoResponse)
	err = json.NewDecoder(response.Body).Decode(respData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to  decode response")
	}

	return respData, nil
}
