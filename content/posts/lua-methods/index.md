---
title: "Lua Methods"
date: 2021-12-26T18:57:01+01:00
draft: true
description: ""
---
Tables also sport some
nice syntactic sugar features such as as the colon sign on methods. Long story
short:

* method definition using colon tells Lua the function expects a hidden first
parameter called `self`
* method calls using colon tells Lua that the parent object must be passed as
the first parameter of the method

{{< source colon.lua >}}
