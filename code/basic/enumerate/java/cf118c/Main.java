import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));

    public static void main(String[] args) throws Exception {
        StringTokenizer nk = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nk.nextToken());
        int k = Integer.parseInt(nk.nextToken());
        String s = in.readLine();
        List<Integer>[] indexes = new List[10];
        for (int i = 0; i < 10; i++) {
            indexes[i] = new ArrayList<>();
        }
        for (int i = 0; i < n; i++) {
            indexes[s.charAt(i) & 15].add(i);
        }
        String mt = "";
        int mc = (int) 1e5;
        for (char b = '0'; b <= '9'; b++) {
            // 替换出k个b
            int cnt = indexes[b & 15].size();
            int cost = 0;
            char[] t = s.toCharArray();
            for (char d = 1; cnt < k; d++) {
                if (b + d <= '9') { // i+d => i 字典序降低，从前往后替换
                    List<Integer> p = indexes[(b + d) & 15];
                    for (int i = 0, m = p.size(); i < m && cnt < k; i++) {
                        t[p.get(i)] = b;
                        cost += d;
                        cnt++;
                    }
                }
                if (b - d >= '0') { // i-d => i 字典序升高，从后往前替换
                    List<Integer> p = indexes[(b - d) & 15];
                    for (int i = p.size() - 1; i >= 0 && cnt < k; i--) {
                        t[p.get(i)] = b;
                        cost += d;
                        cnt++;
                    }
                }
            }
            String str = new String(t);
            if (mc > cost) {
                mc = cost;
                mt = str;
            } else if (mc == cost && mt.compareTo(str) > 0) {
                mt = str;
            }
        }
        out.println(mc);
        out.println(mt);
        out.flush();
        out.close();
        in.close();
    }
}
