---
title: automation is not just a word
date: 2022-01-06T10:55:06+01:00
draft: true
tags:
- DevOps
- SRE
- Soft skills
---

After over a decade working in the IT industry I still occasionally hear an
"automation is just a word" comment. Let's demystify what automation stands for,
and how it relates to DevOps and SRE principles.

<!--more-->

Automation can be defined as a set of technologies and practices that aims to
reduce the need of human intervention in processes. The concept is far older
than IT, and it predates modern society as [history tells us](https://en.wikipedia.org/wiki/Water_clock).

<!-- home automation -->
<!-- financial automation -->

can trigger people to either ranging from those that say it is just a word to
people that want a Wallace and Grommit morning.

{{< youtube o0JIznO-RlI >}}

Some people, depending on how they read or receive the information given on
Google's SRE book, might ask every time someone raises an automation/autonomous
flag: "why should we automate this?", or even "do we need to automate this?" -
which might be a valid question, given the system/operation in question is a
transitive one, as a new system, a disposable PoC or a unique operation that
has no clean pattern or recurrence yet for instance.

Aside from those specific exceptions, operations can be, to some or total extend,
mapped into well-known steps and actions that get repeated every single time.
On worst case scenarios, there will be ROI-oriented arguments, as "it'll cost
more to automate this than to keep a support shift to distribute the toil", or
"creating yet another system will add more toil, as we'll need to maintain yet
another system!". Have you seen any of those arguments backed by a cost analysis
comparing the team hours drained by toil, support and any sort of known tasks
repeated throughout any given time?

Google's SRE does indeed mention to take automations with a grain of salt, which
is very specific on when and why:

- negligent and short-sighted automation planning can have side/negative effects
- quick win solutions can stack up as technical debt if taken as "problem solved"
- side-loading automations to dedicated teams incur in coverage drift and thus toil

Naturally some processes do not need automations. They might have a low recurrence,
low impact and low relevance, and thus does not _need_ and automation as per se.
And that's exactly where the true automation boundary lies in: a system/process
becomes an automation candidate the moment it presents a reproducible workflow
operated by a human and has a non-negligible recurrence, impact (on it's result
and on the team that does execute it) and relevance, not just through sugar-coated
management components as OKRs, KPIs, or the trinity SLI/SLO/SLA.

Speaking of such syntatic sugar, they are supposed to support and in some cases
are a by-product of the SRE work - not the other way around. The adoption of SRE
terminology starting by the result keywords rather than by its tenets, skipping
the essence altogether, is as artificial as The Matrix itself: you get the nice
black suit and sunglasses, while your problems are everywhere, and multiplying
themselves, until "a hero" comes in and solves everything - manually.

Don't get me wrong on OKRs, KPIs and the trinity SLI/SLO/SLA: they _are_ important,
useful and needed. The problem is how it is implemented, specially the order it
comes, and what priorities are established in order to get those alive and kicking.
You can have a service with 5 9s availability for the last 6 months, but then how
many changes have been rolled out on the very same interim? Also how much
operational hours has the team dedicated to operations related to the expected
outcome of such service? And how high is the internal/external users + responsible
team churn rates? 5 9's is easy on cloud nowadays; that doesn't mean the service
has an adequate availability setup, is really reliable or even sustainable on the
long run. Shall we check the disaster recovery measures, procedures and the
feasibility of those?

Last but not least, automations are a step towards the nirvana state-of-art of
operations, or rather _no operations_: autonomous systems/processes. This means
not only tasks are executed automatically, but they also are capable of reacting
and/or maintain state changes, thus responding to a known state accordingly.

With all that said, it is clear that automation, and further on autonomous systems,
are an intrinsic part of the SRE role and expected work. These are all related
to a couple of tenets, which can be further explored individually given the
context and boundaries aforementioned.

- availability
- latency
- performance
- efficiency
- change management
- monitoring
- emergency response
- demand forecasting/capacity planning

## Tenets

### Availability

- level which the users will be happy with
- alternatives to users dissatisfied with the product's availability
- user response to different availability levels

Automations:

- SLO monitoring and alerts
- user NPS
- user feedback vs availability timeline

### Latency

- more human intervention = more latency
- self-healing emergencies = high availability

manual intervention => playbook => runbook (automated healing) => service improvement

Automations:

- self-healing
- rollback
- declarative definitions
- definition sync/status (to rule out manual-change problems)

### monitoring

- alerts
- tickets
- logging

Automations:

- generic alerts
- alert delivery
- alert response (runbooks)
- ticket generation (human needed intervention/follow up actions)
- detailed logging (e.g. debug level during incidents)

### emergency response

- MTTF: mean time to failure
- MTTR: mean time to recovery

more Humans = more Latency

Automations:

- failure detection (less MTTF)
- self-healing (less MTTR)
- troubleshoot/diagnostic runbooks (less human diagnostics = less latency)
- countermeasure runbooks (less human actions = less errors = less latency)

### Change management

- progressive rollouts
- quick and accurate problem detection
- roll back changes safely when a problem happens

Automations:

- automatic integration
  - CI with tests
- automatic delivery
  - deployment
  - rollback on failure
- automatic health check/problem detection
  - service liveness
  - service readiness
  - troubleshoot/diagnostic runbooks

### Demand forecasting/Capacity planning

- organic demand forecast (outside of the natural lead time required to acquire capacity)
  - natural growth (service adoption/usage)
  - more features = slowdown
- inorganic demand forecast
  - spikes due to business-driven actions (marketing, events, feature launch)
  - spikes due to external actions (attacks, security analysis, region outage)
- constant load testing to measure raw capacity
  - hardware
  - network
- provisioning
  - new instances/regions
  - configuration of resources (LB, network, tools)

Automations:

- auto-scaling
- historical data
- inorganic margin
- automated provisioning

### Efficiency and Performance

- target at a specific response speed
- monitor and modify to improve performance = more capacity = more efficiency

Automations:

- scaling by service capacity usage (e.g. RPS instead of raw capacity as CPU/RAM)
- monitoring and alerting on capacity/raw capacity thresholds (loss of
  efficiency/performance)

## Quotes

> "If a human operator needs to touch your system during normal operations, you
> have a bug. The definition of normal changes as your systems grow."
>
> - Carla Geisser, Google SRE

---

> "If we are engineering processes and solutions that are not automatable, we
> continue having to staff humans to maintain the system. If we have to staff
> humans to do the work, we are feeding the machines with the blood, sweat, and
> tears of human beings. Think The Matrix with less special effects and more
> pissed off System Administrators."
>
> - Joseph Bironas, Google SRE

---

> "Besides black art, there is only automation and mechanization."
>
> - Federico García Lorca

---

> "Hope is not a strategy."
>
> - Traditional SRE saying

---

> "SRE hates manual operations, so we obviously try to create systems that don’t
> require them. However, sometimes manual operations are unavoidable."
>
> - SRE book

## Extras

### Operations and support teams

- monitoring and alerting (e.g. automated metrics collection, alert management)
- automated checks (e.g. jobs on Rundeck, Stackstorm)
- homegrown toolkit for troubleshooting and check (e.g. CLI, portal)
- proper issue escalation
  - more automated checks = faster and precise RCA/domain narrowing = less
    people requested/involved = less latency

## Formulas

### Time-based availability

![time-based availability](https://render.githubusercontent.com/render/math?math=availability%20=%20\frac{uptime}{(uptime%20%2B%20downtime)})

### Aggregate availability

![aggregate availability](https://render.githubusercontent.com/render/math?math=availability%20=%20\frac{requests_{successful}}{requests_{total}})
