# @param {String} s
# @return {Integer}
def length_of_last_word(s)
  flg = false
  ans = 0

  (s.length-1).downto(0) do |i|
    flg = s[i] == " " ? false : true

    if flg
      ans += 1
      
      return ans if i == 0 || s[i-1] == " "
    end
  end
end
