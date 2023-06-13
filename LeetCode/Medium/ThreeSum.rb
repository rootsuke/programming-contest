# @param {Integer[]} nums
# @return {Integer[][]}
def three_sum(nums)
  nums.sort!
  ans = []

  0.upto(nums.length-3) do |i|
    return ans if nums[i] > 0

    next if i > 0 && nums[i-1] == nums[i]
    
    left = i + 1
    right = nums.length - 1

    while left < right
      if nums[i] + nums[left] + nums[right] == 0
        ans << [nums[i], nums[left], nums[right]]

        while left < right && nums[left] == nums[left+1]
          left += 1
        end
        left += 1

        while left < right && nums[right-1] == nums[right]
          right -= 1
        end
        right -= 1
        
      elsif nums[i] + nums[left] + nums[right] < 0
        left += 1
      else
        right -= 1
      end
    end
  end

  return ans
end
