# TUNER

Tuner searches and plays songs from YouTube and SoundCloud inside your terminal.

[![CI status](https://ci.notagovernment.agency/api/v1alpha/badges/5784c967-246f-4077-9f4a-12b5aa4f14a7?branch=master)](https://ci.notagovernment.agency/user/Pauloo27/projects/tuner.proj)


_If you want a GUI application with more features, take a look at
[Neptune](https://github.com/Pauloo27/neptune)._

![Screenshot](https://i.imgur.com/REvt9Kw.png)

## Table of Contents

- [Features](#features)
- [Usage](#usage)
- [Quick play with dmenu](#quick-play-with-dmenu)
- [Installing](#installing)
- [Compile](#compile)
- [Screenshots](#screenshots)
- [Album Art (experimental)](#album-art-experimental)
- [MPRIS](#mpris)
- [Discord](#discord)
- [Keybinds](#keybinds)
- [Commands](#commands)
- [Storage](#storage)
- [License](#license)

## Features

- Playlists
- Lyrics
- Discord Rich Presence
- No "search rate limits" (Tuner doesn't use the youtube API)
- Show album art (experimental)
- SoundCloud search support
- "I'm Feeling Lucky" (play the first result when `!` is used as a prefix for the
search query)

## Usage

Using Tuner is simple. Just install it and then run `tuner` inside your terminal.
That will launch a interactive player. If that's not what you want, here are some
alternative modes:

- `tuner play <search term>` (or `tuner p`): plays the first result found for 
the search term (a terminal is required).
- `tuner simple-play <search term>` (or `tuner sp`): plays the first result 
found for the search term but without the interactive player (does not require a 
terminal).

### Quick play with dmenu

You can create a script in `/usr/bin/play` (or any other folder in your path)
with the following content:
```bash
#!/bin/bash

search_query="$@"

# when the search query is not defined
if [ -z "$1" ]; then
  search_query=$(echo "stop" | dmenu -p "Play: ")
fi

# when the search query is stop
if [ "$search_query" = "stop" ]; then
  killall tuner -w 
  exit 0
fi

# when the search query is ok to be played
tuner sp "$search_query"
```

With that, you can use dmenu to quickly play songs, by:
- Creating a bind to `play` (that will open a dmenu prompt asking for the search
 query).
- Running `play <song name>` (where `song name` is not required) in dmenu_run
(usually `super+d`).

## Installing

### Arch Linux

Install from the AUR (`yay -S go-tuner-git`).

### Debian (also Ubuntu, Linux Mint and PopOS)

Use this install script:

> $ sh -c "$(curl -fsSL https://raw.githubusercontent.com/Pauloo27/tuner/master/install-debian.sh)"

## Compile

Tuner is a written in Go, so you will need to install the GoLang compiler. On 
Arch Linux, you should install the `go` package.

### Dependencies

Before running Tuner you need to install 
[MPV](https://github.com/mpv-player/mpv) and
[youtube-dl](https://github.com/ytdl-org/youtube-dl/).

On Arch Linux, the dependencies packages are `mpv youtube-dl`

On Debian based (Ubuntu, PopOs, Linux Mint, etc), install `libmpv-dev` via APT
and `youtube-dl` via PIP (`sudo -H pip3 install --upgrade youtube-dl`).

**Font Awesome 5 is also required to display the icons. You can also customize
the icons in the `icons/icons.go` file.**

### Build

Clone the repository: 
> $ git clone https://github.com/Pauloo27/tuner.git && cd tuner

Install:
> $ make install

## Screenshots

**Tuner + pfetch:**


![Screenshot with MOTD](https://i.imgur.com/9f6XonE.png)


**Tuner search menu:**


![Screenshot search](https://i.imgur.com/glnKkUH.png)

**Tuner + Discord:**

![Screenshot of Discord Rich Presence](https://i.imgur.com/zvHtld8.png)

**Tuner + cava:**


![Screenshot playing with CAVA](https://i.imgur.com/K7Qtozn.png)


**Tuner + lolcat:**


![Screenshot playing with CAVA](https://i.imgur.com/Mb532M7.png)


## Album Art (experimental)

The option to show the song Album Art is disabled by default, here's how to 
enable it:

First install [Überzug](https://github.com/seebye/ueberzug) 
(on Arch Linux, install the `ueberzug` package)...

Then open Tuner and type `/a` in the search bar and restart your Terminal and 
Tuner.

## MPRIS

To enable MPRIS, first install [mpv-mpris](https://github.com/hoyon/mpv-mpris)
(it needs to be placed at `~/.config/mpv/scripts/mpris.so` or 
`/etc/mpv/scripts/mpris.so`).

_if you are using Arch Linux and installed Tuner using the `go-tuner-git` AUR package, it's already installed._

After that, use `/mpris` command and restart tuner.

## Discord

To enable or disable the integration with Discord, type `/discord` in the search bar.

## Keybinds

**If keybinds aren't working, try running Tuner with `TERM="xterm" tuner`.
If it doesn't work, restart you terminal and open a issue on GitHub with 
the output of `echo $TERM`. If it does work, you can also open the issue or set
a alias.**

_TIP: You can see the keybinds inside Tuner by pressing `?`._

- Arrow Left: Seek 5 seconds back.
- Arrow Right: Seek 5 seconds.
- Ctrl C: Stop the player.
- Space: Play/Pause song.
- Arrow Down: Decrease the volume.
- Arrow Up: Increase the volume.
- ?: Toggle keybind list.
- L: Toggle loop.
- P: Toggle lyric.
- W: Scroll lyric up.
- S: Scroll lyric down.
- U: Show video URL.
- B: Save song to playlist and edit current playlist.
- R: Shuffle playlist.
- \>: Next song in playlist.
- <: Previous song in playlist.

## Commands

There are a few commands you can type in the search bar:

- `/discord` or `/dc`: Toggle option to show Tuner in Discord. 
- `/cache` or `/c`: Toggle option to keep cache (default is false).
- `/album` or `/a`: Toggle option to show album art (default is false).
- `/mpris` or `/m`: Toggle option to load mpv-mpris (default is false).
- `/help` or `/h`: List all commands.

## Storage

Tuner data (cache and config) is stored at `~/.cache/tuner`. Tuner doesn't keep
the songs downloaded, the only cached data is the album art (if the feature is
enabled) and the info of playlist entries.

## License

<img src="https://i.imgur.com/AuQQfiB.png" alt="GPL Logo" height="100px" />

This project is licensed under [GNU General Public License v2.0](./LICENSE).

This program is free software; you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.
