#include "point.h"

Point::Point(double x, double y) : _x(x), _y(y)
{

}

Point Point::operator*(double val) const
{
    return Point(_x * val, _y * val);
}

Point Point::operator*(const Point &val) const
{
    return Point(_x *val._x, _y * val._y);
}

Point Point::operator+(const Point &val) const
{
    return Point(_x + val._x, _y + val._y);
}

Point Point::operator-(const Point &val) const
{
    return Point(_x - val._x, _y - val._y);
}

double Point::x() const
{
    return _x;
}

double Point::y() const
{
    return _y;
}
