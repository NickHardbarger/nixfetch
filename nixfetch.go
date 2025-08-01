package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
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

func get_userhost() string {
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

func get_model() string {
	model_name, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_version")
	if err != nil {
		log.Fatal(err)
	}
	// Newline is redundant if using fmt.Println
	return strings.Trim(string(model_name), "\n")
}

func get_flake() string {
	flake_version, err := os.ReadFile("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(flake_version), "\"")[21]
}

func get_kernel() string {
	kernel_version, err := os.ReadFile("/proc/version")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(kernel_version), " ")[2]
}

func get_uptime() string {
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
	const six_months int64 = 15778463
	// NOTE: date must be updated manually
	start_time := time.Date(2025, 8, 1,
		0, 0, 0, 0, time.UTC).Unix()
	current_time := time.Now().Unix()
	time_passage := current_time - start_time
	days_passed := time_passage / 86400
	if time_passage > six_months {
		return red + strconv.Itoa(int(days_passed)) + " days ago!!"
	} else {
		return strconv.Itoa(int(days_passed)) + " days ago"
	}
}

func main() {
	fmt.Print(logo + "\033[6A" + reset_cursor)
	fmt.Println(green + get_userhost() + white)
	fmt.Println(reset_cursor + "┌──────────────────────────────┐")
	fmt.Println(reset_cursor + red + " model " + white + get_model())
	fmt.Println(reset_cursor + blue + " flake " + white + get_flake())
	fmt.Println(reset_cursor + yellow + " kernel " + white + get_kernel())
	fmt.Println(reset_cursor + cyan + " uptime " + white + get_uptime())
	fmt.Println(reset_cursor + purple + " updated " + white + updated())
	fmt.Println(reset_cursor + "└──────────────────────────────┘")
}
