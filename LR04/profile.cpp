#include <iostream>
#include <fstream>
#include <string>

#include "matrix.hpp"

using namespace std;

#define N 20

void experiment(int start, int stop, int step, string fileName) {

    ofstream out(fileName);
    for (int i = start; i <= stop; i+=step) {
        Matrix<int> a(i, i);
        Matrix<int> b(i, i);

        double orderTime = 0;
        double threadsTime[6] = {0};

        for (int j = 0; j < N; ++j) {
            // Обычный
            std::chrono::time_point<std::chrono::high_resolution_clock> start, end;
            start = std::chrono::high_resolution_clock::now();
            Matrix<int> res = Matrix<int>::mul(a, b);
            end = std::chrono::high_resolution_clock::now();
    
            double elapsed = std::chrono::duration_cast<std::chrono::nanoseconds>
                                (end-start).count();
            
            orderTime += elapsed;

            int count = 0;
            for (int k = 1; k <= 32; k *=2, count++) {
                std::chrono::time_point<std::chrono::high_resolution_clock> start, end;
                start = std::chrono::high_resolution_clock::now();
                Matrix<int> res = Matrix<int>::mul(a, b, k);
                end = std::chrono::high_resolution_clock::now();
    
                double elapsed = std::chrono::duration_cast<std::chrono::nanoseconds>
                                (end-start).count();

                threadsTime[count] += elapsed;
            }
        }

        orderTime /= N;
        cout << "SIZE " << i << " ";
        out << i << " ";
        cout << "O:" << (uint_least64_t)orderTime << " ";
        out << (uint_least64_t)orderTime << " ";
        for (int i = 0; i < 6; ++i) {
            threadsTime[i] /= N;
            cout << (uint_least64_t)threadsTime[i] << " ";
            out << (uint_least64_t)threadsTime[i] << " ";
        }
        out << endl;
        out.flush();
        cout << endl;
    }
    out.close();
}

int main(int argc, char **argv) {
    if (argc < 3) {
        cout << "Give please 2 names of file";
        return 1;
    }

    string fileName1(argv[1]);
    string fileName2(argv[2]);

    experiment(100, 1000, 100, fileName1);
    experiment(101, 1001, 100, fileName2);
}