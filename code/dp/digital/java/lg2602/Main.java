import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        StringTokenizer ab = new StringTokenizer(in.readLine());
        long a = Long.parseLong(ab.nextToken());
        long b = Long.parseLong(ab.nextToken());
        high = String.valueOf(b).toCharArray();
        n = high.length;
        String s = String.format("%013d", a);
        low = s.substring(13 - n, 13).toCharArray();
        memo = new long[n][n];
        for (d = '0'; d <= '9'; d++) {
            out.printf("%d ", countDigit());
        }
        out.flush();
        out.close();
        in.close();
    }

    static int n;
    static char[] low, high;
    static long[][] memo;
    static char d;

    private static long countDigit() {
        // memo[i][cnt]记录 f(i, cnt, false, false, false) 的结果
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                memo[i][j] = -1;
            }
        }
        return f(0, 0, true, true, true);
    }

    /**
     * 记忆化搜索
     * 
     * @param i          当前选择从左至右第几位
     * @param cnt        进入第i位之前数字d出现的次数
     * @param zero       进入第i位之前是否未选择非0数字，如果未选择则当前位可以选择0作为前缀0
     * @param limitLower 当前位是否受low对应位的限制，如果受限则当前位只能选择大于等于low[i]的字符
     * @param limitUpper 当前位是否受high对应位的限制，如果受限则当前位只能选择小于等于high[i]的字符
     */
    private static long f(int i, int cnt, boolean zero, boolean limitLower, boolean limitUpper) {
        if (i == n) {
            return cnt;
        } else if (!zero && !limitLower && !limitUpper && memo[i][cnt] != -1) {
            // 受限情况下，不能使用记忆的结果，也不能记忆其结果，因为受限情况下的结果是不完整的
            return memo[i][cnt];
        } else {
            long ans = 0;
            char lower = limitLower ? low[i] : '0';
            char upper = limitUpper ? high[i] : '9';
            for (char c = lower; c <= upper; c++) {
                boolean zo = zero && c == '0';
                boolean ll = limitLower && c == lower;
                boolean lu = limitUpper && c == upper;
                if (!zo && c == d) {
                    ans += f(i + 1, cnt + 1, zo, ll, lu);
                } else {
                    ans += f(i + 1, cnt, zo, ll, lu);
                }
            }
            if (!zero && !limitLower && !limitUpper) {
                memo[i][cnt] = ans;
            }
            return ans;
        }
    }
}
