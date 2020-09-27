# Cybertron - A CLI Theme Compiler

This is a cli tool to compile theme repositories into a single file that can be uploaded to your board host to utilize the theme.

## Supported Platforms

Currently only JCINK is supported. Future plans may include mybb.

## Usage - JCINK

There are 4 components that make up a JCINK theme:

1. Templates - These html files should be placed in a folder (default `html-templates`). Each template should be named the same as the JCINK template name.
1. Macros - These html files should be placed in a folder (default `macros`) and named for the macro that it replaces - i.e. LAST_POST.html
1. Stylesheet - This should be a single css file. The default is `assets/stylesheet.css`
1. Wrapper - This should be a single html file. The default is `wrapper.html`

If all of the files are in the default location, then you can run:

```
cybertron build --name "MyTheme"
```

This will output the `MyTheme.xml` file that can be uploaded as a full JCINK "skin set"
