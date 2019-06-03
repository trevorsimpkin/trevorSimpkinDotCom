// Package main listens and serves on port 3000
package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
        if port == "" {
            port = "3000"
        }

        f,err := os.Create("/home/administrator/go/log/log.txt")
        if err != nil{
            log.Println(err)
        }
        defer f.Close()
        log.SetOutput(f)
        //certManager := autocert.Manager{
        //    Prompt:     autocert.AcceptTOS,
        //    HostPolicy: autocert.HostWhitelist("localhost"),
        //    Cache:      autocert.DirCache("certs"),
        //}
        const indexPage = "public/index.gohtml"
        http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            if r.Method == "POST" {
                if buf, err := ioutil.ReadAll(r.Body); err == nil {
                    log.Printf("Received message: %s\n", string(buf))
                }
            } else {
                log.Printf("Serving %s to %s...\n", indexPage, r.RemoteAddr)
                http.ServeFile(w, r, indexPage)
            }
        })
        http.HandleFunc("/scheduled", func(w http.ResponseWriter, r *http.Request){
            if r.Method == "POST" {
            log.Printf("Received task %s scheduled at %s\n", r.Header.Get("X-Aws-Sqsd-Taskname"), r.Header.Get("X-Aws-Sqsd-Scheduled-At"))
            }
        })
        //server := &http.Server{
        //    Addr: ":https",
        //    TLSConfig: &tls.Config{
        //        GetCertificate: certManager.GetCertificate,
        //    },
        //}
        log.Printf("Listening on port %s\n\n", port)
        //http.ListenAndServe(":http", certManager.HTTPHandler(nil))
        err = http.ListenAndServe(":"+port, nil)
        //if false {
        //    log.Fatal(server.ListenAndServeTLS("", ""))
        //}  //Key and cert are coming from Let's Encrypt
}
