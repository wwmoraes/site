#!/usr/bin/env lua

local t2 = {
  ["bar"] = "qux",
  ["baz"] = "quz",
}

function t2.Bar(tbl)
  return tbl.bar
end

function t2:Baz()
  -- self is implicitly added as the first parameter by the runtime
  return self.baz
end

print(t2.Bar(t2)) -- qux
print(t2:Bar()) -- qux
print(t2.Baz(t2)) -- quz
print(t2:Baz()) -- quz
