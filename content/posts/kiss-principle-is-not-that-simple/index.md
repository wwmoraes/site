---
title: KISS principle is not that simple
description: A refresher on its origins and implications on software engineering
date: 2023-12-03T13:10:27+01:00
resources:
- name: featured-image
  src: featured-image.jpg
categories:
- Principle
tags:
- Analysis
lastmod: 2024-08-11T13:26:03+02:00
---

Software engineers use acronyms to convey certain ideas in a single word to
save time. _DRY_, _YAGNI_, _KISS_ and _SOLID_ are samples of keywords that
summarize lots of meaning. Yet compression of human language leads to context
and even meaning loss. Let's talk about the _KISS_ principle.

<!--more-->

## Down the memory lane

_KISS_ stands for _keep it simple stupid_. It's said that Kelly Johnson, a
renowned aeronautical engineer from Lockheed, coined the term. Johnson is best
known for his planes and jets key to the World War II Allies victory.

Kelly is also known for his management style that became norm in his lifelong
company, Lockheed. He led an area that became known as Skunk Works, the cradle
of all his brilliant creations. As per Ben Rich, author of Johnson's posthumous
biography:

> a concentration of a few good people… applying the simplest, most
> straightforward methods possible to develop and produce new products
{source="Ben Rich, describing the Skunk Works division"}

There's also reports on how Johnson challenged his engineering team to design
solutions that a regular engineer would be able to fix using common tools. This
led to simple solutions that were stupid to fix. Hence the term. How does that
relate to software engineering though?

## Simplicity

Let's elaborate on what simplicity means. The Merriam-Webster dictionary defines
simplicity as:

> 1. the state of being simple, uncomplicated, or uncompounded
> 2. lack of subtlety or penetration
>    1. INNOCENCE, NAIVETÉ
>    2. FOLLY, SILLINESS
> 3. freedom from pretense or guile : CANDOR
> 4. directness of expression : CLARITY
> 5. restraint in ornamentation : AUSTERITY

Out of all options above, the 4th and 5th describe better what the term conveys.
We may say something is simple when its easy to explain (clarity) and devoid of
distractions that'd hinder understanding (austerity).

## Baseline

How simple can something be? I'm able to understand someone explain to me a
concept about engineering in plain English. I'll find it complex and fail
miserably if they try to do the same explanation in say, French or German. That
means simplicity relies on an implicit, common baseline shared between both
ends.

Providing the same explanation by mimicking or drawing will certainly increase
the complexity to convey the same message. And even those two methods also need
their own baseline as well. What a thumbs up mean to one person [may not mean
the same to someone else][gestures].

[gestures]: https://www.rd.com/article/common-hand-gestures-rude-in-other-countries/

Without that common baseline, the receiving end will struggle. I may be able to
understand around 80% of Spanish due to is similarity to Portuguese. Yet some
language constructs or specific words that I'm not familiar with can throw me
off in a second.

Yet we don't account for that complexity thanks to translation or by learning
the idiom. Same goes for programming languages 😉

## Perspective

Simplicity depends on the target audience and their point of view over the
solution. Lets take for instance the Blackbirds as seen by:

- **Pilots**: its simple to handle as they already know other aircraft and the
  controls are at the same position.
- **Mechanic engineers**: simple to fix as the tools required are commonplace
  and its intuitive to access certain parts of the engine that often require
  repair.
- **Aeronautical engineers**: simple design considering their domain knowledge,
  good schematics and well-written documentation.
- **Farmers**: an absolute black box, the utmost complexity they've ever seen.
  No idea about how it works or any of that _Mach_ talk.

> [!NOTE]
> Before I get cancelled: please, I hold no prejudice towards farmers. I don't
> mean to say they're ignorant or misinformed. My point is that the knowledge
> related to aircraft is specific so anyone outside that field will most
> certainly have no more knowledge than any other casual passer-by. Unless
> they're enthusiasts about the topic.

## Less is more - sometimes

A recurring interpretation among engineers is that KISS means doing less. Some
even preach it should be the minimal necessary to make it work. While certainly
one of Johnson's tenets was to produce minimal designs, it had to solve the
problem as best as they could at the time. It also had to be clear and easy for
their engineers and mechanics to understand.

That didn't hold back Johnson's team to do one of the most complex projects the
world ever saw that era. In the 1960's they released the Blackbird, an aircraft
that still holds world records to date. As per Johnson himself:

> The idea of attaining and staying at Mach 3.2 (more than three times the speed
> of sound) over long flights was the toughest job the Skunk Works ever had and
> the most difficult of my career.
>
> Aircraft operating at those speeds would require development of special fuels,
> structural materials, manufacturing tools and techniques, hydraulic fluid,
> fuel tank sealants, paints, plastics, wiring, and connecting plugs. Everything
> about the aircraft had to be invented.
{source="Clarence \"Kelly\" Johnson, about the Blackbirds design"}

They had to invent everything from scratch. Isn't that… _non-KISS_?
Over-engineering even? I'm sure they could solve it with some extra glue and
regular connecting plugs instead… 🌚

Doing less can harm the comprehension of the final design, let alone deliver the
expected results. The Blackbirds are nothing short of an engineering wonder, yet
they were in service for over 30 years for their quality, performance and easy
maintenance. That last part is where Johnson applied KISS consistently: the
product itself may be complex, yet its design, build, operations and maintenance
are not.

The opposite certainly applies: imagine if the Blackbirds required all sorts of
unique screwdrivers, hidden compartments and specific ways to access something.
It'd be tiresome to go through that, as the nuances of it would slip the most
attentive eye. The excessive level of "ornaments" distracts the user and may
lead to a loss of context.

