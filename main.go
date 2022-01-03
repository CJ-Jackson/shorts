package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	var addr, urlLocation string

	{
		fS := flag.NewFlagSet("shorts", flag.ContinueOnError)
		fS.StringVar(&addr, "a", ":8080", "Listening Address")
		fS.StringVar(&urlLocation, "u", "", "Locations of URLs")
		urlLocation = strings.TrimSpace(urlLocation)
		if err := fS.Parse(os.Args[1:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			fS.PrintDefaults()
		} else if urlLocation == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Must have URL location.")
				fS.PrintDefaults()
			}
			urlLocation = home + "/shorts"
		}
		urlLocation = strings.TrimRight(urlLocation, "/")
	}

	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		serve(writer, request, urlLocation)
	})))
}

func serve(writer http.ResponseWriter, request *http.Request, urlLocation string) {
	fileLocation := urlLocation + "/" + strings.Trim(request.URL.Path, "/") + ".txt"
	file, err := os.Open(fileLocation)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(writer, "Not Found")
		return
	}
	defer file.Close()

	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(writer, "Not Allowed")
		return
	}

	b, err := io.ReadAll(file)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, "Could not read file")
		return
	}

	parsedUrl, err := url.Parse(strings.TrimSpace(string(b)))
	if err != nil || parsedUrl.Scheme == "" {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, "Could not parse URL correctly")
		return
	}

	writer.Header().Set("location", parsedUrl.String())
	writer.WriteHeader(http.StatusPermanentRedirect)
}
