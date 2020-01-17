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

void Drawer::flush() {
    _canvas->setPixmap(_pxp);
}
