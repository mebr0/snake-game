# snake-game

Console snake game as test task for SkyCoin made within 3 hours

---
After feedback following was fixed:

- Using `fmt.Scanln` rather than `bufio.Scanner`
- Returning proper error rather than default value
- Removing loop label
---

## How to play

Run

```shell
make build
```

for building project. Then run 

```shell
./main {height} {width}
```

for starting the game. Height and width are integers of the game frame. 
During game enter every move as single letter:

* `l` - left
* `r` - right
* `u` - up
* `d` - down

Height and width as well as round and score count printed at the top. 
Symbols within game:

* `q` - snake head
* `o` - snake body
* `@` - food
* `#` - game frame

## Commands

`mockery --name={interface} --filename=mock.go` - generate mock classes _(in package)_

`make lint` - lint with `golangci-lint`

`make format` - format with `goimport`

`make test` - run unit tests

`make cover` - run unit tests and show coverage report

`make build` - build project
