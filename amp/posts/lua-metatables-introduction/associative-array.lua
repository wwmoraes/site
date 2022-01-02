#!/usr/bin/env lua

local waldo = {}

waldo[2] = "foo"
waldo["bar"] = true

for key,value in pairs(waldo) do
  print(string.format("%s: %s", key, value))
end
-- output
-- 2: foo
-- bar: true
