---
_build:
  list: true
  render: false
description: The 90's solution to portability that needs to rest in peace
radarIndex: 50
radarSection: languages
radarTier: hold
radarX: 464
radarY: -70
table-of-contents: false
title: Java
---

Back in the 90's, the information technology scene was quite fragmented.
Interoperability wasn't in the agenda of most operating system and solution
vendors. The following decades also witnessed the beef about web standards -
something that's far from over, but its far less pronounced these days.

In that context, Java was born in 1995 to solve it all with the promise of
coding once and running everywhere. Even its ominous installer tries to sell
that to this very day.

Java took the market by storm: developers started learning it and graduation
classes soon picked up the trend. Sun Microsystems, then owner of Java's rights,
seized the opportunity to create an entire ecosystem of solutions such as JBoss
Enterprise Application Platform (JBoss EAP), Java Server Pages (JSP) and the
web applets. Life was good.

Fast forward to 2020's, portability isn't an issue anymore. Operating systems
shifted their focus from isolationist strategies to a more open-doors policy,
embracing other technologies in its wake. For instance Microsoft evolved
Windows' POSIX subsystem to an usable state, and in 2016 started the .NET Core
to decouple their CLR from Windows.

What's more, Linux containers changed everything. With Docker's and Apache
Aurora's (part of Mesos) releases in 2013, followed by Kubernetes in 2014,
containers became the new standard for portable solutions.

In that sense, running Java containers have the drawback that each container
needs its own JVM tooling and auxiliary system binaries from for it to work.
Even the benefits of JBoss EAP are nowadays replaced with language-agnostic
solutions such as message streaming with Apache Kafka or RPC frameworks like
gRPC and Apache Avro.

In general, Java holds market share as there's experienced developers in the
market and demand to create Android applications. Conversely even those are
slowly fading away in favor of Kotlin or Dart which converges the code to
provide a single mobile source across OSes.

We should let Java rest in peace, unless someone have an existing setup of
J2EE/JBoss and other fancy solutions from its ecosystem that they're fond of.
