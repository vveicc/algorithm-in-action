import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        List<Integer> chairs = new ArrayList<>();
        List<Integer> peoples = new ArrayList<>();
        int n = Integer.parseInt(in.readLine());
        StringTokenizer tokenizer = new StringTokenizer(in.readLine());
        for (int i = 0; i < n; i++) {
            (tokenizer.nextToken().charAt(0) == '0' ? chairs : peoples).add(i);
        }
        n = peoples.size();
        int[] dp = new int[n + 1];
        for (int j = 1; j <= n; j++) {
            dp[j] = (int) 1e8;
        }
        for (int i = 0, m = chairs.size(); i < m; i++) {
            for (int j = Math.min(n - 1, i); j >= 0; j--) {
                dp[j + 1] = Math.min(dp[j + 1], dp[j] + Math.abs(chairs.get(i) - peoples.get(j)));
            }
        }
        out.println(dp[n]);
        out.flush();
        out.close();
        in.close();
    }
}
