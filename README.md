
A Distributed Key-Value Server
=================================

[![Go Report Card](https://goreportcard.com/badge/github.com/derekargueta/go-kv-server)](https://goreportcard.com/report/github.com/derekargueta/go-kv-server)

Design
=======

This implements a gossip-style protocol in which when one node recieves new
information, it is broadcasted to its neighbors.

Failure Detection
====================
Nodes send a heartbeat to their peers every 10 seconds. If a node does not
hear from a peer within 30 seconds, it is assumed that that peer has crashed.

