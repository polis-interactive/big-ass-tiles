
## Pi Install

- Configure Pi
  - 
- Install libraries
  - vim git tmux libx11-dev xvfb libgl1-mesa-dev cmake xorg-dev
- Install Go
  - https://www.jeremymorgan.com/tutorials/raspberry-pi/install-go-raspberry-pi/
  - Install armv6l
- Install github.com/jgarff/rpi_ws281x
- Add go path to sudo
  - visudo, add secure_path="...:/usr/local/go/bin"
- Build
  - pull repo, use common key
  - grab only rpi-ws281x, periph
  - sudo go build ./cmd/runApplication/main.go
- Make Service
  - https://superuser.com/questions/544399/how-do-you-make-a-systemd-service-as-the-last-service-on-boot
  - Remove nymea from target graphical to target big_ass_tiles, resymlink