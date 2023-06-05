package config

import "flag"

var (
	Port     int    // webserver port number
	Debug    bool   // enable debug logging
	MediaDir string // directory to discover media files from
)

func init() {
	flag.IntVar(&Port, "port", 8080, "Port to serve on.")
	flag.BoolVar(&Debug, "debug", false, "Enable debug logging")
	flag.StringVar(&MediaDir, "media", "", "The directory to scan for media files")

	flag.Parse()
}
