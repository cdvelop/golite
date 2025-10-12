package godev

import (
	"os"
	"sync"

	. "github.com/cdvelop/assetmin"
	"github.com/cdvelop/devbrowser"
	. "github.com/cdvelop/devtui"
	"github.com/cdvelop/devwatch"
	"github.com/cdvelop/goflare"
	"github.com/cdvelop/goserver"
	"github.com/cdvelop/tinywasm"
)

type handler struct {
	rootDir       string  // Application root directory
	config        *Config // Main configuration source
	tui           *DevTUI
	serverHandler *goserver.ServerHandler
	assetsHandler *AssetMin
	wasmHandler   *tinywasm.TinyWasm
	watcher       *devwatch.DevWatch
	browser       *devbrowser.DevBrowser

	deployCloudflare *goflare.Goflare

	exitChan chan bool // Canal global para señalizar el cierre
	// pendingBrowserReload is used by tests to set a custom BrowserReload
	// before the watcher is created. If non-nil it will be applied when
	// AddSectionBUILD creates the watcher.
	pendingBrowserReload func() error
}

func Start(rootDir string, logger func(messages ...any), exitChan chan bool) {
	h := &handler{
		rootDir:  rootDir,
		exitChan: exitChan,
		// goDepFind:  godepfind.New(rootDir),
	}

	// Make the handler available to tests so they can override the
	// BrowserReload callback when needed.
	ActiveHandler = h

	// Validate we're not in system directories
	homeDir, _ := os.UserHomeDir()
	if rootDir == homeDir || rootDir == "/" {
		// Use the provided logger since Translator is not initialized yet
		logger("Cannot run godev in user root directory. Please run in a Go project directory")
		return
	}

	h.tui = NewTUI(&TuiConfig{
		AppName:  "GODEV",
		ExitChan: h.exitChan,
		Color:    DefaultPalette(),
		Logger:   func(messages ...any) { logger(messages...) },
	}) // Initialize AutoConfig FIRST - this will be our configuration source

	// ADD SECTIONS
	h.AddSectionBUILD()
	h.AddSectionDEPLOY()

	var wg sync.WaitGroup
	wg.Add(3)

	// Start the tui in a goroutine
	go h.tui.Start(&wg)

	// Iniciar servidor
	go h.serverHandler.StartServer(&wg)

	// Iniciar el watcher de archivos
	go h.watcher.FileWatcherStart(&wg)

	// Esperar a que todas las goroutines terminen
	wg.Wait()
}
