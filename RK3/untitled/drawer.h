#ifndef DRAWER_H
#define DRAWER_H

#include <QLabel>
#include <QPainter>

class Drawer
{
public:
    Drawer(QLabel *canvas);
    ~Drawer() = default;

    int width() const;
    int height() const;

    void draw_line(int x1, int y1, int x2, int y2);
    void flush();
private:
    QPixmap _pxp;
    QLabel *_canvas;
};

#endif // DRAWER_H
