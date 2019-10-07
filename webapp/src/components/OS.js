// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, {Component} from 'react';
import {T} from './Utils'

class OS extends Component {
    renderOS() {
        return this.props.board.os.map(o => {
            let id=o.type + o.version
            return (
                <div className="p-card os" key={id}  onClick={this.handleClick}>
                    <h3 className="u-vertically-center">
                        {T(o.type)} {T(o.version)}
                    </h3>
                    <a href={'/boards/' + this.props.board.id + '/' + id} className="p-button--neutral is-inline">{T('select-version')}</a>
                </div>
            )
        })
    }

    render() {
        return (
            <div className="first">
                <h2>{T('choose-os')}</h2>
                <p>{T('choose-os-desc')}</p>

                <div className="summary">
                    <p><i className="p-icon--success" />{this.props.board.name}</p>
                </div>

                {this.renderOS()}
            </div>
        );
    }
}


export default OS;