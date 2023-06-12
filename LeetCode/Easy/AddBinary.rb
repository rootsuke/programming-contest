# @param {String} a
# @param {String} b
# @return {String}
def add_binary(a, b)
  i = a.length - 1
  j = b.length - 1

  carry = 0
  ans = ""

  while i >= 0 || j >= 0 || carry == 1
    x = i < 0 ? 0 : a[i].to_i
    y = j < 0 ? 0 : b[j].to_i
    
    bit = (x + y + carry) % 2
    carry = (x + y + carry) / 2

    ans.insert(0, bit.to_s)

    i -= 1 if i >= 0
    j -= 1 if j >= 0
  end

  return ans
end
