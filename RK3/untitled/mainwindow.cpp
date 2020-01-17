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
    int it_count = ui->iterCount->text().toInt();
    FractalDrawer drawer(new Drawer(ui->canvas));
    drawer.draw(it_count, ui->delay->isChecked());
}
