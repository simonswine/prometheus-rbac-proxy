package prometheus

import (
	promlabels "github.com/prometheus/prometheus/pkg/labels"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LabelSelectorToPromLabels converts kubernetes style label selectors to
// prometheus matchers
func LabelSelectorToPromLabels(selector *metav1.LabelSelector) []*promlabels.Matcher {
	if selector == nil {
		return []*promlabels.Matcher{}
	}

	var result []*promlabels.Matcher

	if selector.MatchLabels != nil {
		for name, value := range selector.MatchLabels {
			result = append(result, &promlabels.Matcher{
				Type:  promlabels.MatchEqual,
				Name:  name,
				Value: value,
			})
		}
	}

	return result
}
