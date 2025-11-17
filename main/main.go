package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	q "github.com/robindallier/quizzgo-team-ilzikout"
	"github.com/robindallier/quizzgo-team-ilzikout/web"
)

func openBrowser(url string) error {
	// Windows: use rundll32 to open default browser
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	return cmd.Start()
}

func waitForServer(url string, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(100 * time.Millisecond)
	}
	return false
}

func main() {
	go func() {
		if err := web.Start(":8080"); err != nil {
			log.Printf("web server stopped: %v", err)
		}
	}()

	fmt.Println("Web UI available at: http://localhost:8080")

	// wait for server and open the quiz page automatically
	url := "http://localhost:8080/quiz/cyber"
	if waitForServer("http://localhost:8080/", 3*time.Second) {
		if err := openBrowser(url); err != nil {
			log.Printf("failed to open browser: %v", err)
		}
	} else {
		log.Printf("server did not start within timeout; open %s manually", url)
	}

	q.Menu()
}
