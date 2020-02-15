#include "drawer.h"

Drawer::Drawer(QLabel *canvas) : _canvas(canvas), _pxp(canvas->width(), canvas->height())
{
    _pxp.fill();
}

int Drawer::width() const
{
    return _canvas->width();
}

int Drawer::height() const
{
    return _canvas->height();
}


void Drawer::draw_line(int x1, int y1, int x2, int y2) {
    QPainter painter(&_pxp);
    painter.drawLine(x1, _canvas->height() - y1, x2, _canvas->height() - y2);
}

void Drawer::draw_line(Point p1, Point p2)
{
    draw_line(p1.x(), p1.y(), p2.x(), p2.y());
}

void Drawer::flush() {
    _canvas->setPixmap(_pxp);
}
