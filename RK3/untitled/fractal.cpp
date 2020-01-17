#include "fractal.h"


Fractal::Fractal(FracElem elem)
{
    _elems.push_back(elem);
}

std::vector<FracElem>::iterator Fractal::begin()
{
    return _elems.begin();
}

std::vector<FracElem>::iterator Fractal::end()
{
    return _elems.end();
}

void Fractal::add(const std::vector<FracElem> &elems)
{
    for (auto it = elems.begin(); it != elems.end(); ++it) {
        _elems.push_back(*it);
    }
}

std::vector<FracElem> Fractal::mutate(const FracElem &elem, int num)
{
    Point shift = elem.shift() * 3;
    int signX = 1;
    int signY = 1;
    Point shiftY(0, 1);
    Point shiftX(1, 0);

    if (num == 3 || num == 4 || num == 5) {
        signY = -1;
    }
    if (num == 1 || num == 4 || num == 7) {
        signX = -1;
    }

    if (signY == -1) {
        shift = shift + shiftY * 2;
        shiftY = shiftY * signY;
    }
    if (signX == -1) {
        shift = shift + shiftX * 2;
        shiftX = shiftX * signX;
    }

    std::vector<FracElem> to_ret;
    /*
     * ++|
     * |++
     */


    to_ret.push_back(FracElem({Point(0, 0), Point(0, signY)}, shift));
    to_ret.push_back(FracElem({Point(0, -signY), Point(0, signY)}, shift + shiftY));
    to_ret.push_back(FracElem({Point(0, -signY), Point(0, 0), Point(signX, 0)}, shift + shiftY * 2));
    to_ret.push_back(FracElem({Point(-signX, 0), Point(0, 0), Point(0, -signY)}, shift + shiftY * 2 + shiftX));
    to_ret.push_back(FracElem({Point(0, signY), Point(0, -signY)}, shift + shiftY + shiftX));
    to_ret.push_back(FracElem({Point(0, signY), Point(0, 0), Point(signX, 0)}, shift + shiftX));
    to_ret.push_back(FracElem({Point(-signX, 0), Point(0, 0), Point(0, signY)}, shift + shiftX * 2));
    to_ret.push_back(FracElem({Point(0, -signY), Point(0, signY)}, shift + shiftX * 2 + shiftY));
    to_ret.push_back(FracElem({Point(0, -signY), Point(0, 0)}, shift + shiftX * 2 + shiftY * 2));

    return to_ret;
}
