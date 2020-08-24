---
title: "Third"
date: 2019-10-30T20:46:30+09:00
description: "My third presentation"
tags:
-
series:
-
categories:
-
image: images/feature1/number-three.png
plugins:
- highlight
- zoom
- math
highlightTheme: monokai
revealBackgroundColor: "" # #fff or rgba() or hsl()
revealBackgroundImage: "" # /images/myImage.png   <= static folder path
revealBackgroundPosition: "" # left top, left center, left bottom, right top, right center ...
revealBackgroundRepeat: "" # repeat, repeat-x, repeat-y, no-repeat, inherit
revealBackgroundOpacity: "" # 0~1
revealBackgroundVideo: "" # /videos/myVideo.mp4 <= static folder path, A single video source, or a comma separated list of video sources.
revealBackgroundVideoLoop: false # true, false
revealBackgroundVideoMuted: false # true, false
revealBackgroundSize: "" # cover, contain, ...
revealTheme: league
reveal: 
  - main:
    - sub: 
      - | 
        ## Syntax highlighting
        ```go
        // If an unknown or empty style is provided, AP style is what you get.
        func GetTitleFunc(style string) func(s string) string {
          switch strings.ToLower(style) {
          case "go":
            return strings.Title
          case "chicago":
            return transform.NewTitleConverter(transform.ChicagoStyle)
          default:
            return transform.NewTitleConverter(transform.APStyle)
          }
        }
        ```
  - main:
    - sub: 
      - |
        ## Mathjax 1
        \\[\mathbf{V}_1 \times \mathbf{V}_2 =  \begin{vmatrix}
        \mathbf{i} &amp; \mathbf{j} &amp; \mathbf{k} \\
        \frac{\partial X}{\partial u} &amp;  \frac{\partial Y}{\partial u} &amp; 0 \\
        \frac{\partial X}{\partial v} &amp;  \frac{\partial Y}{\partial v} &amp; 0
        \end{vmatrix}  \\]
  - main:
    - sub: 
      - |
        ## Mathjax 2
        \\[\begin{aligned}
        \dot{x} &amp; = \sigma(y-x) \\
        \dot{y} &amp; = \rho x - y - xz \\
        \dot{z} &amp; = -\beta z + xy
        \end{aligned} \\]
  - main:
    - sub: 
      - |
        ## Mathjax 3
        \\[P(E)   = {n \choose k} p^k (1-p)^{ n-k} \\]
  - main:
    - sub: 
      - |
        ## Mathjax 4
        \\[ \frac{1}{\Bigl(\sqrt{\phi \sqrt{5}}-\phi\Bigr) e^{\frac25 \pi}} =
        1+\frac{e^{-2\pi}} {1+\frac{e^{-4\pi}} {1+\frac{e^{-6\pi}}
        {1+\frac{e^{-8\pi}} {1+\ldots} } } } \\]
---
