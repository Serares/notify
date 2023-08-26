//go:build integration
// +build integration

package notify_test

import (
	"testing"

	"github.com/Serares/notify"
)

func TestSend(t *testing.T) {

	n := notify.New("testing title", "test msg", notify.SeverityNormal)
	err := n.Send()
	if err != nil {
		t.Error(err)
	}
}
