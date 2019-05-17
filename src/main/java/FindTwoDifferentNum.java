public class FindTwoDifferentNum {
    public String solution(int[] ints)
    {
        int ret = ints[0];
        for(int i = 1; i < ints.length; i++) {
            ret ^= ints[i];
        }
        System.out.println("第一次异或结果：" + ret);

        int bit = ret & (-ret);

        System.out.println("bit位结果:" + bit);
        int [] result = new int[2];
        for(int j = 0; j < ints.length; j++) {
            int flag = ints[j] & bit;
            if (flag == 1) {
                result[0] ^= ints[j];
            }else {
                result[1] ^= ints[j];
            }
        }

        return result[0] + " and " + result[1];
    }
}
