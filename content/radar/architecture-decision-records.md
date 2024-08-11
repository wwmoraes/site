---
description: keep track of architecture changes as much as your code
lastmod: 2024-08-11T13:01:53+02:00
radarIndex: 5
radarSection: techniques
radarTier: trial
radarX: -250
radarY: -124
table-of-contents: false
title: Architecture Decision Records (ADRs)
---

True agile projects follow an emergent architecture, changing both functional
and non-functional requirements along its entire lifecycle. Those changes
require decisions and updates to its backing material.

Documentation is good, yet long-winded wikis or multi-paged documents are hardly
maintained, let alone read. Even worse if those decisions are slide decks. Those
are wasteful forms of information dumping that don't pay off.

In 2011 [Michael Nygard][michael-nygard] shared the [idea of creating
architecture decision records (ADRs)][adrs]: plain-text, independent markdown
files with key topics. Those files contains all the context needed to reach a
decision and the consequences of it. They're then versioned and stored the same
way as code.

ADRs allows tracking historical reasoning in a clean and concise way.

You can find the original template and scripts from Nygard plus variations in
the community-driven  [ADR GitHub organization][gh-adr].
I personally like the [MADR template][madr-template] due to its short questions
that reminds you of important content to fill in.

[michael-nygard]: https://cognitect.com/authors/MichaelNygard.html
[adrs]: https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions
[gh-adr]: https://adr.github.io
[madr-template]: https://adr.github.io/madr/#full-template
