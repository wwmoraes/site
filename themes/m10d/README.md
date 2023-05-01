# m10d theme

![Intro](https://github.com/wwmoraes/hugo-theme-m10d/blob/master/images/cover.png)

A Hugo minimalistic theme for bloggers. Forked from the simple m10c by @vaga.

Main features:

- Fully responsive
- Twitter Cards and Open Graph supported (see Hugo docs)
- Custom comments system: Giscus
- Custom analytics system: GoatCounter
- Customizable colors (through SCSS)
- Customizable picture and description
- Customizable menu on sidebar
- Customizable social media links on sidebar
- Optimized for performance 100/100 on Lighthouse
- All feather icons included

## Getting started

### Installation

Create a new Hugo site:

```bash
hugo new site [path]
```

Clone this repository into `themes/` directory:

```bash
cd [path]
git clone https://github.com/wwmoraes/hugo-theme-m10d.git themes/m10c
```

Add this line in the `config.toml` file:

```toml
theme = "m10d"
```

### Configuration

In your `config.toml` file, define the following variables in `params`:

- `description`: Short description of the author
- `menu_item_separator`: Separator between each menu item. HTML allowed
(default " // ")

To add a menu item, add the following lines in `menu`:

```toml
[[menu.main]]
  identifier = "tags"
  name = "Tags"
  url = "/tags/"
```

[Read Hugo documentations](https://gohugo.io/content-management/menus/#readout)
for more informations about menu

To add a social link, use the `author.social` native setting by adding the
following lines in your config:

```toml
[[author.social]]
  github = "wwmoraes"
  linkedin = "wwmoraes"
```

### Styling

To change theme colors, create the `assets/css/_variables.scss` file and
override any of the values from `themes/m10d/assets/css/_default.scss` as you
see fit.

**Note:** Hugo releases come in two versions, `hugo` and `hugo_extended`. You
need `hugo_extended` to automatically compile your scss.

## Acknowledgements

- [feather](https://feathericons.com/) - [MIT](https://github.com/feathericons/feather/blob/master/LICENSE)
- [m10c](https://github.com/vaga/hugo-theme-m10c) - [MIT](https://github.com/vaga/hugo-theme-m10c/blob/master/LICENSE.md)
