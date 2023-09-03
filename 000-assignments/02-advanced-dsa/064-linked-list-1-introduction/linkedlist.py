class LinkedList:
    val = None
    next = None

    def __init__(self, val):
        self.val = val

    def append(self, val):
        if self.next is None:
            self.next = LinkedList(val)
        else:
            self.next.append(val)

    def __str__(self):
        return str(self.val) + " -> " + str(self.next)

    def search(self, key):
        if self == None:
            return False

        temp = self
        while temp.next != None:
            if temp.val == key:
                return True
            temp = temp.next
        return False

    def insert(self, val, pos):
        newNode = LinkedList(val)
        if pos == 0:
            newNode.next = self
            return newNode

        temp = self
        index = 1
        while temp.next != None:
            if pos == index:
                newNode.next = temp.next
                temp.next = newNode
                return self
            temp = temp.next
            index += 1

        temp.next = newNode
        return self

    def delete(self, val):
        if self.val == val:
            self = self.next
            return self

        temp = self
        while temp.next != None:
            if temp.next.val == val:
                temp.next = temp.next.next
                return self
            temp = temp.next
        return self

    def reverse(self):
        prev = None
        curr = self
        while curr != None:
            forw = curr.next
            curr.next = prev
            prev = curr
            curr = forw

        self = prev
        return self

    def is_palindrome(self):
        p1 = self
        p2 = self

        n = 0
        temp = self
        while temp != None:
            n += 1
            temp = temp.next

        mid = n // 2
        temp = self
        while mid > 0:
            p2 = p2.next
            mid -= 1
        print(p1, p2)

        p2 = p2.reverse()
        print(p1, p2)

        while p2 != None:
            if p1.val != p2.val:
                return False
            p1 = p1.next
            p2 = p2.next

        return True


if __name__ == "__main__":
    ll = LinkedList(1)
    ll.append(2)
    ll.append(3)
    ll.append(4)
    ll.append(5)
    print(ll)
    # 1 -> 2 -> 3 -> 4 -> 5 -> None

    print("3 is present? ", ll.search(3))
    # True
    print("6 is present? ", ll.search(6))
    # False

    ll = ll.insert(0, 0)
    print(ll)

    ll = ll.insert(6, 6)
    print(ll)

    print(ll)
    ll = ll.insert(3, 3)
    print(ll)

    ll = ll.delete(0)
    print(ll)

    ll = ll.delete(6)
    print(ll)

    ll = ll.delete(3)
    print(ll)

    ll = ll.reverse()
    print(ll)

    ll = ll.reverse()
    print(ll)

    print(ll.is_palindrome())

    ll = LinkedList(1)
    ll.append(2)
    ll.append(3)
    ll.append(2)
    ll.append(1)
    print(ll)

    print(ll.is_palindrome())

    ll = LinkedList(1)
    ll.append(2)
    ll.append(3)
    ll.append(3)
    ll.append(2)
    ll.append(1)
    print(ll)

    print(ll.is_palindrome())
