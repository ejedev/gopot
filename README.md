# gopot

An SSH honeypot written in Go. Very early in development.

A while back I was working with [Cowrie](https://github.com/micheloosterhof/cowrie-dev) and thought the idea was something I'd like to try and recreate in Go. My goal is to have a *mostly* simulated terminal that will log actions, downloaded files, and other fun things. Currently it does not do much.

GoPot can:

- Accept a connection via SSH and authenticate (currently set to allow any user/pass combination)
- Present a terminal to the user.
- Log logins as well as actions performed by the user (which are currently being echoed back to them. This is not production ready.)

GoPot currently can not:

- Do anything else.