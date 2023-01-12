![gopher](gopher.png)

# Gophercraft

[![Go Reference](https://pkg.go.dev/badge/github.com/Gophercraft/core.svg)](https://pkg.go.dev/github.com/Gophercraft/core)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Chat on discord](https://img.shields.io/discord/556039662997733391.svg)](https://discord.gg/xPtuEjt)

The Gophercraft project provides 100% Go libraries and programs for research and experimentation with MMORPG software.

**⚠️ WARNING: Gophercraft is experimental and prone to all sorts of game-ruining exploits. Don't use it yet for anything serious yet.**

Read [here](https://github.com/Gophercraft/core/wiki) for more information 

## Server setup/installation (Linux)

```bash
# Install packages
sudo apt install git golang mariadb-server

git clone https://github.com/Gophercraft/core gophercraft; cd gophercraft

# Install Gophercraft Core
go install github.com/Gophercraft/core/cmd/gophercraft_wiz
go install github.com/Gophercraft/core/cmd/gophercraft_home
go install github.com/Gophercraft/core/cmd/gophercraft_world

# Generate configuration files and create databases with the Gophercraft Wizard
gophercraft_wiz

# You can edit your configurations in ~/.local/Gophercraft/Home/Home.txt
# and in ~/.local/Gophercraft/<worldserver folder>/World.txt
gophercraft_home

# in a different command prompt
# launch worldserver
gophercraft_world <Name of server>
```

## home server

The center of a Gophercraft network is the Home server.

The home server acts as a central authority in a Gophercraft network, similar to the MaNGOS "realmd". 

Upon creating new World config, an ECDSA keypair is generated. The Home server associates this public key with a new Realm ID.

Now, the world server can post its info back to Home, and that will update the realm list. Hooray!

> Tip: Set OpenRegistration to true in your `Home.txt` if you wish to allow anybody to register a world server. 

## world server (or realm server)

The world server contains the in-game experience. Players connect to it with the IP address posted by the registered world server.

It aims to be highly extensible through the use of datapacks and Go plugins.

