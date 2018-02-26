# swamp
Keeps configuration files safe, organized, and collectivized.

```bash
go get github.com/kbougy/swamp
go install github.com/kbougy/swamp
```

Before getting started, verify that you have your $GOPATH/bin as part of your path.

See here for more information: https://golang.org/doc/code.html

Move the .zshrc file into the .swamp directory
```bash
mv ~/.zshrc ~/.swamp/.zshrc
```

And activate the swamp
```bash
swamp activate
```

Verify the file exists, and that it's a symlink
```bash
ls -l ~/.zshrc
```

Deactivate at any time
```bash
swamp deactivate
```


## Todo
* A user should be able to split their configuration files into organized groups
* A user should be able to download and install someone else's configuration files
