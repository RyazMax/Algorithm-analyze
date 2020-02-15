#ifndef FRACTALDRAWER_H
#define FRACTALDRAWER_H

#include "drawer.h"
#include "fractal.h"

class FractalDrawer
{
public:
    FractalDrawer(Drawer *drawer);
    virtual ~FractalDrawer();

    virtual void draw(size_t iter, bool delay);
protected:
    void _delay(bool delay);

    Drawer *_drawer;
};

#endif // FRACTALDRAWER_H
