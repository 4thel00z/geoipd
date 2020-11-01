# ```geoipd```

![geoipd-tests](https://github.com/4thel00z/geoipd/workflows/Test/badge.svg)

## What this project is about

This is [my](https://github.com/4thel00z) go http service template.
It sports features like:

- validation support (see debug module for example)
- jwt validation support
- module support (see debug module for example)


## How do I install it ?

To create a new project simply invoke this script, (make sure to use your own project name instead of `<project_name>` lol):

```
curl  --proto '=https' --tlsv1.2 -L -sSf https://shortly.fun/boilerplate | bash -s <project_name>
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