Even more so if they decided to add all sorts of extras to the end product. A
small LCD so the pilot could watch their favorite series during a 1:54h trip
from New York to London? Marvelous! What about a popcorn dispenser? Delicious.

## Finding the sweet spot

Here's my definition of the KISS principle, based on its origins: _the art of
designing a solution to a problem using a set of components and tools that leads
to a clear, ease to explain, understand and maintain end product._ That last
part is key to debunk the association _KISS = less_.

> [!quote] Short and sweet
> Keep it simple to understand and stupid easy to maintain.

In coding terms, skipping to isolate a certain functionality does more harm
than good. The fear of "increased complexity" or "excessive ornamentation"
gives birth to functions and scripts with hundreds of lines of code and a high
cyclomatic complexity.

To avoid the use of more capable mechanisms of a language or even a more
capable language altogether leads to hard to understand and maintain code.
Sure, breaking down functionality into smaller pieces will add extra code. Yet
those that follow the domain-driven design principles have a clear intention.
Its easy to understand a dozen clean components than a 200+ lines-long function.

Likewise, overly-short and cryptic mechanisms may look clever, yet they
increase the cognitive load required to understand something. Certain languages
like Haskell and Golang don't even support ternary operators as they
[contribute to create complex expressions more often than not][golang-ternary].

[golang-ternary]: https://go.dev/doc/faq#Does_Go_have_a_ternary_form

> [!info]
> Languages that do not have the ternary operator may still support a stricter
> version of an if/else construct as the right-hand side value.
>
> For instance in Python and Haskell you can initialize values with an expression
> that must have both the if and else blocks returning a value.

Let's do an exercise here with one of my favorite love-hate features of Bash:
variable expansion. Do you know from the top of your head what's happening on
this script?

```bash
FOO=""

echo "${FOO-hello}"
echo "${FOO:-hello}"
echo "${FOO?hello}"
echo "${FOO:?hello}"
echo "${FOO+hello}"
echo "${FOO:+hello}"
echo "${FOO=hello}"
echo "${FOO:=hello}"
```

This code is short and works. Yet its hard to understand and explain, hence far
from KISS. If you were able to guess with confidence and right what each line
does then congratulations, you need some vacations. Urgently. Here's the
reference of how the Bash variable expansions above work.

| Expression      | `FOO="world"` | `FOO=""`    | `unset FOO` |
|-----------------|---------------|-------------|-------------|
| `${FOO:-hello}` | world         | hello       | hello       |
| `${FOO-hello}`  | world         | ""          | hello       |
| `${FOO:=hello}` | world         | `FOO=hello` | `FOO=hello` |
| `${FOO=hello}`  | world         | ""          | `FOO=hello` |
| `${FOO:?hello}` | world         | error, exit | error, exit |
| `${FOO?hello}`  | world         | ""          | error, exit |
| `${FOO:+hello}` | hello         | ""          | ""          |
| `${FOO+hello}`  | hello         | hello       | ""          |

As readable as it can get. And I didn't even touch the prefix, suffix and
replacement expansions. Now let's see how a dual of that feature would look in
Golang:

{{< source variables.go >}}

> [!warning]
> **Please do NOT use this code in real projects!**
>
> This is a silly example, done to port and keep parity with the Bash
> functionality. It lacks proper error handling and even uses the `panic`
> mechanism which is not recommended. There's better ways to handle configuration
> and environment variables in languages such as Golang.
>
> I personally recommend the [spf13/cobra][spf13-cobra] and its peer package
> [spf13/viper][spf13-viper] to handle environment variables, command line flags
> and configuration files seamlessly.
>
> [spf13-cobra]: https://github.com/spf13/cobra
> [spf13-viper]: https://github.com/spf13/viper

I'll grant you that's way more code to get those quirk variable expansions done.
Yet if you read the main function its much clearer what each line does. Even if
you're not familiar with Golang it should be faster to deduct what's going to
happen without deep-diving into its implementation. That's KISS with a dash of
DDD to finish it off.

> [!note]
> My goal here isn't to say Golang is better than Bash, nor that you should
> replace all your scripts with binaries. I used those languages as the latter has
> plenty of unreadable quirks and cryptic design choices, while the former is
> easier to produce a dual with a KISS design.

## Over-engineering

Ah yes, the dreaded over-engineering label. Converting that Bash script to
Golang is certainly an over-engineering, despite how KISS it may be due to the
improved readability. Lets ignore the fact that such conversion happened.
We'll focus on the design of the environment variables expansion rules as a
direct requirement of an existing Golang code base.

Some may consider the `EnvVar` and `EnvVarContext` interfaces "non-KISS" or an
over-engineering as well. Sounds like _You Ain't Gonna Need It_ (YAGNI) instead
of KISS, as the alternatives are arguably more complex:

- **Standalone functions for each use-case**. Lengthy function names and shared
intent. Increased maintenance as functions share part of the logic plus a
side-effect. By separating `getEnv` from each condition logic we apply both the
_Don't Repeat Yourself_ (DRY) and KISS principles, making it easier to
understand and maintain each part alone.

- **inline all the logic directly in the main body**. Plain old spaghetti with
all sorts of logic mixed up in a high cyclomatic complexity function. Also low
to no reusability.

- **write custom operators like Bash**. This is possible in functional languages
that allow you to create operators such as `:?`. Haskell is one of those. Yet
these operators would be specific to a single use case. There's better ways to
functionally represent the conditionals from our example with existing
mechanisms of the language.

## Takeaway

Making your design easy to understand and maintain while delivering the expected
result is the goal of the KISS principle. A design that focus on basic/quirky
languages and primitive constructs is a distorted derivative of the Occam's
Razor principle, not KISS. Be careful, its sharp edges will hurt you at some
point.
