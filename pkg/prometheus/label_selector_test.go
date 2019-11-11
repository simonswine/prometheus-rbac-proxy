package prometheus

import (
	"reflect"
	"testing"

	promlabels "github.com/prometheus/prometheus/pkg/labels"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type testCaseLabelSelector struct {
	input    *metav1.LabelSelector
	expected []*promlabels.Matcher
}

func checkTestCaseLabelSelector(t *testing.T, tc []testCaseLabelSelector) {
	for _, x := range tc {
		actual := LabelSelectorToPromLabels(x.input)

		if len(x.expected) == 0 && len(actual) == 0 {
			continue
		}

		if !reflect.DeepEqual(actual, x.expected) {
			t.Errorf("Unexpected result input=%+#v act=%+#v exp=%+#v", x.input, actual, x.expected)
		}
	}
}

func TestLabelSelectorToPromLabelsEmpty(t *testing.T) {
	checkTestCaseLabelSelector(t, []testCaseLabelSelector{
		{
			&metav1.LabelSelector{},
			[]*promlabels.Matcher{},
		},
		{
			nil,
			[]*promlabels.Matcher{},
		},
	})
}

func TestLabelSelectorToPromLabelsMatchLabels(t *testing.T) {
	checkTestCaseLabelSelector(t, []testCaseLabelSelector{
		{
			&metav1.LabelSelector{
				MatchLabels: map[string]string{
					"team": "infrastructure",
					"env":  "prod",
				},
			},
			[]*promlabels.Matcher{},
		},
	})
}
