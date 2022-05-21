public class FindTheMiddleIndexInArray {
    public int findMiddleIndex(int[] nums) {
        int sum = 0;
        for (int num : nums) {
            sum += num;
        }
        return FindPivotIndex.pivotIndexRec(sum, 0, 0, nums);
    }
}
