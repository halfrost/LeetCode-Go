## Zdoc theme for Hugo

Thank you for click me!. Zdoc theme is a simple documentation theme powered by Hugo

## Table of contents

* [Features](#features)
* [Minimum Hugo version](#minimum-hugo-version)
* [Installation](#installation)
* [Updating](#updating)
* [Run example site](#run-example-site)
* [Configuration](#configuration)
* [How to make Doc](#how-to-make-doc)
* [Multi Language](#multi-language)
* [Favicon](#favicon)
* [Shortcodes](#shortcodes)

## Features

* Dark mode
* A mobile menu
* Simple documentation
* Search
* Search Engine Optimization(SEO)
* Multilingual (i18n)
* Responsive design

## Minimum Hugo version

Hugo version 0.60.0 or higher is required.

## Installation

First of all, You need to add config files.
Follow the [Configuration](#configuration) step.

Then, You can download and unpack the theme manually from Github but it's easier to use git to clone the repo.

From the root of your site:

```bash
$ git clone https://github.com/zzossig/hugo-theme-zdoc.git themes/zdoc
```

If you use git to version control your site, highly recommended, it's best to add the zdoc theme as a submodule.

From the root of your site:

```bash
git submodule add https://github.com/zzossig/hugo-theme-zdoc.git themes/zdoc
```

## Updating

From the root of your site:

```bash
git submodule update --remote --merge
```

## Run example site

From the root of themes/zdoc/exampleSite:

```bash
hugo server --themesDir ../..
```

## Configuration

0. From the root of your site: delete config.toml file and add the files below

1. config folder structure. Keep in mind the underscore on the `_default` folder

```bash
root
├── config
│   ├── _default
│   │   ├── config.toml
│   │   ├── languages.toml
│   │   ├── menus.en.toml
│   │   ├── params.toml
```

2. config.toml

```bash
baseURL = "http://example.org"
title = "Hugo ZDoc Theme"
theme = "zdoc"

defaultContentLanguage = "en"
defaultContentLanguageInSubdir = true
hasCJKLanguage = true

copyright = "&copy;{year}, All Rights Reserved"
timeout = 10000
enableEmoji = true
paginate = 13
rssLimit = 100

googleAnalytics = ""

[markup]
  [markup.goldmark]
    [markup.goldmark.renderer]
      hardWraps = true
      unsafe = true
      xHTML = true
  [markup.highlight]
    codeFences = true
    lineNos = true
    lineNumbersInTable = true
    noClasses = false
  [markup.tableOfContents]
    endLevel = 4
    ordered = false
    startLevel = 2

[outputs]
  home = ["HTML", "RSS", "JSON"]

[taxonomies]
  tag = "tags"
```

3. languages.toml

```bash
[en]
  title = "Hugo ZDoc Theme"
  languageName = "English"
  weight = 1

[ko]
  title = "Hugo ZDoc Theme"
  languageName = "한국어"
  weight = 2
```

4. menus.en.toml

You shoud make your own menu.

```bash
[[main]]
  identifier = "docs"
  name = "Docs"
  url = "docs"
  weight = 1

[[main]]
  identifier = "updates"
  name = "Updates"
  url = "updates"
  weight = 2

[[main]]
  identifier = "blog"
  name = "Blog"
  url = "blog"
  weight = 3
...
```

5. params.toml

```bash
logo = true # Logo that appears in the site navigation bar.
logoText = "ZDoc" # Logo text that appears in the site navigation bar.
logoType = "short" # long, short
description = "The ZDoc theme for Hugo example site." # for SEO

themeOptions = ["light", "dark"] # select options for site color theme

useFaviconGenerator = true # https://www.favicon-generator.org/

enableSearch = true
enableLangChange = true
enableDarkMode = true
enableBreadcrumb = true
enableToc = true
enableMenu = true
enableNavbar = true
enableFooter = true
showPoweredBy = true

paginateWindow = 1 # setting it to 1 gives 7 buttons, 2 gives 9, etc. If set 1: [1 ... 4 5 6 ... 356] [1 2 3 4 5 ... 356] etc
taxoPaginate = 13 # items per page
taxoGroupByDate = "2006" # "2006-01": group by month, "2006": group by year

github = "https://github.com/zzossig/hugo-theme-zdoc"
```

## How to make doc

1. Make a folder in the `content` folder. The folder will be appeared in the menu. I'm going to make `doc` folder.

2. Make a `_index.md` file in the `doc` folder.

```yaml
---
title: "Documentation"
description: "test doc index"
date: 2020-01-11T14:09:21+09:00
---

The content here is appeared when you click the manu. So called overview page.
```

3. There are two types of pages. One is a single page and the other is collapsible page.

- single page - Just make a md file in the `doc` folder

    ```yaml
    ---
    title: "Content Formats"
    description: "test post"
    date: 2020-01-28T00:38:51+09:00
    draft: false
    weight: 1
    ---

  *Markdown here*

  ```

- collapsible page - We need to make a new folder inside the doc folder. I'll make a folder named `gettingstarted`. And then, make a `_index.md` file.

    `root/content/doc/gettingstarted/_index.md`

    ```yaml
    ---
    title: "Getting started"
    description: "test post index"
    date: 2020-01-28T00:36:39+09:00
    draft: false
    weight: 2
    collapsible: true
    ---

    ```
    
    The weight defines the order of the post. If the `collapsible` param set `true`, you can see the menu that can be collapsible.

- Make more pages in the collapsible section. Something like `Getting Started`, `Installation`, `Basic usage`, etc... For example, make a file at `root/content/doc/gettingstarted/installation.md`

    ```yaml
    ---
    title: "Frontmatter"
    description: "test post"
    date: 2020-01-28T00:36:14+09:00
    draft: false
    ---

    *Markdown here*

    ```

4. Finally, make a menu in the file at `root/config/_default/menus.en.toml`.

You should make your own menu.

```toml
[[main]]
  identifier = "docs"
  name = "Docs"
  url = "docs"
  weight = 1

[[main]]
  identifier = "updates"
  name = "Updates"
  url = "updates"
  weight = 2

[[main]]
  identifier = "blog"
  name = "Blog"
  url = "blog"
  weight = 3
```

## Multi Language

The default language of this theme is English. If you want to use another language, follow these steps

1. Make a menu file.

```bash 
root
├── config
│   ├── _default
│   │   ├── ...
│   │   ├── menus.ko.toml
```

```bash
config/_default/menus.ko.toml

[[main]]
  identifier = "about"
  name = "about"
  url = "/about/"
  weight = 1

[[main]]
    identifier = "archive"
    name = "archive"
    url = "/archive/"
    weight = 2
...
```

2. Make a content file. Add your language code before the md extension.

```bash
hugo new about/index.ko.md
hugo new posts/markdown-syntax.ko.md
...
```

3. Make an i18n file.

```bash
i18n/ko.toml

[search-placeholder]
other = "검색..."

[summary-dateformat]
other = "2006년 01월 02일"

[tags]
other = "태그"

...
```

4. Edit config.toml file.

```bash
defaultContentLanguage = "ko"
defaultContentLanguageInSubdir = true
hasCJKLanguage = true
```

## Favicon

Put your `favicon.ico` file under the static folder. The filename should be `favicon` and the extension should be `ico`.

### Using favicon-genarator

If you want to support mobile favicon, use [favicon-generator](https://www.favicon-generator.org/).

- Make favicons from favicon-generator site.
- Make a folder at `root/static/favicon`
- Unzip the generated favicon to that folder.
- Set the config param `useFaviconGenerator` to `true`

## Shortcodes

### alert

```bash
{{< alert theme="warning" >}} # warning, success, info, danger
**this** is a text
{{< /alert >}}
```

### expand

```bash
{{< expand "Expand me" >}}Some Markdown Contents{{< /expand >}}
```

### notice

```bash
{{< notice success "This is title" >}} # success, info, warning, error
success
{{< /notice >}}
```
