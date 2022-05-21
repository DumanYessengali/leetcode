public class FindPivotIndex {
    public int pivotIndex(int[] nums) {
        int sum = 0;
        for (int num : nums) {
            sum += num;
        }
        return pivotIndexRec(sum,0,0,nums);
    }

    public static int pivotIndexRec(int sum, int leftSum, int i, int[] nums) {
        if (i < nums.length) {
            sum -= nums[i];
            if (sum == leftSum) {
                return i;
            }
            leftSum += nums[i];
        }
        if (i == nums.length-1){
            return -1;
        }
        return pivotIndexRec(sum, leftSum, ++i, nums);
    }
}
