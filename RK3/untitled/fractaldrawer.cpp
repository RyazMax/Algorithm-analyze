#include "fractaldrawer.h"
#include <QApplication>

FractalDrawer::FractalDrawer(Drawer *drawer) : _drawer(drawer)
{
}

FractalDrawer::~FractalDrawer() {
    delete _drawer;
}

void FractalDrawer::draw(size_t iter, bool delay) {
    timespec t1, t2;
    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &t1);

    // Инциализация массива итераций
    int it_count[iter];
    double abs_shift = 1;
    for (size_t i = 0; i < iter; ++i) {
        it_count[i] = 0;
        abs_shift /= 3;
    }

    Point shift = Point(abs_shift, abs_shift);
    Point inv_x = Point(-1, 1), inv_y = Point(1, -1);

    Point scale = Point(_drawer->height(), _drawer->height());
    Point last = Point(0, 0);
    Point pos = Point(0, 0);

    for (;it_count[iter - 1] == 0;) {
        shift = shift * 0.5;
        Point shift_x = Point(shift.x(), 0), shift_y = Point(0, shift.y());
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
        last = pos + shift + shift_x * 4 + shift_y * 4;
        shift = shift * 2;
        pos = pos + shift * 3;

        // Обновление массива итераций
        for (size_t i = 0; i < iter; ++i) {
            it_count[i]++;
            if (it_count[i] == 9) {
                it_count[i] = 0;
            } else {
                if (it_count[i] == 3 || it_count[i] == 6) {
                    shift = shift * inv_y;
                } else {
                    shift = shift * inv_x;
                }
                break;
            }
        }
    }
    _drawer->flush();

    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &t2);

    qDebug("%u", t2.tv_nsec - t1.tv_nsec);
}

void FractalDrawer::_delay(bool delay)
{
    if (delay) {
        QApplication::processEvents(QEventLoop::AllEvents, 50);
        _drawer->flush();
    }
}
