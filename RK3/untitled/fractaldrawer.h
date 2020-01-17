#ifndef FRACTALDRAWER_H
#define FRACTALDRAWER_H

#include "drawer.h"
#include "fractal.h"

class FractalDrawer
{
public:
    FractalDrawer(Drawer *drawer);
    virtual ~FractalDrawer();

    void draw(size_t iter, bool delay);
protected:
    Drawer *_drawer;
};

#endif // FRACTALDRAWER_H
