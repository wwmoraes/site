#!/usr/bin/env lua

local function tprint (tbl, indent)
  if not indent then indent = 0 end
  local toprint = string.rep(" ", indent) .. "{\r\n"
  indent = indent + 2
  for k, v in pairs(tbl) do
    toprint = toprint .. string.rep(" ", indent)
    if (type(k) == "number") then
      toprint = toprint .. "[" .. k .. "] = "
    elseif (type(k) == "string") then
      toprint = toprint  .. k ..  "= "
    end
    if (type(v) == "number") then
      toprint = toprint .. v .. ",\r\n"
    elseif (type(v) == "string") then
      toprint = toprint .. "\"" .. v .. "\",\r\n"
    elseif (type(v) == "table") then
      toprint = toprint .. tprint(v, indent + 2) .. ",\r\n"
    else
      toprint = toprint .. "\"" .. tostring(v) .. "\",\r\n"
    end
  end
  toprint = toprint .. string.rep(" ", indent-2) .. "}"
  return toprint
end

--- @generic T1 : table, T2 : table
--- @param t1 T1
--- @param t2 T2
--- @return T1|T2
function table.merge(t1,t2)
  for key, value in pairs(t2) do
    if type(value) == "table" and type(t1[key] or false) == "table" then
      table.merge(t1[key], t2[key])
    else
      t1[key] = value
    end
  end
  return t1
end

function table.mergeLeft(...)
  local target = select(1, ...)
  for i = 2, select("#", ...), 1 do
    table.merge(target, select(i, ...))
  end
  return target
end

function table.mergeRight(...)
  local target = select(select("#", ...), ...)
  for i = select("#", ...)-1, 1, -1 do
    table.merge(target, select(i, ...))
  end
  return target
end

local PreferencesManager = {}

function PreferencesManager:IsAccountWide()
  return self.active == self.account
end

function PreferencesManager:SetAccountWide(toggle)
  self.character.accountWide = toggle
  if toggle == true then
    self.active = self.account
  else
    self.active = self.character
  end
end

function PreferencesManager:__index(key)
  if PreferencesManager[key] then
    return PreferencesManager[key]
  end

  return self.active[key]
end

function PreferencesManager:__newindex(key, value)
  -- set both if this is a new variable
  if self.account[key] == nil and self.character[key] == nil then
    self.account[key] = value
    self.character[key] = value
  else
    self.active[key] = value
  end
end

function PreferencesManager.New(account,character)
  local account = account or {}
  local character = character or {}
  local preferences = {
    account = table.mergeLeft({}, account),
    character = table.mergeLeft({}, account, { accountWide = false }, character),
  }

  if preferences.character.accountWide == true then
    preferences.active = preferences.account
  else
    preferences.active = preferences.character
  end

  setmetatable(preferences, PreferencesManager)

  return preferences
end

local preferences = PreferencesManager.New({
  foo = 0,
  sub = {
    val = true,
  },
})

print("account wide:", preferences:IsAccountWide()) -- false
print("foo:", preferences.foo) -- 0
print("bar:", preferences.bar) -- nil

print("foo = 1")
preferences.foo = 1
print("foo:", preferences.foo) -- 1

print("bar = 3")
preferences.bar = 3
print("bar:", preferences.bar) -- 3

preferences:SetAccountWide(true)
print("account wide:", preferences:IsAccountWide()) -- true
print("foo:", preferences.foo) -- 0
print("bar:", preferences.bar) -- 3

print("foo = 2")
preferences.foo = 2
print("foo:", preferences.foo) -- 2

preferences:SetAccountWide(false)
print("account wide:", preferences:IsAccountWide()) -- false
print("foo:", preferences.foo) -- 1
print("bar:", preferences.bar) -- 3

print(tprint(preferences))
