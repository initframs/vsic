# vsic
*very* simple internet chat
## features
- simple text-based protocol
- tiny and fast
- light and customizable
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
motd = "~/.config/vsicd/motd.txt"

[moderation]
banlist = "~/.config/vsicd/bans.txt"
modcmd = "python3 ~/.config/vsicd/moderation.py"

max_conns_per_ip = 4
max_msgs_per_sec = 1
max_msg_size = 4096
max_keepalive_timeout = 120

[server.tcp]
enabled = false # both are disabled by default

[server.tls]
enabled = true
port = 4570
tls_cert = "/etc/certs/example.com.crt"
tcp_key = "/etc/certs/example.com.key"
```
## vs2c
a simple and light cli-based vsic client, also written in go. simple file-based configuration
```toml
# ~/.config/vs2c/config.toml

[servers.coolserver]
name = "cool server"
port = 4570
tls = true
nick = "initframs"

[servers.offtopic]
name = "offtopic lounge"
port = 4572
tls = false
nick = "ilovevsic2642"
```