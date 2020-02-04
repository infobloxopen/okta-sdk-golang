package okta

import (
	"fmt"

	"github.com/okta/okta-sdk-golang/okta/query"
)

type OktaSignOnPolicyResource resource

func (m *OktaSignOnPolicyResource) GetPolicy(policyId string, qp *query.Params) (*OktaSignOnPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policy *OktaSignOnPolicy
	resp, err := m.client.requestExecutor.Do(req, &policy)
	if err != nil {
		return nil, resp, err
	}
	return policy, resp, nil
}
func (m *OktaSignOnPolicyResource) UpdatePolicy(policyId string, body OktaSignOnPolicy) (*OktaSignOnPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)
	req, err := m.client.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policy *OktaSignOnPolicy
	resp, err := m.client.requestExecutor.Do(req, &policy)
	if err != nil {
		return nil, resp, err
	}
	return policy, resp, nil
}
func (m *OktaSignOnPolicyResource) DeletePolicy(policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)
	req, err := m.client.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *OktaSignOnPolicyResource) ListPolicies(qp *query.Params) ([]*OktaSignOnPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policy []*OktaSignOnPolicy
	resp, err := m.client.requestExecutor.Do(req, &policy)
	if err != nil {
		return nil, resp, err
	}
	return policy, resp, nil
}
func (m *OktaSignOnPolicyResource) CreatePolicy(body Policy, qp *query.Params) (*OktaSignOnPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policy *OktaSignOnPolicy
	resp, err := m.client.requestExecutor.Do(req, &policy)
	if err != nil {
		return nil, resp, err
	}
	return policy, resp, nil
}
func (m *OktaSignOnPolicyResource) ActivatePolicy(policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/lifecycle/activate", policyId)
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *OktaSignOnPolicyResource) DeactivatePolicy(policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/lifecycle/deactivate", policyId)
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *OktaSignOnPolicyResource) ListPolicyRules(policyId string) ([]*OktaSignOnPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules", policyId)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policyRule []*OktaSignOnPolicyRule
	resp, err := m.client.requestExecutor.Do(req, &policyRule)
	if err != nil {
		return nil, resp, err
	}
	return policyRule, resp, nil
}
func (m *OktaSignOnPolicyResource) AddPolicyRule(policyId string, body OktaSignOnPolicyRule, qp *query.Params) (*OktaSignOnPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules", policyId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *OktaSignOnPolicyRule
	resp, err := m.client.requestExecutor.Do(req, &policyRule)
	if err != nil {
		return nil, resp, err
	}
	return policyRule, resp, nil
}
func (m *OktaSignOnPolicyResource) DeletePolicyRule(policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)
	req, err := m.client.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *OktaSignOnPolicyResource) GetPolicyRule(policyId string, ruleId string) (*OktaSignOnPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *OktaSignOnPolicyRule
	resp, err := m.client.requestExecutor.Do(req, &policyRule)
	if err != nil {
		return nil, resp, err
	}
	return policyRule, resp, nil
}
func (m *OktaSignOnPolicyResource) UpdatePolicyRule(policyId string, ruleId string, body OktaSignOnPolicyRule) (*OktaSignOnPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)
	req, err := m.client.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *OktaSignOnPolicyRule
	resp, err := m.client.requestExecutor.Do(req, &policyRule)
	if err != nil {
		return nil, resp, err
	}
	return policyRule, resp, nil
}
func (m *OktaSignOnPolicyResource) ActivatePolicyRule(policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v/lifecycle/activate", policyId, ruleId)
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *OktaSignOnPolicyResource) DeactivatePolicyRule(policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v/lifecycle/deactivate", policyId, ruleId)
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type OktaSignOnPolicyRuleResource resource

func (m *OktaSignOnPolicyRuleResource) UpdatePolicyRule(policyId string, ruleId string, body OktaSignOnPolicyRule) (*OktaSignOnPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)
	req, err := m.client.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *OktaSignOnPolicyRule
	resp, err := m.client.requestExecutor.Do(req, &policyRule)
	if err != nil {
		return nil, resp, err
	}
	return policyRule, resp, nil
}
func (m *OktaSignOnPolicyRuleResource) DeletePolicyRule(policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)
	req, err := m.client.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
