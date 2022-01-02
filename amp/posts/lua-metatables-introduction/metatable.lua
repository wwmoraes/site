#!/usr/bin/env lua

local meta = {}

local waldo = {}

print(getmetatable(meta)) -- nil
print(getmetatable(waldo)) -- nil

setmetatable(waldo, meta) -- set meta as the metatable of waldo

print(getmetatable(meta)) -- nil
print(getmetatable(waldo)) -- table: 0x...
print(getmetatable(waldo) == meta) -- true
