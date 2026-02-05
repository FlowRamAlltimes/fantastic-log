package main // ok guys i have stared to create my own sowtware fo saving your pc if somebody wants to DDoS you with logs

import (
	"flag"
	"fmt"     // XD
	"log"     // log
	"os/exec" // working with cmd
	"syscall" // system calls(files, etc.)
	"time"    // and time also
)

func main() { // i hope it will profitly from experience side
	path := "/"

	maxPercentage := flag.Float64("p", 75, "need for something great") // sorry i'll make in tool view
	flag.Parse()
	var Size syscall.Statfs_t
	for {
		err := syscall.Statfs(path, &Size)
		if err != nil {
			fmt.Println("Error while size inicialized!")
			log.Fatal(err)
		}

		all := Size.Blocks * uint64(Size.Bsize) // counting all sizee of disk
		free := Size.Bavail * uint64(Size.Bsize) // not use bsize as bavail can see percentage earlier
		used := all - free

		usedPercentage := float64(used) / float64(all) * 100

		if usedPercentage > maxPercentage { // most popular log files
			exec.Command("truncate", "-s", "0", "/var/log/syslog").Run() // 1
			exec.Command("truncate", "-s", "0", "/var/log/auth.log").Run() // 2
			exec.Command("truncate", "-s", "0", "/var/log/kern.log").Run() // 3
			exec.Command("journalctl", "--vacuum-size=100M").Run()
			exec.Command("logrotate", "-f", "/etc/logrotate.d/rsyslog").Run()
			time.Sleep(5 * time.Second)
		} else {
			time.Sleep(300 * time.Second)
		}
}
