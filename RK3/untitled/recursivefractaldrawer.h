#ifndef RECURSIVEFRACTALDRAWER_H
#define RECURSIVEFRACTALDRAWER_H

#include <QApplication>

#include "fractaldrawer.h"

class RecursiveFractalDrawer : public FractalDrawer
{
public:
    RecursiveFractalDrawer(Drawer *drawer);
    ~RecursiveFractalDrawer() = default;

    void draw(size_t iter, bool delay);
private:
    Point _draw(size_t iter, bool delay, Point pos, Point shift, Point last);
};

#endif // RECURSIVEFRACTALDRAWER_H
