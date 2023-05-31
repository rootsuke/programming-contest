# @param {String[]} strs
# @return {String}
def longest_common_prefix(strs)
  strs.sort!

  first = strs.first
  last = strs.last

  ans = ""

  (0...[first.length, last.length].min).each do |i|
    if first[i] != last[i]
      return ans
    end
  end

  return ans
end
