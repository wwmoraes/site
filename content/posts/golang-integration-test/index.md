---
title: Instrumented testing in Golang
description: Say goodbye to thousands of unit test lines of code
date: 2023-08-24T18:35:53+02:00
categories:
- Code
tags:
- Golang
- Testing
lastmod: 2024-08-11T13:07:57+02:00
---

Imagine you create a Golang application, run it and get a test coverage report
from it as a bonus. No unit tests written, no hundreds of mocks/stubs nor
workarounds. Sounds too good?

<!--more-->

## Golang and tests

Golang developers have outstanding support from the tooling to test their code.
The [testing][testing] standard package and its companion [go test][go-test-cmd]
command provides a solid foundation to test code. If you're up to writing unit
test cases, that is.

I won't preach to you nor start yet another flame war about types of tests.
Suffice it to say that unit tests aren't my personal preference these days. They
lead the developer to focus on technical aspects and not the real use cases of
the code. It's also a chore to maintain when done solely for the sake of
coverage.

> [!note]
> My exception to this standing is property testing in functional languages. Even
> more so when coupled with random data generators like the `Gen` monad from
> `Test.QuickCheck` in Haskell. They may sound like unit testing, yet they ensure
> a defined law of a type holds water with randomized input.

## Integration tests

What's the next best thing then? Integration testing. That means running the
code as unmodified as possible. This may trigger some to think about running
their entire stack of dependencies. That's what some call an end-to-end test,
and while nice to have it's often not practical to do constantly or locally.

When I say _as unmodified as possible_ I mean to isolate the invariants. If your
code follows a design-driven-based principle such as clean, hexagonal, onion or
similar architecture pattern then you're set. Isolating IO drivers and replacing
them with reproducible ones should be a breeze.

## Instrumenting code

With that in place, we then arrive at the Golang 1.20 [coverage profiling
automatic instrumentation](https://go.dev/testing/coverage/). This gem allows
you to build an instrumented version of your application and run it to generate
coverage data.

In practice, you may need a separate binary to change IO drivers or make all
method/endpoint calls to cover your use cases. With that in place you can then
run:

```shell
GOCOVERDIR=coverage
export GOCOVERDIR

go run -cover ./path/to/your/app/pkg/...
go tool covdata textfmt -i="${GOCOVERDIR}" -o=gcov.txt
go tool cover -func=gcov.txt
```

You can optionally select specific packages to report coverage for with `-pkgs`.
That's useful to exclude test/telemetry packages or other binary sources that
don't belong in a coverage report.

## Results

I got an outstanding [72% coverage][coverage-sample] over my
[anilistarr][anilistarr] project on my first try using 80 lines of code. With
this new option, I don't need to imagine testing anymore; I can isolate my code
as usual and get tests, dead code reports and coverage with little effort. 🖤

What do you think of this option provided by Golang? 😊

[testing]: https://pkg.go.dev/testing
[go-test-cmd]: https://pkg.go.dev/cmd/go#hdr-Testing_functions
[anilistarr]: https://github.com/wwmoraes/anilistarr/
[coverage-sample]: https://github.com/wwmoraes/anilistarr/actions/runs/5957377534/job/16159944308#step:8:61
