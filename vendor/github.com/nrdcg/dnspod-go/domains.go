package dnspod

import (
	"encoding/json"
	"fmt"
	"strconv"
	// "time"
)

const (
	methodDomainList   = "Domain.List"
	methodDomainCreate = "Domain.Create"
	methodDomainInfo   = "Domain.Info"
	methodDomainRemove = "Domain.Remove"
)

// DomainInfo handles domain information.
type DomainInfo struct {
	DomainTotal   json.Number `json:"domain_total,omitempty"`
	AllTotal      json.Number `json:"all_total,omitempty"`
	MineTotal     json.Number `json:"mine_total,omitempty"`
	ShareTotal    json.Number `json:"share_total,omitempty"`
	VipTotal      json.Number `json:"vip_total,omitempty"`
	IsMarkTotal   json.Number `json:"ismark_total,omitempty"`
	PauseTotal    json.Number `json:"pause_total,omitempty"`
	ErrorTotal    json.Number `json:"error_total,omitempty"`
	LockTotal     json.Number `json:"lock_total,omitempty"`
	SpamTotal     json.Number `json:"spam_total,omitempty"`
	VipExpire     json.Number `json:"vip_expire,omitempty"`
	ShareOutTotal json.Number `json:"share_out_total,omitempty"`
}

// Domain handles domain.
type Domain struct {
	ID               json.Number `json:"id,omitempty"`
	Name             string      `json:"name,omitempty"`
	PunyCode         string      `json:"punycode,omitempty"`
	Grade            string      `json:"grade,omitempty"`
	GradeTitle       string      `json:"grade_title,omitempty"`
	Status           string      `json:"status,omitempty"`
	ExtStatus        string      `json:"ext_status,omitempty"`
	Records          string      `json:"records,omitempty"`
	GroupID          json.Number `json:"group_id,omitempty"`
	IsMark           string      `json:"is_mark,omitempty"`
	Remark           string      `json:"remark,omitempty"`
	IsVIP            string      `json:"is_vip,omitempty"`
	SearchenginePush string      `json:"searchengine_push,omitempty"`
	UserID           string      `json:"user_id,omitempty"`
	CreatedOn        string      `json:"created_on,omitempty"`
	UpdatedOn        string      `json:"updated_on,omitempty"`
	TTL              string      `json:"ttl,omitempty"`
	CNameSpeedUp     string      `json:"cname_speedup,omitempty"`
	Owner            string      `json:"owner,omitempty"`
	AuthToAnquanBao  bool        `json:"auth_to_anquanbao,omitempty"`
}

type domainListWrapper struct {
	Status  Status     `json:"status"`
	Info    DomainInfo `json:"info"`
	Domains []Domain   `json:"domains"`
}

type domainWrapper struct {
	Status Status     `json:"status"`
	Info   DomainInfo `json:"info"`
	Domain Domain     `json:"domain"`
}

// DomainsService handles communication with the domain related methods of the dnspod API.
//
// dnspod API docs: https://www.dnspod.cn/docs/domains.html
type DomainsService struct {
	client *Client
}

// List the domains.
//
// dnspod API docs: https://www.dnspod.cn/docs/domains.html#domain-list
func (s *DomainsService) List() ([]Domain, *Response, error) {
	payload := s.client.CommonParams.toPayLoad()

	returnedDomains := domainListWrapper{}

	res, err := s.client.post(methodDomainList, payload, &returnedDomains)
	if err != nil {
		return nil, res, err
	}

	if returnedDomains.Status.Code != "1" {
		return nil, nil, fmt.Errorf("could not get domains: %s", returnedDomains.Status.Message)
	}

	return returnedDomains.Domains, res, nil
}

// Create a new domain.
//
// dnspod API docs: https://www.dnspod.cn/docs/domains.html#domain-create
func (s *DomainsService) Create(domainAttributes Domain) (Domain, *Response, error) {
	payload := s.client.CommonParams.toPayLoad()
	payload.Set("domain", domainAttributes.Name)
	payload.Set("group_id", domainAttributes.GroupID.String())
	payload.Set("is_mark", domainAttributes.IsMark)

	returnedDomain := domainWrapper{}

	res, err := s.client.post(methodDomainCreate, payload, &returnedDomain)
	if err != nil {
		return Domain{}, res, err
	}

	return returnedDomain.Domain, res, nil
}

// Get fetches a domain.
//
// dnspod API docs: https://www.dnspod.cn/docs/domains.html#domain-info
func (s *DomainsService) Get(id int) (Domain, *Response, error) {
	payload := s.client.CommonParams.toPayLoad()
	payload.Set("domain_id", strconv.Itoa(id))

	returnedDomain := domainWrapper{}

	res, err := s.client.post(methodDomainInfo, payload, &returnedDomain)
	if err != nil {
		return Domain{}, res, err
	}

	return returnedDomain.Domain, res, nil
}

// Delete a domain.
//
// dnspod API docs: https://dnsapi.cn/Domain.Remove
func (s *DomainsService) Delete(id int) (*Response, error) {
	payload := s.client.CommonParams.toPayLoad()
	payload.Set("domain_id", strconv.Itoa(id))

	returnedDomain := domainWrapper{}

	return s.client.post(methodDomainRemove, payload, &returnedDomain)
}
