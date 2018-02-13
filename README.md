
# DevOps Assignment
#### CI/CD flow for AWLESS (open source)
![flow](https://i.imgur.com/gijyL8w.png)

#### Status: [![Build Status](https://api.travis-ci.org/talkimhi/awless.svg?branch=master)](https://travis-ci.org/talkimhi/awless)

# About AWLESS
`awless` is a powerful, innovative and small surface command line interface (CLI) to manage Amazon Web Services.


`awless` stands out by providing the following features:

- small and hierarchical set of commands
- create and revert fully-fledged infrastructures through a new simple and powerful templating language
- local log of all your cloud modifications done through `awless`
- exploration of your cloud infrastructure and resources relations, **even offline** using a local graph storage
- greater output's readability with numerous machine and human friendly formats
- ensure smart defaults & security best practices
- connect easily using awless' smart SSH to your private & public instances

### Demo
<p align="center">
  <a href="https://raw.githubusercontent.com/wiki/wallix/awless/apng/awless-demo.png"><img src="https://raw.githubusercontent.com/wiki/wallix/awless/apng/awless-demo.png" alt="video of a few awless commands"></a>
<br/>
<em>Note that the video above is in <a href="https://en.wikipedia.org/wiki/APNG">APNG</a> and requires a recent browser.</em>
</p>

#### Awards : [Top 50 Developer Tools of 2017](https://stackshare.io/posts/top-developer-tools-2017), [InfoWorld Bossie Awards 2017](https://www.infoworld.com/article/3227920/cloud-computing/bossie-awards-2017-the-best-cloud-computing-software.html#slide12)
#### More information:
Take the tour at [Getting Started (wiki)](https://github.com/wallix/awless/wiki/Getting-Started).
Or read the [introductory blog post about awless](https://medium.com/@hbbio/awless-io-a-mighty-cli-for-aws-a0d48bdb59a4).

More articles:

   - [Simplified user management for AWS](https://medium.com/@awlessCLI/simplified-user-management-for-aws-6f828ccab387)
   - [InfoWorld: Production-grade deployment of WordPress](https://www.infoworld.com/article/3230547/cloud-computing/awless-tutorial-try-a-smarter-cli-for-aws.html)
   - [Easy create & tear down of a multi-AZ CockroachDB cluster](https://github.com/talkimhi/awless-templates/tree/master/cockroachdb)


# DevOps Assignment - About
#### My assignment objectives were:
1. Choose and "containerize" an existing known open source project which doesn’t have a Docker repository.
2. build a CI/CD Flow for the project (must have some level of unit testing).

#### Prerequisites:
1. Github Account
2. Travis-CI Account
3. Ansible
4. Docker
5. AWS Account


####
# Step-By-Step

### #Install Awless
`awless` is written in Golang, so you need to install go in order to use `awless`.
 You can install `awless` using the binaries, or by installing Go and using the `go get` command.

Choose one of the following options:
1. With curl (macOS/Linux), run:  `curl https://raw.githubusercontent.com/talkimhi/awless/master/getawless.sh | bash`
2. Download the latest `awless` binaries (Windows/Linux/macOS) [from here](https://github.com/talkimhi/awless/releases/latest)
2. If you have Golang already installed, install from the source with: `go get -u github.com/talkimhi/awless`






#  Apache License

`awless` is an open source project created by Henri Binsztok, Quentin Bourgerie, Simon Caplette and François-Xavier Aguessy at WALLIX.
`awless` is released under the Apache License and sponsored by [Wallix](https://github.com/wallix).

    Disclaimer: Awless allows for easy resource creation with your cloud provider;
    we will not be responsible for any cloud costs incurred (even if you create a
    million instances using awless templates).

Contributors are welcome! Please head to [Contributing (wiki)](https://github.com/wallix/awless/wiki/Contributing) to learn more.
Note that `awless` uses [triplestore](https://github.com/wallix/triplestore) another project developped at WALLIX.
