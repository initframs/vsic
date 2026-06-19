# vsic
![GitHub last commit](https://img.shields.io/github/last-commit/initframs/vsic) ![GitHub License](https://img.shields.io/github/license/initframs/vsic) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/initframs/vsic) ![Static Badge](https://img.shields.io/badge/binary%20size-5.47MB-blue)                                                                               

*very* simple internet chat

this part of the project, `libvsic`, handles framing, parsing, and other useful vsic protocol stuff. if you're looking to run a vsic server, check out [vsicd](https://github.com/initframs/vsicd).

## features

- minimal text-based protocol- no client needed, it's simple enough to be used over just plain `nc`
- tiny (5.47M) and fast (starts in 0.004s, stops in 0.005s*)
- easily customizable (toml config!)
  > *note: start time measured as average over 128 starts/stops. excluding fixed sleep times in code*

## libvsic

libvsic is a collection of utils to help keep protocol behavior repeatable and to allow the protocol to be reused in other client/server implementations. 

the vsic protocol is designed to be human readable, even over the raw tcp connection. to achieve this, the general command structure is very simple:
```
COMMAND param1 param2 [etc etc] \n
```
this structure is used both by the client and the server.
> *note: while the actual vsic protocol works this way, web clients like [lwvc](https://github.com/initframs/lwvc) most likely use a different, json-framed protocol due to the need to communicate with a ws-tcp or wss-tls compatibility layer (example: [wssc](https://github.com/initframs/wssc))