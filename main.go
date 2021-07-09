package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
)

func main() {
    port := os.Getenv("HEALTH_CHECKER_PORT")

    list := "MICRO_A;MICRO_B"

    if port == "" {
        panic("HEALTH_CHECKER_PORT env var not set")
    }

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        msg := fmt.Sprintf("Health-checker service is OK")
        w.Write([]byte(msg))
    })

    http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        responseString := ""
        for _, service := range strings.Split(list, ";") {
            hostKey := fmt.Sprintf("%s_HOST", service)
            portKey := fmt.Sprintf("%s_SERVICE_PORT", service)
            host := os.Getenv(hostKey)
            port := os.Getenv(portKey)
            url := fmt.Sprintf("http://%s:%s/health", host, port)
            log.Println(url)
            resp, err := http.Get(url)
            if err != nil {
                log.Println(err)
            }
            //We Read the response body on the line below.
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                log.Println(err)
            }
            responseString = fmt.Sprintf("%s\n", body)
            w.Write([]byte(responseString))
        }
        w.WriteHeader(http.StatusOK)
        //msg := fmt.Sprintf("Runnable %s is OK", name)

    })

    addr := fmt.Sprintf(":%s", port)
    http.ListenAndServe(addr, nil)
}