---
title: "Katex support"
date: 2019-11-15T12:00:06+09:00
description: "KaTeX is a fast, easy-to-use JavaScript library for TeX math rendering on the web."
draft: false
enableToc: false
enableTocContent: false
tags:
- 
series:
-
categories:
- math
libraries:
- katex
image: images/feature2/mathbook.png
---

The following

$$ \int_{a}^{b} x^2 dx $$

Is an integral

$$ \varphi = 1+\frac{1} {1+\frac{1} {1+\frac{1} {1+\cdots} } } $$

Enable Katex in the config file by setting the `katex` param to `true`. This will import the necessary Katex CSS/JS. 

See the online reference of [supported TeX functions](https://katex.org/docs/supported.html). 

**Note:** For inline math to render correctly, your content file extension must be `.mmark`. See the [official mmark site](https://mmark.nl/). 

```
Inline math: $ \varphi = \dfrac{1+\sqrt5}{2}= 1.6180339887… $
```

Inline math: $ \varphi = \dfrac{1+\sqrt5}{2}= 1.6180339887… $

```
Block math:

$$ \varphi = 1+\frac{1} {1+\frac{1} {1+\frac{1} {1+\cdots} } } $$
```

Block math:

$$ \varphi = 1+\frac{1} {1+\frac{1} {1+\frac{1} {1+\cdots} } } $$