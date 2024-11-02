
package main

import (
  "os/exec"
)


func RunNav() {
  exec.Command("pkill", "ffmpeg").Run();
  exec.Command("x-www-browser", "http://localhost:8070").Run();
  defer exec.Command("pkill", "ffmpeg").Run();

}

func main() {
  go RunNav();
  exec.Command("pkill","tomp3").Run();
  exec.Command("./tomp3").Run();
}
