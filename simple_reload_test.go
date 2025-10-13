package godev

import (
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestSimpleBrowserReload creates a single file, waits long enough for timer to expire
func TestSimpleBrowserReload(t *testing.T) {
	tmp := t.TempDir()

	var reloadCount int64

	logger := func(messages ...any) {
		var msg string
		for i, m := range messages {
			if i > 0 {
				msg += " "
			}
			msg += fmt.Sprint(m)
		}
		logIfVerbose(t, "LOG: %s", msg)
	}

	// Start godev
	exitChan := make(chan bool)
	go Start(tmp, logger, newUiMockTest(logger), exitChan)

	time.Sleep(200 * time.Millisecond)

	// Set up browser reload tracking
	SetWatcherBrowserReload(func() error {
		count := atomic.AddInt64(&reloadCount, 1)
		logIfVerbose(t, "*** BROWSER RELOAD CALLED! Count: %d ***", count)
		return nil
	})

	time.Sleep(100 * time.Millisecond)

	// Create and modify ONE file, then wait a long time
	jsFile := filepath.Join(tmp, "modules", "test", "simple.js")
	require.NoError(t, os.MkdirAll(filepath.Dir(jsFile), 0755))
	require.NoError(t, os.WriteFile(jsFile, []byte("console.log('initial');"), 0644))

	logIfVerbose(t, "=== File created, waiting for initial processing ===")
	time.Sleep(500 * time.Millisecond)

	initialCount := atomic.LoadInt64(&reloadCount)
	logIfVerbose(t, "Reload count after initial creation: %d", initialCount)

	// Modify the file ONCE
	logIfVerbose(t, "=== Single modification ===")
	require.NoError(t, os.WriteFile(jsFile, []byte("console.log('modified');"), 0644))

	// Wait long enough for timer to definitely expire (much longer than 100ms debounce)
	logIfVerbose(t, "=== Waiting 1 second for timer to expire ===")
	time.Sleep(1 * time.Second)

	finalCount := atomic.LoadInt64(&reloadCount)
	logIfVerbose(t, "Final reload count: %d", finalCount)

	exitChan <- true

	if finalCount > initialCount {
		t.Logf("✓ Browser reload was called %d times", finalCount-initialCount)
	} else {
		t.Errorf("Browser reload was never called even with single file modification and 1 second wait")
	}
}
