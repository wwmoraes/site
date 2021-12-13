#!/usr/bin/env lua

local systemDefaults = {
  ["foo"] = "1",
}

systemDefaults.__index = systemDefaults

local userDefaults = {
  ["foo"] = "x",
  ["bar"] = "2",
}

userDefaults.__index = function(_, key)
  return userDefaults[key] or ""
end

setmetatable(userDefaults, systemDefaults)

local waldo = {
  ["baz"] = "3",
}

setmetatable(waldo, userDefaults)

print(waldo["foo"]) -- x
print(waldo["bar"]) -- 2
print(waldo["baz"]) -- 3
print(waldo["qux"]) -- (an empty string)
