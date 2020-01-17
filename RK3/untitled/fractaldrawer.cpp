#include "fractaldrawer.h"
#include <QApplication>

FractalDrawer::FractalDrawer(Drawer *drawer) : _drawer(drawer)
{
}

FractalDrawer::~FractalDrawer() {
    delete _drawer;
}

void FractalDrawer::draw(size_t iter, bool delay) {

    // Generate
    Fractal frac(FracElem({Point(0.5, 0), Point(0.5, 1)}));
    for (size_t i = 0; i < iter; ++i) {
        Fractal next;
        int count = 0;
        for (auto it = frac.begin(); it != frac.end(); ++it, ++count) {
            auto new_elems = Fractal::mutate(*it, count % 9);
            next.add(new_elems);
        }
        frac = next;
    }

    // Draw
    // Screen size
    int width = _drawer->width() - 1;
    int height = _drawer->height();
    Point scale(height, height);

    // Shifts
    double inner_scale = 2. / (pow(3, iter));
    // Draw edges

    Point last;
    for (auto it = frac.begin(); it != frac.end(); ++it) {
        Point shift = it->shift();
        Point p1 = ((*(it->begin()) + shift) * inner_scale + Point(1, 0.1)) * scale * 0.5;
        if (it != frac.begin()) _drawer->draw_line(last.x(), last.y(), p1.x(), p1.y());

        for (auto in_it = it->begin()+1; in_it != it->end(); ++in_it) {
            Point p2 = ((*in_it + shift) * inner_scale + Point(1, 0.1)) * scale * 0.5;
            _drawer->draw_line(p1.x(), p1.y(), p2.x(), p2.y());
            p1 = p2;
        }

        last = p1;
        if (delay) {
            QApplication::processEvents(QEventLoop::AllEvents, 50);
            _drawer->flush();
        }
    }
    _drawer->flush();
}
