---
title: ESO addon and Lua fun
description: LOL
date: 2021-12-13T09:38:13+01:00
draft: true
categories:
- gaming
- scripting
tags:
- lua
---

## Prologue: My moon cycle

{{< admonition type=note >}}
you can [skip ahead]({{< relref "#saved-variables-proxy" >}}) to the technical
stuff if storytelling is not your thing 😄
{{< /admonition >}}

I've had no real use for Lua in the past, despite my ~~frustrated~~ attempts.
The last one was ages ago, with Trunk Notes, a now-defunct iOS notes app. It had
wiki features and support for creating Lua snippets to further customize the
notes, akin to what TiddlyWiki does. I loved the app and the concept, even with
its limited API and library access.

Fast-forward to 2020, and I've started looking for better ways to automate my
workflows on MacOS. AppleScript is awesome in what it provides, however the
language and tooling consumes a lot of time testing and troubleshooting. Even
JXA JavaScript for Automation is way to cryptic to be used consistently.

Then I've found [Hammerspoon][hs], a beautiful MacOS app that exposes a ton of
OS APIs through easy-to-use Lua objects. I'll write a bit more about it on
another post. Feel free to jump straight into [my code][hs-init] and/or
[_spoon_ plugins][spoons] I've created so far if you already know about it.

[hs]: http://www.hammerspoon.org
[hs-init]: https://github.com/wwmoraes/dotfiles/tree/master/.systems/darwin/hammerspoon/.hammerspoon
[spoons]: https://github.com/wwmoraes/spoons

Workflow automations aside, I've started playing The Elder Scrolls Online. Apart
from the game story and gameplay being awesome, Zenimax Online Studios decided
to use Lua for their internal API, and also expose part of it so players can
create add-ons. They even [open-sourced the UI][esoui] together with a raw API
documentation, and update it regularly.

[esoui]: https://github.com/esoui/esoui/

Despite the community being quite advanced on
[add-on and library offerings][eso-addons], I still decided to try and create
something both for fun, but also functional enough to justify the effort.

[eso-addons]: https://www.esoui.com/addons.php

### ESO and Lua

I won't deep dive on the ESO API as that's not the focus of this post. Suffice
to say that most of it is accessible either through global objects and events.
Most addon interactions are done in a responsive way, by listening for
server-emitted events. There's events from combat actions to bank interactions,
inventory and even when you try to use a skill on an action slot with the wrong
weapon equipped. Some of the more advanced addons available completely replace
the standard UI and change the way you interact with features such as inventory
and map. It's simply beautiful.

### Use case: auto status addon

One interesting game feature is your player status. You can set yourself as
either online, away, do-not-disturb or offline. By default, you're online
whenever you log-in, and stay as online throughout your playtime. You can
manually switch your state through the friends list panel.

TODO INSERT IMAGE

The first thing I've missed is a set of slash commands to switch status. MMOs
usually sport dozens of slash commands to easily change configurations, switch
chat channels, whisper to friends and gesture. The second one was the lack of
use for the other statuses besides online and offline. So I decided to create
both the missing commands, and automate status changes depending on the player
activity:

* when player logs in/right before log out
* when they enter/exit hostile environments (such as dungeons/trials/PvP)
* when changing window focus (e.g. switched to a IM/browser app)

I also want to provide the user options to decide which of those events they
want to trigger a status change, and also which status they want to switch to.

### addons and persistence

Addons on ESO can save state and preferences through an officially provided API
that returns an empty table for you to use as you see fit. This table is then
_automagically_ dumped into a Lua file for your addon, and you can load it back
through the same API.

These saved variables, as they are known as, can be either character-specific or
account-wide. Most addons even provide a toggle for users to choose how they
want to persist their settings. However, checking everywhere on your code
whether the user wants to use character or account-wide is not efficient. That's
a perfect use-case for metatables!

## Saved variables proxy

TODO

## Wrap-up

TODO
