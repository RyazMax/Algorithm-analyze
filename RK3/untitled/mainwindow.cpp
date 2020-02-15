#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);

    QPixmap pxp(ui->canvas->width(), ui->canvas->height());
    pxp.fill();
    ui->canvas->setPixmap(pxp);
}

MainWindow::~MainWindow()
{
    delete ui;
}


void MainWindow::on_drawButton_clicked()
{
    bool ok;
    int it_count = ui->iterCount->text().toInt(&ok);
    if (!ok || it_count <= 0) {
        QMessageBox::warning(this, "Неверные данные", "Число итераций должно быть натуральным числом");
        return;
    }

    FractalDrawer *drawer;

    if (ui->iterButton->isChecked()) {
        drawer = new FractalDrawer(new Drawer(ui->canvas));
    } else {
        drawer = new RecursiveFractalDrawer(new Drawer(ui->canvas));
    }
    drawer->draw(it_count, ui->delay->isChecked());

    delete drawer;
}
