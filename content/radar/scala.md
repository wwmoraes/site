---
description: lorem ipsum
radarIndex: 77
radarSection: languages
radarTier: hold
radarX: 360
radarY: -313
table-of-contents: false
title: Scala
---

Scala, a JVM-based language born in 2004, is an spiritual successor to Java.
Its designed to be fully interoperable with existing Java code. It also tries to
solve [some of the criticisms about Java and its cousin C#][scala-overview].

In general, Scala follows my diagnosis about Java: the JVM isn't fit for a
containerized world. Some may point out Scala Native, an attempt to use LLVM to
compile the language to native code + interoperate with C. That's experimental
as of Q4 2023.

Aside from its roots in the JVM, there's more reasons to steer clear from the
language: even though it sports both functional and object-oriented paradigms,
the mix-and-match causes more harm than good. There's even jokes about Scala
being a "Java without semicolons" or "Perl in the JVM".

The fact that the language allows such loose mixing and optional immutability
may lead to what it set out to destroy: hard to understand and debug code,
which is one of the features pure functional programming languages have.
Even [upstream recommendations][thoughtworks-scala] points out to adopt only its
"good parts", which is vague and prone to drift.

[scala-overview]: https://www.scala-lang.org/docu/files/ScalaOverview.pdf
[thoughtworks-scala]: https://www.thoughtworks.com/radar/languages-and-frameworks/scala-the-good-parts
