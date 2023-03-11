class Solution {
    public int countTriplets(int[] nums) {
        int[] cnt = new int[1 << 16];
        for (int x : nums) {
            for (int y : nums) {
                cnt[x & y]++;
            }
        }
        int ans = 0;
        for (int a = 0; a < 1 << 16; a++) {
            for (int z : nums) {
                if ((a & z) == 0) {
                    ans += cnt[a];
                }
            }
        }
        return ans;
    }
}
