# go-isolation-as-container

A simple code in golang isolating UTS, PID and FileSystem.
Presentation [link](https://docs.google.com/presentation/d/18N1ndb_RIC4sa54WaEb49bI-3sVqjxcH2O2xRzrFXMQ/edit?usp=sharing) 


## How to run

If you are using LINUX system you can run 
```
# just list current directory
go run main.go ls -la
# create a new bash from here
go run main.go bash
```

You can run `ps aux` or `top` to see running processes

----------

If you are running a non LINUX system, you can create a container simulating a LINUX kernel to do this. You can run this command
```
docker run -it -v $(pwd):/app --workdir /app -v /var/run/docker.sock:/var/run/docker.sock --name go_bemug --privileged golang
```

You need `--privileged` flag to access all capabilities on container "as you are a host"

## Future

Implement a full isolation of file system and usage of images

## Thanks

This project is based on Wellington Silva press on TDC SP 2018. @wsilva presentation is based on @lizrice and @julz code/presentation.

Thank you by allowing me to do this presentation on community.

- Wellington Silva - @wsilva
- Liz Rice - @lizrice
- Julian Friedman - @julz

