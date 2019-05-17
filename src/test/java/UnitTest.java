import org.junit.Test;

public class UnitTest {
    @Test
    public void testLongestPalindRome() {
        LongestPalindRome palindRome = new LongestPalindRome();
        System.out.println(palindRome.findLongestPalindRome("abbad"));
    }

    @Test
    public void testFindTwoDiffNum() {
        FindTwoDifferentNum findTwoDifferentNum = new FindTwoDifferentNum();
        int[] nums = new int[]{1,2,2,1,3,7};
        System.out.println(findTwoDifferentNum.solution(nums));

        int result = 3 ^ 7;
        System.out.println("result=" + result);
    }
}
