// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, {Component} from 'react';
import {T, saveSelection, getSelection} from './Utils'
import Snaps from './Snaps'

class Applications extends Component {
    constructor(props) {
        super(props)
        this.state = {
            snaps: [],
        };

        let selected = getSelection()
        if (selected && selected.snaps && selected.snaps.length > 0) {
            this.state.snaps = selected.snaps
        }
    }
    renderSelected() {
        return (
            <div className="row summary">
                <p><i className="p-icon--success" />{this.props.board.name}</p>
                <p><i className="p-icon--success" />Ubuntu {T(this.props.os.type)} {this.props.os.version}</p>
            </div>
        )
    }

    handleSnapInstall = (snap) => {
        // Check we don't have the snap
        if (this.state.snaps.indexOf(snap) >= 0) {
            return
        }

        let snaps = this.state.snaps
        snaps.push(snap)
        this.setState({snaps: snaps})
    }

    handleSnapOnChange= (e) => {
        e.preventDefault();
        this.setState({snapName: e.target.value});
    }

    handleDialogCancel = (e) => {
        e.preventDefault();
    }

    handleRemove = (e) => {
        e.preventDefault();
        let snap = e.target.getAttribute('data-key')
        let snaps = this.state.snaps.filter(s => {
            return s !== snap
        })
        this.setState({snaps: snaps});
    }

    handleConfirm = (e) => {
        e.preventDefault();

        // Save the settings as JSON string to local storage
        let board = {
            board: this.props.board,
            os: this.props.os,
            snaps: this.state.snaps
        }
        saveSelection(board)

        // Navigate to the confirmation page
        window.location.href = '/confirm'
    }

    render() {
        return (
            <div className="first">

                <h2>{T('choose-applications')}</h2>
                <p>{T('choose-applications-desc')}</p>

                {this.renderSelected()}

                <div className="row">
                    <Snaps message={T('find-snaps')}
                           board={this.props.board}
                           os={this.props.os}
                           handleTextChange={this.handleSnapOnChange}
                           handleInstallClick={this.handleSnapInstall}
                           handleCancelClick={this.handleDialogCancel} />
                    <div className="col-4">
                        <p className="p-muted-heading">{T('pre-installed-snaps')}</p>
                        <hr className="lean" />
                        <ul className="p-list">
                        {this.state.snaps.map(s => {
                            return (
                                <li className="p-list__item hover">
                                    {s} &nbsp;
                                    <button data-key={s} className="p-button--neutral small hover__hover" title={T("remove-application")} onClick={this.handleRemove}>
                                        {T('remove')}
                                    </button>
                                </li>
                            )
                        })}
                        </ul>
                    </div>
                </div>

                <div>
                    <button className="p-button--brand" onClick={this.handleConfirm} disabled={this.state.snaps.length===0 ? 'disabled': ''}>{T('confirm')}</button>
                </div>
            </div>
        )
    }
}

export default Applications