//go:build !dev && !app
// +build !dev,!app

package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

//go:embed assets/*
var embeddedFiles embed.FS
var headless bool
var devmode bool
var port int
var logger zerolog.Logger

func init() {
	// Define the headless flag
	flag.BoolVar(&devmode, "devmode", false, "Run with passthrough vite server")
	flag.BoolVar(&headless, "headless", false, "Run in headless mode without opening the browser")
	flag.IntVar(&port, "port", 0, "Run in headless mode without opening the browser")
	flag.Parse()

	logger = zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	if(!headless) {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func main() {
	// Signal catching for clean shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	done := make(chan struct{})
	failed := make(chan struct{})

	logger.Info().Bool("headless mode", !headless).Msg("the app is starting")

	assets, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		logger.Info().Err(err).Msg("Failed to locate embedded assets")
		return
	}

	// Get the port the server is listening on.
	// Listen on a random port.
	listenHost := "localhost"
	if(headless) {
		listenHost = "0.0.0.0"
	}

	listenAddr := fmt.Sprintf("%s:%d", listenHost, port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to listen on a port: %v")
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	serverURL := fmt.Sprintf("http://localhost:%d", port)
	logger.Info().Str("serverUrl", serverURL).Msgf("Production server listening on %s", serverURL)

    // Register API routes.
	RegisterAPIRoutes()
	
	if !headless {
		logger.Info().Msg("Opening browser")
		go func() {
			waitForServer(serverURL)
			openbrowser(serverURL)
		}()
	}

	// passthrough vite server
	// todo get port randomized
	if(devmode) {
		viteServerURL, _ := url.Parse("http://localhost:5173")
		proxy := httputil.NewSingleHostReverseProxy(viteServerURL)

		// Handle all other requests by forwarding them to the Vite server
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})
	} else {
		// Use the http.FileServer to serve the embedded assets.
		http.Handle("/", http.FileServer(http.FS(assets)))
	}

	// Start API in a separate goroutine.
	go func() {
		err := http.Serve(listener, nil)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Warn().Err(err).Msg("Closed Server")
			close(failed)
		}
	}()

	select {
		case <-sig:
			logger.Info().Msg("MyApp stopping")
		case <-done:
			logger.Info().Msg("MyApp stopped")
		case <-failed:
			logger.Info().Msg("MyApp process was aborted")
	}

	// If we receive a second interrupt signal, exit immediately.
	go func() {
		<-sig
		logger.Warn().Msg("forcing exit")
		os.Exit(1)
	}()
}

func waitForServer(url string) {
	for {
		// Attempt to connect to the server.
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close() // Don't forget to close the response body.
			logger.Info().Msg("App is Running. CTRL+C to quit.")
			return
		}
		// Close the unsuccessful response body to avoid leaking resources.
		if resp != nil {
			resp.Body.Close()
		}
		// Wait for a second before trying again.
		time.Sleep(1 * time.Second)
	}
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		logger.Fatal().Err(err)
	}

}