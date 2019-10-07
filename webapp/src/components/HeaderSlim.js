// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, { Component } from 'react';

class HeaderSlim extends Component {
    render() {
        return (
            <header id="navigation" class="p-navigation header-slim">
                <div className="p-navigation__banner row">
                    <div className="p-navigation__logo">
                        <div className="u-vertically-center">
                            <img src="/static/images/logo.png" width="150px"  />
                        </div>
                    </div>
                </div>
            </header>
        );
    }
}

export default HeaderSlim;
