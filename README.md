# vsic
*very* simple internet chat
## features
- simple text-based protocol
- tiny (5.2M) and fast (starts in 0.004s, stops in 1.005s*)
- light and customizable (toml config!)
> *note: start time measured as average over 128 starts/stops on Arch Linux (i7-4790k and 16GB DDR3)
## vsicd
a vsic server written in go, designed to be completely async and have an ultra small footprint. only 1 config file and 3 commands.
```bash
vsicd start # start the vsic daemon
vsicd stop # stop the vsic daemon
vsicd info # see stats like ram/cpu usage, connected clients, and more
```
```toml
# ~/.config/vsicd/config.toml
name = "my cool server"
motd = """
i love vsicd
irc still good though
""" # vsicd automatically trims blank lines so your config can look better :D

[moderation]
# banlist = "~/.config/vsicd/bans.txt" (in progress)
# modcmd = "python3 ~/.config/vsicd/moderation.py" (in progress)

max_conns_per_ip = 4
max_msgs_per_sec = 1
max_msg_size = 4096
max_keepalive_timeout = 120

[server.tcp]
enabled = false # both are disabled by default, you don't need to explicitly disable them

[server.tls]
enabled = true
port = 4570
tls_cert = "/etc/certs/example.com.crt"
tcp_key = "/etc/certs/example.com.key"
```
### building vsicd
vsicd is pure go, so there's no need for C interop support :)

`-ldflags="-s -w"` removes the DWARF debugging info and the symbol table. this isn't recommended for dev, since DWARF lets you know what broke, but it reduces binary size. when reporting errors/bugs, don't use these options.
```bash
CGO_ENABLED=0 go build -ldflags="-s -w" -o vsicd
```