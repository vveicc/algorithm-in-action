import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        StringTokenizer uv = new StringTokenizer(in.readLine());
        long u = Long.parseLong(uv.nextToken());
        long v = Long.parseLong(uv.nextToken());
        if (u > v) {
            out.println(-1);
        } else if (u == v) {
            if (u == 0) {
                out.println(0);
            } else {
                out.println(1);
                out.println(u);
            }
        } else {
            long d = v - u;
            if ((d & 1) == 1) {
                out.println(-1);
            } else {
                d >>= 1;
                if ((d & u) == 0) {
                    out.println(2);
                    out.printf("%d %d", u | d, d);
                } else {
                    out.println(3);
                    out.printf("%d %d %d", u, d, d);
                }
            }
        }
        out.flush();
        out.close();
        in.close();
    }
}
