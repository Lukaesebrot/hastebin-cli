# hastebin-cli
A cross-platform CLI for hastebin with private server support


## Installation
If you have Go installed, you can just use the `go install` command:
```
go install github.com/Lukaesebrot/hastebin-cli
```

If not, feel free to download a binary from the releases tab.


## Configuration
The CLI will create a configuration folder and file in your home directory (ex.: `C:\Users\Lukas\.hastebincli\config.json`) which looks like this:
```json
{
  "autoClip": true,
  "instance": "https://hasteb.in"
}
```

If `autoClip` is set to `true`, the CLI will automatically copy the paste URL to your clipboard.


## Usage
You can simply pipe in some text using stdin:

**Windows:**
```
type test.txt | hastebin-cli
```

---

If you disabled the auto-clip configuration, you may want to copy the output to your clipboard manually:

**Windows:**
```
type test.txt | hastebin-cli | clip
```


## Issues
If you find any bugs or if you have a feature request, feel free to open an issue and/or pull request!