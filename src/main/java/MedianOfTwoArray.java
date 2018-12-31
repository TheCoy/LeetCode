public class MedianOfTwoArray
{
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int i = 0, j = 0;
        int len1 = nums1.length;
        int len2 = nums2.length;

        int mid = (nums1.length + nums2.length) / 2;
        double[] result = new double[len1 + len2];
        int k = 0;

        while (i < len1 && j < len2 && k <= mid) {
            if (nums1[i] < nums2[j]) {
                result[k++] = nums1[i++];
            }else {
                result[k++] = nums2[j++];
            }
        }

        if (i < len1 && k <= mid) {
            while (i < len1) {
                result[k++] = nums1[i++];
            }
        }else if (j < len2 && k <= mid){
            while (j < len2) {
                result[k++] = nums2[j++];
            }

        }

        if ((len1 + len2) % 2 == 0) {
            return (double) (result[mid] + result[mid-1]) / 2;
        }else {
            return result[mid];
        }
    }
}