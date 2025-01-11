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
lastmod: 2024-01-28T14:25:12+01:00
---

Old cartoons can become timeless. The Flintstones sports a scene where the main
character presses a button on a remote to turn the TV on. A bird then comes out
of it, flies towards the TV and bickers the power on button. How does that
relate to modern automation?

<!--more-->

Of all romanticized elements of such utopian ancient times, that remote scene
stuck with me. It represents a perfect analogy for modern pseudo-automations,
more than 60 years after the show debut.

## Classifying processes

Company processes fall into two categories: unattended and monitored. The first
one represents workflows that requires no check or direct interaction with an
operator. The second one requires some check of the input information.

Unattended processes, known by some as events, denotes non-blocking happenings
that businesses register for audit and/or side-effect actions. Conversely
monitored ones require a check to proceed. Those require a human operator to
review and approve them.

Request for access represents one sample monitored process. For compliance
purposes the request must go through a check to ensure the user has the least
amount of access needed to perform their work. Requests for peripherals or new
firewall rules also apply monitoring more often than not.

## Bird flight

Monitored requests need a check against a pre-established set of rules. For
instance, engineers responsible for the servers of a cluster should not need
access to database data. A business user should not have privileged access to
the underlying server hosts.

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
