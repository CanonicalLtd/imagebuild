// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, {Component} from 'react';
import {T} from './Utils'

class Boards extends Component {
    renderBoards() {
        if (!this.props.boards) return;

        return this.props.boards.map(brd => {
            return (
                <div className="p-card board">
                    <div className="p-card__title">
                        <img src={'/static/images/' + brd.id + '.png'} />
                    </div>
                    <div className="p-card__content">
                        <hr />
                        <h4 >{brd.name}</h4>
                        <a href={'/boards/' + brd.id} className="p-button--neutral is-inline">{T('select-board')}</a>
                    </div>
                </div>
            )
        })
    }

    render() {
        return (
            <div className="first">
                <h2>{T('choose-board')}</h2>
                <p>{T('choose-board-desc')}</p>

                {this.renderBoards()}
            </div>
        );
    }
}

export default Boards;