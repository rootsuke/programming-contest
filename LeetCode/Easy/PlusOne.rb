def plus_one(digits)
	lastIndex = digits.length - 1
	carry = 1
	ans = []

	lastIndex.downto(0) do |i|
		digit = (digits[i] + carry) % 10
		ans = ans.unshift(digit)

		carry = (digits[i] + carry) / 10

    if i == 0 && carry == 1
      ans = ans.unshift(carry)
    end
	end

	return ans
end
