# TUNER

Tuner search for YouTube videos to play in MPV.

![Screenshot](https://i.imgur.com/v5GbcaL.png)

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
[MPV](https://github.com/mpv-player/mpv),
[youtube-dl](https://github.com/ytdl-org/youtube-dl/) and 
[mpv-mpris](https://github.com/hoyon/mpv-mpris).

_On Arch Linux, the dependencies packages are `mpv youtube-dl mpv-mpris-git`
(mpv-mpris-git come from AUR)._

### Build

Clone the repository: 
> $ git clone https://github.com/Pauloo27/tuner.git && cd tuner

Install:
> $ make install

## Screenshots
![Screenshot search](https://i.imgur.com/7KRlSnS.jpg)
![Screenshot playing with CAVA](https://i.imgur.com/YGhMcwK.jpg)

## Album Art (experimental)

The option to show the song Album Art is disabled by default, here's how to 
enable it:

First install [Überzug](https://github.com/seebye/ueberzug) 
(on Arch Linux, install the `ueberzug` package)...

Then open Tuner and type `/a` in the search bar and restart Tuner.

## Keybinds

_TIP: You can see the keybinds inside Tuner by pressing `?`._

Keybinds:
- Arrow Left: Seek 5 seconds back
- Arrow Right: Seek 5 seconds
- Ctrl C: Stop the player
- Space: Play/Pause song
- Arrow Down: Decrease the volume
- Arrow Up: Increase the volume
- ?: Toggle keybind list
- L: Toggle loop
- P: Toggle lyric
- W: Scroll lyric up
- S: Scroll lyric down
- U: Show video URL
- B: Save song to playlist
- \>: Next song in playlist
- <: Previous song in playlist



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
