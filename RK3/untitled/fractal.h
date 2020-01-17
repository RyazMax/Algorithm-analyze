#ifndef FRACTAL_H
#define FRACTAL_H

#include "fracelem.h"

class Fractal
{
public:
    Fractal() = default;
    Fractal(FracElem elem);
    ~Fractal() = default;

    std::vector<FracElem>::iterator begin();
    std::vector<FracElem>::iterator end();

    void add(const std::vector<FracElem> &elems);

    static std::vector<FracElem> mutate(const FracElem &elem, int num);
private:
    std::vector<FracElem> _elems;
};

#endif // FRACTAL_H
