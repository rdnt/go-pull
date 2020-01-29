# go-pull
Compiles and starts a go program, automatically pulls updates and restarts its process. I was bored, okay?

*Note:* Currently only working on Windows. If you want Unix support, you are free to implement it :D

## Setup

`go build`

## Usage

`go-pull <repo-path>`

You can provide a relative or absolute path to a git repository, and if there is a `main.go` file inside, it will automatically compile and launch it, and update and restart when new commits are pushed to the repository.

*Note:* There must be a remote called `origin` set up and by default go-pull tracks the `master` branch. Additionally, your working directory **must not** have any unstaged stages.


## Contributing
You are free and actively encouraged to contribute to this project by either contributing code, creating issues, reporting bugs, highlighting vulnerabilities, proposing improvements or helping maintain the documentation.

If you would like to submit code changes, create a new branch from the *master* branch with the name of the feature you are implementing  and submit a pull request to the *master* branch after you make your changes. Click [here](https://gist.github.com/Chaser324/ce0505fbed06b947d962#doing-your-work) for a how-to guide.

In case you want to submit a bug report, please add as many details as possible regarding how the error occurred and include the steps required to reproduce it, if that is at all possible. It will help a lot in testing, finding the cause and implementing fixes.


## Changelogs
Changelogs for each and every release can be found [here](https://github.com/SHT/go-pull/releases).


## Copyright
Any reproductions of this project *must* include a link to this repository and the following copyright notice, along with the project's license.

© 2020 Tasos Papalyras - All Rights Reserved
