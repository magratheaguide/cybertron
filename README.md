# Cybertron: CLI Theme Compiler

This is a command-line interface (CLI) tool to compile (or _transform_) theme repositories into a single file that can be uploaded to your board host to utilize the theme.

> Cybertron is the homeworld of the Transformers.

## Supported Platforms

Currently only Jcink is supported. Future plans may include mybb.

## Usage: Jcink

There are 4 components that make up a Jcink theme:

1. **Wrapper:** A single HTML file
1. **Stylesheet:** A single CSS file
1. **Templates:** A collection of HTML files, each is optional and will fallback to Jcink's default for that template if not provided
1. **Macros:** A collection of HTML files, each is optional and will fallback to Jcink's default for that macro if not provided

Your folders and files should be structured and named as follows:

```
assets/*
    stylesheet.css*
html-templates/*
    board_stats.html
    cat_head.html
    forum_row.html
    profile.html
    memberlist_head.html
    memberlist_row.html
    mini_profile.html
    post_row.html
    redirect_row.html
    subf_head.html
    topic_list_head.html
    topic_row.html
macros/*
    <MACRO_KEY>.html
    ...
wrapper.html*
```

Note that folder and files names markd with a trailing asterisk above are Cybertron defaults and may be configured.

The names of each `html-template` and macro file, however, come from conventions used in Jcink's own XML for imports and exports and  **must match exactly, including capitalization** or else they won't be recognized on import. Be advised, in Jcink's conventions, template names are all lowercase and macro names are all uppercase.

For the [list of available macro keys and their descriptions](https://jcink.com/main/wiki/jfb-skinning-macros), see the Jcink documentation.

If all of the files are in the default location, then you can run:

```
cybertron build --name "MyTheme"
```

This will output the `MyTheme.xml` file that can be uploaded as a full Jcink "skin set."


## Installation

### OSX

You can find the latest release [on the releases page](https://github.com/magratheaguide/cybertron/releases/latest)

1. Download the latest binary for Darwin
1. Rename to `cybertron`
1. Move that to somewhere in in your PATH. Likely `/usr/local/bin` is a safe choice
1. Add execute permissions via `chmod +x /usr/local/bin/cybertron`

If you get a security error trying to run Cybertron:

1. Go to your "Security and Privacy" settings
1. Choose the "General" tab
1. Where it mentions that Cybertron was blocked, click the "Open Anyway" button

### Windows

Download the latest .exe file from the [releases page](https://github.com/magratheaguide/cybertron/releases/latest). Run it from there and use the `-d` and `-o` flags.
