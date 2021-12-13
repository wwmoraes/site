#!/usr/bin/env lua

local waldo = {"qux", "quz"}

print(waldo[1]) -- qux

for key,value in ipairs(waldo) do
  print(string.format("%s: %s", key, value))
end

-- output
-- 1: qux
-- 2: quz
