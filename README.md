## Mole

Mole is an experiment as I'm learning golang.  I wanted to write something with
completely separate frontend/backend.  So a websocket golang API server pulling
a Twitter feed was something fairly simple to get a steady stream of data.

Running Mole is fairly straight forward, the `website` directory is self contained
frontend.  Its all static HTML and can be placed anywhere you have a webserver
or you can run a single command python webserver with `make webserver` which under
the hood is running:

```
cd website && python -m SimpleHTTPServer 8080
```

API Server is written in Go, designed to stream a few topics off twitter and publish
to a channel, which is then pushed to all websocket clients.  All clients receive
the same content, relative to when they connected to the socket.  Depending on where
you start this server, you will need to change the websocket URL in the index.html
as part of your static assets.  API Server requires 5 Environment Variables and a
6th if you want to change it from its default port 5000. I've provided an example
docker-compose.yml file to help with launching it.  I also have an example .env
file that can be used to supply the environment variables as well, it will read
both.

And you would be able to run it with the command `docker-compose up`


I'll Update this README when I have a demo of it running somewhere...
