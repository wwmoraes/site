---
title: Lua metatables introduction
description: what are they and how do they work?
date: 2022-01-02T14:21:06+01:00
resources:
- name: featured-image
  src: featured-image.png
categories:
- coding
tags:
- lua
- automation
- gaming
---

If you ever used Lua, then you heard about its tables and the metatable feature.
It is a mix of a powerful, simple and yet confusing mechanism for newcomers.
Let's walk through how they work and some examples of how useful they can be.
You can also find real use cases by checking my [Hammerspoon spoons][spoons].

<!--more-->

## Definitions

Lua has, beyond its primitive types, only one object-like type called table.
Tables are associative arrays that can have any other type as key (except nil)
and value. The runtime even uses them to load packages. That means a call
to `io["open"]()` works as the usual `io.open()`. Here's a sample of how to
use tables:

{{< source "associative-array.lua" >}}

You can also use them as conventional arrays:

{{< source "conventional-array.lua" >}}

{{< admonition note >}}
Lua arrays are 1-indexed. I.e. they start on index one instead of zero like some
common programming languages.
{{< /admonition >}}

### Metatables

Metatables are tables assigned as a sort of controller for another table. We
can describe them as a class akin to how classes work in other object-oriented
languages. Another analogy is that meta tables act as a proxy of the target
table.

You can assign a metatable to another table using the `setmetatable`:

{{< source "metatable.lua" >}}

Neat! Both `setmetatable` and `getmetatable` are standard library functions to
either apply or remove a metatable from another table, respectively. You can
use both tables in the previous example as usual, given the metatable has no
metamethods. A metatable without metamethods is an ordinary table. It does not
change the behavior of other tables in any way. Thus metatables need at least
one metamethod to do something useful.

### Metamethods

A metamethod is a value - usually a function or table - assigned to reserved
key names on a metatable. The runtime looks them up when executing a certain
operation on a table. They all start with two underscore characters `__`, so
they're easy to spot.

A good example of a metamethod is the `__index`. The default behavior of a table
is to return `nil` if a requested key is not set:

{{< source "table-index.lua" >}}

The runtime flow is:

{{< diagram "table-index.svg" >}}

When an index is not found on the target table, it falls back to its metatable's
`__index` metamethod. If present, Lua uses it to return a value instead of
`nil`. That means you can create a custom logic for unknown index retrieval:

{{< source "table-meta-index.lua" >}}

What Lua tries to do in this case is:

{{< diagram "table-meta-index.svg" >}}

The flow above is by no means authoritative, as I left out a few more nuances to
this logic to keep it simple.[^metaindex] It does cover the most common use
cases though, so let's keep it that way.

{{< admonition note >}}
I've written `__index(tbl,key)` instead of `__index(table,key)` to prevent
confusing the `waldo` table instance (or any metatable looked up on this
loop) with the standard library `table` global object.
{{< /admonition >}}

[^metaindex]: The [documentation][docs-meta] states that any value that can
resolve the `__index` metamethod works. This means userdata (C objects) can also
act as a metatable.

## Use cases

Here's some simple use cases to get used to how metatables work.

### Default values

As the `__index` meta method can be a table, it may serve default values. This
makes it dead simple to prevent checking/hard-coding alternative values
everywhere:

{{< source "default-values.lua" >}}

### Fallback values

Like the default values case, you can use a function as the `__index` metamethod
to return a fallback value. This works for any key, not only those set on the
metatable. This allows you to use many fallback tables, or always return an
specific value type. This again contributes to a DRY code:

{{< source "fallback-values.lua" >}}

### Queue

The metatable `__index` together with the colon operator allows the
implementation of methods. These resemble object-oriented methods thanks to
Lua's syntactic sugar. You can then create a model-controller pattern between a
table and its metatable:

{{< source "queue.lua" >}}

Beautiful. I'll dive into the colon operator on a separate post later on.

{{< admonition note >}}
The example above is a rather naive implementation to keep it simple. It doesn't
prevent direct CRUD operations, non-numerical or out-of-order keys.

Another reason is that Lua by design is extensible. That means it doesn't
provide as many ways to lock down tables as it offers to extend them. That
doesn't mean you can't make things harder to tamper. You can set the `__newindex`
and `__index` functions to block assignment and retrieval. In this case, both
`push` and `pull` methods must use rawset and rawget to bypass them.

If you have a strong need to enforce data integrity then you should consider
using an userdata value. These are defined using the C API.
{{< /admonition >}}

## Conclusion

I've only scratched the surface of Lua's metatables.
There's [dozens of supported metamethods][docs-meta].
Yet the indexing ones are more than enough to show the power of this mechanism.
They are the main drive that makes Lua excellent to extend and prototype fast.
This alone explains why the gaming industry has a wide use for Lua.

{{< admonition edit >}}
Updated diagrams, as I don't use Mermaid anymore. In fact, I used it only in this blog while using PlantUML and Mingrammer's diagrams everywhere else. Now they're all consistent.
{{< /admonition >}}

[spoons]: https://github.com/wwmoraes/spoons
[docs-meta]: https://www.lua.org/manual/5.4/manual.html#2.4
