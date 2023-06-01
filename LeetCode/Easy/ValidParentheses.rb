# @param {String} s
# @return {Boolean}
def is_valid(s)
  stack = []

  s.each_char do |v|
    case v
    when "(", "[", "{"
      stack << v
    when ")"
      return false unless stack.pop == "("
    when "]"
      return false unless stack.pop == "["
    when "}"
      return false unless stack.pop == "{"
    else
      raise StandardError
    end
  end

  stack.empty?
end
