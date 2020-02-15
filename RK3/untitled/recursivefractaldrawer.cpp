#include "recursivefractaldrawer.h"
#include <ctime>


RecursiveFractalDrawer::RecursiveFractalDrawer(Drawer *drawer) : FractalDrawer(drawer)
{

}

void RecursiveFractalDrawer::draw(size_t iter, bool delay) {

    timespec t1, t2;
    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &t1);
    _draw(iter, delay, Point(0, 0), Point(1. / 3, 1. / 3), Point(0, 0));
    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &t2);

    qDebug("%d", t2.tv_nsec - t1.tv_nsec);
    _drawer->flush();
}

Point RecursiveFractalDrawer::_draw(size_t iter, bool delay, Point pos, Point shift, Point last)
{
    if (iter == 1) {
        shift = shift * 0.5;
        Point shift_x = Point(shift.x(), 0), shift_y = Point(0, shift.y());
        Point scale = Point(_drawer->height(), _drawer->height());

        _drawer->draw_line(last * scale, (pos + shift) * scale);
        _delay(delay);
        _drawer->draw_line((pos + shift) * scale, (pos + shift + shift_y * 4) * scale);
        _delay(delay);
        _drawer->draw_line((pos + shift + shift_y * 4) * scale, (pos + shift + shift_y * 4 + shift_x * 2) * scale);
        _delay(delay);
        _drawer->draw_line((pos + shift + shift_y * 4 + shift_x * 2) * scale, (pos + shift + shift_x * 2) * scale);
        _delay(delay);
        _drawer->draw_line((pos + shift + shift_x * 2) * scale, (pos + shift + shift_x * 4) * scale);
        _delay(delay);
        _drawer->draw_line((pos + shift + shift_x * 4) * scale, (pos + shift + shift_x * 4 + shift_y * 4) * scale);
        _delay(delay);

        return pos + shift + shift_x * 4 + shift_y * 4;
    } else {
        size_t niter = iter - 1;
        Point inv_x(-1, 1), inv_y(1, -1);
        Point shift_x(shift.x(), 0), shift_y(0, shift.y());
        shift = shift * (1. / 3);

        last = _draw(niter, delay, pos, shift, last);
        last = _draw(niter, delay, pos + shift_x + shift_y, shift * inv_x, last);
        last = _draw(niter, delay, pos + shift_y * 2, shift, last);
        last = _draw(niter, delay, pos + shift_x + shift_y * 3, shift * inv_y, last);
        last = _draw(niter, delay, pos + shift_x * 2 + shift_y * 2, shift * inv_y * inv_x, last);
        last = _draw(niter, delay, pos + shift_x + shift_y, shift * inv_y, last);
        last = _draw(niter, delay, pos + shift_x * 2, shift, last);
        last = _draw(niter, delay, pos + shift_x * 3 + shift_y, shift * inv_x, last);
        last = _draw(niter, delay, pos + shift_x * 2 + shift_y * 2, shift, last);

        return last;
    }
}
