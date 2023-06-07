# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end
# @param {ListNode} l1
# @param {ListNode} l2
# @return {ListNode}
def add_two_numbers(l1, l2)
  carry = 0
  ans = make_node(l1, l2, carry)
end

def make_node(l1, l2, carry)
  if l1.nil? && l2.nil? && carry == 0
    return nil
  end
  
  a = l1&.val || 0
  b = l2&.val || 0
  sum = a + b + carry
  carry = sum / 10
  digit = sum % 10

  l1 = l1&.next || nil
  l2 = l2&.next || nil
  next_node = make_node(l1, l2, carry)
  ListNode.new(digit, next_node)
end
