#ifndef POINT_H
#define POINT_H

#include <cmath>

class Point
{
public:
    Point(double x = 0, double y = 0);
    ~Point() = default;

    Point operator*(double val) const;
    Point operator*(const Point &val) const;

    Point operator+(const Point &val) const;
    Point operator-(const Point &val) const;

    double x() const;
    double y() const;
private:
    double _x;
    double _y;
};

#endif // POINT_H
