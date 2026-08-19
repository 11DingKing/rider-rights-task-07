package dispatch

import (
	"testing"

	"riderguard/internal/domain"
)

func TestEscalationLeadUsesRequestedLevel(t *testing.T) {
	adj := &Adjudicator{}
	if got := adj.SelectEscalationLead("朱家湖服务站", 1); got != "朱家湖服务站-supervisor-l1" {
		t.Fatalf("unexpected escalation lead: %q", got)
	}
	if got := domain.EscalationDepartmentName("朱家湖服务站", 2); got != "朱家湖服务站-supervisor-l2" {
		t.Fatalf("unexpected domain escalation lead: %q", got)
	}
}
