// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, { Component } from 'react';
import HeaderSlim from "./HeaderSlim";
import {T} from './Utils'


class Header extends Component {
  render() {
    return (
      <div>
          <HeaderSlim />
          <section className="p-strip--image is-dark header">
              <div className="row">
                    <div className="col-5 title">
                        <h1>{T('title')}</h1>
                        <p>{T('subtitle')}</p>
                    </div>
              </div>
          </section>
      </div>
    );
  }
}

export default Header;
