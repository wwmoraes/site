---
title: Automation is not just a word
date: 2022-01-06T10:55:06+01:00
draft: true
categories:
- Architecture
tags:
- DevOps
- SRE
lastmod: 2024-08-11T13:29:08+02:00
---

After over a decade working in the IT industry I still occasionally hear an
"automation is just a word" comment. Let's demystify what automation stands for,
and how it relates to DevOps and SRE principles.

<!--more-->

We can define automation as a set of technologies and practices that aims to
reduce the need of human intervention in processes. The concept is far older
than IT, and it predates modern society as [history tells
us](https://en.wikipedia.org/wiki/Water_clock).

can trigger people to either ranging from those that say it is just a word to
people that want a Wallace and Gromit morning.

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
is specific on when and why:

- negligent and short-sighted automation planning can have side/negative effects
- quick win solutions can stack up as technical debt if taken as "problem solved"
- side-loading automations to dedicated teams incur in coverage drift and thus toil

Naturally some processes do not need automations. They might have a low
recurrence, low impact and low relevance, and thus does not _need_ an
automation. And that's exactly where the true automation boundary lies in: a
system/process becomes an automation candidate the moment it presents a
reproducible workflow operated by a human and has a non-negligible recurrence,
impact (on it's result and on the team that does execute it) and relevance, not
just through sugar-coated management components as OKRs, KPIs, or the trinity
SLI/SLO/SLA.

Speaking of such syntactic sugar, they support and are a by-product of the SRE
work - not the other way around. The adoption of SRE terminology starting by the
result keywords rather than by its tenets, skipping the essence altogether, is
as artificial as The Matrix itself: you get the nice black suit and sunglasses,
while your problems are everywhere, and multiplying themselves, until "a hero"
comes in and solves everything - manually.

Don't get me wrong on OKRs, KPIs and the trinity SLI/SLO/SLA: they _are_
important and useful. The problem is how one implements them, specially the
order it comes and what priorities they establish to get those. Imagine you have
a service with 5 9s availability for the last 6 months; then question it:

- volume of deployed changes in the same interim
- operational hours the team spent related to the expected outcome
- internal/external users + responsible team churn rates

5 9's is easy on cloud nowadays. That doesn't mean the service has an adequate
availability setup. Neither that its reliable nor sustainable on the long run.
Check the disaster recovery measures, procedures and the feasibility of those,
it might surprise you.

Automations are a step towards the nirvana state-of-art of operations, or rather
_no operations_: autonomous systems/processes. This means tasks capable of
reacting and/or maintain state changes and to act automatically.

Automation and autonomous systems are an intrinsic part of the SRE role and
their expected work. These are all related to a couple of tenets, which can be
further explored individually given the context and boundaries mentioned so far.

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
{source="Joseph Bironas, Google SRE"}

---

> "Besides black art, there is only automation and mechanization."
{source="Federico García Lorca"}

---

> "Hope is not a strategy."
{source="Traditional SRE saying"}

---

> "SRE hates manual operations, so we obviously try to create systems that don’t
> require them. However, sometimes manual operations are unavoidable."
{source="Site Reliability Engineering: How Google Runs Production Systems"}

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

{{< diagram "time-based-availability.svg" >}}

### Aggregate availability

{{< diagram "aggregate-availability.svg" >}}
