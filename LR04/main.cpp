#include <iostream>
#include <chrono>

#include "matrix.hpp"

using std::cin;
using std::cout;
using std::endl;

int main() {
    int n, m, q;
    cout << "Input size of first matrix" << endl;
    cin >> n >> m;

    cout << "Input first matrix" << endl;
    Matrix<int> m1(n, m);
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j) {
            cin >> m1[i][j];
        }
    }

    cout << "Input number of columns of second matrix" << endl;
    cin >> q;
    Matrix<int> m2(m, q);

    cout << "Input second matrix" << endl;
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < q; ++j) {
            cin >> m2[i][j];
        }
    }

    int threadsNum;
    cout << "Input threads num" << endl;
    cin >> threadsNum;

    if (threadsNum < 1) {
        cout << "Threads should be positive" <<endl;
        return 1;
    }
    
    std::chrono::time_point<std::chrono::high_resolution_clock> start, end;
    start = std::chrono::high_resolution_clock::now();
    Matrix<int> res = Matrix<int>::mul(m1, m2, threadsNum);
    end = std::chrono::high_resolution_clock::now();
 
    int elapsed_seconds = std::chrono::duration_cast<std::chrono::nanoseconds>
                             (end-start).count();
    std::time_t end_time = std::chrono::high_resolution_clock::to_time_t(end);
 
    std::cout << "Execution time: " << elapsed_seconds << "ns\n";
    

    cout << "RESULT" << endl;
    for (int i = 0; i < res.rows(); ++i) {
        for (int j = 0; j < res.cols(); ++j) {
            cout << res[i][j] << " ";
        }
        cout << "\n";
    }
    return 0;
}