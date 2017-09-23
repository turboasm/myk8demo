package handlers

import (
	"net/http"
	"testing"

	"github.com/turboasm/myk8demo/pkg/config"
	"github.com/turboasm/myk8demo/pkg/logger"
	"github.com/turboasm/myk8demo/pkg/logger/standard"
	"github.com/turboasm/myk8demo/pkg/router/bitroute"
)

func TestHealth(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Health)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
