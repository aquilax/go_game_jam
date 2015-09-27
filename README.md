## Number Crusher

![Game screenshot](http://i.imgur.com/QVp6dMI.png)

**Number crusher** is a game inspired by the old time classic [Number Munchers](https://en.wikipedia.org/wiki/Munchers).
The rules are simple for each level, You must find all equations, which result to the level number. While calculating, You should watch for the Math police, which are always for unauthorized calculators. Try to get as many points as possible.

Each successfully completed level, gives you one live, and score bonus, equal to the lives you had when finishing the level. Each level has 10 correct solutions, which you need to find.

### Installation

To install the game you'll need to have [go](https://golang.org/) installed. 

```bash
$ go get github.com/aquilax/number_crusher
```

To run the game:
```bash
$ $GOPATH/bin/number_crusher
```

The game uses the [termloop](https://github.com/JoelOtter/termloop) library, which is based on [Termbox](https://github.com/nsf/termbox-go) to control the terminal. 

### Control
```
▲ - move up
▼ - move down
◀ - move left
▶ - move right
Space - select solution
Esc - exit game
```

[![asciicast](https://asciinema.org/a/5h9v1095b8y7wrggc8xa6to7f.png)](https://asciinema.org/a/5h9v1095b8y7wrggc8xa6to7f)
