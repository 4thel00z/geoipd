# ```geoipd```

![geoipd-tests](https://github.com/4thel00z/geoipd/workflows/Test/badge.svg)
![geoipd-logo](https://github.com/4thel00z/geoipd/raw/assets/logo.png)

## What this project is about



This is the geoipd tool.
It's intention is to provide a simple HTTP server which uses the MaxMind GeoLite2 Database to:

- translate between IPs to geo locations
- render IPs on a map

[![asciicast](https://asciinema.org/a/R1C7LZcAjOluuhI57CTngzbPM.svg)](https://asciinema.org/a/R1C7LZcAjOluuhI57CTngzbPM)

All the usages of this software have to comply with [maxmind's end user license agreement.](https://www.maxmind.com/en/end-user-license-agreement) 
I distance myself from usages that do not comply with the aforementioned end user license agreement.

## Example rendering

This is the example rendering that was rendered by locating the ip adress `1.1.1.1`:

![geoipd-rendering](https://github.com/4thel00z/geoipd/raw/assets/out.png)

## How do I install it ?

Pick one of the [release tarballs](https://github.com/4thel00z/geoipd/releases/latest).

If you have linux machine with amd64, you can use this oneliner instead:
```
curl -s https://api.github.com/repos/4thel00z/geoipd/releases/latest |grep "browser_download_url"| cut -d '"' -f 4| grep linux| grep amd64| wget -qi -
```

## How do I run it ?

After creating a new project like above you can simply run:

```
make run
```

or

```
just run
```

if you have [just](https://github.com/casey/just) installed.
Running `make help` will show you the rest of the targets.

## License

This project is licensed under the GPL-3 license.
