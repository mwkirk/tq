# tq - A simple transcoding queue

## WARNING: This is a personal learning project and is not intended for production use!

This is a toy project built around the idea of a transcoding queue using FFMPEG. It's an excuse for me to noodle 
around with various Go concepts, protobufs, gRPC, etc. and generally scratch a few itches. There are many questionable 
decisions and things that are not implemented or won't be implemented. This repo is only public because I 
occasionally point folks to it. I'm not accepting PRs or looking at issues ;)

### Overview

There are three binaries:

* tq_srv - a basic queue controller
* tq_worker - worker node. Upon starting, it will contact the controller to register itself.
* tq - CLI to submit, list, and cancel jobs




