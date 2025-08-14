package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

const (
	white  string = "\033[0m"
	red    string = "\033[31m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	blue   string = "\033[34m"
	purple string = "\033[35m"
	cyan   string = "\033[36m"
	grey   string = "\033[90m"

	reset_cursor string = "\033[15G"

	// TODO: make logo multicolor like in fetsh
	logo string = blue +
		`  \\  \\ //  
 ==\\__\\/ // 
   //   \\//  
==//     //== 
 //\\___//     
// /\\  \\==   
  // \\  \\`
)

func userhost() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return user.Username + "@" + host
}

func model() string {
	model_name, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_version")
	if err != nil {
		log.Fatal(err)
	}
	// Newline is redundant if using fmt.Println
	return strings.Trim(string(model_name), "\n")
}

func flake() string {
	flake_version, err := os.ReadFile("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(flake_version), "\"")[21]
}

func kernel() string {
	kernel_version, err := os.ReadFile("/proc/version")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(kernel_version), " ")[2]
}

func uptime() string {
	uptime, err := os.ReadFile("/proc/uptime")
	if err != nil {
		log.Fatal(err)
	}

	seconds, err := strconv.ParseFloat(
		strings.Split(string(uptime), " ")[0],
		64)
	if err != nil {
		log.Fatal(err)
	}

	hours := strconv.FormatFloat(
		math.Floor(seconds/3600),
		'f', 0, 64)
	minutes := strconv.FormatFloat(
		math.Floor(math.Mod(seconds, 3600)/60),
		'f', 0, 64)

	if len(minutes) < 2 {
		minutes = "0" + minutes
	}

	return hours + ":" + minutes
}

func updated() string {
	cmd := exec.Command("date", "+%Y-%m-%d", "-r", "/home/nh/nix/flake.lock")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(string(out), "\n")
}

func main() {
	fmt.Print(logo + "\033[6A" + reset_cursor)
	fmt.Println(green + userhost() + white)
	fmt.Println(reset_cursor + "┌──────────────────────────────┐")
	fmt.Println(reset_cursor + red + " model " + white + model())
	fmt.Println(reset_cursor + blue + " flake " + white + flake())
	fmt.Println(reset_cursor + yellow + " kernel " + white + kernel())
	fmt.Println(reset_cursor + cyan + " uptime " + white + uptime())
	fmt.Println(reset_cursor + purple + " updated " + white + updated())
	fmt.Println(reset_cursor + "└──────────────────────────────┘")
}
