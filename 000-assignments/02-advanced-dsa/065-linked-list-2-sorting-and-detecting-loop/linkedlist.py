class linked_list:
    val = None
    next = None

    def __init__(self, val):
        self.val = val

    def append(self, val):
        if self.next is None:
            self.next = linked_list(val)
        else:
            self.next.append(val)

    def __str__(self):
        return str(self.val) + " -> " + str(self.next)

    def mid(self):
        if self == None:
            return self

        slow = self
        fast = self

        while fast.next != None and fast.next.next != None:
            slow = slow.next
            fast = fast.next.next

        return slow

    def merge(self, ll_2):
        if self == None:
            return ll_2
        if ll_2 == None:
            return self

        head = None
        if self.val < ll_2.val:
            head = self
            self = self.next
        else:
            head = ll_2
            ll_2 = ll_2.next

        temp = head
        while self != None and ll_2 != None:
            if self.val < ll_2.val:
                temp.next = self
                self = self.next
            else:
                temp.next = ll_2
                ll_2 = ll_2.next
            temp = temp.next

        if self != None:
            temp.next = self
        if ll_2 != None:
            temp.next = ll_2

        return head

    def merge_sort(self):
        if self == None or self.next == None:
            return self

        mid = self.mid()
        next_to_mid = mid.next
        mid.next = None

        left = self.merge_sort()
        right = next_to_mid.merge_sort()

        return left.merge(right)

    def detect_loop(self):
        slow = self
        fast = self
        while fast != None and fast.next != None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                return True
        return False

    def remove_loop(self):
        slow = self
        fast = self
        while fast != None and fast.next != None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                break

        if slow == fast:
            slow = self
            while slow.next != fast.next:
                slow = slow.next
                fast = fast.next
            fast.next = None

        return self


if __name__ == "__main__":
    ll = linked_list(1)
    ll.append(2)
    ll.append(3)
    ll.append(4)
    ll.append(5)
    print(ll)

    print(ll.mid().val)

    ll = linked_list(1)
    ll.append(2)
    ll.append(3)
    ll.append(4)
    ll.append(5)
    ll.append(6)
    print(ll)

    print(ll.mid().val)

    h1 = linked_list(1)
    h1.append(3)
    h1.append(5)
    h1.append(7)
    h1.append(9)
    print(h1)

    h2 = linked_list(2)
    h2.append(4)
    h2.append(6)
    h2.append(8)
    h2.append(10)
    print(h2)

    print(h1.merge(h2))

    ll = linked_list(5)
    ll.append(4)
    ll.append(3)
    ll.append(2)
    ll.append(1)
    print(ll)

    print(ll.merge_sort())

    ll = linked_list(1)
    ll.append(2)
    ll.append(3)
    ll.append(4)
    ll.append(5)

    print(ll.detect_loop())

    ll.next.next.next.next.next = ll.next
    print(ll.detect_loop())

    ll = linked_list(1)
    ll.append(2)
    ll.append(3)
    ll.append(4)
    ll.append(5)

    ll.next.next.next.next.next = ll.next
    print(ll.detect_loop())
    print(ll.remove_loop())
    print(ll.detect_loop())
