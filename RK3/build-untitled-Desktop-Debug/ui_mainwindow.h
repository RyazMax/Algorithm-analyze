/********************************************************************************
** Form generated from reading UI file 'mainwindow.ui'
**
** Created by: Qt User Interface Compiler version 5.14.0
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_MAINWINDOW_H
#define UI_MAINWINDOW_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QButtonGroup>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QMainWindow>
#include <QtWidgets/QMenuBar>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QRadioButton>
#include <QtWidgets/QStatusBar>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_MainWindow
{
public:
    QWidget *centralwidget;
    QLabel *canvas;
    QLabel *label;
    QLineEdit *iterCount;
    QRadioButton *iterButton;
    QLabel *label_2;
    QRadioButton *recButton;
    QPushButton *drawButton;
    QRadioButton *delay;
    QMenuBar *menubar;
    QStatusBar *statusbar;
    QButtonGroup *buttonGroup;

    void setupUi(QMainWindow *MainWindow)
    {
        if (MainWindow->objectName().isEmpty())
            MainWindow->setObjectName(QString::fromUtf8("MainWindow"));
        MainWindow->resize(1231, 698);
        centralwidget = new QWidget(MainWindow);
        centralwidget->setObjectName(QString::fromUtf8("centralwidget"));
        canvas = new QLabel(centralwidget);
        canvas->setObjectName(QString::fromUtf8("canvas"));
        canvas->setGeometry(QRect(20, 20, 951, 661));
        label = new QLabel(centralwidget);
        label->setObjectName(QString::fromUtf8("label"));
        label->setGeometry(QRect(1020, 20, 191, 25));
        iterCount = new QLineEdit(centralwidget);
        iterCount->setObjectName(QString::fromUtf8("iterCount"));
        iterCount->setGeometry(QRect(1010, 50, 201, 33));
        iterButton = new QRadioButton(centralwidget);
        buttonGroup = new QButtonGroup(MainWindow);
        buttonGroup->setObjectName(QString::fromUtf8("buttonGroup"));
        buttonGroup->addButton(iterButton);
        iterButton->setObjectName(QString::fromUtf8("iterButton"));
        iterButton->setGeometry(QRect(1020, 120, 181, 31));
        label_2 = new QLabel(centralwidget);
        label_2->setObjectName(QString::fromUtf8("label_2"));
        label_2->setGeometry(QRect(1080, 90, 51, 25));
        recButton = new QRadioButton(centralwidget);
        buttonGroup->addButton(recButton);
        recButton->setObjectName(QString::fromUtf8("recButton"));
        recButton->setGeometry(QRect(1020, 150, 181, 31));
        drawButton = new QPushButton(centralwidget);
        drawButton->setObjectName(QString::fromUtf8("drawButton"));
        drawButton->setGeometry(QRect(1060, 380, 130, 34));
        delay = new QRadioButton(centralwidget);
        delay->setObjectName(QString::fromUtf8("delay"));
        delay->setGeometry(QRect(1040, 280, 160, 31));
        MainWindow->setCentralWidget(centralwidget);
        menubar = new QMenuBar(MainWindow);
        menubar->setObjectName(QString::fromUtf8("menubar"));
        menubar->setGeometry(QRect(0, 0, 1231, 30));
        MainWindow->setMenuBar(menubar);
        statusbar = new QStatusBar(MainWindow);
        statusbar->setObjectName(QString::fromUtf8("statusbar"));
        MainWindow->setStatusBar(statusbar);

        retranslateUi(MainWindow);

        QMetaObject::connectSlotsByName(MainWindow);
    } // setupUi

    void retranslateUi(QMainWindow *MainWindow)
    {
        MainWindow->setWindowTitle(QCoreApplication::translate("MainWindow", "MainWindow", nullptr));
        canvas->setText(QCoreApplication::translate("MainWindow", "TextLabel", nullptr));
        label->setText(QCoreApplication::translate("MainWindow", "\320\247\320\270\321\201\320\273\320\276 \320\270\321\202\320\265\321\200\320\260\321\206\320\270\320\271", nullptr));
        iterButton->setText(QCoreApplication::translate("MainWindow", "\320\230\321\202\320\265\321\200\320\260\321\202\320\270\320\262\320\275\321\213\320\271", nullptr));
        label_2->setText(QCoreApplication::translate("MainWindow", "\320\242\320\270\320\277", nullptr));
        recButton->setText(QCoreApplication::translate("MainWindow", "\320\240\320\265\320\272\321\203\321\200\321\201\320\270\320\262\320\275\321\213\320\271", nullptr));
        drawButton->setText(QCoreApplication::translate("MainWindow", "\320\240\320\270\321\201\320\276\320\262\320\260\321\202\321\214", nullptr));
        delay->setText(QCoreApplication::translate("MainWindow", "\320\227\320\260\320\264\320\265\321\200\320\266\320\272\320\260", nullptr));
    } // retranslateUi

};

namespace Ui {
    class MainWindow: public Ui_MainWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_MAINWINDOW_H
