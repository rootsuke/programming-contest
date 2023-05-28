# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  completion = {}
  nums.each_with_index do |n, i|
    if completion[n]
      return completion[n], i
    end

    key = target - n
    completion[key] = i
  end
end
