#!/usr/bin/env lua

local meta = {
  bar = 2,
}

meta.__index = function(tbl, key)
  return meta[key] or "unset"
end

local waldo = {
  foo = 1,
}

setmetatable(waldo, meta)

print(waldo.foo) -- 1
print(waldo.bar) -- 2
print(waldo.qux) -- unset
