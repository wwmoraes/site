---
title: Are you a DevOps engineer?
description: how a way of working became a catch-all role
date: 2023-06-28T08:00:00+01:00
resources:
- name: featured-image
  src: devops.jpg
categories:
- Opinion
tags:
- Career
- DevOps
---

A search for the term DevOps on LinkedIn will yield countless posts and hundreds of jobs. The market for DevOps engineers is plentiful and has been for at least a decade now. Everyone wants to join the hype and either calls themselves DevOps engineers or adds the DevOps keyword in their CVs for that extra oomph. Is that the right thing to do?

{{< figure "Perhaps?" >}}
![http://reactiongifs.com/?p=24986](http://www.reactiongifs.com/r/dsori.gif)
{{< /figure >}}

<!--more-->

## Prologue

I had DevOps references across my LinkedIn profile and curriculum vitae up until mid-2022. My description depicted me as a DevOps engineer looking to contribute to automation. Not anymore. The reason is simple: the market overloads and marginalizes the term. Recruiters and managers generally associate it with distorted practices and solutions I don't agree with.

![http://reactiongifs.com/?p=22693](http://www.reactiongifs.com/r/wsod.gif)

## What is DevOps

I won't preach to you about the term. My understanding is that DevOps means a body of knowledge and best practices to close the gap between development and operations. It includes the reuse of tools and the art of automation to glue processes together. How to glue is up to the engineer and business, and this acts as a double-edged sword.

![http://reactiongifs.com/?p=10562](http://www.reactiongifs.com/wp-content/uploads/2013/06/flesh-wound.gif)

This definition is intentionally vague. Bodies of knowledge and processes' best practices are so. This allows great flexibility when applying them. You can still reap great results granted you preserve the main tenets. If you know agile you'll also recognize the same concept from its manifesto here. It allows small iterations over the processes to further improve them.

## Automation

We deem a system or process as automated when it requires minimal to no human intervention to make decisions. Note that this differs from *automatic*, which means a process that acts involuntarily. A car wash is automatic; a home light that turns itself off 30 minutes before dawn is automated.

The degree and scope of automation can vary between a small, atomic process and a bigger flow. An integration pipeline in a source code repository does not mean automation as per see. Let's say it has a step to run a lint tool and check the code. Regardless of the lint output, the merge requires a human to approve and execute it.

![http://reactiongifs.com/?p=19982](http://www.reactiongifs.com/r/mgc.gif)

In such a scenario the linting is *automatic*, and the integration itself is manual. We'd have an *automated* integration if a merge happened after a successful lint run with no issues found.

Before the four-eyes-principle priests start throwing rocks at me, mind you that this is a simplistic example to illustrate my argument. Certainly, we'd need testing, checking the author against an authorized group and other more sophisticated checks like security and secret scanning to securely automate such flow.

![http://reactiongifs.com/?p=24159](http://www.reactiongifs.com/r/jj.gif)

## The DevOps role

If DevOps is knowledge and best practices, then the role means people that follow those, as you'd expect. The catch is how broad the term is: again, it acts as a double-edged sword here. It's not clear what's the exact set of soft and hard skills nor the technologies a DevOps must know; instead, each company/industry has its wishlist printed on a DevOps-themed page.

Aside from the technology stack, which is arguably subjective, one expects the required soft and hard skills to remain. For instance, the need for an engineer with programming and architecture experience should be consistent. It's not.

![http://reactiongifs.com/?p=21184](http://www.reactiongifs.com/r/W28tx6T.gif)

### Expectation

As a software engineer joining a position that mentions DevOps and automation, I expect a company looking for ways to, well, *automate*. That means investment into proper architecture and software development to handle certain domains within the business. This also leads to connecting those solutions in a way that certain decisions do not require human intervention.

Those automations in turn would allow improving processes, or free up time to focus on other processes. Maintenance and operations then would derive from business rule changes and fine-tuning decisions the automation has to make.

Is there a margin for over-engineering here? Yes, there's always. The point is to solve a problem *without* creating another (dozen) one(s). And trust me, a pipeline that calls another pipeline that uses a CLI to trigger an API to do Odin knows what is *far* from the word stable, let alone testable. Yet some people prefer to deal with shell script quote shenanigans than to code in a properly typed and structured language. Beats me why.

![http://reactiongifs.com/?p=18567](http://www.reactiongifs.com/r/but-why.gif)

### Reality

What happens when you focus on gluing together dozens of scripts, CLIs and raw HTTP calls together? You risk losing the why and numb yourself on how to best tackle the problem. You may even forget about how the big picture performs. Flagship solutions become the target, not the means to an end.

This means classic system administrators and operators are at home: this kind of specialized professional knows the in and out of specific solutions or vendor suites in the market. They can - and trust me, they will - get you onboard on all sorts of tools.

![http://reactiongifs.com/?p=4694](http://www.reactiongifs.com/wp-content/uploads/2012/12/more.gif)

Such is the reality of DevOps positions. Most of the time they require low to no programming languages, listing scripting ones such as Bash, PowerShell and Python. Then focus on dozens of branded solutions you must have years of experience. Those are constantly misused and work on top of terrible integrations, which you have to provide support for until the next migration exodus happens.

The day-to-day of DevOps engineers then becomes the maintenance of *magic* scripts, pipelines and sorting out integration problems between solutions. As those integrations are black boxes operated by the DevOps engineers, the users won't know a thing about how to solve them. This leads to the ownership and operations of the DevOps team. Then the snowball begins.

![http://reactiongifs.com/?p=24900](http://www.reactiongifs.com/r/idwt.gif)

Granted, when done right, gluing tools can give awesome results. Kubernetes is a successful example of this. The difference between such a project and an average pipeline solution is simple: Kubernetes offers an extensive abstraction API layer on top of all the tools it uses. It smooths over differences between options so you don't have to. It also relies on a strongly typed and compiled language to provide both the coding security and functional testing needed.

Don't get me wrong about vendor solutions. They deliver good value when you use them right and to their fullest. I have yet to see a company fully use all features of Jira, such as components, instead of relying on tags or external spreadsheets to sync up.

![http://reactiongifs.com/?p=22317](http://www.reactiongifs.com/r/emb1.gif)

## What about SRE?

Although part of the market uses both roles interchangeably, they differ in key areas. Some good sources for in-depth understanding are the books by Google about the topic. They are [available online for free](https://sre.google/books/).  You can get a paperback or ebook copy in most stores as well. I'll briefly touch on the points I feel are the most important.

### A superset

SRE is a superset/extends DevOps. Some may say an SRE does abide by the DevOps practices and mindset. In other terms, SRE defines new mechanisms and constraints to turn the abstract ideals of DevOps into a concrete way of working.

### business OS

SREs tackle operational problems as if its software. Developers know that the worst enemy of any software is *input*. What's a good source of random input? Humans.

Integrating systems is easier as you can make both ends agree on the protocol, messages and content statically, and fully test those. Breaking changes are then caught during tests to avoid rolling them out into production. Granted you're writing tests, that is.

![https://tenor.com/en-GB/view/dev-developer-bugs-bug-qa-gif-25713242](https://media.tenor.com/CZZJVKwzRCMAAAAd/dev-developer.gif)

Reducing human interaction reduces uncertainty. That in turn reduces the amount of data transformations and validations needed to convert a human representation into a machine one. This means systems are easier to build and maintain when they connect solely with other systems.

If you don't believe it, ask a fellow frontend developer about the nightmares they have about the QA analyst feedback, and how they got the most unexpected bugs. I bet it won't be a short list.

![https://tenor.com/en-GB/view/qa-gif-26507223](https://media.tenor.com/S-CxC0jhfrMAAAAC/qa.gif)

Sure, at the end of the day, a business actor still needs to click something somewhere to keep the business running. This happens in a system specialized in receiving such input, the frontend - anything beyond it talks machine language. Clean architecture anyone?

### Operation is not the job

This to me is the biggest difference. Organizations often organize DevOps teams in anti-pattern topology. Some are:

- independent DevOps team - bridges development and operations teams, a sort of DevOps-as-a-Service (DaaS)
- "We all are DevOps engineers"/"DevOps is not needed" - developers either handle some superficial operations and offload the rest to the operations team, or outright switch to IaaS/SaaS/DaaS solutions
- DevOps tooling team - a development team creates tools to interface the operations team's way of working and solutions. The other developers then rely on those
- System Administrator 2.0 ™️ - the operations team either hires DevOps engineers or adopts solutions related to DevOps
- Ops as a Dev subset - there's no operations team; the development reserves time to do all operations
- System Administrator 3.0 ™️ - same as the 2.0 + now they call themselves SRE

![https://tenor.com/en-GB/view/i-mean-yolo-right-you-only-live-once-right-you-got-to-go-for-it-you-have-to-risk-it-debby-ryan-gif-15835808](https://media.tenor.com/YXnqEtz2FCsAAAAC/i-mean-yolo-right-you-only-live-once-right.gif)

A sharp reader will notice that all those patterns still have the classic operations team in direct contact with the developers, regardless of who wears the DevOps badge. If you want to know more about good and bad patterns then I recommend the outstanding Team Topologies book, by Matthew Skelton.

In the SRE model, the operations team works in complete isolation from the developers. They focus on maintaining the infrastructure and underlying solutions that shoulder the business. The SRE team then is an extension and interface that intakes solutions from the development teams and run them on top of it. Do they accept anything that's thrown at them? **No, and here lies the biggest difference.**

![https://tenor.com/en-GB/view/gandalf-the-grey-lord-of-the-rings-ian-mckellen-mines-of-moria-fellowship-of-the-ring-gif-24036179](https://media.tenor.com/EgvXcIbZLqgAAAAd/gandalf-the-grey-lord-of-the-rings.gif)

*An SRE team can autonomously deny running a solution that doesn't meet their quality standards.* For instance, minimal test coverage - I'm not talking about those fully mocked unit tests to snatch an easy 100% coverage - and meaningful logging are minimal requirements. This ensures they keep the operational levels of the entire platform and related services.

They treat each solution as a black box that needs to run; it doesn't matter what or how it does each of them does its job, as long as it can pull its weight. It also enables the SRE team to reduce outages and problems on their own, without either the development or operations teams unless strictly needed.

This means they focus on the thin layer of glue needed to run and the entire catalog of services under their watch on top of the company's infrastructure. Automation to run, observe and recover faulty services automatically, falling back to alerts and incident management, are then put in place by them to keep the business running.

![https://tenor.com/en-GB/view/talk-about-credibility-michael-che-saturday-night-live-credible-validity-gif-16232625](https://media.tenor.com/ZxezqP2XjnQAAAAC/talk-about-credibility-michael-che.gif)

### Fail fast, recover faster

Reliability is subjective. Most people consider the rate of failure as the sole indicator of it. What people forget to consider is who perceives failure. If a customer never sees your solution down then it seems like a 100% uptime. How to make such an illusion possible is the real magic here.

For SREs, failure is inevitable. Taking hours to nurse a system back to health isn't. A system that takes a short time to recover gives the development team the confidence to increase the pace at which they deliver changes, which reduces the size of each delivery, making them easier to test and fix.

![https://tenor.com/en-GB/view/so-pure-rare-sarcasm-john-oliver-gif-9404210]

In comparison, DevOps lacks the authority and autonomy of the SRE to assert quality and deny low standards solutions. That means all checks and validations must run in the so-called *CI/CD pipelines*. A single project build and push may take dozens of minutes due to the level of hoops and repeated checks made each time.

It's needless to mention all the trampoline gymnastics needed to trigger protected pipelines and whatnot, then poll their status to proceed. The end result? Teams stack changes to do a big-bang release to reduce the wait time.

![https://tenor.com/en-GB/view/enter-dev-developer-commit-git-gif-21009539](https://media.tenor.com/M0na3YR-rTcAAAAd/enter-dev.gif)

## And platform engineering?

If done wrong, it's the same as the system administrators 2.0. Check the [operation is not the job](#operation-is-not-the-job) above. Otherwise, they are the precursor of SRE.

![https://tenor.com/en-GB/view/same-different-but-still-gif-18224441](https://media.tenor.com/BRckVcpUYlUAAAAC/same-different.gif)

## Developer experience

As far as I've seen, DevOps and platform teams tend to offer the same set of solutions:

- **Out-of-tree infrastructure management**. Businesses mandate product developers to use Infrastructure as Code (IaC) for their solutions, yet developers need to request certain underlying resources through services that don't integrate with any IaC tooling. That either happens through an internal website, sending an instant message to one of the DevOps engineers, using a pipeline provided by them or creating a pipeline with provided templates and access to restricted service connections
- **Operational support**. The DevOps team invests most time documenting and supporting the use of their abstractions due to their divergence from any commonly known patterns and solutions available in the market, and also bringing those same solutions up to speed due to changes in allowed options or internal use cases
- **Flimsy integrations**. The DevOps team relies on scripting languages - such as Bash, Python and PowerShell - and trampoline solutions such as pipelines calling other pipelines to deal with limitations related to security or the perceived low technical value of a refined solution. This reflects the operational support as it adds a considerable amount of time to troubleshoot when one of those moving parts fails

![https://tenor.com/en-GB/view/burn-in-hell-elmo-fire-flame-gif-8764555](https://media.tenor.com/X1UBzspDL3kAAAAC/burn-in-hell-elmo.gif)

On top of that, big companies may try to *empower* the product teams by delegating to them their infrastructure management (see again [operation is not the job](#operation-is-not-the-job)). While the business believes this is automation and best practices, developers lose considerable time to learn, maintain and troubleshoot the custom solutions they need to handle.

The platform team then has to deal with a surge of developers asking for help to troubleshoot the black box they call the platform, and it snowballs. Ultimately they will create a support team to block those, following the ITIL practices. They may also document everything on their terms to justify rejecting any requests for help.

![https://tenor.com/en-GB/view/the-it-crowd-moss-the-it-crowd-the-it-crowd-fire-moss-the-it-crowd-fire-richard-ayoade-gif-15210949](https://media.tenor.com/PRN-EHOCuHwAAAAd/the-it-crowd-moss-the-it-crowd.gif)

## Aftermath

Misuse of the DevOps term and its partial implementation leads to a poor developer experience. Managers think process automation works fine, while the reality is that developers struggle to figure out how to consume the solutions provided. Troubleshooting becomes a drag. Operations then soar up, scaling proportionally with the engineering area size.

Flimsy integrations add up to that. Due to what I consider a misinterpretation of agility, managers may see a proper integration as something that doesn't deliver value to the external customer and is not worth the investment. This means a quick win using a convoluted and error-prone script running on a hidden pipeline sounds better than a properly designed and engineered solution.

![https://tenor.com/en-GB/view/srg-bülach-weighing-things-gif-13323220](https://media.tenor.com/nRder1OJlxwAAAAd/srg-bülach.gif)

Both problems lead to the Fear, Uncertainty and Doubt (FUD) syndrome: any change to the underlying resources faces high resistance within the developer team as they see the DevOps solutions as a disconnected black box. DevOps engineers conversely resist introducing any changes due to the fragile dependencies and hard-to-test pipelines and scripts.

In the end, this kind of DevOps initiative becomes what it sought to destroy: a stale, maintenance-ridden platform solution that is slow and resistant to evolution. We need to break the misconceptions before it's too late.

![https://tenor.com/en-GB/view/send-help-help-help-me-help-me-im-poor-bad-situation-gif-15097358](https://media.tenor.com/oUI0zMDbavcAAAAC/send-help-help.gif)

I love the practice and what it can bring to the development cycle. I also love to automate a process the right way, without shoehorning some Bash or Python around. Hence why I changed my profile.

What about you? Do you have a vastly different experience when it comes to DevOps? Or perhaps you don't fully agree with something here? Please let me know in the comments, I'd love to further discuss about it.
