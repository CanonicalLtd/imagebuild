// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, {Component} from 'react';
import {T} from './Utils'

class Index extends Component {
    constructor(props) {
        super(props)
        this.state = {
            token: props.token || {},
        }
    }

    render() {
        return (
            <div className="row">

                <section className="row">
                    <div className="row">
                        <div className="first">
                            <h2>{T('get-started')}</h2>
                            <ul className="p-list">
                                <li className="p-list__item is-ticked">{T('site-description1')}</li>
                                <li className="p-list__item is-ticked">{T('site-description2')}</li>
                                <li className="p-list__item is-ticked">{T('site-description3')}</li>
                            </ul>
                            <p>{T('build')}</p>
                            <a className="p-button--brand" href="/boards" alt="">{T('choose-board')}</a>
                        </div>
                    </div>
                </section>
            </div>
        );
    }
}

export default Index;