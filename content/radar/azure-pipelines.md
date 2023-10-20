---
description: YAML templates and includes everywhere
radarIndex: 7
radarSection: platforms
radarTier: hold
radarX: -121
radarY: 468
table-of-contents: false
title: Azure Pipelines
---

Microsoft acquired GitHub in 2018. Since that Azure DevOps, the service which
hosts Azure Pipelines, reduced its rate of features and improvements release.
Their [feedback community][ado-feedback] is collecting dust with a variety of
highly voted, 5 years-old issues that are still open.

In other words, its an abandoned project that's in an unannounced maintenance
mode.

Functionally speaking, as Uncle Ben once told Peter Parker, with great power
there must also come great responsibility. Azure Pipelines offers template
mechanisms that allows building pipeline definitions themselves. This allows
complex and conditional logic to dynamically change the structure of a pipeline.

Misuse of features is an user rather than tool problem. Yet in practice Azure
Pipelines users rely on those build-time templates as the main solution. I've
seen entire solutions with business logic and toggles developed with such
mechanism which are overly complex, slow and cumbersome to deal with.

Programmable CI and CD pipelines and YAML templates are harmful as they lead to
a fragmentation in the local tooling vs the pipeline. What's more, custom tasks
rely on the old Visual Studio Team Services model of extensions, which aren't
easy to create or enable.

[ado-feedback]: https://developercommunity.visualstudio.com/AzureDevOps?space=21&sort=votes
