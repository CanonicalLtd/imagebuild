// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, { Component } from 'react';
import Footer from './components/Footer';
import Header from './components/Header';
import HeaderSlim from './components/HeaderSlim';
import Index from './components/Index';
import Boards from './components/Boards';
import OS from './components/OS';
import api from './components/api';
import {parseRoute, getSelection} from './components/Utils'
import Applications from "./components/Applications";
import Confirm from "./components/Confirm";

import createHistory from 'history/createBrowserHistory'
import Breadcrumbs from "./components/Breadcrumbs";
const history = createHistory()

class App extends Component {
    constructor(props) {
        super(props)
        this.state = {
            location: history.location,
            token: props.token || {},
            boards: [],
        }

        this.getBoards()
    }

    getBoards() {
        api.boardsList().then(response => {
            this.setState({boards: response.data.boards})
        })
    }

    renderBoards(sectionId, subsection) {
        if (!sectionId) {
            return <Boards boards={this.state.boards} />
        }

        // Get the board
        let board = this.state.boards.filter( b => {
            return b.id === sectionId
        })
        if (board.length === 0) {
            return
        }
        if (!subsection) {
            return <OS board={board[0]}/>
        }

        // Get the OS
        let os = board[0].os.filter( o => {
            return o.id === subsection
        })
        if (os.length === 0) {
            return
        }
        return <Applications board={board[0]} os={os[0]} />
    }

    renderConfirm(sectionId, subsection) {
        let board = getSelection()
        return <Confirm board={board} />
    }

    render() {
        const r = parseRoute()
        console.log(r)

        return (
          <div className="App">
              {r.section===''? <Header /> : ''}
              {r.section!==''? <HeaderSlim /> : ''}
              {r.section!==''? <Breadcrumbs route={r} /> : ''}

              <div className="content row">
                  {r.section===''? <Index /> : ''}
                  {r.section==='boards'? this.renderBoards(r.sectionId, r.subsection) : ''}
                  {r.section==='confirm'? this.renderConfirm(r.sectionId, r.subsection) : ''}
              </div>

              <Footer />
          </div>
        );
    }
}

export default App;
