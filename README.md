# logmon

A simple console program that monitors HTTP traffic on your machine.

## Setup & Usage

1/ Install go

2/ Run ```go get github.com/gabsn/logmon```

3/ Use it to monitor your log files:
```bash
logmon [logPath] [threshold]
```
4/ You can go to the [loggen](loggen) directory and type ```make run``` to generate a log.txt file

## Configuration

You can edit your preferences concerning the period of refreshment in the [config/global.go](config/global.go) file.

## Testing

Run ```go test ./...``` in the root directory to run all test units.
You will find the test of the alert logic in the [models](models) directory.

## Project Architecture

![logmon architecture](docs/logmon_architecture.png)

I used two goroutines to manage the two main concurent tasks:

1/ The first one represented by the [__feeder package__](feeder) is to consume the log file, parse it into a Hit struct and then send it to the circular buffer.

2/ The second one represented by the [__controller package__](controller) is to monitor and alert the user if the number of hits becomes greater than the threshold.

Finally, I used a [__circular buffer__](models/circularbuffer.go) to store information about hits. Since the two goroutines are concurrently talking to the circular buffer I used a mutex to protect it against race conditions.

## Why go ?

The programming languages I enjoy coding with at the moment are Go, Python and C++.

Though I love Python for one-shot apps, it does not have built-in concurrency and has rather poor performances in comparison with Go and C++.

Concerning C++, I wanted to avoid having to deal with POSIX threads and memory management...

That's why I chose Go: for its lightweight easy-to-use goroutines, its efficient garbage collector and also because I enjoy working with it.

## Possible Improvements

- Build an external package in charge of parsing all type of logs (I searched for one but I decided to build the parser myself because I didn't find it. That's why the parser is not very robust...)

- Build a graphical visualization tool, because it's not very clear to print all the stats directly in the console

- Add some missing features built on top of fields provided by the W3C format (for example, a feature to detect the client who sent the most of requests by using the c-ip field)

- Create new alerts based on status code for example


