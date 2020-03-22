/**
 * Example 1:
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 * Example 2:
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 * we will return 0 when needle is an empty string.
 */
public class ImplementStr {
    /* Library Method*/
    public int strStr(String haystack, String needle) {
        return haystack.indexOf(needle) >= 0 ? haystack.indexOf(needle) : -1;
    }

    /* Native Method*/
    public int indexOf(String sourceStr, String targetStr) {
        if(sourceStr.length() < 1 || sourceStr.length() < targetStr.length()) {
            return -1;
        }

        if(targetStr.length() < 1) {
            return 0;
        }

        char first = targetStr.charAt(0);
        int end = sourceStr.length() - targetStr.length() + 1;

        for (int i = 0; i < end; i++) {
            if (sourceStr.charAt(i) != first) {
                while(++i < end && sourceStr.charAt(i) != first);
            }

            if (i < end) {
                int j = i + 1;
                for (int k = 1 ; j < i + targetStr.length(); j++, k++) {
                    if (sourceStr.charAt(j) != targetStr.charAt(k)) {
                        break;
                    }
                }

                if(j == i + targetStr.length()) {
                    return i;
                }
            }
        }
        return -1;
    }
}
