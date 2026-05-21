---
title: The Flintstones Automation
description: Lorem ipsum?
date: 2023-12-16T14:05:30+01:00
draft: true
resources:
  - name: featured-image
    src: featured-image.jpg
    title: Featured image placeholder
    params:
      copyright: Image by William Artero
      background: '#000'
categories:
- Architecture
tags:
- Opinion
- Principle
lastmod: 2024-01-28T14:25:12+01:00
---

Old cartoons are timeless gems. Hannah-Barbera's _The Flintstones_ show, set
in the stone age, sports a scene where the main character presses a button on a
TV remote to turn the TV on. A bird then comes out of it, flies towards the TV
and bickers the power on button. Call it an operator and we can start to draw a
parallel with reality.

<!--more-->

Of all romanticized elements of such utopian ancient times, that remote scene
stuck with me. It represents a perfect analogy for modern pseudo-automations,
more than 60 years after the show's debut.

## Classifying processes

I often separate company processes into two categories: unattended and
monitored. The former represents workflows that requires no check or direct
interaction with/by an operator. The latter requires some check and/or an
authoritative action, such as an approval.

Unattended processes denotes non-blocking happenstances that requires little
to no manual actions - from a business perspective - to support it. User
registration and sending a confirmation email are good examples.

Conversely monitored ones requires some form of manual action to proceed. Those
require a human operator to review and approve them.

Request for access represents one sample monitored process. For compliance
purposes the request must go through a check to ensure the user has the least
amount of access needed to perform their work. Requests for peripherals or
new firewall rules also apply monitoring more often than not due to the cost
involved.

## Bird flight

Monitored requests need a check against a pre-established set of rules. For
instance, engineers responsible for the servers of a cluster should not need
access to production database data. A business user should not have access to
the underlying server hosts environment.

![](monitored-flow-1.svg)

That's an average approval flow, nothing fancy or new here. it works with a
few trade-offs:

1. dependency on an operator (the "bird")
2. overhead
3. operational bias

The first point, dependency on an operator, is about time-to-action. The request
has a high jitter time due to the operator availability. What if its their lunch
time or something longer, such as a public holiday? No birds will fly that day.

The second one, overhead, is subjective. Information Technology Service
Management (ITSM) processes and solutions (I'm looking at you, ServiceNow) are
notorious for how big and bloated they are. A seasoned support engineer sure
knows their way through that mess; a newcomer not so much. Users even less.

The last one, operational bias, is the biggest and the case in point of this
writing. Edge cases require a human operator to check the input and make a
decision. The problem is such decision isn't straightforward; its prone to bias
and context knowledge.

## Removing noise

One way to solve operational bias and reduce decision drift is to have
well-documented conditions and their respective decisions. Think a decision
tree with simple binary checks (yes/no). Then provide it to all operators to
follow it to the letter. Granted that the checks are unambiguous, the process
outcomes will converge to a consistent state. Or will they?

Having a decision tree checked by human operators in a repeated fashion
introduces a novel problem: fatigue. Kahneman (2011)[^kahneman-2011] elaborates
on how the human brain has two operating "modes": the system 1 and 2. The first
is quick, emotional, instinctive; the second is slow, logical, objective.

System 1, the faster mode, is what our brains want to operate on most of the
time. After all, _thinking is hard_. This often leads to our brains creating
shortcuts on any slow processes to develop intuition about it to avoid the slow
repetition.

In a decision tree analysis, an operator may think certain checks are _granted_
for a particular case because it _feels like it_ — often because a relevant
proportion of past cases resulted in the same outcome. That's the fatigue
kicking in, skipping the complete check and taking an instinctive, quick
decision.

Instinct is non-deterministic, i.e. a source of entropy, which re-introduces
inconsistencies into the process. It in turn feeds right back into actions
to mitigate them such as tighter controls or more checks. Suffice to say this
becomes a snowball of bureaucracy more often than not.

How to avoid such fatigue then? Is there any entity that doesn't suffer from
that fast/slow system psychological behaviour that would always follow the
instructions thoroughly?

Yes, there is. That's called a computer.

## Automation

Decision trees are great candidates for an automatic approval flow. That's why
services became a staple in companies after the IT boom: they codify decisions
in a way for computers to consistently evaluate them without a fatigue factor,
let alone handle orders of magnitude more cases than it is humanly possible.

This frees operators up and avoids _instinctive_ — to not say outright wrong —
decisions over requests. Happy operators, happy users.

Those operators can then concentrate their efforts on handling edge-cases to
which there's no clear decision established yet. These always require logical
thinking; yet for being novel they don't risk fatigue. The idea is that the
first decision made becomes the default for future ones, a _stare decisis_
principle that allows evolving decisions as you go.

![](monitored-flow-2.svg)

Such approach has the benefits of:

- consistent outcomes
- decreased time-to-resolutions
- reduced pressure on operators

## Design processes out of existence

Ousterhout (2018)[^ousterhout-2018] presents the idea to _define errors out
of existence_. Errors introduce optionality which requires extra handling
and cognitive load. Depending on the tool in use — in particular programming
languages with support for throwing exceptions —, the complexity increases
non-linearly in regards to the variety of errors that may occur. The argument is
to then craft a solution that doesn't require an error.

The book provides a few examples of that, such as file deletion. In the Windows
operating system, deleting a file in use by other processes results in an error.
Conversely on Unixes, a file deletion does not error, even if it's in use by
other processes. The kernel instead marks the file for deletion, removing its
data once all handles close.

We can apply that same principle to processes. By designing an activity from a
different perspective we may avoid the requirement for an entire check/approve
process.

One example I came across myself in practice was network firewall rules in a
platform context. The original process intuition was that a network team and a
CISO team would need to approve all rules requested by users. That's an entire
process with not one but two _birds_ into its hot path.

The alternative? Ingress and egress rules. Instead of central parties approving
a single bi-directional rule, sender and receiver each would need to configure a
rule to match the other end. This eliminates the entire process and the need for
two extra actors.

This also has the added benefit of enabling frictionless automation: since this
model requires no manual action, an automated rule creation solution can handle
those rules, such as declarative infrastructure-as-code tools.

> [!edit]-
>
> 1. Yet another diagram change. This time I'm migrating to D2, which provides
>    a great syntax for diagrams and a standalone CLI to generate good-looking
>    SVGs.

[^kahneman-2011]: Kahneman, Daniel. 2011. Thinking, Fast and Slow. Farrar, Straus and Giroux.
[^ousterhout-2018]: Ousterhout, John K. 2018. A Philosophy of Software Design. Second edition. Yaknyam Press.
