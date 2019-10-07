// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import Messages from './Messages'

export function T(message) {
    const msg = Messages[message] || message;
    return msg
}

// URL is in the form:
//  /section
//  /section/sectionId
//  /section/sectionId/subsection
export function parseRoute() {
    const parts = window.location.pathname.split('/')

    switch (parts.length) {
        case 2:
            return {section: parts[1]}
        case 3:
            return {section: parts[1], sectionId: parts[2]}
        case 4:
            return {section: parts[1], sectionId: parts[2], subsection: parts[3]}
        default:
            return {}
    }
}

export function saveSelection(board) {
    let s = JSON.stringify(board)
    localStorage.setItem('board', s)
}

export function getSelection() {
    let s = localStorage.getItem('board')
    let b = JSON.parse(s)
    return b
}
