#include <iostream>
#include <vector>

using namespace std;

const int INF = 1000000000;

void deikstra(vector<vector<pair<int, int>>> &g, vector<int> &d, vector<int> &p, int s, int n) { // 1
    d[s] = 0;                                       // 2
	vector<char> u (n);                             // 3
	for (int i=0; i<n; ++i) {                       // 4
		int v = -1;                                 // 5
		for (int j=0; j<n; ++j)                     // 6
			if (!u[j] && (v == -1 || d[j] < d[v]))  // 7
				v = j;                              // 8
		if (d[v] == INF)                            // 9
			break;                                  // 10
		u[v] = true;                                // 11
 
		for (size_t j=0; j<g[v].size(); ++j) {      // 12
			int to = g[v][j].first;                 // 13
			int	len = g[v][j].second;               // 14
			if (d[v] + len < d[to]) {               // 15
				d[to] = d[v] + len;                 // 16
				p[to] = v;                          // 17
			}
		}
	}
}

int main() {
	int n;

    // Cписок смежности
	vector < vector < pair<int,int> > > g (n);
	// чтение графа ...
	int s = 0; // стартовая вершина
 
	vector<int> d (n, INF),  p (n);
	deikstra(g, d, p, s, n);
}