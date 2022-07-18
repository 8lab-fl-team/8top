/*

rtop - the remote system monitoring utility

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

const VERSION = "1.0"
const DEFAULT_REFRESH = 5 // default refresh interval in seconds

var currentUser *user.User
var interval int
var globalKey string
var dist string
var configFile string = "config.yml"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "8top",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntVarP(&interval, "interval", "i", DEFAULT_REFRESH, "refresh interval in seconds")
	rootCmd.Flags().StringVarP(&globalKey, "key", "k", "", "path to private key")
	rootCmd.Flags().StringVarP(&dist, "dist", "d", "", "distribution name")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.yml", "config file")
}

type Target struct {
	Host string `yaml:"host,omitempty"`
	Port int    `yaml:"port,omitempty"`
	User string `yaml:"user,omitempty"`
	Key  string `yaml:"key,omitempty"`
}

func parseYamlConfig(filename string) []Target {
	var targets []Target
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&targets)
	if err != nil {
		log.Fatal(err)
	}
	return targets
}

//----------------------------------------------------------------------------

func run() {

	configs := parseYamlConfig(configFile)
	interval := time.Second * time.Duration(interval)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	for _, config := range configs {
		host := config.Host
		port := config.Port
		username := config.User
		key := config.Key
		if key == "" {
			key = globalKey
		}

		var err error
		currentUser, err = user.Current()
		if err != nil {
			log.Print(err)
			return
		}

		// fill in still-unknown ones with defaults
		if port == 0 {
			port = 22
		}
		if len(username) == 0 {
			username = currentUser.Username
		}
		if len(key) == 0 {
			idrsap := filepath.Join(currentUser.HomeDir, ".ssh", "id_rsa")
			if _, err := os.Stat(idrsap); err == nil {
				key = idrsap
			}
		}

		addr := fmt.Sprintf("%s:%d", host, port)
		client := sshConnect(username, addr, key)

		timer := time.NewTicker(interval)
		defer timer.Stop()
		go func() {
			logStats(client)
			for range timer.C {
				logStats(client)
			}
		}()
	}
	for range sig {
		os.Exit(0)
	}

}

func getIP(addr string) string {
	return strings.Split(addr, ":")[0]
}

func logStats(client *ssh.Client) {
	stats := Stats{
		Time:   time.Now().Unix(),
		HostIP: getIP(client.Conn.RemoteAddr().String()),
	}
	getAllStats(client, &stats)

	b, err := json.Marshal(&stats)
	if err != nil {
		return
	}
	filename := fmt.Sprintf("%s.json", stats.HostIP)
	if dist != "" {
		os.MkdirAll(dist, 0664)
		filename = path.Join(dist, filename)
	}
	log.Println("fetch stats:", stats.HostIP)

	wt, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer wt.Close()
	wt.WriteString(string(b) + "\n")
}
