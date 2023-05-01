---
title: Are you a DevOps engineer?
description: how a way of working became a catch-all role
date: 2023-02-04T20:31:45+01:00
draft: true
resources:
- name: featured-image
  src: devops.jpg
categories:
- Opinion
tags:
- Career
- DevOps
---

A search for the term DevOps on LinkedIn will yield countless posts and hundreds of jobs. The market for DevOps engineers is plentiful and has been for at least a decade now. Everyone wants to join the hype and either call themselves DevOps engineers or add the DevOps keyword in their CVs for that extra oomph. Is that the right thing to do?

<!--more-->

## Prologue

I had DevOps references across my LinkedIn profile and curriculum vitae up until mid-2022. My description depicted me as a DevOps engineer looking to contribute to automation. Not anymore. The reason is simple: the market overloads and marginalizes the term. Recruiters and managers generally associate it with distorted practices and solutions I don't agree with.

## What is DevOps

I won't preach to you about the term. My understanding is that DevOps means a body of knowledge and best practices to close the gap between development and operations. It includes the reuse of tools and the art of automation to glue processes together. How to glue is up to the engineer and business, and this acts as a double-edged sword.

This definition is intentionally vague. Bodies of knowledge and processes' best practices are so. This allows great flexibility when applying them. You can still reap great results granted you preserve the main tenets. If you know agile you'll also recognize the same concept from its manifesto here. It allows small iterations over the processes to further improve them.

## Automation

We deem a system or process as automated when it requires minimal to no human intervention to make decisions. Note that this differs from _automatic_, which means a process that acts involuntarily. A car wash is automatic; a home light that turns itself on before dawn is automation.

The degree and scope of automation can vary between a small, atomic process and a bigger collection. An integration pipeline in a source code repository does not mean automation as per see. Let's say it has a step to run a lint tool and check the code. Now, regardless of the lint output, the merge requires a human to approve and execute it.

In such a scenario the linting is _automatic_, and the integration itself is manual. We'd have an _automated_ integration if the pipeline checked the lint output and automatically merged the changes. And before the four-eyes-principle priests start throwing rocks, mind you that this is a simplistic and rather unpractical example. We didn't mention testing or checking the author, nor doing any other more sophisticated checks.

## The DevOps role

If DevOps is knowledge and best practices, then the role means people that follow those right? Yes, and it does. The catch is how broad the term is: as mentioned before, it acts as a double-edged sword here. It's not clear what's the set of hard skills and technologies a DevOps must know; instead, each company/industry has its wishlist printed on a DevOps-themed page.

### Expectation

As a software engineer joining a position that mentions DevOps and automations, I expect a company looking for ways to, well, _automate_. That means investment into connecting disparate solutions and data sources in a way that certain decisions do not require human intervention.

Those automations in turn would allow improving their processes, or free up time to focus on other processes. Maintenance and operations here would translate business rule changes into new or different decisions the automation has to make.

### Reality

What happens when you focus on gluing together dozens of solutions, each tackling a specific problem domain? You risk losing touch with how the problem can be best solved, and forget about how the big picture performs. Flagship solutions become the target, not the means to an end.

This means classic system administrators and operators are at home: this kind of specialized professional knows the in and out of specific solutions or vendor suites in the market. They can - and trust me, they will - get you onboard on all sort of tools.

Such is the reality of DevOps positions. Most of the time they require low to no programming languages, listing scripting ones such as Bash, PowerShell and Python. Then focus on dozens of branded solutions you must have years of experience. Those are constantly misused and work on top of terrible integrations, which you have to provide support until the next migration exodus happens.

The day-to-day of DevOps engineers then becomes to both maintain scripts on pipelines, sort out integration problems between solutions and worst of all, act as the operators when it comes to problems. As all those integrations are black boxes, the users won't know a thing on how to solve them. Then the snowball begins.

### TODO

Such flexible definition leads to different interpretations throughout the industry. This translates into distinct practical implementations between segments and companies, which is fine. Even with some nuances, common patterns became synonymous with DevOps along the way. Teams responsible for DevOps commonly:

- Provide developers a way to consume infrastructure, either through an internal website, pipelines by them or a workflow using the company ITSM solution of choice, also known as a request for fulfillment form
- Act as the first responder on any infrastructure-related issues, even when using a cloud solution provider
- Create abstractions on top of raw offerings to restrict what developers can do due to either security concerns or a higher architectural board decision
- Document all internal solutions and provide support on their usage
- Operate requests from developers through any gaps between the front end provided to them and the real technical procedure to deliver it

Seasoned professionals will notice that those patterns were the job of infrastructure engineers or system administrators. This means DevOps is aggregating those functions and improving them, which is fine. Now as the proverb goes, the _devil is in the details_, and zooming into the patterns mentioned above to see how they are technically implemented reveals:

- **Out-of-tree infrastructure management**. Businesses mandate product developers to use Infrastructure as Code (IaC) for their solutions, yet developers need to request certain underlying resources through solutions that don't integrate with any IaC tooling. That either happens through an internal website, sending an instant message to one of the DevOps engineers, using a pipeline provided by them or creating their pipeline using templates provided
- **Operational support**. The DevOps team invests most time documenting and supporting the use of their abstractions due to their divergence from any commonly known patterns and solutions available in the market, and also bringing those same solutions up to speed due to changes in allowed options or internal use cases
- **Flimsy integrations**. The DevOps team relies on scripting languages - such as Bash, Python and PowerShell - and trampoline solutions such as pipelines calling other pipelines to deal with limitations related to security or the perceived low technical value of a refined solution. This reflects the operational support as it adds a considerable amount of time to troubleshoot when one of those moving parts fails

### Aftermath

The first problem is the low developer experience. Managers think process automation works fine, while the reality is that developers struggle to figure out how to consume the solutions provided by the DevOps team and maintain the resources they need through the same. This adds up to the time consumed by both ends during operations, further reducing the time spent on improving the entire stack. The troubleshooting of the abstraction layers also adds up to the time consumed.

The second problem is the flimsy integrations. Due to what I consider a misinterpretation of agility, managers see a proper integration as something that doesn't deliver value to the external customer and is not worth the investment. This means a quick win using a convoluted and error-prone script running on a hidden pipeline sounds better than a properly designed and engineered solution.

Both problems lead to the Fear, Uncertainty and Doubt (FUD) syndrome: any change to the underlying resources faces high resistance within the developer team as they see the DevOps solutions as a disconnected black box; it's not rare to see also developers with accumulated bad past experiences, increasing their uncertainty. DevOps engineers conversely resist introducing any changes due to the fragile dependencies and custom-tailored integrations.

In the end, this kind of DevOps initiative becomes what it sought to destroy: a stale, maintenance-ridden platform solution that is slow and resistant to evolution.

```plantuml
@startuml
actor Developer
file "Infrastructure Code" as IaC
database "Source Control" as SourceControl
actor Approver
agent Pipeline

Developer -> IaC
IaC -> SourceControl
@enduml
```

## Other roles

TODO

### Site Reliability Engineer

TODO

### Platform Engineer

TODO

### Software Engineer

TODO

## The dream

TODO
