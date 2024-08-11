---
categories:
- Architecture
date: 2023-11-20T00:24:04.867706+01:00
description: a tale about the blind spot of enterprise-grade solutions
publishDate: 2023-11-22T10:42:08.741538+01:00
resources:
- name: featured-image
  src: featured-image.jpg
tags:
- Opinion
title: Scripts don't scale; they give you scriptitis
lastmod: 2024-08-11T13:30:17+02:00
---

Have you ever had to wait for an "automated" process to unblock you that took
hours? Did this process use a pipeline and a bunch of scripts underneath to "get
the job done"? Welcome to a new chronic disease of modern enterprises, what I
call _scriptitis_.

<!--more-->

## Glue language

The concept of script languages first appeared in the 1960's, when computers
began to shift from the batch processing to the time sharing model. That allowed
users to interact (gasp!) with the machine without the need to create a big-bang
workflow and pray it'd work end-to-end.

![https://tenor.com/en-GB/view/chris-pratt-andy-dwyer-omg-shocked-face-meme-gif-25585329](https://media.tenor.com/9CJaHEmyKPAAAAAC/chris-pratt-andy-dwyer.gif)

This new interactive style paved the way for small components that solve a
specific problem to arise. Users then were able to connect or "pipe" them to
get the result in a specific format, or send to another tool for further
processing. Such way-of-working led to the Unix Philosophy and the GNU tools
that are commonplace nowadays.

![https://tenor.com/en-GB/view/linux-trash-linuxbad-gif-18671901](https://media.tenor.com/JFVk98vql5gAAAAd/linux-trash.gif)

The goal of script languages is to glue such components. They help connect
tools and solutions without an expensive or complicated code. The downside
should be clear: they lack stronger interfaces and data contracts. Changes on
any of the chained elements aren't easy to check to ensure compatibility, and
troubleshooting script output is far from great.

For such reasons, they're are also known as glue code or glue languages.

## Glueing glue

What happens when someone overuses glue? They end up submerging the objects
they were trying to glue and lose sight of them. When the glue dries up they
have a new, disfigured object. This new "thing" then is the one that needs most
maintenance or worse, glueing elsewhere.

![https://tenor.com/en-GB/view/sticky-mess-glue-shove-shovel-gif-16740698](https://media.tenor.com/LxKLh3-BsoEAAAAd/sticky-mess.gif)

In case my analogy didn't hammer home my point, scripts are fine as long as
they're:

- kept simple and solve one problem well
- short and sweet (worst case: hundreds of lines of code, less than 1k)
- do something the users can locally do without it, but is faster/easier
- don't become a dependency of another system

They turn into a system of its own when they overstay their welcome. You then
have to maintain a new shiny system written with glue instead of a
strongly-typed and testable programming language.

{{< admonition warning "Testing scripts" >}}
There's testing solutions for scripts as well, such as BATS for Bash and Pester
for PowerShell. They are hacks at best, and you won't see them around for any
vanilla use case of scripts.

Those tools serve for nothing but prove my point that scripts aren't meant to
write systems. Yet you see this kind of tooling appear due to the untreated
_scriptitis_ in the IT industry.
{{< /admonition >}}

## Drawing a line in the sand

How to separate useful from harmful glue? Here's some good uses for scripts:

- chaining tools into custom commands used during manual operations
- small functions to solve minor inconveniences (Bash join anyone?)
- wrappers of more complex commands to simplify or isolate one use case
- auto-completion definitions (decent tools even generate those for you)
- throwaway prototypes and PoCs of future, properly coded solutions
- host configuration (specially when tapping OS-specific settings)

Anything beyond these use cases are at the verge of becoming its own system.

### What about pipelines/CI/CD?

All pipeline/automation solutions I've seen so far have an option to run scripts
such as shells or Python. Those are fine as long as the scripts they use are the
same the developers use locally. If your pipeline has its own scripting separate
from the developer code then I feel sorry for you: you have an anti-pattern.

![https://tenor.com/en-GB/view/30rock-comfort-feel-better-jack-donaghy-gif-12796475](https://media.tenor.com/fLhCWlXe5Q4AAAAC/30rock-comfort.gif)

Programming in a pipeline to do something that the user cannot do locally is the
prime excuse from security advocates to do all sorts of compliance checks.
Those solutions rely on roadblocks such as merge/pull requests and branch
policies to enforce them at the wrong place. This leads to yet another set of
anti-patterns, namely continuous isolation and CI theatre. Those are extensive
topics I'll save for another day.

The bottom line is: Pipeline's broken? Uh oh, time to switch context from the
main solution code to troubleshoot the pipeline code. Split brain right there.
Then you realize you need an external team to change the templates as you don't
have the rights to update them…

![https://tenor.com/en-GB/view/barney-stinson-neil-patrick-harris-himym-how-i-met-your-mother-gif-5353868](https://media.tenor.com/C45MBZAcrlwAAAAC/barney-stinson.gif)

## Script ambivalence

Don't get me wrong: I love scripting languages. I use them for common tasks on
both personal and work environments. My teammates at some point get used to see
_a wild `Makefile`_ that appears on each repository I ever touch. Another good
example is how I use [chezmoi][chezmoi] to deploy [preference files][dotfiles]
and run dozens scripts for the "last mile" setup on my machines.

[chezmoi]: https://www.chezmoi.io
[dotfiles]: https://github.com/wwmoraes/dotfiles

Anything beyond that becomes the DevOps-certified ™️ version of spaghetti code.
Script A uses script B/tool C, cleverly glued together in a pipeline more often
than not. It commonly requires secrets or dependencies non-reproducible locally
as well. Why? Because that's what "continuous integration" is all about, right?
A castle of cards that no one dares to touch unless it breaks.

![https://tenor.com/en-GB/view/right-natalie-portman-star-wars-rd_btc-gif-24051918](https://media.tenor.com/Wza_7q92YIQAAAAC/right-natalie-portman.gif)

## Glue replacement

What's the alternative then? Software engineering using a strong typed and
preferably compiled language. Those allow developers to better provide and
interact with distinct solution APIs. It also doubles down to create a
decoupled internal architecture where solutions communicate based on known data
contracts.

A case in point is your cloud provider of preference. You'll notice that all
services they provide to you have an API. Sure, you also have a graphical
portal as well, which relies on the API to relay your requests. Infrastructure
code also depends on these APIs to make the magic happen. Why don't they provide
you "scripts" to get this done instead?

{{< admonition note "cloud vendor scripts ™️" >}}
Matter of fact cloud service vendors **do** provide you scripts to interact
with their APIs. The big three (AWS, GCP, Azure) for instance have their
official CLI tools written in Python. Then people go and make all sorts of
atrocities in pipelines with it. 😰

From a developer perspective, most of those CLI tools rely on generated code
from the API specs and are as mechanical as they can get. They have tests, yet
[simple fixes face a bumpy road and take its sweet time][azcli-fix] to get
through due to how the test becomes the problem.

[azcli-fix]: https://github.com/Azure/azure-cli/pull/26013#issuecomment-1651158706
{{< /admonition >}}

A clean architecture won't take more time than any thousand-line-sized script.
The trade off is where the speed slope is: scripts are faster to create and
slow to maintain, while a proper application is slower to create and faster
to test and thus change.

![https://tenor.com/en-GB/view/there-is-a-trade-off-matt-ginsberg-startalk-compromise-risk-gif-20160753](https://media.tenor.com/LlK2_K0paXgAAAAd/there-is-a-trade-off-matt-ginsberg.gif)

Companies that rely on such "system scripts" waste plenty of engineering
potential. One hour of a glued pipeline execution siphons at least one engineer
time that could've been better invested elsewhere. And even if not, it'd at
least avoid the deadline snowball.
