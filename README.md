# untemis
untemis is a small utility for downloading all your repositories from the [TUM Artemis](https://artemis.ase.in.tum.de)' [Bitbucket server](https://bitbucket.ase.in.tum.de), but it should also work for other Bitbucket servers.


### Installation & Setup
Make sure that [Go](https://go.dev) is installed, you can just "go install" the binary:
```
go install github.com/xarantolus/untemis@latest
```

Then you can create a configuration file based on the [example configuration file](config.example.yml):
```yaml
server: bitbucket.ase.in.tum.de

username: <artemis login user>
password: <artemis login password>
```

Either rename your configuration file to `config.yml` or specify the path to the configuration file via the `-cfg` command-line option.

Now just run the binary to download all your repositories:
```
untemis
```

You can also run the binary with the `--help` flag to see all available options.

### [License](LICENSE)
This is free as in freedom software. Do whatever you like with it.
