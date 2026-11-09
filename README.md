# gobble 

![gobble demo](https://files.stevedylan.dev/gobble-demo.png)

A simple systems fetch in ~300 lines of Go with no dependencies

## Quickstart 

1. Make sure Go and Git are installed

```bash
go version
git --version
```

2. Clone repo and install

```bash
git clone https://github.com/stevedylandev/gobble
cd gobble
go install . 
```

3. Run and Debug

```bash
gobble
```

This program likely won't work for your system out of the box. Run it, figure out the issue, and fix it. 

>[!WARNING]
>Running a clanker on this will take away the satisfaction of fixing it yourself; proceed with caution and think twice

## Reasoning

Yes it's another fetch tool, but I promise this one is (kinda) different. I wanted to use this project as an exercise to practice writing some Go by hand without AI assistance. It turned out to be a rewarding experience, and I kinda hope others use it as an opportunity for themselves. Gobble was built on a machine running Alpine Linux, so the chances of it working out of the box on another system is unlikely. Instead of expanding it to support other distros or adding a config, I decided to leave it as is. This is my version of gobble, but I hope others fork it and use it to build their own version, and learn something along the way. 

## License

[GNU GPL v3.0](LICENSE)
