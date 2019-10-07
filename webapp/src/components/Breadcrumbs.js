// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, {Component} from 'react';
import {T, getSelection} from "./Utils";

class Breadcrumbs extends Component {
    render() {
        let r = this.props.route;
        let boardsActive = r.section==='boards' && !r.sectionId ? 'active' : ''
        let osActive = r.section==='boards' && r.sectionId && !r.subsection ? 'active' : ''
        let appsActive = r.section==='boards' && r.sectionId && r.subsection ? 'active' : ''

        if (r.section==='confirm') {
            let board =  getSelection();
            console.log(board)
            r.section = 'boards'
            r.sectionId = board.board.id
            r.subsection = board.os.type + board.os.version
            r.subsubsection = 'confirm'
        }

        return (
            <div className="breadcrumbs">
                <div className="row">
                    <ul className="p-breadcrumbs">
                        <li className="p-breadcrumbs__item"><a href="/">{T('home')}</a></li>
                        <li className={'p-breadcrumbs__item ' + boardsActive}><a href="/boards">{T('boards')}</a></li>
                        {osActive || appsActive || r.subsubsection ?
                            <li className={'p-breadcrumbs__item ' + osActive}><a href={'/' + r.section + '/' + r.sectionId}>{T('os')}</a></li>
                            : ''}
                        {appsActive || r.subsubsection ?
                            <li className={'p-breadcrumbs__item ' + appsActive}><a href={'/' + r.section + '/' + r.sectionId + '/' + r.subsection}>{T('applications')}</a></li>
                            : ''}
                        {r.subsubsection ? <li className="p-breadcrumbs__item active"><a href="/confirm">{T('confirm')}</a></li> : ''}
                    </ul>
                </div>
            </div>
        );
    }
}

export default Breadcrumbs;