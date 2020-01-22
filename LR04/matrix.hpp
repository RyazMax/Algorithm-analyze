#include <vector>
#include <thread>

#include "matrixexception.h"

using std::vector;

template <class T>
class Matrix {
public:
    Matrix(size_t n, size_t m);

    ~Matrix() = default;

    vector<T>& operator[](size_t i);
    const vector<T>& operator[](size_t i) const;

    static Matrix<T> mul(const Matrix<T> &left, const Matrix<T> &right);
    static Matrix<T> mul(const Matrix<T> &left, const Matrix<T> &right, size_t threads);

    size_t rows() const;
    size_t cols() const;
private:
    static void rowsMul(std::vector<T> &mulH, const Matrix<T> &left, size_t i, size_t threads);
    static void colsMul(std::vector<T> &mulV, const Matrix<T> &right, size_t i, size_t threads);
    static void mainMull(const Matrix<T> &left, const Matrix<T> &right, Matrix<T>  &res,
                         const std::vector<T> &mulH, const std::vector<T> &mulV, size_t i, size_t threads);

    static void joinThreads(std::vector<std::thread> &vec);

    size_t _n;
    size_t _m;

    vector<vector<T>> _rows;
};

template <class T>
Matrix<T>::Matrix(size_t n, size_t m) : _rows(n), _n(n), _m(m) {
    for (size_t i = 0; i < n; ++i) {
        _rows[i] = vector<T>(m);
    }
}

template <class T>
size_t Matrix<T>::rows() const {
    return _n;
}

template <class T>
size_t Matrix<T>::cols() const {
    return _m;
}

template <class T>
vector<T>& Matrix<T>::operator[](size_t i) {
    return _rows[i];
}

template <class T>
const vector<T>& Matrix<T>::operator[](size_t i) const {
    return _rows[i];
}

template <class T>
void Matrix<T>::rowsMul(std::vector<T> &mulH, const Matrix &left, size_t i, size_t threads) {
    int nMinus = left.cols() - 1;
    int m = left.rows();

    for (; i < m; i+=threads) {
        for (int j = 0; j < nMinus; j+=2){
            mulH[i] -= left[i][j] * left[i][j+1];
        }
    }
}

template <class T>
void Matrix<T>::colsMul(std::vector<T> &mulV, const Matrix &right, size_t i, size_t threads) {
    int nMinus = right.rows() -1;
    int q = right.cols();
    
    for (; i < q; i+=threads) {
        for (int j = 0; j < nMinus; j += 2) {
            mulV[i] -= right[j][i] * right[j+1][i];
        }
    }
}

template <class T>
void Matrix<T>::mainMull(const Matrix<T> &left, const Matrix<T> &right, Matrix<T>  &res,
    const std::vector<T> &mulH, const std::vector<T> &mulV, size_t i, size_t threads) {
    
    int m = left.rows();
    int n = right.rows();
    int q = right.cols();
    int nMinus = n -1;
    bool isOdd = n & 1;

    for (; i < m; i += threads) {
        for (int j = 0; j < q; ++j) {
            T buf = mulH[i] + mulV[j];
            for (int k = 0; k < nMinus; k += 2) {
                buf += (left[i][k] + right[k+1][j]) * (left[i][k+1] + right[k][j]);
            }
            if (isOdd) {
                buf += left[i][nMinus] * right[nMinus][j];
            }
            res[i][j] = buf;
        }
    }                    
}

template <class T>
void Matrix<T>::joinThreads(std::vector<std::thread> &vec) {
    for (auto &t : vec) {
        t.join();
    }
}

template <class T>
Matrix<T> Matrix<T>::mul(const Matrix &left, const Matrix &right, size_t threads) {
    int m = left.rows();
    int n = right.rows();
    int q = right.cols();
    Matrix<T> res(m, q);

    std::vector<std::thread> trd(threads);

    std::vector<T> mulH(m);
    for (size_t i = 0; i < threads; ++i) {
        trd[i] = std::thread(rowsMul, std::ref(mulH), std::ref(left), i, threads);    
    }

    joinThreads(trd);

    std::vector<T> mulV(q);
    for (size_t i = 0; i < threads; ++i) {
        trd[i] = std::thread(colsMul, std::ref(mulV), std::ref(right), i, threads);
    }

    joinThreads(trd);

    for (int i = 0; i < threads; ++i) {
        trd[i] = std::thread(mainMull, std::ref(left), std::ref(right), std::ref(res), std::ref(mulH), std::ref(mulV), i, threads);
    }

    joinThreads(trd);

    return res;
}

template <class T>
Matrix<T> Matrix<T>::mul(const Matrix &left, const Matrix &right) {
    int m = left.rows();
    int n = right.rows();
    int q = right.cols();
    int nMinus = n -1;
    Matrix<T> res(m, q);


    std::vector<T> mulH(m);
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < nMinus; j+=2){
            mulH[i] -= left[i][j] * left[i][j+1];
        }
    }

    std::vector<T> mulV(q);
    for (int i = 0; i < q; ++i) {
        for (int j = 0; j < nMinus; j += 2) {
            mulV[i] -= right[j][i] * right[j+1][i];
        }
    }

    bool isOdd = n % 2 == 1;
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < q; ++j) {
            T buf = mulH[i] + mulV[j];
            for (int k = 0; k < nMinus; k += 2) {
                buf += (left[i][k] + right[k+1][j]) * (left[i][k+1] + right[k][j]);
            }
            if (isOdd) {
                buf += left[i][nMinus] * right[nMinus][j];
            }
            res[i][j] = buf;
        }
    }

    return res;
}