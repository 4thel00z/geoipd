# ```geoipd```

![geoipd-tests](https://github.com/4thel00z/geoipd/workflows/Test/badge.svg)

## What this project is about

This is the geoipd tool.
It's intention is to provide a simple HTTP server which uses the MaxMind GeoLite2 Database to:

- translate between IPs to geo locations
- render IPs on a map (not done yet)

All the usages of this software have to comply with [maxmind's end user license agreement.](https://www.maxmind.com/en/end-user-license-agreement) 
I distance myself from usages that do not comply with the aforementioned end user license agreement.

## How do I install it ?

Pick one of the release tarballs.

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
