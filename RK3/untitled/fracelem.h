#ifndef FRACELEM_H
#define FRACELEM_H

#include <vector>

#include "point.h"

class FracElem
{
public:
    FracElem(const std::vector<Point> &cords, const Point &shift = Point());
    ~FracElem() = default;

    std::vector<Point>::iterator begin();
    std::vector<Point>::iterator end();

    Point shift() const;

private:
    std::vector<Point> _cords;
    Point _shift;
};

#endif // FRACELEM_H
