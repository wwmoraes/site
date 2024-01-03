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
    background: "#000"
categories:
- TODO
tags:
- TODO
lastmod: 2023-12-17T15:06:05+01:00
---

<!-- TODO intro story/situation -->

Old cartoons are timeless. Back when I was a kid I watched The Flintstones, a
Hanna-Barbera production produced 30 years before. The show intro has one scene
where Fred uses a remote to turn the TV on. A bird then comes out of it, flies
towards the TV and bickers the power on button.

<!--more-->

Of all romanticized elements of such utopian Stone Age, that remote scene stuck
with me as one of the most plausible and clever ones. It also struck me as the
perfect analogy for modern automations, more than 60 years after the show debut.

## Classifying processes

Company processes fall into two categories: unattended and monitored. The first
one relates to workflows where the user doesn't need any sort of check or direct
interaction with another operator. Conversely, the second one relates to the
ones that require some form of validation of the input information.

Unattended processes are sometimes called events. They are non-blocking and the
business merely registers it for audit and/or side-effect actions. Monitored
ones require a check/validation to proceed. Those more often than not require a
human operator to give the green light.

One example of monitored process is the request for access. For compliance
purposes such request must go through a check to ensure the user has the least
amount of access needed to perform their work. Other examples are the request
for peripherals or a network permission such as a new firewall rule.

## Bird flight

Monitored requests need a check against a pre-established set of rules. For
instance, engineers responsible for the servers of a cluster should not need
access to database data. Conversely, a business user should not have privileged
access to the underlying cluster hosts. Hence the approval needed.

{{< diagram "monitored-flow-1.png" >}}

That's an average approval flow, nothing fancy or new here. it works in general,
yet I see a some trade-offs:

1. dependency on an operator (the "bird")
2. overhead
3. operational bias

The first point is about time-to-action. The request has a high jitter time due
to the operator availability. What if its their lunch time or worse, public
holiday? No birds will fly that day.

The second one is subjective. Information Technology Service Management (ITSM)
are notorious for how big and bloated they are. A seasoned support engineer sure
knows their way through that mess; a newcomer not so much. Users even less.

The last one is the biggest and the case in point of this writing. Edge cases
require a human operator to check the input and make a decision. The problem is
such decision isn't straightforward; its prone to bias and context knowledge.

## Removing noise

Decisions based on well-documented conditions are great candidates for an
automatic approval flow. this can be in any form, such as a service that
processes the input and gives the thumbs up/down, or something embedded directly
into your process automation solution of choice.

This frees birds up and avoids a wrong decision over requests. Happy operators,
happy users.

If the process isn't clear like Fred's remote control, then a second-best option
is a memoization solution. You may know it as "learn-as-you-go". The idea is
that the first decision made becomes the default for future ones.

{{< diagram "monitored-flow-2.png" >}}

Such approach has the benefits of:

- consistent outcomes
- decreased time-to-solution
- reduced pressure on operators

<!-- TODO offer solution -->

<!-- TODO wrap up -->

<!-- TODO link back to story/situation -->
