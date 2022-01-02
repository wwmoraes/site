#!/usr/bin/env lua

local Queue = {}

Queue.__index = Queue

function Queue.new(target)
  local target = target or {}
  setmetatable(target, Queue)
  return target
end

function Queue:push(value)
  table.insert(self, value)
end

function Queue:pop()
  return table.remove(self, 1)
end

function Queue:length()
  return #self
end

local jobs = Queue.new()

print(jobs:length()) -- 0
print(jobs:pop()) -- nil

jobs:push("foo")

print(jobs:length()) -- 1

jobs:push("bar")

print(jobs:length()) -- 2
print(jobs:pop()) -- foo
print(jobs:length()) -- 1

jobs:push("qux")

print(jobs:length()) -- 2

-- jobs is still a plain table
for key,value in ipairs(jobs) do
  print(string.format("%s: %s", key, value))
end
-- output
-- 1: bar
-- 2: qux
