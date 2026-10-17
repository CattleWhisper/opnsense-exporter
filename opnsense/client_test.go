package opnsense

import (
	"compress/gzip"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AthennaMind/opnsense-exporter/internal/options"
)

// The client only knows how to decode gzip, so it must not advertise
// any other encoding, and it must decode a gzip body.
func TestClientRequestsAndDecodesGzipOnly(t *testing.T) {
	var gotEncoding string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotEncoding = r.Header.Get("Accept-Encoding")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		_, _ = gz.Write([]byte(`{"memory": {"states": 400000}}`))
		_ = gz.Close()
	}))
	defer server.Close()

	client, err := NewClient(options.OPNSenseConfig{
		Protocol:  "http",
		Host:      strings.TrimPrefix(server.URL, "http://"),
		APIKey:    "test",
		APISecret: "test",
	}, "test", slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		t.Fatalf("failed to build client: %v", err)
	}

	limits, apiErr := client.FetchPFMemoryLimits()
	if apiErr != nil {
		t.Fatalf("unexpected error: %s", apiErr.Error())
	}
	if gotEncoding != "gzip" {
		t.Errorf("expected Accept-Encoding \"gzip\", got %q", gotEncoding)
	}
	if limits.Limits["states"] != 400000 {
		t.Errorf("unexpected limits: %+v", limits.Limits)
	}
}
