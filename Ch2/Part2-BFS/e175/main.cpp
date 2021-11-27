#include <bits/stdc++.h>

using namespace std;

#define x first
#define y second

typedef pair< int, int > PII;

const int N = 510, M = N * N;

int  n, m;
char g[N][N];
int  dist[N][N];
bool st[N][N];

int bfs() {
    char cs[5] = "\\/\\/"; /* 四个字符\/\/,对于每个位置来说正确的方向,从左上顺时针4个方向  */
    int  dx[4] = {-1, -1, 1, 1}, dy[4] = {-1, 1, 1, -1};
    int  ix[4] = {-1, -1, 0, 0}, iy[4] = {-1, 0, 0, -1};

    deque< PII > q;
    memset(st, 0, sizeof st);
    memset(dist, 0x3f, sizeof dist);

    dist[0][0] = 0;
    q.push_back({0, 0});  //最开始0,0 在

    while (!q.empty()) {
        // 取队头
        auto t = q.front();
        q.pop_front();

        int x = t.x, y = t.y;
        if (x == n && y == m) {  // 如果是结果
            return dist[x][y];
        }
        if (st[x][y]) {  // 如果访问过
            continue;
        }

        st[x][y] = true;

        for (int i = 0; i < 4; ++i) {
            int a = x + dx[i], b = y + dy[i];
            if (a < 0 || a > n || b < 0 || b > m) {  // 出界，因为点要比n,m多1，所以等于不算出界
                continue;
            }
            int ga = x + ix[i], gb = y + iy[i];
            int w = (g[ga][gb] != cs[i]);
            int d = dist[x][y] + w;

            if (d < dist[a][b]) {
                dist[a][b] = d;

                if (w) {
                    q.push_back({a, b});
                } else {
                    q.push_front({a, b});
                }
            }
        }
    }
    return -1;
}

int main() {
    int T;
    scanf("%d", &T);

    while (T--) {
        scanf("%d%d", &n, &m);

        for (int i = 0; i < n; ++i) {
            scanf("%s", g[i]);
        }

        if ((n + m) & 1) {
            // 说明终点是奇点，一定走不到，因为起点第一步必须是偶点
            puts("NO SOLUTION");
        } else {
            printf("%d\n", bfs());
        }
    }

    return 0;
}
