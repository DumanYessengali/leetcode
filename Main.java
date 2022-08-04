public class Main {
    public static void main(String[] args) {
        ListNode list = new ListNode(3);
        list = ListNode.listPushBack(list, 2);
        list = ListNode.listPushBack(list, 0);
        list = ListNode.listPushBack(list, -4);
        LinkedListCycle l = new LinkedListCycle();
        System.out.println(l.hasCycle(list));
    }
}
