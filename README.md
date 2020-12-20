# TUNER

Tuner search for YouTube videos to play in MPV.

![Screenshot](https://i.imgur.com/v5GbcaL.png)

## Table of Contents

- [Features](#features)
- [Compile](#compile)
- [Screenshots](#screenshots)
- [Album Art (experimental)](#album-art-experimental)
- [MPRIS](#mpris)
- [Keybinds](#keybinds)
- [Commands](#commands)
- [Storage](#storage)
- [License](#license)

## Features

- Playlists
- Lyrics
- No "search rate limits" (Tuner doesn't use the youtube API)
- Show album art (experimental)

## Compile

Tuner is a written in Go, so you will need to install the GoLang compiler. On 
Arch Linux, you should install the `go` package.

### Dependencies

Before running Tuner you need to install 
[MPV](https://github.com/mpv-player/mpv) and
[youtube-dl](https://github.com/ytdl-org/youtube-dl/).

On Arch Linux, the dependencies packages are `mpv youtube-dl`


**Font Awesome 5 is also required to display the icons. You can also customize
the icons in the `icons/icons.go` file.**

### Build

Clone the repository: 
> $ git clone https://github.com/Pauloo27/tuner.git && cd tuner

Install:
> $ make install

## Screenshots

![Screenshot with MOTD](https://i.imgur.com/9f6XonE.png)
![Screenshot search](https://i.imgur.com/7KRlSnS.jpg)
![Screenshot playing with CAVA](https://i.imgur.com/YGhMcwK.jpg)

## Album Art (experimental)

The option to show the song Album Art is disabled by default, here's how to 
enable it:

First install [Überzug](https://github.com/seebye/ueberzug) 
(on Arch Linux, install the `ueberzug` package)...

Then open Tuner and type `/a` in the search bar and restart Tuner.

## MPRIS

In version `v0.0.1`, Tuner used mpris to comunicate with MPV. In newer version
Tuner uses libmpv instead. You can enable mpris by running the command `/mpris`.
The script file should be placed at `~/.config/mpv/scripts/mpris.so`.

If you installed the package `mpv-mpris-git` from the AUR, you need to copy the
file `/usr/share/mpv/scripts/mpris.so` to `~/.config/mpv/scripts/mpris.so`.

## Keybinds

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
- B: Save song to playlist.
- \>: Next song in playlist.
- <: Previous song in playlist.

## Commands

There are a few commands you can type in the search bar:

- `/cache` or `/c`: Toggle option to keep cache (default is false).
- `/album` or `/a`: Toggle option to show album art (default is false).
- `/mpris or `/m`: Toggle option to load mpv-mpris (default is false).
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
