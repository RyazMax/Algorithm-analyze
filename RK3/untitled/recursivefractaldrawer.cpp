#include "recursivefractaldrawer.h"

RecursiveFractalDrawer::RecursiveFractalDrawer(Drawer *drawer) : FractalDrawer(drawer)
{

}

void RecursiveFractalDrawer::draw(size_t iter, size_t size) {
    if (iter == 1) {
        _drawer->draw_line(0, 0, 100, 100);
    } else {
        // DO RECURSIVE
        draw(iter - 1, size);
    }
}
