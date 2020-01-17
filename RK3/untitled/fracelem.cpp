#include "fracelem.h"

FracElem::FracElem(const std::vector<Point> &cords, const Point &shift) : _cords(cords), _shift(shift)
{

}

std::vector<Point>::iterator FracElem::begin()
{
    return _cords.begin();
}

std::vector<Point>::iterator FracElem::end()
{
    return _cords.end();
}

Point FracElem::shift() const
{
    return _shift;
}
