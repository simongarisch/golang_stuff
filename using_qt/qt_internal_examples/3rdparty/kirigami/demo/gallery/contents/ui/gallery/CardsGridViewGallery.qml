/*
 *   Copyright 2018 Marco Martin <mart@kde.org>
 *
 *   This program is free software; you can redistribute it and/or modify
 *   it under the terms of the GNU Library General Public License as
 *   published by the Free Software Foundation; either version 2 or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU Library General Public License for more details
 *
 *   You should have received a copy of the GNU Library General Public
 *   License along with this program; if not, write to the
 *   Free Software Foundation, Inc.,
 *   51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */

import QtQuick 2.6
import QtQuick.Controls 2.0 as Controls
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.4 as Kirigami

Kirigami.ScrollablePage {
    id: page

    title: "Grid view of Cards"

        actions.main: Kirigami.Action {
        iconName: "documentinfo"
        text: "Info"
        checkable: true
        onCheckedChanged: sheet.sheetOpen = checked;
        shortcut: "Alt+I"
    }

    Kirigami.OverlaySheet {
        id: sheet
        onSheetOpenChanged: page.actions.main.checked = sheetOpen
        header: RowLayout {
            Kirigami.Heading {
                Layout.fillWidth: true
                text: "Cards Grid View"
            }
            Controls.ToolButton {
                text: "HIG..."
                enabled: false
                onClicked: Qt.openUrlExternally("")
            }
            Controls.ToolButton {
                text: "Source code..."
                onClicked: Qt.openUrlExternally("https://cgit.kde.org/kirigami.git/tree/examples/gallery/contents/ui/gallery/CardsGridViewGallery.qml")
            }
        }

        Controls.Label {
            property int implicitWidth: Kirigami.Units.gridUnit * 25
            wrapMode: Text.WordWrap
            text: "The Kirigami types AbstractCard and Card are used to implement the popular Card pattern used on many mobile and web platforms that is used to display a collection of information or actions.\n Besides the Card components, Kirigami offers also 3 kinds of views and positioners to help to present cards with beautiful and responsive layouts.\n\nIn this page, CardsGridView shows an example on how to put cards in a grid view, generated by a Qt model.\nThe behavior is same as CardsLayout, and it allows cards to be put in one or two columns depending from the available width.\nCardsGridView has the limitation that every Card must have the same exact height, so cellHeight must be manually set to a value in which the content fits for every item.\nIf possible use cards only when you don't need to instantiate that many and use CardsLayout intead."
        }
    }

    Component.onCompleted: {
        for (var i = 0; i < 50; ++i) {
            mainModel.append({"title": "Item " + i,
                "image": "../banner.jpg",
                "text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id risus id augue euismod accumsan. Nunc vestibulum placerat bibendum.",
                "actions": [{text: "Action 1", icon: "add-placemark"},
                            {text: "Action 2", icon: "address-book-new-symbolic"}]
            })
        }
    }
    Kirigami.CardsGridView {
        id: view
        model: ListModel {
            id: mainModel
        }

//property Component delegate
        delegate:Kirigami.Card {
            id: card
            banner {
                title: model.title
                imageSource: model.image
            }
            contentItem: Controls.Label {
                wrapMode: Text.WordWrap
                text: model.text
            }
            //HACK: this instantiator hack is just for demonstration purposes, normally the actions objects should be embedded as a role of a QAbstractItemModel, either as QActions or just QObjects with the proper properties and signals (the new qqc2 Action should ideally become a public c++ type)
            property var actionsModel: model.actions
            Instantiator {
                model: actionsModel
                delegate: Kirigami.Action {
                    text: model.text
                    icon.name: model.icon
                    onTriggered: showPassiveNotificaton(model.text + " triggered")
                }
                onObjectAdded: {
                    card.actions.push(object)
                }
            }
        }
    }
}
