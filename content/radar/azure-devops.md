---
description: When will Microsoft sunset it to focus on GitHub?
radarIndex: 6
radarSection: platforms
radarTier: hold
radarX: -362
radarY: 299
table-of-contents: false
title: Azure DevOps
---

Azure DevOps is the spiritual successor of the old Team Foundation Services.
Its the favorite solution of Microsoft-heavy companies and teams to manage
boards, repositories and pipelines. It shines on its wiki feature, which uses a
git repository + markdown files for all pages while providing a WYSIWYG editor
directly in the browser.

Their approach for other features is confusing to say the least. Components
such as pipelines and builds have their own sections independently of the
source repository. Its the opposite of similar tools like GitHub Actions or
GitLab where the repository has a clear listing of all pipelines and builds in
its scope.

Unfortunately the entire kit is poorly glued together, and in practice I've
seen companies using it in parallel with JIRA, which is a better option for
project management anyway.
