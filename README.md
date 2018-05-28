# Games
A compilation of various games written in [GO](https://golang.org/).

- [Chess](chess/)

## Goal
To better my own personal knowledge of Adversarial Search and other AI techniques that can be applied in a game fashion.

## References
A lot of my work in Adversarial Search is based off of Chapter 5 of [Artificial Intelligence: A Modern Approach](http://aima.cs.berkeley.edu/) (Third Edition) by [Stuart Russell](http://www.cs.berkeley.edu/~russell/) and [Peter Norving](http://www.norvig.com/).  

## Interface
Currently, the only available interface is via the command line, but with additional efforts, I hope to allow a web based interface as well for each individual game.

## Install
```
go get github.com/bign8/games
games
```

*Eventually, I will be releasing pre-built binaries that can be directly downloaded without the need of installing GO.*

## Contributions
Bug reports, suggestions and code contributions are welcome.  Just be sure to follow best go and git practices along with only committing code that is covered with tests and benchmarks.  Additionally, this is a GO repository, lets keep it that way!

- Resources
  - [Effective GO](https://golang.org/doc/effective_go.html)
  - [GO Practices](https://peter.bourgon.org/go-best-practices-2016/)
  - [Commit Messages](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html)
  - [GO Proverbs](http://go-proverbs.github.io/) [video](https://www.youtube.com/watch?v=PAAkCSZUG1c)
- Additional Rules
  - Prefer the `%q` and `%+v` directives over the `%v` directive in errors and logging messages.  Justification: `%v` can result in ambiguous error messages when strings are used because the value is not encapsulated.
  - Be careful when implementing the magic `String()` method.  If one accidentally uses `Printf` the wrong way, the application could crash due to infinite recursion.
  - Don't use [dot imports](http://stackoverflow.com/a/6478990); They pollute the namespace and confuse language tooling.
  - If you know you are leaving a feature un-implemented, please make a comment prefaced with `TODO:`.  This enables easy grep-ing for tasks that remain to be completed.

### Development Suggestions
I develop GO in an [Atom](https://atom.io/) editor in conjunction with the [Delve](https://github.com/derekparker/delve) Debugger and the following plugins.

- [go-debug](https://atom.io/packages/go-debug)
- [go-config](https://atom.io/packages/go-config)
- [go-plus](https://atom.io/packages/go-plus)

### Running the service

*Aside
I ran into a few problems getting Docker to start on AWS EC2 instances.
Needed to add a config that sounds terrible when starting `dockerd`: `--storage-opt dm.override_udev_sync_check=true`
Thanks to [thomas15v](https://github.com/docker-library/docker/issues/19#issuecomment-298835023) for figuring out a short term solution.
(Yes, I know this is terrible, but these containers are currently ephemeral until I get stateless persistence, scaling and reconnecting working)*

```
docker run -d --name games --publish 4000:4000 bign8/games:latest
docker service create --name games --publish 4000:4000 bign8/games:latest
```
