# @param {Integer} x
# @return {Boolean}
def is_palindrome(x)
  return false if x < 0

  reversed = x.to_s.reverse
  x.to_s == reversed
end
