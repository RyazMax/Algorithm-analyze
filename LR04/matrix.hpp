#include <vector>

#include "matrixexception.h"

using std::vector;

template <class T>
class Matrix {
public:
    Matrix(size_t n, size_t m);

    ~Matrix() = default;

    vector<T>& operator[](size_t i);
    const vector<T>& operator[](size_t i) const;

    Matrix<T> mul(const Matrix<T> &right, size_t threads = 1) const;
private:
    vector<vector<T>> _rows;
};

template <class T>
Matrix<T>::Matrix(size_t n, size_t m) : _rows(n) {
    for (size_t i = 0; i < n; ++i) {
        _rows[i] = vector<T>(m);
    }
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
Matrix<T> Matrix<T>::mul(const Matrix &right, size_t threads) const {
    return Matrix(this->_rows.size(), right._rows.size());
}