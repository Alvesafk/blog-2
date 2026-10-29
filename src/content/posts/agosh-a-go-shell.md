---
title: "Agosh, a go shell"
description: "Hello! In the last few days i got kinda of hooked on the Go Programming Language, so, to understand it better (also for funsies) i decided to do a small project, that project is agosh."
tags: ["Programming", "Go", "Shell"]
createdAt: "May 16 2026"
---

# Agosh, a go shell

Hello! In the last few days i got kinda of hooked on the Go Programming Language, so, to understand it better (also for funsies) i decided to do a small project, that project is agosh.

Originaly i called gosh but obviously in the world of programming you don't really create nothing, so, a go shell project called gosh already exists, =( so i changed to a go shell, but whatever!

The thing is Go is an amazing language! I'm kinda baffled actually, the speed and ease that you can develop stuff, so as a learning project i did agosh, honestly never thought on doing a shell, and i really thought it would be harder but shells on their essence are simple, they are an abstraction on functions of the kernel.
Interesting stuff

Working on a shell actually makes you learn a bit on how:

- Terminal works
- Syscalls
- How bash/zsh/fish or whatever actually are amazing software

You really get to admire this almost ancient software that are shells, one thing that really baffled me actually is the fact that the command cd, one of the most, if not, the most common and used command on shells, it's an alias, this actualy blew my mind, i never really thought about some of this commands but the fact is, cd is not a binary it's a system call, on the case of Unix like systems cd is an aliases to the chdir syscall. And it doesn't do just that! It can do some other stuff, it's not an alias only! Let me show how this command is working on the my latest commit on agosh:

```go
switch args[0] {
    case "cd":

        if len(args) < 2 {
            last_working_directory = getWorkingDirectory()
            return os.Chdir(getUserHomeDir())
        }

        if args[1] == "-" {
            fmt.Println(last_working_directory)
            wd_to_go := last_working_directory
            last_working_directory = getWorkingDirectory()
            return os.Chdir(wd_to_go)
        }

        last_working_directory = getWorkingDirectory()

        return os.Chdir(args[1])
...
```

On this case the variable args is the arguments passed when you do a command, it gets stored on a slice type, index 0 is the command or binary (cd, nvim, ls, mkdir), so for implementing aliases to syscalls (like cd or exit) you have to check this values, with a switch statement this is easily done! In this case the switch get's cd and then checks for: 1. if the length of args is smaller than 2
This means the user typed only cd, the normal behaviour that the other shells do is take you to your home directory, so on my code i first change the last_working_directory var to the directory that the user is now, and then use the chdir syscall using the os standard lib of Go. 2. if the second argument is -
In this case when the second argument of cd is the minus sign you get back to the last directory that you was, so chdir is called with the last_working_directory variable, actually with the wd_to_go, for one reason only, if you use cd - more than one time you see that you're cycling between 2 dirs, so to replicate this behaviour i use other variable, because if not, the second cd - wouldn't get you nowhere! 3. if none of this is true
Then we just call chdir with the second argument, your path! Maybe you're wondering about the ~ argument, this is a aliases to your home dir, and i actually treat this when i'm cleaning the user input, let me show you how it looks.

```go
if len(args) >= 2 {
    args[1] = strings.Replace(args[1], "~", getUserHomeDir(), 1)
}
```

It's just this, the Golang stdlib is actually amazing and has a lot of amazing functions and methods, let me breakdown this code for you, if the amount of arguments it's more or equal to 2, we get the second argument and use the strings.Replace() method to replace ~ with your home dir. Actually now that i'm reading this code again, it is kinda shit, it would not really work with multiple home aliases, im going to refactor this!

Trying to remake the functionalites of the terminal i hit a wall when i got the Arrow problem.

The thing is, usually when we are making CLI applications (which is basically my shell) your input on the terminal is already "clean" when it arrives on the code, so the first thing that got me was trying to get arrows inputs. When a user input a argument it's only "send" to the code after the user clicks Enter ("\n"), and thats a problem with arrows, because the functionalty of arrows are triggered on click and not after a Enter keystroke, so,to actually get this value and do it's function without printing shit to the terminal i had to rewrite the input reading function, to basically hijack the input.

```go
func readInput() (string, error) {
    var (
        buf      []rune
        result   string
        inputErr error
    )

    data_raw, err := os.ReadFile(history_absolute_path)
    if err != nil {
        log.Fatal("Could not read from the history file")
        return "", err
    }

    data := reverseStringArray(strings.Split(string(data_raw), "\n"))

    keyboard.Listen(func(key keys.Key) (stop bool, err error) {
        switch key.Code {
        case keys.CtrlC, keys.CtrlD:
            inputErr = io.EOF
            buf_position = 0
            return true, nil

        case keys.Enter:
            fmt.Print("\r\n")
            result = string(buf) + "\n"
            buf_position = 0
            return true, nil

        case keys.Backspace:
            if len(buf) > 0 || buf_position > 0 {
                buf = buf[:len(buf)-1]
                fmt.Printf("\b \b")
                buf_position--
            }

        case keys.Right:
            if buf_position <= len(buf) {
                fmt.Printf("\033[%dC", 1)
                buf_position++
            }

        case keys.Left:
            if len(buf) > 0 || buf_position > 0 {
                fmt.Printf("\033[%dD", 1)
                buf_position--
            }

        case keys.Up:
            clearStdin()
            command_index++
            buf = []rune(data[command_index])
            fmt.Print(data[command_index])

        case keys.Down:
            clearStdin()
            if command_index > 0 {
                command_index--
                buf = []rune(data[command_index])
                fmt.Print(data[command_index])
            }

        default:
            if key.Code == keys.RuneKey || key.Code == keys.Space {
                buf = append(buf, key.Runes[0])
                fmt.Print(string(key.Runes[0]))
                buf_position++
            }
        }

        return false, nil
    })

    return result, inputErr
}
```

Previous to that problem this shell used to be full on stdlib and nothing else, but to get the raw input of the terminal i had to use a package called atomicgo.dev.

So this function is not finished actually, it's kinda bad, the left and right arrows + backspace functions don't work in a really good way, they are very buggy, but i will work on this!

So that's it, that's agosh, it's very simple and buggy but it was my first Golang project and i had a lot of fun so far doing it, im going to still fix a lot of stuff and turn this into a better project.

Thanks for the attention, sincerely, Alvesafk.
