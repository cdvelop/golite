package godev

import (
	"os"
	"sync"

	. "github.com/cdvelop/assetmin"
	"github.com/cdvelop/devbrowser"
	"github.com/cdvelop/devwatch"
	"github.com/cdvelop/goflare"
	"github.com/cdvelop/goserver"
	"github.com/cdvelop/tinywasm"
)

// handler contains application state and dependencies
// CRITICAL: This struct does NOT import DevTUI
type handler struct {
	rootDir   string
	config    *Config
	tui       TuiInterface // Interface defined in GODEV, not DevTUI
	exitChan  chan bool

	// Build dependencies
	serverHandler *goserver.ServerHandler
	assetsHandler *AssetMin
	wasmHandler   *tinywasm.TinyWasm
	watcher       *devwatch.DevWatch
	browser       *devbrowser.DevBrowser

	// Deploy dependencies
	deployCloudflare *goflare.Goflare

	// Test hooks
	pendingBrowserReload func() error
}

// Start is called from main.go with UI passed as parameter
// CRITICAL: UI instance created in main.go, passed here as interface
func Start(rootDir string, logger func(messages ...any), ui TuiInterface, exitChan chan bool) {
	h := &handler{
		rootDir:  rootDir,
		tui:      ui, // UI passed from main.go
		exitChan: exitChan,
	}

	ActiveHandler = h

	// Validate directory
	homeDir, _ := os.UserHomeDir()
	if rootDir == homeDir || rootDir == "/" {
		logger("Cannot run godev in user root directory. Please run in a Go project directory")
		return
	}

	// ADD SECTIONS using the passed UI interface
	h.AddSectionBUILD()
	h.AddSectionDEPLOY()

	var wg sync.WaitGroup
	wg.Add(3)

	// Start the UI (passed from main.go)
	go h.tui.Start(&wg)

	// Start server
	go h.serverHandler.StartServer(&wg)

	// Start file watcher
	go h.watcher.FileWatcherStart(&wg)

	wg.Wait()
}