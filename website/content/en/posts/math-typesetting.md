---
author: Hugo Authors
title: Math Typesetting
date: 2019-12-17T12:00:06+09:00
description: A brief guide to setup KaTeX
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: Park
authorEmoji: 👽
libraries:
- katex
---

{{< box >}}
We need goldmark katex entension which is not yet we have: 
[https://github.com/gohugoio/hugo/issues/6544](https://github.com/gohugoio/hugo/issues/6544)
{{< /box >}}

Mathematical notation in a Hugo project can be enabled by using third party JavaScript libraries.
<!--more-->

In this example we will be using [KaTeX](https://katex.org/)

- Create a partial under `/layouts/partials/math.html`
- Within this partial reference the [Auto-render Extension](https://katex.org/docs/autorender.html) or host these scripts locally.
- Include the partial in your templates like so:  

```
{{ if or .Params.math .Site.Params.math }}
{{ partial "math.html" . }}
{{ end }}
```  
- To enable KaTex globally set the parameter `math` to `true` in a project's configuration
- To enable KaTex on a per page basis include the parameter `math: true` in content files.

**Note:** Use the online reference of [Supported TeX Functions](https://katex.org/docs/supported.html)

### Examples

Inline math: $$ \varphi = \dfrac{1+\sqrt5}{2}= 1.6180339887… $$

Block math:

$$
 \varphi = 1+\frac{1} {1+\frac{1} {1+\frac{1} {1+\cdots} } } 
$$