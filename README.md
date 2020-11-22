# Webhook Service

## Build
```
go build -o ./bin/cli ./cli/main.go 
go build -o ./bin/server ./server/main.go
```

## Usage
### Cli
You can use `./cli command` to run specified command.

Or you can use `./cli` to get into the cli's shell and execute commands.

#### Commands
1. n / new: create new webhook.
2. d / delete id: delete specified webhook.
3. m / modify id: modify existed webhook.
4. e / execute id: execute specified webhook.
5. l / list: list all webhooks.
6. s / search keyword: search webhooks by a keyword in name or description.
7. p / print id: print detail information of specified webhooks.
8. h / help: print help information. 
9. q / quit: quit cli shell.

### Server
Start server by run `./server` or `./server port`, the default port is 8080.

Start server with pm2: `pm2 start ./server --name webhook-service -- 8080`.