#!/usr/bin/env lua

local defaults = {
  ["bar"] = true
}

defaults.__index = defaults

local waldo = {}

setmetatable(waldo, defaults)

print(waldo["bar"]) -- true

for key,value in pairs(waldo) do
  print(string.format("%s: %s", key, value))
end
-- no output

waldo.bar = false

for key,value in pairs(waldo) do
  print(string.format("%s => %s", key, value))
end
-- output:
-- bar => false

print(waldo["bar"]) -- false
