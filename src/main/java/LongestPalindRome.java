public class LongestPalindRome {
    public String findLongestPalindRome(String s) {
        if (s == null || s.length() < 1) {
            return "";
        }

        int start = 0, end = 0;
        for (int i = 0; i < s.length(); i++) {
            int len1 = expandString(s, i, i);
            int len2 = expandString(s, i, i + 1);
            int len = Math.max(len1, len2);
            if (len > (end - start)) {
                start = i - (len - 1) / 2;
                end   = i + len / 2;
            }
        }

        //System.out.format(" start = %d，end =%d", start, end);

        return s.substring(start, end + 1);
    }

    private int expandString(String s, int i, int j) {
        int left = i, right = j;
        while(left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            left--;
            right++;
        }
        return right - left - 1;
    }
}
