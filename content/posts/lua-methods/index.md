---
title: Lua Methods
date: 2021-12-26T18:57:01+01:00
draft: true
categories:
- Code
tags:
- Lua
description: ""
lastmod: 2024-08-11T13:08:28+02:00
---
Tables also sport some nice syntactic sugar features such as the colon sign on
methods. Long story short:

* method definition using colon tells Lua the function expects a hidden first
parameter called `self`
* method calls using colon tells Lua to pass the parent object as the first
parameter of the method

{{< source colon.lua >}}
