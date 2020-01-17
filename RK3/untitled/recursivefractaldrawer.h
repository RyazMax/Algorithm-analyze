#ifndef RECURSIVEFRACTALDRAWER_H
#define RECURSIVEFRACTALDRAWER_H

#include "fractaldrawer.h"

class RecursiveFractalDrawer : public FractalDrawer
{
public:
    RecursiveFractalDrawer(Drawer *drawer);
    ~RecursiveFractalDrawer() = default;

    void draw(size_t iter, size_t size);
};

#endif // RECURSIVEFRACTALDRAWER_H
