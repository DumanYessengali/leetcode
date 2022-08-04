public class LinkedListCycle {
    public boolean hasCycle(ListNode head) {
        if (head==null || head.next==null) return false;
        ListNode slow = head;
        ListNode fast = head;

        while(fast != null && fast.next != null){
            if(fast.next.equals(slow)){
                return true;
            }
            fast = fast.next.next;
            slow = slow.next;
        }

        return false;
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
        next = null;
    }

    public static ListNode listPushBack(ListNode node, int data) {
        ListNode list = new ListNode(data);
        if (node == null) return list;
        ListNode N  =node;
        while(N.next!=null) {
            N= N.next;
        }
        N.next = list;
        return node;
    }
}