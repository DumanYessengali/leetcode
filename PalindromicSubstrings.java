import java.util.HashSet;

public class PalindromicSubstrings {
    public int countSubstrings(String s) {
        int count = 0;
        for (int i = 0; i < s.length(); i++) {
            StringBuilder subStr = new StringBuilder();

            for (int j = i; j < s.length(); j++) {
                subStr.append(s.charAt(j));
                if (isPalindrome(subStr)) {
                    count++;
                }
            }

        }
        return count;
    }

    public boolean isPalindrome(StringBuilder str) {
        int i = 0;
        int j = str.length() - 1;
        while (i <= j) {
            if (str.charAt(i) != str.charAt(j)) {
                return false;
            }
            i++;
            j--;
        }
        return true;
    }
}
