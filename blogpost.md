To follow along in this post, you'll need to log into https://zeit.co/ with GitHub or another Git option, and install the "now" utility on your development machine.

First, we're going to write a simple config for now. This will set up the builders that generate our serverless functions, as well as point routes to specific functions. In the root of our repo we'll add a new file called "now.json" with the following data:

now.json
```
{
  "name": "now-c2",
  "version": 2,
  "builds": [
      { "src": "*.go", "use": "@now/go" }
  ],
  "routes": [
      {
        "src": "/command",
        "dest": "/command.go"
      },
      {
        "src": "/control",
        "dest": "/control.go"
      }
  ]
}
```

This will point the "/command" route to code in "command.go", and the "/control" route to code in "control.go", as well as build our serverless functions using the "@now/go" builder. Next, we can write the skeleton code for these two files:

command.go
```
package main

import (
	"fmt"
	"net/http"
)

func Command(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Someday, I'll reply with command data!")
}
```

control.go
```
package main

import (
	"fmt"
	"net/http"
)

func Control(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Someday, I'll handle new data to issue as commands.")
}
```

We now how everything we need to deploy our code. Deploying a new version of your code is extremely simple, just type "now" in the root directory of your project. You should see something similar to the following:

```
> Deploying ~/repositories/now-c2 under milkshak3s
> Using project now-c2
> Synced 3 files [118ms]
> https://now-c2-o8u93zdz7.now.sh [2s]
> Ready! Deployed to https://now-c2.milkshak3s.now.sh [in clipboard] [22s]
```

Push this new code to your master branch and check your zeit.co dashboard, you'll notice that you have a nice little display of your current deployments:

![zeit.co dashboard]
(dashboard.png)