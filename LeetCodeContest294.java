import java.util.*;

public class LeetCodeContest294 {


    public static int totalStrength(int[] strength) {
        HashMap<Integer, List<Integer>> map = new HashMap<>();
        int m = 0;

        for (int i = 0; i < strength.length; i++) {

            for (int j = i; j < strength.length; j++) {
                List<Integer> list = new ArrayList<>();

                for (int k = i; k <= j; k++)
                    list.add(strength[k]);
                Collections.sort(list);
                map.put(m, list);
                m++;
            }
        }
        int result = 0;
        for (int i = 0; i < m; i++) {
            List<Integer> list = map.get(i);
            int sum = list.stream().mapToInt(Integer::intValue).sum();
            int min = list.get(0);
            result += (min * sum);
        }
        return result;
    }

    public static void main(String[] args) {
        int[] arr = {747, 812, 112, 1230, 1426, 1477, 1388, 976, 849, 1431, 1885, 1845, 1070, 1980, 280, 1075, 232, 1330, 1868, 1696, 1361, 1822, 524, 1899, 1904, 538, 731, 985, 279, 1608, 1558, 930, 1232, 1497, 875, 1850, 1173, 805, 1720, 33, 233, 330, 1429, 1688, 281, 362, 1963, 927, 1688, 256, 1594, 1823, 743, 553, 1633, 1898, 1101, 1278, 717, 522, 1926, 1451, 119, 1283, 1016, 194, 780, 1436, 1233, 710, 1608, 523, 874, 1779, 1822, 134, 1984};

//        int[] arr = {1, 3, 1, 2};
        System.out.println(totalStrength(arr));
//        System.out.println(percentageLetter("foobar", 'o'));
//        int[] arr1 = {54,18,91,49,51,45,58,54,47,91,90,20,85,20,90,49,10,84,59,29,40,9,100,1,64,71,30,46,91};
//        int[] arr2 = {14,13,16,44,8,20,51,15,46,76,51,20,77,13,14,35,6,34,34,13,3,8,1,1,61,5,2,15,18};
//        System.out.println(maximumBags(arr1, arr2, 77));

    }

    public static int maximumBags(int[] capacity, int[] rocks, int additionalRocks) {
//        for (int i=0;i<rocks.length-1;++i){
//
//            for(int j=0;j<rocks.length-i-1; ++j){
//
//                if(capacity[j+1]<capacity[j]){
//
//                    int swap = capacity[j];
//                    capacity[j] = capacity[j+1];
//                    capacity[j+1] = swap;
//
//                    int m = rocks[j];
//                    rocks[j] = rocks[j+1];
//                    rocks[j+1] = m;
//                }
//            }
//        }

        for (int i = 0; i < capacity.length; i++) {
            if (additionalRocks == 0) {
                break;
            }
            if (capacity[i] != rocks[i] && additionalRocks > 0 && additionalRocks >= capacity[i] - rocks[i]) {
                additionalRocks = additionalRocks - (capacity[i] - rocks[i]);

                rocks[i] = capacity[i];
            }
        }
        int k = 0;
        for (int i = 0; i < capacity.length; i++) {
            if (capacity[i] == rocks[i]) {
                k++;
            }
        }
        System.out.println(Arrays.toString(capacity));
        System.out.println(Arrays.toString(rocks));
        return k;
    }

    public static int percentageLetter(String s, char letter) {
        int m = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == letter) {
                m++;
            }
        }
        double a = m;
        double k = s.length();
        double n = a / k * 100;
        return (int) n;
    }
}
