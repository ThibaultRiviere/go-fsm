# go-fsm
FSM library

## Documentation

See: [go-fsm Documentation](https://godoc.org/github.com/ThibaultRiviere/go-fsm)

## Getting Started

### Installing

To start using go-fsm, install Go and run `go get`:

```sh
$ go get github.com/ThibaultRiviere/go-fsm
```

### Creating a new go-fsm

```go

import (
    "fmt"
    "github.com/ThibaultRiviere/go-fsm"
)

func main() {

    states := []string{"locked", "close", "open"}
    defaultState := "close"

    door, err := fsm.NewFsm([]states, defaultState)    
}

```

### Adding a new transition

```go
    // transition name, state needed, next state, handler
    door.AddTransition("unlock door", "locked", "close", func() {
        fmt.Println("The door have been unlock")
    })
```

### Handler a transition

```go
    currentState, err := door.HandleTransition("unlock door")
    if err != nil {
        fmt.Println("Couldn't unlock the door because the state is ", currentState)        
    } else {
        fmt.Prinln("The door state now is ", currentState)        
    }
```

If the current state is ```locked```, then the door will be unlock and the current state will be change to ```close```.
In case where the current state is not lock then the transition ```unlock door``` will failed returning an error and the current state of the door.


### Adding a new Action

```go
    // action name, state needed, handler
    door.AddAction("travers door", "open", func() {
        fmt.Prinln("Someone go through the door")    
    })
```

### Handler an action

```go
    err := door.HandleAction("travers door")
    if err != nil {
        fmt.Prinln("Impossible to travers the door"
    } else {
        fmt.Prinln("New people in the room")        
    }
```
If the current state of the door is ```open```, then it's possible to go through the door and enter in the room.

## License

MIT
