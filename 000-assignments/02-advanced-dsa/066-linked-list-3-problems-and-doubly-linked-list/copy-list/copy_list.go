package copylist

type RandomListNode struct {
	value  int
	next   *RandomListNode
	random *RandomListNode
}

func CopyList(head *RandomListNode) *RandomListNode {
    if head == nil {
        return nil
    }

    hc := head
    for hc != nil {
        //Insert value at alternate place
        nextNode := hc.next
        hc.next = &RandomListNode{hc.value, nextNode, nil}
        hc = nextNode
    }

    //attach random pointer of deep LL accroding to original LL
    temp1 := head
    for temp1 != nil {
        //Edge case as every node dont have random pointer
        if temp1.random != nil {
            temp1.next.random = temp1.random.next
        }

        temp1 = temp1.next.next
    }

    //detach the original LL
    h1 := head
    h2 := head.next
    newHead := head.next
    for h1 != nil {
        h1.next = h1.next.next
        if h2.next != nil {
            h2.next = h2.next.next
        }
        h1 = h1.next
        h2 = h2.next
    }
    return newHead

}