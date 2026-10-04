# tagless-bonsai

**A website making a bonsai tree without any HTML tags!**

![Hackatime tracked time](https://hackatime.hackclub.com/api/v1/badge/U0APWD8KNPP/Ankumeah/tagless-bonsai)

![Hackatime language stats](https://github-readme-stats.hackclub.dev/api/wakatime?username=72035&api_domain=hackatime.hackclub.com&theme=catppuccin_mocha&custom_title=Hackatime+Stats&layout=compact&cache_seconds=0&langs_count=8)

> Quick question. If anyone has any idea what the "other" language
up there is, please do tell me on a github issue or something

> [!WARNING]
> I suck at spelling so be prepaired to burn
your eyes after you see who knows what horrors lie below

This is a small website I made for
[Hackclub tagless](https://tagless.hackclub.com)

## What does it do?

Makes a cool bonsai tree!

## How does it work without HTML tags?

Well that's kind of a lie (dissapointing, ik).
It does use some tags like `<script>` and `<head>` but other then those
basic things, we hand it off to javascript to load a wasm (Web Assembly)
file made in go, where we use the `syscall/js` library to execute
some javascript to interface with the html and do all the magic

## Could it have been done with only HTML and javascript?

Undoubtably!

## Then why did I not use just HTML and javascript?

Because I though it would be cool to use wasm through go

## Why is the bonsai so ugly?

Cus I suck at maths
