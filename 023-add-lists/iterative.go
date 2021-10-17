package addlists

func Iterative(head1 *Node, head2 *Node) *Node {
	var (
		carry int
		head  *Node
		prev  *Node
	)

	for (head1 != nil) || (head2 != nil) || (carry > 0) {
		var (
			v1 int
			v2 int
		)

		if head1 != nil {
			v1 = head1.Value
		}

		if head2 != nil {
			v2 = head2.Value
		}

		total := v1 + v2 + carry

		//nolint:gomnd
		newNode := NewNode(total % 10)
		//nolint:gomnd
		carry = total / 10

		if head == nil {
			head = newNode
		}

		if prev != nil {
			prev.Next = newNode
		}

		if head1 != nil {
			head1 = head1.Next
		}

		if head2 != nil {
			head2 = head2.Next
		}

		prev = newNode
	}

	return head
}
