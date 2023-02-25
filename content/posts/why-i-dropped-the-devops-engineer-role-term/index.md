---
date: 2023-02-04T20:31:45+01:00
draft: true
resources:
- name: featured-image
  src: featured-image.png
title: Why I dropped the DevOps Engineer role term
categories:
- career
tags:
- DevOps
description: why filtering keywords matter to attract the desired jobs
---

A search for the term DevOps on LinkedIn will yield hundreds of thousands of posts and hundreds of jobs. The market for so-called DevOps engineers is plentiful and has been for over a decade now. Why would someone want to dissociate themselves from it? Seems illogical, I know.

<!--more-->

## Prologue

I've removed most of the DevOps references from my profile and curriculum vitae since mid-2022. I used to describe myself as a DevOps engineer or looking forward to contributing to the business as one. The reason is simple: the market overloads and marginalizes the term. Recruiters and managers associate it with practices and solutions I don't agree with.

Let's touch on my understanding of what DevOps is, and what being a DevOps engineer means. Then I'll share my personal experience working as a DevOps engineer, and why I'm moving away from such a designation. This will help you understand where my opinion comes from, and form your own.

## DevOps

### Definition

TODO

The definitions above are vague, and that's intentional. Bodies of knowledge and processes' best practices are so. This allows great flexibility when applying them. You can still reap great results granted you preserve the main tenets. To me, this is a subset of the agile movement embedded into the DevOps concepts. It allows small iterations over the processes to further improve them.

### Role

TODO

### Expectation

TODO

### Reality

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

## Alternatives

TODO

### Site Reliability Engineer

TODO

### Platform Engineer

TODO

### Software Engineer

TODO
